package models4contactus

import (
	"github.com/sneat-co/sneat-go-modules/modules/contactus/const4contactus"
	"github.com/sneat-co/sneat-go-modules/modules/linkage/models4linkage"
)

func NewContactRef(teamID, contactID string) models4linkage.TeamModuleDocRef {
	return models4linkage.NewTeamModuleDocRef(teamID, const4contactus.ModuleID, const4contactus.ContactsCollection, contactID)
}
