package api4scrumus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/scrumus/facade4scrumus"
	"net/http"
)

var setMetric = facade4scrumus.SetMetric

// httpPostSetMetric is an API endpoint that sets metric value
func httpPostSetMetric(w http.ResponseWriter, r *http.Request) {
	ctx, userContext, err := apicore.VerifyRequestAndCreateUserContext(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: 10 * apicore.KB,
		AuthRequired:     true,
	})
	if err != nil {
		return
	}
	var request facade4scrumus.SetMetricRequest
	if err = apicore.DecodeRequestBody(w, r, &request); err != nil {
		return
	}
	response, err := setMetric(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, response)
}
