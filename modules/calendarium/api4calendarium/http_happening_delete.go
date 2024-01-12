package api4calendarium

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/facade4calendarium"
	"net/http"
)

// httpDeleteHappening deletes happening
func httpDeleteHappening(w http.ResponseWriter, r *http.Request) {
	var request = getHappeningRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = facade4calendarium.DeleteHappening(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, nil)
}
