package api4scrumus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/httpserver"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dto4teamus"
	"github.com/sneat-co/sneat-go-modules/meetingus/facade4meetingus"
	"github.com/sneat-co/sneat-go-modules/scrumus/facade4scrumus"
	"net/http"
)

var deleteTask = facade4scrumus.DeleteTask

// httpDeleteTask is an API endpoint that delete a task
func httpDeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{AuthRequired: true})
	if err != nil {
		httpserver.HandleError(err, "httpDeleteTask", w, r)
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
		httpserver.HandleError(err, "httpDeleteTask", w, r)
		return
	}
	err = deleteTask(ctx, userContext, request)
	apicore.IfNoErrorReturnOK(ctx, w, r, err)
}