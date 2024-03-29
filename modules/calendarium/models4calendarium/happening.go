package models4calendarium

import (
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/const4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/teamus/dal4teamus"
)

// HappeningsCollection defines recurring happening's collection name
const HappeningsCollection = "happenings"

//const SingleHappeningsCollection = "single_happenings"

// NewHappeningKey creates a new happening key
func NewHappeningKey(teamID, happeningID string) *dal.Key {
	return dal4teamus.NewTeamModuleItemKey(teamID, const4calendarium.ModuleID, const4calendarium.HappeningsCollection, happeningID)
}

// HappeningType is either "recurring" or "single"
type HappeningType = string

const (
	// HappeningTypeRecurring = "recurring"
	HappeningTypeRecurring HappeningType = "recurring"

	// HappeningTypeSingle = "single"
	HappeningTypeSingle HappeningType = "single"
)

const (
	HappeningStatusActive   = "active"
	HappeningStatusArchived = "archived"
	HappeningStatusCanceled = "canceled"
	HappeningStatusDeleted  = "deleted"
)

// IsKnownHappeningStatus detects if a string is a know happening status
func IsKnownHappeningStatus(status string) bool {
	switch status {
	case HappeningStatusActive, HappeningStatusArchived, HappeningStatusCanceled, HappeningStatusDeleted:
		return true
	}
	return false
}

func NewHappeningContext(teamID, happeningID string) (v HappeningContext) {
	return NewHappeningContextWithDto(teamID, happeningID, new(HappeningDto))
}

func NewHappeningContextWithDto(teamID, happeningID string, dto *HappeningDto) (v HappeningContext) {
	v.ID = happeningID
	v.Key = NewHappeningKey(teamID, happeningID)
	v.Dto = dto
	v.Record = dal.NewRecordWithData(v.Key, dto)
	return
}
