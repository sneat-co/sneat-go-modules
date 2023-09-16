package models4assetus

import (
	dbmodels2 "github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/sneat-co/sneat-go-core/modules/contactus/briefs4contactus"
	"github.com/sneat-co/sneat-go-core/validate"
	"github.com/sneat-co/sneat-go-modules/assetus/briefs4assetus"
)

type AssetSpecificData interface {
	Validate() error
}

type AssetMain interface {
	Validate() error
	AssetMainData() *AssetMainDto
	SpecificData() AssetSpecificData
	SetSpecificData(AssetSpecificData)
}

type AssetCreationData interface {
	Validate() error
	AssetMainData() *AssetMainDto
	SpecificData() AssetSpecificData
}

// AssetDbData defines mandatory fields & methods on an asset record
type AssetDbData interface {
	AssetMain
	AssetExtraData() *AssetExtraDto
}

// AssetMainDto was intended to be used in both AssetBaseDto and request to create an asset,
// but it was not a good idea as not clear how to manage module specific fields
type AssetMainDto struct {
	briefs4assetus.AssetBrief
	briefs4assetus.WithMultiTeamAssetBriefs
	dbmodels2.WithTags
	briefs4contactus.WithMultiTeamContactIDs
	dbmodels2.WithCustomFields
	AssetDates
}

func (v *AssetMainDto) AssetMainData() *AssetMainDto {
	return v
}

func (v *AssetMainDto) Validate() error {
	if err := v.AssetBrief.Validate(); err != nil {
		return err
	}
	if err := v.WithMultiTeamAssetBriefs.Validate(); err != nil {
		return err
	}
	if err := v.WithTags.Validate(); err != nil {
		return err
	}
	if err := v.WithMultiTeamContactIDs.Validate(); err != nil {
		return err
	}
	if err := v.WithCustomFields.Validate(); err != nil {
		return err
	}
	if err := v.AssetDates.Validate(); err != nil {
		return err
	}
	return nil
}

// AssetExtraDto defines extra fields on an asset record that are not passed in create asset request
type AssetExtraDto struct {
	dbmodels2.WithModified
	dbmodels2.WithUserIDs
	dbmodels2.WithTeamIDs
}

func (v *AssetExtraDto) AssetExtraData() *AssetExtraDto {
	return v
}

func (v *AssetExtraDto) Validate() error {
	if err := v.WithModified.Validate(); err != nil {
		return err
	}
	if err := v.WithUserIDs.Validate(); err != nil {
		return err
	}
	if err := v.WithTeamIDs.Validate(); err != nil {
		return err
	}
	return nil
}

// AssetDates defines dates of an asset - TODO: consider refactoring to custom fields?
type AssetDates struct {
	DateOfBuild       string `json:"dateOfBuild,omitempty" firestore:"dateOfBuild,omitempty"`
	DateOfPurchase    string `json:"dateOfPurchase,omitempty" firestore:"dateOfPurchase,omitempty"`
	DateInsuredTill   string `json:"dateInsuredTill,omitempty" firestore:"dateInsuredTill,omitempty"`
	DateCertifiedTill string `json:"dateCertifiedTill,omitempty" firestore:"dateCertifiedTill,omitempty"`
}

// Validate returns error if not valid
func (v *AssetDates) Validate() error {
	if v.DateOfBuild != "" {
		if _, err := validate.DateString(v.DateOfBuild); err != nil {
			return err
		}
	}
	if v.DateOfPurchase != "" {
		if _, err := validate.DateString(v.DateOfPurchase); err != nil {
			return err
		}
	}
	if v.DateInsuredTill != "" {
		if _, err := validate.DateString(v.DateInsuredTill); err != nil {
			return err
		}
	}
	if v.DateCertifiedTill != "" {
		if _, err := validate.DateString(v.DateCertifiedTill); err != nil {
			return err
		}
	}
	return nil
}
