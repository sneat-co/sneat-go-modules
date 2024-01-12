package dal4calendarium

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/calendarium/const4calendarium"
	"github.com/sneat-co/sneat-go-modules/calendarium/models4calendarium"
	"github.com/sneat-co/sneat-go-modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/teamus/dto4teamus"
)

func RunCalendariumTeamWorker(
	ctx context.Context,
	user facade.User,
	request dto4teamus.TeamRequest,
	worker func(ctx context.Context, tx dal.ReadwriteTransaction, params *CalendariumTeamWorkerParams) (err error),
) error {
	return dal4teamus.RunModuleTeamWorker(ctx, user, request, const4calendarium.ModuleID, new(models4calendarium.CalendariumTeamDto), worker)
}
