package api4scrumus

import (
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/sneatfb"
	"github.com/sneat-co/sneat-go-modules/scrumus/facade4scrumus"
	"net/http"
)

var getScrum = facade4scrumus.GetScrum

// httpGetScrum is an API endpoint that returns scrum data
func httpGetScrum(w http.ResponseWriter, r *http.Request) {
	authContext, err := sneatfb.NewAuthContext(r)
	if err != nil {
		return
	}
	ctx := r.Context()
	response, err := getScrum(ctx, authContext, facade.IDRequest{ID: r.Header.Get("id")})
	apicore.ReturnJSON(ctx, w, r, http.StatusOK, err, response)
}
