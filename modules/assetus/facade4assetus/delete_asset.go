package facade4assetus

import (
	"context"
	"errors"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/assetus/const4assetus"
	"github.com/sneat-co/sneat-go-modules/modules/assetus/dal4assetus"
	"github.com/sneat-co/sneat-go-modules/modules/assetus/models4assetus"
	"github.com/sneat-co/sneat-go-modules/modules/teamus/dal4teamus"
)

// DeleteAsset deletes an asset
func DeleteAsset(ctx context.Context, user facade.User, request dal4teamus.TeamItemRequest) (err error) {
	if err = request.Validate(); err != nil {
		return
	}
	if user == nil || user.GetID() == "" {
		return errors.New("no user context")
	}
	input := dal4teamus.TeamItemRunnerInput[*models4assetus.AssetusTeamDto]{
		IsTeamRecordNeeded: true,
		Counter:            "assets",
		ShortID:            request.ID,
		TeamItem:           dal.NewRecord(dal.NewKeyWithID(dal4assetus.AssetsCollection, request.ID)),
	}
	err = dal4teamus.DeleteTeamItem(ctx, user, input, const4assetus.ModuleID, new(models4assetus.AssetusTeamDto), func(ctx context.Context, tx dal.ReadwriteTransaction, params *dal4teamus.TeamItemWorkerParams) (err error) {
		return errors.New("not implemented")
	})
	return
}
