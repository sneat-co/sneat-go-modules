package models4schedulus

import (
	"github.com/sneat-co/sneat-core-modules/contactus/briefs4contactus"
)

// HappeningBrief DTO
type HappeningBrief struct {
	HappeningBase

	// TODO: document why we need each additional field of HappeningBrief, e.g. why we can't get rid of HappeningBase
	briefs4contactus.WithMultiTeamContactIDs
}

// Validate returns error if not valid
func (v HappeningBrief) Validate() error {
	if err := v.HappeningBase.Validate(); err != nil {
		return err
	}
	if err := v.WithMultiTeamContactIDs.Validate(); err != nil {
		return err
	}
	return nil
}
