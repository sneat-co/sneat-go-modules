package api4scrumus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-core/httpserver"
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/facade4meetingus"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/facade4scrumus"
	"github.com/sneat-co/sneat-go-modules/modules/teamus/dto4teamus"
	"net/http"
)

var deleteTask = facade4scrumus.DeleteTask

// httpDeleteTask is an API endpoint that delete a task
func httpDeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, verify.DefaultJsonWithAuthRequired)
	if err != nil {
		httpserver.HandleError(ctx, err, "httpDeleteTask", w, r)
		return
	}
	query := r.URL.Query()
	request := facade4scrumus.DeleteTaskRequest{
		Request: facade4meetingus.Request{
			TeamRequest: dto4teamus.TeamRequest{
				TeamID: query.Get("team"),
			},
			MeetingID: query.Get("date"),
		},
		Task:      query.Get("id"),
		Type:      query.Get("type"),
		ContactID: query.Get("members"),
	}

	if err := request.Validate(); err != nil {
		httpserver.HandleError(ctx, err, "httpDeleteTask", w, r)
		return
	}
	err = deleteTask(ctx, userContext, request)
	apicore.IfNoErrorReturnOK(ctx, w, r, err)
}
