package facade4retrospectus

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-core-modules/memberus/briefs4memberus"
	"github.com/sneat-co/sneat-go-modules/meetingus/facade4meetingus"
	"github.com/sneat-co/sneat-go-modules/retrospectus/models4retrospectus"
	"github.com/strongo/validation"
	"time"
)

// StartRetroReview starts review
func StartRetroReview(ctx context.Context, userContext facade.User, request RetroRequest) (response RetrospectiveResponse, err error) {
	if err = request.Validate(); err != nil {
		return
	}
	uid := userContext.GetID()
	err = runRetroWorker(ctx, uid, request,
		func(ctx context.Context, tx dal.ReadwriteTransaction, params facade4meetingus.WorkerParams) error {
			retrospective := params.Meeting.Record.Data().(*models4retrospectus.Retrospective)
			retrospective.Stage = models4retrospectus.StageReview
			now := time.Now()
			retrospective.TimeLastAction = &now

			var teamRetroUpdates []dal.Update
			if teamRetroUpdates, err = moveRetroItemsFromUsers(ctx, tx, params); err != nil {
				return err
			}

			teamRetroUpdates = append(teamRetroUpdates,
				dal.Update{Field: "stage", Value: retrospective.Stage},
				dal.Update{Field: "timeLastAction", Value: retrospective.TimeLastAction},
			)

			//retrospetiveKey := dal.NewKeyWithID("api4meetingus", ret)
			if err = txUpdateRetrospective(ctx, tx, params.Meeting.Key, retrospective, teamRetroUpdates); err != nil {
				return err
			}

			response.ID = request.MeetingID
			response.Data = retrospective
			return err
		})
	return
}

func moveRetroItemsFromUsers(ctx context.Context, tx dal.ReadwriteTransaction, params facade4meetingus.WorkerParams) (teamRetrosUpdates []dal.Update, err error) {
	retrospective := params.Meeting.Record.Data().(*models4retrospectus.Retrospective)

	//wg := sync.WaitGroup{}
	userRetroRecords := make([]dal.Record, len(retrospective.Contacts))
	for _, member := range retrospective.Contacts {
		if member.UserID != "" && member.HasRole(briefs4memberus.TeamMemberRoleContributor) {
			userRetroRecords = append(userRetroRecords, getUserRetroRecord(member.UserID, params.Team.ID, new(models4retrospectus.Retrospective)))
		}
		//}
	}
	if err = tx.GetMulti(ctx, userRetroRecords); err != nil {
		return
	}
	retroItems := make([]interface{}, 0)

	countsByMemberAndType := make(map[string]map[string]int, 0)
	for i, ur := range userRetroRecords {
		if ur == nil {
			continue
		}
		userRetro := ur.Data().(*models4retrospectus.Retrospective)

		if len(userRetro.Items) == 0 {
			continue
		}
		userRetroRecord := userRetroRecords[i]
		uid := userRetroRecord.Key().Parent().ID.(string)
		userCounts := make(map[string]int)
		countsByMemberAndType[uid] = userCounts
		for i, retroItem := range userRetro.Items {
			newItemErr := func(message string) error {
				return validation.NewErrBadRecordFieldValue(
					fmt.Sprintf("items[%v]{id=%v}", i, retroItem.ID),
					message,
				)
			}
			if retroItem.Type == "" {
				err = newItemErr("user's retro item has no type")
				return
			}
			if len(retroItem.Children) > 0 {
				err = newItemErr("user's retro item has child items")
				return
			}
			userCounts[retroItem.Type]++
			retroItems = append(retroItems, retroItem)
		}
		if err = txUpdate(ctx, tx, userRetroRecords[i].Key(), []dal.Update{
			{Field: "items", Value: firestore.Delete},
			{Field: "countsByMemberAndType", Value: firestore.Delete},
		}); err != nil {
			return
		}
	}
	if len(retroItems) > 0 {
		teamRetrosUpdates = []dal.Update{
			{Field: "items", Value: firestore.ArrayUnion(retroItems...)},
		}
	}
	return
}
