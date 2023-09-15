package dal4retrospectus

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/retrospectus/models4retrospectus"
)

const RetrospectusModuleID = "retrospectus"

type RetroTeam = record.DataWithID[string, *models4retrospectus.RetroTeamDto]

func NewRetroTeamKey(id string) *dal.Key {
	teamKey := dal4teamus.NewTeamKey(id)
	return dal.NewKeyWithParentAndID(teamKey, dal4teamus.Collection, RetrospectusModuleID)
}

func NewRetroTeam(id string) RetroTeam {
	key := NewRetroTeamKey(id)
	return record.NewDataWithID(id, key, new(models4retrospectus.RetroTeamDto))
}

func GetRetroTeam(ctx context.Context, tx dal.ReadTransaction, id string) (RetroTeam, error) {
	retro := NewRetroTeam(id)
	return retro, tx.Get(ctx, retro.Record)
}
