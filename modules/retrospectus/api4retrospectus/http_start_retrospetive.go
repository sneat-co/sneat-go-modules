package api4retrospectus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-modules/modules/retrospectus/facade4retrospectus"
	"net/http"
)

var startRetrospective = facade4retrospectus.StartRetrospective

// httpPostStartRetrospective an API endpoint that starts retrospective
func httpPostStartRetrospective(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := verifyRequest(w, r, verify.DefaultJsonWithAuthRequired)
	if err != nil {
		return
	}
	request := facade4retrospectus.StartRetrospectiveRequest{}
	if err := apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	response, isNew, err := startRetrospective(ctx, userContext, request)
	var statusCode int
	if isNew {
		statusCode = http.StatusCreated
	} else {
		statusCode = http.StatusOK
	}
	apicore.ReturnJSON(ctx, w, r, statusCode, err, response)
}
