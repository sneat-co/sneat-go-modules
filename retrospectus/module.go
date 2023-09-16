package retrospectus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/retrospectus/api4retrospectus"
	"github.com/sneat-co/sneat-go-modules/retrospectus/const4retrospectus"
)

func Module() modules.Module {
	return modules.NewModule(const4retrospectus.ModuleID, api4retrospectus.RegisterHttpRoutes)
}
