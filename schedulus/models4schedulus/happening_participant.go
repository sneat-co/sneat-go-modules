package models4schedulus

import (
	"fmt"
	"github.com/strongo/validation"
	"strings"
)

type HappeningParticipant struct {
	Roles []string `json:"roles,omitempty"`
}

func (v HappeningParticipant) Validate() error {
	for i, role := range v.Roles {
		if strings.TrimSpace(role) == "" {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("roles[%d]", i), "role is empty")
		}
	}
	return nil
}
