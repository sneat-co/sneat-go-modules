package api4schedulus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/facade4schedulus"
	"net/http"
)

var revokeHappeningCancellation = facade4schedulus.RevokeHappeningCancellation

// httpRevokeHappeningCancellation marks happening as canceled
func httpRevokeHappeningCancellation(w http.ResponseWriter, r *http.Request) {
	var happeningRequest = getHappeningRequestParamsFromURL(r)
	request := dto4schedulus.CancelHappeningRequest{
		HappeningRequest: happeningRequest,
	}
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.DefaultMaxJSONRequestSize,
		AuthRequired:     true,
	}, &request)
	if err != nil {
		return
	}
	err = revokeHappeningCancellation(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
