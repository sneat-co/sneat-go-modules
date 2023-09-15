package models4assets

import (
	"github.com/sneat-co/sneat-go-modules/assetus/briefs4assets"
)

type AssetusTeamDto struct {
	briefs4assets.WithMultiTeamAssets[*briefs4assets.AssetBrief]
	Assets []*briefs4assets.AssetBrief `json:"assets,omitempty" firestore:"assets,omitempty"`
}

func (v AssetusTeamDto) Validate() error {
	return nil
}
