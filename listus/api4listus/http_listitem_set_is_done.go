package api4listus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/listus/facade4listus"
	"net/http"
)

var setListItemsIsDone = facade4listus.SetListItemsIsDone

// httpPostSetListItemsIsDone marks list items as completed
func httpPostSetListItemsIsDone(w http.ResponseWriter, r *http.Request) {
	var request facade4listus.ListItemsSetIsDoneRequest
	request.ListRequest = getListRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	if err = request.Validate(); err != nil {
		apicore.ReturnError(r.Context(), w, r, err)
		return
	}
	err = setListItemsIsDone(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
