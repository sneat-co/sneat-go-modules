package api4schedulus

import (
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-modules/schedulus/dto4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/facade4schedulus"
	"net/http"
)

var deleteSlot = facade4schedulus.DeleteSlots

// httpDeleteHappening deletes happening
func httpDeleteSlots(w http.ResponseWriter, r *http.Request) {
	var request = dto4schedulus.DeleteHappeningSlotRequest{
		HappeningSlotRefRequest: dto4schedulus.HappeningSlotRefRequest{
			HappeningRequest: getHappeningRequestParamsFromURL(r),
		},
	}
	ctx, userContext, err := apicore.VerifyAuthenticatedRequestAndDecodeBody(w, r, endpoints.VerifyRequest{
		MinContentLength: apicore.MinJSONRequestSize,
		MaxContentLength: apicore.DefaultMaxJSONRequestSize,
		AuthRequired:     true,
	}, &request)
	if err != nil {
		return
	}
	err = deleteSlot(ctx, userContext, request)
	apicore.ReturnJSON(ctx, w, r, http.StatusCreated, err, nil)
}
