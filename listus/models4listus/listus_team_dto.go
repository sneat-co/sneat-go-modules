package models4listus

import (
	"github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/strongo/validation"
)

type ListusTeamDto struct {
	dbmodels.WithCreated
	Lists map[string]*ListBrief `json:"lists,omitempty" firestore:"lists,omitempty"`
	//ListGroups []*ListGroup          `json:"listGroups,omitempty" firestore:"listGroups,omitempty"`
}

func (v ListusTeamDto) Validate() error {
	if err := validateListBriefs(v.Lists); err != nil {
		return validation.NewErrBadRecordFieldValue("lists", err.Error())
	}
	if err := v.WithCreated.Validate(); err != nil {
		return err
	}
	return nil
}
