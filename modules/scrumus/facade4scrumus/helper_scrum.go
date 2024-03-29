package facade4scrumus

import (
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/facade4meetingus"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/dal4scrumus"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/models4scrumus"
	"github.com/sneat-co/sneat-go-modules/modules/teamus/models4teamus"
)

type worker = func(ctx context.Context, tx dal.ReadwriteTransaction, params facade4meetingus.WorkerParams) (err error)

func runScrumWorker(ctx context.Context, userContext facade.User, request facade4meetingus.Request, worker worker) error {
	return facade4meetingus.RunMeetingWorker(ctx, userContext.GetID(), request, MeetingRecordFactory{}, worker)
}

// UpdateLastScrumIDIfNeeded updates scrum if needed
func UpdateLastScrumIDIfNeeded(
	ctx context.Context,
	tx dal.ReadwriteTransaction,
	params facade4meetingus.WorkerParams,
) (err error) {

	scrumTeamUpdates := make([]dal.Update, 0, 1)
	scrumID := params.Meeting.GetID()
	scrum := params.Meeting.Record.Data().(*models4scrumus.Scrum)

	var scrumTeam dal4scrumus.ScrumTeam
	scrumTeam, err = dal4scrumus.GetScrumTeam(ctx, tx, params.Team.ID)
	if err != nil && !dal.IsNotFound(err) {
		return
	}
	if lastScrum := scrumTeam.Data.Last; lastScrum != nil && lastScrum.ID != "" && lastScrum.ID < scrumID {
		if scrum.ScrumIDs == nil {
			scrum.ScrumIDs = &models4scrumus.ScrumIDs{}
		}
		scrum.ScrumIDs.Prev = lastScrum.ID
		prevScrumKey := dal.NewKeyWithParentAndID(params.Team.Key, "api4meetingus", lastScrum.ID)
		prevScrum := new(models4scrumus.Scrum)
		prevScrumRecord := dal.NewRecordWithData(prevScrumKey, prevScrum)
		if err = tx.Get(ctx, prevScrumRecord); err != nil {
			return
		}
		if prevScrum.ScrumIDs == nil || prevScrum.ScrumIDs.Next == "" {
			if err = prevScrum.Validate(); err != nil {
				return
			}
			prevScrumUpdates := []dal.Update{{Field: "scrumIds.next", Value: scrumID}}
			if err = tx.Update(ctx, prevScrumKey, prevScrumUpdates); err != nil {
				return
			}
		}
	}
	if scrumTeam.Data.Last == nil || scrumTeam.Data.Last.ID < scrumID {
		scrumTeam.Data.Last = &models4teamus.TeamMeetingInfo{
			ID:       scrumID,
			Stage:    "planning",
			Started:  scrum.Started,
			Finished: scrum.Finished,
		}
		scrumTeamUpdates = append(scrumTeamUpdates, dal.Update{
			Field: "last",
			Value: scrumTeam.Data.Last,
		})
	}
	if len(scrumTeamUpdates) > 0 {
		if err = scrumTeam.Data.Validate(); err != nil {
			return
		}
		if scrumTeam.Record.Exists() {
			if err = tx.Update(ctx, scrumTeam.Key, scrumTeamUpdates); err != nil {
				return fmt.Errorf("failed to update scrum team record: %w", err)
			}
		} else if err = tx.Insert(ctx, scrumTeam.Record); err != nil {
			return fmt.Errorf("failed to insert new scrum team record")
		}
	}
	return
}
