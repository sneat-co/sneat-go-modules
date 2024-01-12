package dal4contactus

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/contactus/const4contactus"
	"github.com/sneat-co/sneat-go-modules/contactus/models4contactus"
	"github.com/sneat-co/sneat-go-modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/teamus/dto4teamus"
)

type ContactusTeamWorkerParams = dal4teamus.ModuleTeamWorkerParams[*models4contactus.ContactusTeamDto]

func NewContactusTeamWorkerParams(userID, teamID string) *ContactusTeamWorkerParams {
	teamWorkerParams := dal4teamus.NewTeamWorkerParams(userID, teamID)
	return dal4teamus.NewTeamModuleWorkerParams(const4contactus.ModuleID, teamWorkerParams, new(models4contactus.ContactusTeamDto))
}

func RunReadonlyContactusTeamWorker(
	ctx context.Context,
	user facade.User,
	request dto4teamus.TeamRequest,
	worker func(ctx context.Context, tx dal.ReadTransaction, params *ContactusTeamWorkerParams) (err error),
) error {
	return dal4teamus.RunReadonlyModuleTeamWorker(ctx, user, request, const4contactus.ModuleID, new(models4contactus.ContactusTeamDto), worker)
}

func RunContactusTeamWorker(
	ctx context.Context,
	user facade.User,
	request dto4teamus.TeamRequest,
	worker func(ctx context.Context, tx dal.ReadwriteTransaction, params *ContactusTeamWorkerParams) (err error),
) error {
	return dal4teamus.RunModuleTeamWorker(ctx, user, request, const4contactus.ModuleID, new(models4contactus.ContactusTeamDto), worker)
}

func RunContactusTeamWorkerTx(
	ctx context.Context,
	tx dal.ReadwriteTransaction,
	user facade.User,
	request dto4teamus.TeamRequest,
	worker func(ctx context.Context, tx dal.ReadwriteTransaction, params *ContactusTeamWorkerParams) (err error),
) error {
	return dal4teamus.RunModuleTeamWorkerTx(ctx, tx, user, request, const4contactus.ModuleID, new(models4contactus.ContactusTeamDto), worker)
}
