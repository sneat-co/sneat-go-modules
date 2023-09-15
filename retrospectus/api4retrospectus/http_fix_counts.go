package api4retrospectus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/retrospectus/facade4retrospectus"
	"net/http"
)

var fixCounts = facade4retrospectus.FixCounts

// httpPostFixCounts is an API endpoint that triggers fixing of counters in a retrospective
func httpPostFixCounts(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := verifyRequest(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: 1 * apicore.KB,
		AuthRequired:     true,
	})
	if err != nil {
		return
	}
	request := facade4retrospectus.FixCountsRequest{}
	if err := apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	err = fixCounts(ctx, userContext, request)
	apicore.IfNoErrorReturnOK(ctx, w, r, err)
}
