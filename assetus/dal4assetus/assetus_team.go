package dal4assetus

import (
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/models4assets"
)

type AssetusTeamContext = record.DataWithID[string, *models4assets.AssetusTeamDto]

// AssetsCollection is a name of a collection in DB
const AssetsCollection = "assets"

var AssetusRootKey = dal.NewKeyWithID(dal4teamus.Collection, assetus.ModuleID)
