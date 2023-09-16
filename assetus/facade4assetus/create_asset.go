package facade4assetus

import (
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	dbmodels2 "github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/sneat-co/sneat-core-modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-core-modules/teamus/dto4teamus"
	"github.com/sneat-co/sneat-go-modules/assetus/const4assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/dal4assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/models4assetus"
	"github.com/strongo/random"
	"time"
)

// AssetSummary DTO
type AssetSummary struct {
	RegNumber      string `json:"number,omitempty"`
	DateOfBuild    string `json:"dateOfBuild,omitempty"`
	DateOfPurchase string `json:"dateOfPurchase,omitempty"`
}

type CreateAssetData struct {
	models4assetus.AssetMainDto
	models4assetus.AssetSpecificData
}

// CreateAssetRequest is a DTO for creating an asset
type CreateAssetRequest struct {
	dto4teamus.TeamRequest
	Asset  models4assetus.AssetCreationData `json:"asset"`
	DbData models4assetus.AssetDbData       `json:"-"`
}

// Validate returns error if not valid
func (v CreateAssetRequest) Validate() error {
	if err := v.TeamRequest.Validate(); err != nil {
		return err
	}
	if err := v.Asset.Validate(); err != nil {
		return err
	}
	return nil
}

// CreateAssetResponse DTO
type CreateAssetResponse struct {
	ID   string                     `json:"id"`
	Data models4assetus.AssetDbData `json:"data"`
}

// CreateAsset creates an asset
func CreateAsset(ctx context.Context, user facade.User, request CreateAssetRequest) (response CreateAssetResponse, err error) {
	if err = request.Validate(); err != nil {
		return
	}
	err = dal4teamus.CreateTeamItem(ctx, user, "assets", request.TeamRequest, const4assetus.ModuleID,
		func(ctx context.Context, tx dal.ReadwriteTransaction, params *dal4teamus.ModuleTeamWorkerParams[*models4assetus.AssetusTeamDto]) (err error) {
			modified := dbmodels2.Modified{
				By: user.GetID(),
				At: time.Now(),
			}
			response.ID = random.ID(7) // TODO: consider using incomplete key with options?
			assetExtraData := request.DbData.AssetExtraData()
			assetExtraData.UserIDs = []string{user.GetID()}
			assetExtraData.WithModified = dbmodels2.NewWithModified(modified.At, modified.By)
			assetExtraData.WithTeamIDs = dbmodels2.WithSingleTeamID(request.TeamRequest.TeamID)

			//assetMainData := request.DbData.AssetMainDto()
			//assetMainData.ContactIDs = []string{"*"} // Required value, TODO: use constant
			//assetMainData.AssetIDs = []string{"*"}   // Required value, TODO: use constant
			assetKey := dal.NewKeyWithParentAndID(dal4assetus.AssetusRootKey, dal4assetus.AssetsCollection, response.ID)
			assetRecord := dal.NewRecordWithData(assetKey, request.DbData)
			if err = tx.Insert(ctx, assetRecord); err != nil {
				return fmt.Errorf("failed to insert response record")
			}
			return nil
		},
	)
	response.Data = request.DbData
	return
}
