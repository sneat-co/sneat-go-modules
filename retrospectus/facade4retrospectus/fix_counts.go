package facade4retrospectus

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-core-modules/teamus/dal4teamus"
	models4userus2 "github.com/sneat-co/sneat-core-modules/userus/models4userus"
	"github.com/sneat-co/sneat-go-modules/retrospectus/dal4retrospectus"
	"github.com/sneat-co/sneat-go-modules/retrospectus/models4retrospectus"
	"time"
)

// FixCounts fixes counts
func FixCounts(ctx context.Context, userContext facade.User, request FixCountsRequest) (err error) {
	uid := userContext.GetID()
	db := facade.GetDatabase(ctx)
	return db.RunReadwriteTransaction(ctx, func(ctx context.Context, tx dal.ReadwriteTransaction) error {
		now := time.Now()
		userRef := models4userus2.NewUserKey(uid)
		team := dal4teamus.NewTeamContext(request.TeamID)
		var retroTeam dal4retrospectus.RetroTeam
		retroTeam, err = dal4retrospectus.GetRetroTeam(ctx, tx, request.TeamID)
		user := new(models4userus2.UserDto)
		userRecord := dal.NewRecordWithData(userRef, user)

		if err := tx.GetMulti(ctx, []dal.Record{userRecord, team.Record}); err != nil {
			return err
		}
		if retroTeam.Data.UpcomingRetro == nil {
			retroTeam.Data.UpcomingRetro = &models4retrospectus.RetrospectiveCounts{
				ItemsByUserAndType: make(map[string]map[string]int),
			}
		}
		teamInfo := user.GetUserTeamInfoByID(request.TeamID)
		updates := make([]dal.Update, 0, 1)
		if teamInfo == nil {
			if _, ok := retroTeam.Data.UpcomingRetro.ItemsByUserAndType[uid]; ok {
				delete(retroTeam.Data.UpcomingRetro.ItemsByUserAndType, uid)
				if len(retroTeam.Data.UpcomingRetro.ItemsByUserAndType) == 0 {
					retroTeam.Data.UpcomingRetro = nil
					updates = append(updates, dal.Update{Field: "upcomingRetro", Value: firestore.Delete})
				} else {
					path := fmt.Sprintf("upcomingRetro.itemsByUserAndType.%v", uid)
					updates = append(updates, dal.Update{Field: path, Value: firestore.Delete})
				}
			}
		} else {
			//for itemType, items := range teamInfo.RetroItems {
			//	count := len(items)
			//	if v, ok := team.Data.UpcomingRetro.ItemsByUserAndType[uid][itemType]; !ok || v != count {
			//		path := fmt.Sprintf("upcomingRetro.itemsByUserAndType.%v.%v", uid, itemType)
			//		if count == 0 {
			//			delete(team.Data.UpcomingRetro.ItemsByUserAndType[uid], itemType)
			//			updates = append(updates, dal.Update{Field: path, Value: firestore.Delete})
			//		} else {
			//			team.Data.UpcomingRetro.ItemsByUserAndType[uid][itemType] = count
			//			updates = append(updates, dal.Update{Field: path, Value: count})
			//		}
			//	}
			//}
			if len(retroTeam.Data.UpcomingRetro.ItemsByUserAndType[uid]) == 0 {
				delete(retroTeam.Data.UpcomingRetro.ItemsByUserAndType, uid)
				updates = []dal.Update{{Field: "upcomingRetro", Value: firestore.Delete}}
			}
		}
		if len(updates) > 0 {
			if err = txUpdateTeam(ctx, tx, now, team, updates); err != nil {
				return err
			}
		}
		return nil
	})
}
