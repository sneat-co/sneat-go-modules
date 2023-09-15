package models4assets

import "github.com/strongo/validation"

type DwellingData struct {
	BedRooms  int `json:"bedRooms,omitempty" firestore:"bedRooms,omitempty"`
	BathRooms int `json:"bathRooms,omitempty" firestore:"bathRooms,omitempty"`
}

func (v DwellingData) Validate() error {
	if v.BedRooms < 0 {
		return validation.NewErrBadRecordFieldValue("bedRooms", "negative value")
	}
	if v.BathRooms < 0 {
		return validation.NewErrBadRecordFieldValue("bathRooms", "negative value")
	}
	return nil
}

var _ AssetMain = (*DwellingAssetMainDto)(nil)

type DwellingAssetMainDto struct {
	AssetMainDto
	DwellingData
}

func (v *DwellingAssetMainDto) SpecificData() AssetSpecificData {
	return &v.DwellingData
}

func (v *DwellingAssetMainDto) SetSpecificData(data AssetSpecificData) {
	v.DwellingData = data.(DwellingData)
}

func (v *DwellingAssetMainDto) Validate() error {
	if err := v.AssetMainDto.Validate(); err != nil {
		return err
	}
	if err := v.DwellingData.Validate(); err != nil {
		return err
	}
	return nil
}

var _ AssetDbData = (*DwellingAssetDbData)(nil)

func NewDwellingAssetDbData() *DwellingAssetDbData {
	return &DwellingAssetDbData{
		DwellingAssetMainDto: new(DwellingAssetMainDto),
		AssetExtraDto:        new(AssetExtraDto),
	}
}

type DwellingAssetDbData struct {
	*DwellingAssetMainDto
	*AssetExtraDto
}

func (v DwellingAssetDbData) Validate() error {
	if err := v.DwellingAssetMainDto.Validate(); err != nil {
		return err
	}
	if err := v.AssetExtraDto.Validate(); err != nil {
		return err
	}
	return nil
}
