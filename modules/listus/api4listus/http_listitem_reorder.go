package api4listus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/listus/facade4listus"
	"net/http"
)

var reorderListItem = facade4listus.ReorderListItem

// httpPostReorderListItem reorders list items
func httpPostReorderListItem(w http.ResponseWriter, r *http.Request) {
	var request facade4listus.ReorderListItemsRequest
	request.ListRequest = getListRequestParamsFromURL(r)
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, verify.DefaultJsonWithAuthRequired, &request)
	if err != nil {
		return
	}
	if err = request.Validate(); err != nil {
		apicore.ReturnError(r.Context(), w, r, err)
		return
	}
	err = reorderListItem(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, nil)
}
