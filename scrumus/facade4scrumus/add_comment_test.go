package facade4scrumus

import (
	"context"
	"encoding/json"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dto4teamus"
	"github.com/sneat-co/sneat-go-core/tests"
	"github.com/sneat-co/sneat-go-modules/meetingus/facade4meetingus"
	"strings"
	"testing"
)

func TestAddCommentRequest_Validate(t *testing.T) {
	request := AddCommentRequest{}
	err := request.Validate()
	if err == nil {
		t.Fatal("expected to get error on empty request")
	}
	request = AddCommentRequest{
		TaskRequest: TaskRequest{
			ContactID: "m1",
			Type:      "done",
			Task:      "task1",
			Request: facade4meetingus.Request{
				TeamRequest: dto4teamus.TeamRequest{
					TeamID: "team1",
				},
				MeetingID: "2020-12-13",
			},
		},
		Message: "message 1",
	}
	if err = request.Validate(); err != nil {
		t.Fatalf("unexpected error on valid request: %v", err)
	}
}

func TestAddComment(t *testing.T) {
	//userContext := facade4meetingus.NewUser("user1")

	db := tests.NewMockDB(t, tests.WithProfile1())

	facade.GetDatabase = func(ctx context.Context) dal.DB {
		return db
	}

	t.Run("add 1st comment", func(t *testing.T) {
		body := []byte(strings.Replace(strings.Replace(`{
	"teamID":"team1",
	"meetingID":"2019-11-22",
	"type":"done",
	"task":"d1","
	members":"m1",
	"message":"msg1"
}`, "\n", "", -1), "\t", "", -1))
		var request AddCommentRequest
		if err := json.Unmarshal(body, &request); err != nil {
			t.Fatal(err)
		}

		//comment, err := AddComment(ctx, userContext, request)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//if comment.InviteID == "" {
		//	t.Error("InviteID is not set")
		//}
	})
}
