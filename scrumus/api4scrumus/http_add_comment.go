package api4scrumus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/scrumus/facade4scrumus"
	"net/http"
)

var addComment = facade4scrumus.AddComment

// httpPostAddComment adds a comment
func httpPostAddComment(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.KB,
		AuthRequired:     true,
	})
	if err != nil {
		return
	}
	var request facade4scrumus.AddCommentRequest
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	comment, err := addComment(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, comment)
}