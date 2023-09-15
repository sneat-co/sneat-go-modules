package api4retrospectus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/retrospectus/facade4retrospectus"
	"net/http"
)

var addRetroItem = facade4retrospectus.AddRetroItem

// httpPostAddRetroItem adds an item to a retrospective
func httpPostAddRetroItem(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := verifyAuthorizedJSONRequest(w, r, apicore.MinJSONRequestSize, 10*apicore.KB)
	if err != nil {
		return
	}
	request := facade4retrospectus.AddRetroItemRequest{}
	if err := apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	response, err := addRetroItem(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, response)
}
