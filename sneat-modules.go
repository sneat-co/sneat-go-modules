package sneatgomodules

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/assetus"
	"github.com/sneat-co/sneat-go-modules/calendarium"
	"github.com/sneat-co/sneat-go-modules/contactus"
	"github.com/sneat-co/sneat-go-modules/generic"
	"github.com/sneat-co/sneat-go-modules/invitus"
	"github.com/sneat-co/sneat-go-modules/listus"
	"github.com/sneat-co/sneat-go-modules/retrospectus"
	"github.com/sneat-co/sneat-go-modules/scrumus"
	"github.com/sneat-co/sneat-go-modules/sportus"
	"github.com/sneat-co/sneat-go-modules/teamus"
	"github.com/sneat-co/sneat-go-modules/userus"
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
