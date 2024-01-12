package sneatgomodules

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/modules/assetus"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/contactus"
	"github.com/sneat-co/sneat-go-modules/modules/generic"
	"github.com/sneat-co/sneat-go-modules/modules/invitus"
	"github.com/sneat-co/sneat-go-modules/modules/listus"
	"github.com/sneat-co/sneat-go-modules/modules/retrospectus"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus"
	"github.com/sneat-co/sneat-go-modules/modules/sportus"
	"github.com/sneat-co/sneat-go-modules/modules/teamus"
	"github.com/sneat-co/sneat-go-modules/modules/userus"
)

func Modules() []modules.Module {
	return []modules.Module{
		calendarium.Module(),
		contactus.Module(),
		invitus.Module(),
		teamus.Module(),
		userus.Module(),
		assetus.Module(),
		listus.Module(),
		scrumus.Module(),
		retrospectus.Module(),
		sportus.Module(),
		generic.Module(),
	}
}
