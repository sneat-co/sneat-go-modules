package models4listus

import (
	"fmt"
	dbmodels2 "github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/strongo/validation"
	"strings"
)

// TeamListsCollection defines collection name
const TeamListsCollection = "lists"

// ListType list type
type ListType = string

const (
	// ListTypeGeneral = "general"
	ListTypeGeneral ListType = "general"

	// ListTypeToBuy = "to-by"
	ListTypeToBuy ListType = "to-buy"

	// ListTypeToDo = "to-do"
	ListTypeToDo ListType = "to-do"

	// ListTypeToWatch = "to-watch"
	ListTypeToWatch ListType = "to-watch"
)

// IsKnownListType checks if it is a known list type
func IsKnownListType(v string) bool {
	switch v {
	case ListTypeGeneral, ListTypeToBuy, ListTypeToWatch, ListTypeToDo:
		return true
	}
	return false
}

// GetFullListID returns full list ContactID
func GetFullListID(listType ListType, listID string) string {
	return fmt.Sprintf("%v:%v", listType, listID)
}

// ListBase DTO
type ListBase struct {
	Type  ListType `json:"type" firestore:"type"`
	Emoji string   `json:"emoji,omitempty" firestore:"emoji,omitempty"`
	// Title should be unique across owning team/company/group/etc
	Title string `json:"title" firestore:"title"`
}

// Validate returns error if not valid
func (v ListBase) Validate() error {
	if v.Type == "" {
		return validation.NewErrRecordIsMissingRequiredField("type")
	}
	if !IsKnownListType(v.Type) {
		return validation.NewErrBadRecordFieldValue("type", "unknown value: "+v.Type)
	}
	if strings.TrimSpace(v.Title) == "" {
		return validation.NewErrRecordIsMissingRequiredField("title")
	}
	return nil
}

// ListGroup DTO
type ListGroup struct {
	Type  string       `json:"type" firestore:"type"`
	Title string       `json:"title" firestore:"title"`
	Lists []*ListBrief `json:"lists,omitempty" firestore:"lists,omitempty"`
}

// Validate returns error if not valid
func (v ListGroup) Validate() error {
	if v.Type == "" {
		return validation.NewErrRecordIsMissingRequiredField("type")
	}
	if v.Title == "" {
		return validation.NewErrRecordIsMissingRequiredField("title")
	}
	//if l := len(v.Emoji); l > 4 {
	//	return validation.NewErrBadRecordFieldValue("emoji", fmt.Sprintf("too long: %v", l))
	//}
	for i, b := range v.Lists {
		if err := b.Validate(); err != nil {
			return fmt.Errorf("invalid list brief at index %v: %w", i, err)
		}
	}
	return nil
}

// ListBrief DTO
type ListBrief struct {
	ID string `json:"id" firestore:"id"`
	ListBase
}

// Validate returns error if not valid
func (v ListBrief) Validate() error {
	if strings.TrimSpace(v.ID) == "" {
		return validation.NewErrRecordIsMissingRequiredField("id")
	}
	if err := v.ListBase.Validate(); err != nil {
		return err
	}
	return nil
}

// ListDto DTO
type ListDto struct {
	ListBase
	dbmodels2.WithModified
	dbmodels2.WithUserIDs
	dbmodels2.WithTeamIDs

	Items []*ListItemBrief `json:"items,omitempty" firestore:"items,omitempty"`
	Count int              `json:"count" firestore:"count"`
}

// Validate returns error if not valid
func (v ListDto) Validate() error {
	if err := v.WithTeamIDs.Validate(); err != nil {
		return err
	}
	if err := v.WithUserIDs.Validate(); err != nil {
		return err
	}
	if err := v.ListBase.Validate(); err != nil {
		return err
	}
	if v.Count < 0 {
		return validation.NewErrBadRecordFieldValue("count", fmt.Sprintf("should be positive, got: %v", v.Count))
	}
	for i, item := range v.Items {
		if err := item.Validate(); err != nil {
			return validation.NewErrBadRecordFieldValue(fmt.Sprintf("items[%v]", i), err.Error())
		}
	}
	if count := len(v.Items); count != v.Count {
		return validation.NewErrBadRecordFieldValue("count", fmt.Sprintf("count != len(items): %v != %v", v.Count, count))
	}
	return nil
}
