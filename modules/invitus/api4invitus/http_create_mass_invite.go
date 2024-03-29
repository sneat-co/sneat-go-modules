package api4invitus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/invitus/facade4invitus"
	"net/http"
)

var createMassInvite = facade4invitus.CreateMassInvite

// httpPostCreateMassInvite is an API endpoint to create a mass-invite
func httpPostCreateMassInvite(w http.ResponseWriter, r *http.Request) {
	ctx, err := apicore.VerifyRequest(w, r, verify.DefaultJsonWithAuthRequired)
	if err != nil {
		return
	}
	var request facade4invitus.CreateMassInviteRequest
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	response, err := createMassInvite(ctx, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, response)
}
