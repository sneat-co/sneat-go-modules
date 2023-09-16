package dal4assetus

import (
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/assetus/const4assetus"
	"github.com/sneat-co/sneat-go-modules/assetus/models4assetus"
)

type AssetusTeamContext = record.DataWithID[string, *models4assetus.AssetusTeamDto]

// AssetsCollection is a name of a collection in DB
const AssetsCollection = "assets"

var AssetusRootKey = dal.NewKeyWithID(dal4teamus.Collection, const4assetus.ModuleID)
