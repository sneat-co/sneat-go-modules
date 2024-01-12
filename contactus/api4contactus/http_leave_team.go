package api4contactus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/contactus/facade4contactus"
	"github.com/sneat-co/sneat-go-modules/teamus/dto4teamus"
	"net/http"
)

var leaveTeam = facade4contactus.LeaveTeam

// httpPostLeaveTeam is an API endpoint that removes user from a team by his/here request
func httpPostLeaveTeam(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, verify.DefaultJsonWithAuthRequired)
	if err != nil {
		return
	}
	var request dto4teamus.LeaveTeamRequest
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	err = leaveTeam(ctx, userContext, request)
	apicore.IfNoErrorReturnOK(ctx, w, r, err)
}
