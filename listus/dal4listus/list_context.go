package dal4listus

import (
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/listus/models4listus"
)

type ListContext struct {
	record.WithID[string]
	Dto *models4listus.ListDto
}

// NewTeamListKey creates new list key
func NewTeamListKey(teamID, id string) *dal.Key {
	teamKey := dal4teamus.NewTeamKey(teamID)
	return dal.NewKeyWithParentAndID(teamKey, models4listus.TeamListsCollection, id)
}

func NewTeamListContext(teamID, listID string) (list ListContext) {
	key := NewTeamListKey(teamID, listID)
	list.ID = listID
	list.FullID = teamID + ":" + listID
	list.Key = key
	list.Dto = new(models4listus.ListDto)
	list.Record = dal.NewRecordWithData(key, list.Dto)
	return
}
