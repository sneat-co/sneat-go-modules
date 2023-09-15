package models4schedulus

import (
	"fmt"
	"github.com/dal-go/dalgo/dal"
)

// HappeningsCollection defines recurring happening's collection name
const HappeningsCollection = "happenings"

//const SingleHappeningsCollection = "single_happenings"

// NewHappeningKey creates a new happening key
func NewHappeningKey(id string) *dal.Key {
	return dal.NewKeyWithID(HappeningsCollection, id)
}

// HappeningType is either "recurring" or "single"
type HappeningType = string

const (
	// HappeningTypeRecurring = "recurring"
	HappeningTypeRecurring HappeningType = "recurring"

	// HappeningTypeSingle = "single"
	HappeningTypeSingle HappeningType = "single"
)

// CreateHappeningDto a DTO object
type CreateHappeningDto struct {
	HappeningBase
}

// Validate returns error if not valid
func (v CreateHappeningDto) Validate() error {
	if err := v.HappeningBase.Validate(); err != nil {
		return fmt.Errorf("failed validation of HappeningBase: %w", err)
	}
	return nil
}

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

func NewHappeningContext(id string) (v HappeningContext) {
	return NewHappeningContextWithDto(id, new(HappeningDto))
}

func NewHappeningContextWithDto(id string, dto *HappeningDto) (v HappeningContext) {
	v.ID = id
	v.Key = NewHappeningKey(id)
	v.Dto = dto
	v.Record = dal.NewRecordWithData(v.Key, dto)
	return
}
