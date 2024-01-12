package assetus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/modules/assetus/api4assetus"
	"github.com/sneat-co/sneat-go-modules/modules/assetus/const4assetus"
)

func Module() modules.Module {
	return modules.NewModule(const4assetus.ModuleID, api4assetus.RegisterHttpRoutes)
}
