package models4contactus

import (
	"github.com/sneat-co/sneat-go-modules/contactus/briefs4contactus"
)

type ContactusTeamDto struct {
	briefs4contactus.WithSingleTeamContactsWithoutContactIDs[*briefs4contactus.ContactBrief]
}

func (v *ContactusTeamDto) Validate() error {
	return nil
}
