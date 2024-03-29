package api4listus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/listus/facade4listus"
	"net/http"
)

var deleteList = facade4listus.DeleteList

// httpDeleteList deletes a list
func httpDeleteList(w http.ResponseWriter, r *http.Request) {
	var request = getListRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	err = deleteList(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, nil)
}
