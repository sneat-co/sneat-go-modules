package generic

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/generic/api"
)

// Register HTTP handle
func Register(handle modules.HTTPHandleFunc) {
	api.RegisterHandlers(handle)
}
