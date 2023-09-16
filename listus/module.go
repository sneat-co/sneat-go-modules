package listus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/listus/api4listus"
	"github.com/sneat-co/sneat-go-modules/listus/const4listus"
)

func Module() modules.Module {
	return modules.NewModule(const4listus.ModuleID, api4listus.RegisterHttpRoutes)
}
