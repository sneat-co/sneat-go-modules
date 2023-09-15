package dal4listus

import (
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-modules/listus/models4listus"
)

type ListusTeamContext = record.DataWithID[string, *models4listus.ListusTeamDto]
