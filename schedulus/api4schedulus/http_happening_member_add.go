package api4schedulus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/facade4schedulus"
	"net/http"
)

var addMemberToHappening = facade4schedulus.AddContactToHappening

func httpAddMemberToHappening(w http.ResponseWriter, r *http.Request) {
	var request dto4schedulus.HappeningContactRequest
	request.HappeningRequest = getHappeningRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = addMemberToHappening(ctx, userContext.GetID(), request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
