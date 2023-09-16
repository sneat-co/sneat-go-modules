package schedulus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/schedulus/api4schedulus"
	"github.com/sneat-co/sneat-go-modules/schedulus/const4schedulus"
)

func Module() modules.Module {
	return modules.NewModule(const4schedulus.ModuleID, api4schedulus.RegisterHttpRoutes)
}
