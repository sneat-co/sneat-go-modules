package models4schedulus

import (
	"github.com/sneat-co/sneat-core-modules/contactus/briefs4contactus"
	"github.com/strongo/validation"
	"strings"
)

// HappeningBrief DTO
type HappeningBrief struct {
	ID string `json:"id" firestore:"id"`
	HappeningBase
	briefs4contactus.WithMultiTeamContactIDs
}

// Validate returns error if not valid
func (v HappeningBrief) Validate() error {
	if id := strings.TrimSpace(v.ID); id == "" {
		return validation.NewErrRecordIsMissingRequiredField("id")
	}
	if err := v.HappeningBase.Validate(); err != nil {
		return err
	}
	return nil
}
