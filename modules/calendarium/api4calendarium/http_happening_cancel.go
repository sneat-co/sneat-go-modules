package api4calendarium

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/dto4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/facade4calendarium"
	"net/http"
)

var cancelHappening = facade4calendarium.CancelHappening

// httpCancelHappening marks happening as canceled
func httpCancelHappening(w http.ResponseWriter, r *http.Request) {
	var happeningRequest = getHappeningRequestParamsFromURL(r)
	request := dto4calendarium.CancelHappeningRequest{
		HappeningRequest: happeningRequest,
	}
	ctx, user, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = cancelHappening(ctx, user, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
