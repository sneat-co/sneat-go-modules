package sport

import (
	"github.com/sneat-co/sneat-go-core/modules"
	entities2 "github.com/sneat-co/sneat-go-modules/generic/entities"
	"github.com/sneat-co/sneat-go-modules/sport/api2sport"
)

// Register registers HTTP handlers
func Register(handle modules.HTTPHandleFunc) {
	api2sport.RegisterRoutes(handle)
	entities2.Register(
		entities2.Entity{Name: "Spot", AllowCreate: true, AllowUpdate: true},
	)
}
