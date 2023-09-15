package facade4assets

import (
	"context"
	"errors"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/dal4assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/models4assets"
)

// DeleteAsset deletes an asset
func DeleteAsset(ctx context.Context, user facade.User, request dal4teamus.TeamItemRequest) (err error) {
	if err = request.Validate(); err != nil {
		return
	}
	if user == nil || user.GetID() == "" {
		return errors.New("no user context")
	}
	input := dal4teamus.TeamItemRunnerInput[*models4assets.AssetusTeamDto]{
		IsTeamRecordNeeded: true,
		Counter:            "assets",
		ShortID:            request.ID,
		TeamItem:           dal.NewRecord(dal.NewKeyWithID(dal4assetus.AssetsCollection, request.ID)),
	}
	err = dal4teamus.DeleteTeamItem(ctx, user, input, assetus.ModuleID, nil)
	return
}
