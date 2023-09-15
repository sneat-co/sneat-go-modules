package api4listus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/listus/facade4listus"
	"net/http"
)

var createListItems = facade4listus.CreateListItems

// httpPostCreateListItems creates list items
func httpPostCreateListItems(w http.ResponseWriter, r *http.Request) {
	var request facade4listus.CreateListItemsRequest
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.DefaultMaxJSONRequestSize,
		AuthRequired:     true,
	}, &request)
	if err != nil {
		return
	}
	response, err := createListItems(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, &response)
}