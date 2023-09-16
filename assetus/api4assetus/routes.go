package api4assetus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"net/http"
)

// RegisterHttpRoutes registers asset routes
func RegisterHttpRoutes(handle modules.HTTPHandleFunc) {
	handle(http.MethodPost, "/v0/assets/create_asset", httpPostCreateAsset)
	handle(http.MethodDelete, "/v0/assets/delete_asset", httpDeleteAsset)
}
