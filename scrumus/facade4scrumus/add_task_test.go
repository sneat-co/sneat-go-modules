package facade4scrumus

import (
	"context"
	"github.com/sneat-co/sneat-core-modules/teamus/dto4teamus"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/meetingus/facade4meetingus"
	testdb "github.com/sneat-co/sneat-go-testdb"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	userContext := facade.NewUser("user1")

	_ = testdb.NewMockDB(t, testdb.WithProfile1())

	t.Run("empty request", func(t *testing.T) {
		if _, err := AddTask(context.Background(), userContext, AddTaskRequest{}); err == nil {
			t.Fatal("should fail on empty request")
		}
	})

	t.Run("valid_requests", func(t *testing.T) {
		now := time.Now()
		request := AddTaskRequest{
			TaskRequest: TaskRequest{
				Request: facade4meetingus.Request{
					TeamRequest: dto4teamus.TeamRequest{
						TeamID: "team1",
					},
					MeetingID: now.Format("2006-01-02"),
				},
				ContactID: "m1",
				Type:      "done",
				Task:      "done1",
			},
			Title: "Test task",
		}

		t.Run("create_new_scrum", func(t *testing.T) {
			if _, err := AddTask(context.Background(), userContext, request); err != nil {
				t.Fatalf("should not fail on valid request, got: %v", err)
			}
		})

		//t.Run("update_existing_scrum", func(t *testing.T) {
		//	if _, err := AddTask(context.Background(), userContext, request); err != nil {
		//		t.Fatalf("should not fail on valid request, got: %v", err)
		//	}
		//})
	})
}
