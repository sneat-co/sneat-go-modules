package api4schedulus

import (
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"net/http"
)

func getHappeningRequestParamsFromURL(r *http.Request) (request dto4schedulus.HappeningRequest) {
	query := r.URL.Query()
	request.TeamID = query.Get("teamID")
	request.HappeningID = query.Get("happeningID")
	request.ListType = query.Get("happeningType")
	return
}
