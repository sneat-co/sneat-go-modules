package api4schedulus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"net/http"
)

// RegisterSchedulusRoutes register schedule routes
func RegisterSchedulusRoutes(handle modules.HTTPHandleFunc) {
	handle(http.MethodPost, "/v0/happenings/create_happening", httpPostCreateHappening)
	handle(http.MethodDelete, "/v0/happenings/delete_happening", httpDeleteHappening)
	handle(http.MethodDelete, "/v0/happenings/delete_slots", httpDeleteSlots)
	handle(http.MethodPost, "/v0/happenings/cancel_happening", httpCancelHappening)
	handle(http.MethodPost, "/v0/happenings/revoke_happening_cancellation", httpRevokeHappeningCancellation)
	handle(http.MethodPost, "/v0/happenings/add_member", httpAddMemberToHappening)
	handle(http.MethodPost, "/v0/happenings/remove_member", httpRemoveMemberFromHappening)
	handle(http.MethodPost, "/v0/happenings/update_slot", httpUpdateSlot)
	handle(http.MethodPost, "/v0/happenings/adjust_slot", httpAdjustSlot)
}
