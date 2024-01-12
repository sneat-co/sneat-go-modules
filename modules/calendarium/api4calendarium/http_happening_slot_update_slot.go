package api4calendarium

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/dto4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/facade4calendarium"
	"net/http"
)

var updateSlot = facade4calendarium.UpdateSlot

func httpUpdateSlot(w http.ResponseWriter, r *http.Request) {
	var request dto4calendarium.HappeningSlotRequest
	request.HappeningRequest = getHappeningRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = updateSlot(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
