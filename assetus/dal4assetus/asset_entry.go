package dal4assetus

import (
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	core "github.com/sneat-co/sneat-go-core"
	"github.com/sneat-co/sneat-go-modules/assetus/const4assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/models4assetus"
	"github.com/sneat-co/sneat-go-modules/teamus/dal4teamus"
)

func NewAssetEntryWithData[D any](teamID, contactID string, data D) (asset record.DataWithID[string, D]) {
	key := NewAssetKey(teamID, contactID)
	asset.ID = contactID
	asset.FullID = teamID + ":" + contactID
	asset.Key = key
	asset.Data = data
	asset.Record = dal.NewRecordWithData(key, data)
	return
}

func NewAssetKey(teamID, assetID string) *dal.Key {
	if !core.IsAlphanumericOrUnderscore(assetID) {
		panic(fmt.Errorf("assetID should be alphanumeric, got: [%v]", assetID))
	}
	teamModuleKey := dal4teamus.NewTeamModuleKey(teamID, const4assetus.ModuleID)
	return dal.NewKeyWithParentAndID(teamModuleKey, models4assetus.TeamAssetsCollection, assetID)
}
