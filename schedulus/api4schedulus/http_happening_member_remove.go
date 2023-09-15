package api4schedulus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/facade4schedulus"
	"net/http"
)

var removeMemberFromHappening = facade4schedulus.RemoveMemberFromHappening

func httpRemoveMemberFromHappening(w http.ResponseWriter, r *http.Request) {
	var request dto4schedulus.HappeningContactRequest
	request.HappeningRequest = getHappeningRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.DefaultMaxJSONRequestSize,
		AuthRequired:     true,
	}, &request)
	if err != nil {
		return
	}
	err = removeMemberFromHappening(ctx, userContext.GetID(), request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
