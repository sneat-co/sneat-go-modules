package models4schedulus

import (
	"fmt"
	dbmodels2 "github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/sneat-co/sneat-core-modules/contactus/briefs4contactus"
	"github.com/strongo/validation"
	"strings"
)

// HappeningDto DTO
type HappeningDto struct {
	HappeningBase
	dbmodels2.WithTags
	dbmodels2.WithUserIDs
	dbmodels2.WithTeamDates
	briefs4contactus.WithMultiTeamContacts[*briefs4contactus.ContactBrief]
}

// Validate returns error if not valid
func (v *HappeningDto) Validate() error {
	if err := v.HappeningBase.Validate(); err != nil {
		return err
	}
	if err := v.WithUserIDs.Validate(); err != nil {
		return err
	}
	if err := v.WithTeamDates.Validate(); err != nil {
		return err
	}
	if err := v.WithTags.Validate(); err != nil {
		return err
	}
	if len(v.TeamIDs) == 0 {
		return validation.NewErrRecordIsMissingRequiredField("teamIDs")
	}
	if err := v.WithMultiTeamContacts.Validate(); err != nil {
		return err
	}
	for i, level := range v.Levels {
		if l := strings.TrimSpace(level); l == "" {
			return validation.NewErrRecordIsMissingRequiredField(
				fmt.Sprintf("levels[%v]", i),
			)
		} else if l != level {
			return validation.NewErrBadRecordFieldValue(
				fmt.Sprintf("levels[%v]", i),
				fmt.Sprintf("whitespaces at beginning or end: [%v]", level),
			)
		}
	}
	if err := v.WithMultiTeamContactIDs.Validate(); err != nil {
		return err
	}
	for i, teamContactID := range v.ContactIDs {
		ids := strings.Split(string(teamContactID), ":")
		if len(ids) != 2 {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("contactIDs[%v]", i), fmt.Sprintf("expected to have an id in '{TEAM_ID}:{CONTACT_ID}' format, got: '%v'", teamContactID))
		}
		teamID, contactID := ids[0], ids[1]
		if teamID == "" {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("contactIDs[%v]", i), fmt.Sprintf("missing TEAM_ID part in '{TEAM_ID}:{CONTACT_ID}' format, got: '%v'", teamContactID))
		}
		if contactID == "" {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("contactIDs[%v]", i), fmt.Sprintf("missing CONTACT_ID part in  '{TEAM_ID}:{CONTACT_ID}' format, got: '%v'", teamContactID))
		}
	}
	switch v.Type {
	case "":
		return validation.NewErrRecordIsMissingRequiredField("type")
	case HappeningTypeSingle:
		if count := len(v.Slots); count > 1 {
			return validation.NewErrBadRecordFieldValue("slots", fmt.Sprintf("single time happening should have only single 'once' slot, got: %v", count))
		}
		if len(v.Dates) == 0 {
			return validation.NewErrRecordIsMissingRequiredField("dates")
		}
		if len(v.TeamDates) == 0 {
			return validation.NewErrRecordIsMissingRequiredField("teamDates")
		}
	case HappeningTypeRecurring:
		if len(v.Dates) > 0 {
			return validation.NewErrBadRequestFieldValue("dates", "should be empty for 'recurring' happening")
		}
	default:
		return validation.NewErrBadRecordFieldValue("type", "unknown value: "+v.Type)
	}
	//if v.Role == HappeningTypeRecurring && v.Status == HappeningStatusCanceled {
	//	for _, slot := range v.Slots {
	//		if slot.Status != SlotStatusCanceled {
	//
	//		}
	//	}
	//}
	return nil
}
