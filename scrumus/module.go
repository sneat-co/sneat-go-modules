package scrumus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/scrumus/api4scrumus"
	"github.com/sneat-co/sneat-go-modules/scrumus/const4srumus"
)

func Module() modules.Module {
	return modules.NewModule(const4srumus.ModuleID, api4scrumus.RegisterHttpRoutes)
}
