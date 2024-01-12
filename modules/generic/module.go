package generic

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/modules/generic/api4generic"
	"github.com/sneat-co/sneat-go-modules/modules/generic/const4generic"
)

func Module() modules.Module {
	return modules.NewModule(const4generic.ModuleID, api4generic.RegisterHttpRoutes)
}
