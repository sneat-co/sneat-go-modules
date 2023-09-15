package api4assets

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"net/http"
)

// RegisterAssetRoutes registers asset routes
func RegisterAssetRoutes(handle modules.HTTPHandleFunc) {
	handle(http.MethodPost, "/v0/assets/create_asset", httpPostCreateAsset)
	handle(http.MethodDelete, "/v0/assets/delete_asset", httpDeleteAsset)
}
