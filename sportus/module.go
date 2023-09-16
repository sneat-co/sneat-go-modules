package sportus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/generic/entities"
	"github.com/sneat-co/sneat-go-modules/sportus/api4sportus"
	"github.com/sneat-co/sneat-go-modules/sportus/const4sportus"
)

func Module() modules.Module {
	entities.Register(
		entities.Entity{Name: "Spot", AllowCreate: true, AllowUpdate: true},
	)
	return modules.NewModule(const4sportus.ModuleID, api4sportus.RegisterRoutes)
}
