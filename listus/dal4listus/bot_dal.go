package dal4listus

import (
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-modules/listus/models4listus"
)

// ListusChat is not used by bots framework
type ListusChat = record.DataWithID[string, models4listus.ListusChatData]
