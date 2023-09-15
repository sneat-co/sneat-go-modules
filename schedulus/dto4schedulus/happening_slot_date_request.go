package dto4schedulus

import (
	"github.com/sneat-co/sneat-go-core/validate"
	"github.com/sneat-co/sneat-go-modules/schedulus/models4schedulus"
	"github.com/strongo/validation"
	"strings"
)

// HappeningSlotDateRequest updates slot
type HappeningSlotDateRequest struct {
	HappeningRequest
	Slot models4schedulus.HappeningSlot `json:"slot"`
	Date string                         `json:"date"`
}

// Validate returns error if not valid
func (v HappeningSlotDateRequest) Validate() error {
	if err := v.HappeningRequest.Validate(); err != nil {
		return err
	}
	if err := v.Slot.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(v.Date) == "" {
		return validation.NewErrRecordIsMissingRequiredField("date")
	}
	if _, err := validate.DateString(v.Date); err != nil {
		return validation.NewErrBadRequestFieldValue("date", err.Error())
	}
	return nil
}
