package api4calendarium

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/dto4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/facade4calendarium"
	"net/http"
)

func httpAddParticipantToHappening(w http.ResponseWriter, r *http.Request) {
	var request dto4calendarium.HappeningContactRequest
	request.HappeningRequest = getHappeningRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = facade4calendarium.AddParticipantToHappening(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusNoContent, err, nil)
}
