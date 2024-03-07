package api4calendarium

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/dto4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/facade4calendarium"
	"net/http"
)

var adjustSlot = facade4calendarium.AdjustSlot

func httpAdjustSlot(w http.ResponseWriter, r *http.Request) {
	var request dto4calendarium.HappeningSlotDateRequest
	request.HappeningRequest = getHappeningRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = adjustSlot(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}