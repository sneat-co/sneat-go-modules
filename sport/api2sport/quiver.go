package api2sport

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/httpserver"
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/sport/facade2sport"
	"github.com/sneat-co/sneat-go-modules/sport/models2sport"
	"net/http"
)

const (
	quiverPathPrefix = "/v0/quiver/"
	//quiverMyPathPrefix     = quiverPathPrefix + "my/"
	quiverWantedPathPrefix = quiverPathPrefix + "wanted/"
)

func registerQuiverHandlers(handle modules.HTTPHandleFunc) {
	handle(http.MethodPost, quiverWantedPathPrefix+"create_wanted", createWantedItem)
	handle(http.MethodPut, quiverWantedPathPrefix+"update_wanted", updateWantedItem)
	handle(http.MethodDelete, quiverWantedPathPrefix+"delete_wanted", deleteWantedItem)
}

func createWantedItem(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.KB * 100,
		AuthRequired:     false,
	})
	if err != nil {
		return
	}
	wanted := models2sport.Wanted{}
	request := facade2sport.CreateWantedRequest{
		Wanted: wanted,
	}
	if err := apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	var id string
	if id, err = facade2sport.CreateWanted(ctx, userContext, request); err != nil {
		httpserver.HandleError(err, "createWantedItem", w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(id))
}

func updateWantedItem(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func deleteWantedItem(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.KB * 100,
		AuthRequired:     false,
	})
	if err != nil {
		return
	}
	request := facade2sport.DeleteWantedRequest{ID: r.URL.Query().Get("id")}
	if err = facade2sport.DeleteWanted(ctx, userContext, request); err != nil {
		httpserver.HandleError(err, "deleteWantedItem", w, r)
	}
	w.WriteHeader(http.StatusOK)
}
