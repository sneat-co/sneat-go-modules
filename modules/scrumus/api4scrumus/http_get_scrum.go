package api4scrumus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/facade4scrumus"
	"net/http"
)

var getScrum = facade4scrumus.GetScrum

// httpGetScrum is an API endpoint that returns scrum data
func httpGetScrum(w http.ResponseWriter, r *http.Request) {
	ctx, user, err := apicore.VerifyRequestAndCreateUserContext(w, r, verify.Request(verify.AuthenticationRequired(true)))
	if err != nil {
		return
	}
	response, err := getScrum(ctx, user, facade.IDRequest{ID: r.Header.Get("id")})
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, response)
}
