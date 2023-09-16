package models4assetus

import "github.com/sneat-co/sneat-go-modules/assetus/briefs4assetus"

type AssetusTeamDto struct {
	briefs4assetus.WithMultiTeamAssets[*briefs4assetus.AssetBrief]
	Assets []*briefs4assetus.AssetBrief `json:"assets,omitempty" firestore:"assets,omitempty"`
}

func (v AssetusTeamDto) Validate() error {
	return nil
}
