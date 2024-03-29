package api4listus

import (
	"github.com/sneat-co/sneat-go-modules/modules/listus/facade4listus"
	"net/http"
)

func getListRequestParamsFromURL(r *http.Request) (request facade4listus.ListRequest) {
	query := r.URL.Query()
	request.TeamID = query.Get("teamID")
	request.ListID = query.Get("listID")
	request.ListType = query.Get("listType")
	return
}
