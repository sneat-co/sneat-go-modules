package facade4scrumus

import (
	"context"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/models4scrumus"
)

// GetScrum returns scrum data
func GetScrum(_ context.Context, user facade.User, _ facade.IDRequest) (scrum models4scrumus.Scrum, err error) {
	return
}
