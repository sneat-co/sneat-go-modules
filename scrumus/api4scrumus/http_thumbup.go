package api4scrumus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/scrumus/facade4scrumus"
	"net/http"
)

var thumbUp = facade4scrumus.ThumbUp

// httpPostThumbUp add a thumb up
func httpPostThumbUp(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.KB,
		AuthRequired:     true,
	})
	if err != nil {
		return
	}
	var request facade4scrumus.ThumbUpRequest
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	err = thumbUp(ctx, userContext, request)
	apicore.IfNoErrorReturnOK(ctx, w, r, err)
}
