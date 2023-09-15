package briefs4assets

import (
	dbmodels2 "github.com/sneat-co/sneat-go-core/models/dbmodels"
)

// WithMultiTeamAssets // TODO: should be moved to assetus module?
type WithMultiTeamAssets[T interface {
	dbmodels2.RelatedAs
	Equal(v T) bool
}] struct {
	dbmodels2.WithMultiTeamAssetIDs
	Assets map[string]T `json:"assets,omitempty" firestore:"assets,omitempty"`
}

// Validate returns error if not valid
func (v *WithMultiTeamAssets[T]) Validate() error {
	if err := v.WithMultiTeamAssetIDs.Validate(); err != nil {
		return err
	}
	return dbmodels2.ValidateWithIdsAndBriefs("assetIDs", "relatedAssets", v.AssetIDs, v.Assets)
}

type WithMultiTeamAssetBriefs = WithMultiTeamAssets[*AssetBrief]
