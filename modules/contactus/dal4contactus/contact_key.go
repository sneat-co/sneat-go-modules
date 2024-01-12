package dal4contactus

import (
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core"
	"github.com/sneat-co/sneat-go-modules/modules/contactus/const4contactus"
	"github.com/sneat-co/sneat-go-modules/modules/contactus/models4contactus"
	"github.com/sneat-co/sneat-go-modules/modules/teamus/dal4teamus"
)

// NewContactKey creates a new contact's key in format "teamID:memberID"
func NewContactKey(teamID, contactID string) *dal.Key {
	if !core.IsAlphanumericOrUnderscore(contactID) {
		panic(fmt.Errorf("contactID should be alphanumeric, got: [%v]", contactID))
	}
	teamModuleKey := dal4teamus.NewTeamModuleKey(teamID, const4contactus.ModuleID)
	return dal.NewKeyWithParentAndID(teamModuleKey, models4contactus.TeamContactsCollection, contactID)
}
