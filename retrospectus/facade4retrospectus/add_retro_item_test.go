package facade4retrospectus

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dto4teamus"
	"github.com/sneat-co/sneat-go-core/tests"
	"github.com/sneat-co/sneat-go-modules/meetingus/facade4meetingus"
	"testing"
)

func TestAddRetroItem(t *testing.T) {

	_ = tests.NewMockDB(t, tests.WithProfile1())

	userContext := facade.NewUser("user1")
	t.Run("should_succeed", func(t *testing.T) {
		t.Run("upcoming_retrospective", func(t *testing.T) {
			newTeamKey = func(id string) *dal.Key {
				return dal.NewKeyWithID(dal4teamus.TeamsCollection, id)
			}

			request := AddRetroItemRequest{
				RetroItemRequest: RetroItemRequest{
					Request: facade4meetingus.Request{
						TeamRequest: dto4teamus.TeamRequest{
							TeamID: "team1",
						},
						MeetingID: UpcomingRetrospectiveID,
					},
					Type: "good",
				},
				Title: "Good # 1",
			}

			ctx := context.Background()

			_, _ = AddRetroItem(ctx, userContext, request)
			//if _, _ = AddRetroItem(ctx, userContext, request); false {
			// TODO: t.Fatalf("failed to add retro item: %v", err)
			//}
		})
	})
}
