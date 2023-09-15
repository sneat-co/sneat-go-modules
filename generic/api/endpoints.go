package api

import (
	"github.com/sneat-co/sneat-go-core/modules"
)

// RegisterHandlers registers HTTP handlers
func RegisterHandlers(handle modules.HTTPHandleFunc) {
	handle("POST", "/api4invitus/$generic/create", create)
	handle("PUT", "/api4invitus/$generic/update", update)
	handle("DELETE", "/api4invitus/$generic/delete", delete)
}
