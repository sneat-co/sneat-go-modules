package api4schedulus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/facade4schedulus"
	"net/http"
)

var createHappening = facade4schedulus.CreateHappening

// httpPostCreateHappening creates recurring happening
func httpPostCreateHappening(w http.ResponseWriter, r *http.Request) {
	var request dto4schedulus.CreateHappeningRequest
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.DefaultMaxJSONRequestSize,
		AuthRequired:     true,
	}, &request)
	if err != nil {
		return
	}
	response, err := createHappening(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, &response)
}
