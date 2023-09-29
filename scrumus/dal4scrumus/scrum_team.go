package dal4scrumus

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-core-modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/scrumus/models4scrumus"
)

const ScrumusModuleID = "scrumus"

type ScrumTeam = record.DataWithID[string, *models4scrumus.ScrumTeamDto]

func NewScrumTeamKey(id string) *dal.Key {
	teamKey := dal4teamus.NewTeamKey(id)
	return dal.NewKeyWithParentAndID(teamKey, dal4teamus.TeamModulesCollection, ScrumusModuleID)
}

func NewScrumTeam(id string) ScrumTeam {
	key := NewScrumTeamKey(id)
	return record.NewDataWithID(id, key, new(models4scrumus.ScrumTeamDto))
}

func GetScrumTeam(ctx context.Context, tx dal.ReadTransaction, id string) (ScrumTeam, error) {
	retro := NewScrumTeam(id)
	return retro, tx.Get(ctx, retro.Record)
}
