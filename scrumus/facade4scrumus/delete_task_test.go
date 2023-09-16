package facade4scrumus

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-core-modules/teamus/dto4teamus"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/meetingus/facade4meetingus"
	testdb "github.com/sneat-co/sneat-go-testdb"
	"testing"
	"time"
)

func TestDeleteTask(t *testing.T) {

	db := testdb.NewMockDB(t, testdb.WithProfile1())

	facade.GetDatabase = func(ctx context.Context) dal.DB {
		return db
	}

	userContext := facade.NewUser("user1")

	ctx := context.Background()

	t.Run("empty_request", func(t *testing.T) {
		if err := DeleteTask(ctx, userContext, DeleteTaskRequest{}); err == nil {
			t.Fatal("Should fail on empty request")
		}
	})

	t.Run("valid_request", func(t *testing.T) {
		now := time.Now()
		request := DeleteTaskRequest{
			Request: facade4meetingus.Request{
				TeamRequest: dto4teamus.TeamRequest{
					TeamID: "team1",
				},
				MeetingID: now.Format("2006-01-02"),
			},
			ContactID: "m1",
			Type:      "done",
			Task:      "d1",
		}

		t.Run("no_tasks", func(t *testing.T) {
			if err := DeleteTask(ctx, userContext, request); err != nil {
				t.Error(err)
			}
		})

		t.Run("existing_task", func(t *testing.T) {
			if err := DeleteTask(ctx, userContext, request); err != nil {
				t.Error(err)
			}
		})
	})

}
