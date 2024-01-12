package calendarium

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/calendarium/api4calendarium"
	"github.com/sneat-co/sneat-go-modules/calendarium/const4calendarium"
)

func Module() modules.Module {
	return modules.NewModule(const4calendarium.ModuleID, api4calendarium.RegisterHttpRoutes)
}
