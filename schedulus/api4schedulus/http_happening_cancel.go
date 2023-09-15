package api4schedulus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/facade4schedulus"
	"net/http"
)

var cancelHappening = facade4schedulus.CancelHappening

// httpCancelHappening marks happening as canceled
func httpCancelHappening(w http.ResponseWriter, r *http.Request) {
	var happeningRequest = getHappeningRequestParamsFromURL(r)
	request := dto4schedulus.CancelHappeningRequest{
		HappeningRequest: happeningRequest,
	}
	ctx, user, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.DefaultMaxJSONRequestSize,
		AuthRequired:     true,
	}, &request)
	if err != nil {
		return
	}
	err = cancelHappening(ctx, user, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}