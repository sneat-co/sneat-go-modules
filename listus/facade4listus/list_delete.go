package facade4listus

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/listus"
	"github.com/sneat-co/sneat-go-modules/listus/dal4listus"
	"github.com/sneat-co/sneat-go-modules/listus/models4listus"
	"github.com/strongo/validation"
)

// DeleteList deletes list
func DeleteList(ctx context.Context, user facade.User, request ListRequest) (err error) {
	if err = request.Validate(); err != nil {
		return
	}
	uid := user.GetID()
	if uid == "" {
		return validation.NewErrRecordIsMissingRequiredField("user.ContactID()")
	}
	id := models4listus.GetFullListID(request.ListType, request.ListID)
	key := dal4listus.NewTeamListKey(request.TeamID, id)
	input := dal4teamus.TeamItemRunnerInput[*models4listus.ListusTeamDto]{
		Counter:       "lists",
		TeamItem:      dal.NewRecord(key),
		BriefsAdapter: briefsAdapter(request.ListType, request.ListID),
	}
	err = dal4teamus.DeleteTeamItem(ctx, user, input, listus.ModuleID, nil)
	return
}
