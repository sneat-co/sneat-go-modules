package api4retrospectus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/retrospectus/facade4retrospectus"
	"net/http"
)

var voteItem = facade4retrospectus.VoteItem

// httpPostVoteItem is an API endpoint that cast a vote for a retrospective item
func httpPostVoteItem(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := verifyRequest(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: 10 * apicore.KB,
		AuthRequired:     true,
	})
	if err != nil {
		return
	}
	request := facade4retrospectus.VoteItemRequest{}
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	err = voteItem(ctx, userContext, request)
	apicore.ReturnStatus(ctx, w, r, http.StatusNoContent, err)
}
