package api4retrospectus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/retrospectus/facade4retrospectus"
	"net/http"
)

var moveRetroItem = facade4retrospectus.MoveRetroItem

// httpPostMoveRetroItem is an API endpoint that changes position of retrospective item
func httpPostMoveRetroItem(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: 10 * apicore.KB,
		AuthRequired:     true,
	})
	if err != nil {
		return
	}
	request := facade4retrospectus.MoveRetroItemRequest{}
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	err = moveRetroItem(ctx, userContext, request)
	apicore.IfNoErrorReturnOK(ctx, w, r, err)
}
