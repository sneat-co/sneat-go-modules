package briefs4assets

import (
	"fmt"
	"github.com/sneat-co/sneat-go-core/geo"
	"github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/sneat-co/sneat-go-modules/assetus/const4assets"
	"github.com/strongo/slice"
	"github.com/strongo/validation"
	"strings"
)

// AssetBrief keeps main props of an asset
type AssetBrief struct {
	IsRequest  bool                         `json:"isRequest,omitempty" firestore:"isRequest,omitempty"` // This is used to flag that struct is part of a request and some validations should be skipped
	Title      string                       `json:"title" firestore:"title"`                             // Should be required if the make, model & reg number are not provided
	Status     const4assets.AssetStatus     `json:"status,omitempty" firestore:"status,omitempty"`
	Category   const4assets.AssetCategory   `json:"category" firestore:"category"`
	Type       const4assets.AssetType       `json:"type" firestore:"type"`
	Possession const4assets.AssetPossession `json:"possession" firestore:"possession"`
	CountryID  geo.CountryAlpha2            `json:"countryID"  firestore:"countryID"` // intentionally not omitempty so can be used in queries
	Make       string                       `json:"make" firestore:"make"`            // intentionally not omitempty so can be used in queries
	Model      string                       `json:"model" firestore:"model"`          // intentionally not omitempty so can be used in queries
	RegNumber  string                       `json:"regNumber"  firestore:"regNumber"` // intentionally not omitempty so can be used in queries
	dbmodels.WithOptionalRelatedAs
}

func (v *AssetBrief) Equal(v2 *AssetBrief) bool {
	return *v == *v2
}

// Validate returns error if not valid
func (v *AssetBrief) Validate() error {
	if !v.IsRequest && v.Make == "" && v.Model == "" && v.RegNumber == "" && strings.TrimSpace(v.Title) == "" {
		return validation.NewErrRecordIsMissingRequiredField("title")
	}
	if err := v.WithOptionalRelatedAs.Validate(); err != nil {
		return err
	}
	if !const4assets.IsValidAssetStatus(v.Status) {
		return validation.NewErrBadRecordFieldValue("status", fmt.Sprintf("unknown status: %s", v.Status))
	}
	if strings.TrimSpace(v.CountryID) == "" {
		return validation.NewErrRecordIsMissingRequiredField("countryID")
	}
	checkType := func(types []string) error {
		switch v.Type {
		case "":
			return validation.NewErrRecordIsMissingRequiredField("type")
		default:
			if slice.Index(const4assets.AssetVehicleTypes, v.Type) < 0 {
				return validation.NewErrBadRecordFieldValue("type", fmt.Sprintf("unknown %s type: %s", v.Category, v.Type))
			}
		}
		return nil
	}
	switch v.Category {
	case "":
		return validation.NewErrRecordIsMissingRequiredField("category")
	case const4assets.AssetCategoryVehicle:
		if err := checkType(const4assets.AssetVehicleTypes); err != nil {
			return err
		}
	case const4assets.AssetCategoryRealEstate:
		if err := checkType(const4assets.AssetRealEstateTypes); err != nil {
			return err
		}
	case const4assets.AssetCategorySportGear:
		if err := checkType(const4assets.AssetSportGearTypes); err != nil {
			return err
		}
	case const4assets.AssetCategoryDocument:
		if err := checkType(const4assets.AssetDocumentTypes); err != nil {
			return err
		}
	default:
		return validation.NewErrBadRecordFieldValue("category", "unknown asset category: "+string(v.Category))
	}

	if strings.TrimSpace(v.Make) == "" {
		return validation.NewErrRecordIsMissingRequiredField("make")
	}
	if strings.TrimSpace(v.Model) == "" {
		return validation.NewErrRecordIsMissingRequiredField("model")
	}
	if err := const4assets.ValidateAssetPossession(v.Possession, true); err != nil {
		return err
	}
	return nil
}
