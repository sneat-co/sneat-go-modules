package facade4calendarium

import (
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/dto4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/contactus/dal4contactus"
)

func getHappeningContactRecords(ctx context.Context, tx dal.ReadwriteTransaction, request *dto4calendarium.HappeningContactRequest, params *happeningWorkerParams) (contact dal4contactus.ContactEntry, err error) {
	if request.Contact.TeamID == "" {
		request.Contact.TeamID = request.TeamID
	}
	contact = dal4contactus.NewContactEntry(request.Contact.TeamID, request.Contact.ID)

	if err = tx.GetMulti(ctx, []dal.Record{params.Happening.Record, params.TeamModuleEntry.Record, contact.Record}); err != nil {
		return contact, fmt.Errorf("failed to get records: %w", err)
	}
	if !params.TeamModuleEntry.Record.Exists() {
		return contact, fmt.Errorf("happening not found: %w", params.TeamModuleEntry.Record.Error())
	}
	if !contact.Record.Exists() {
		return contact, fmt.Errorf("contact not found: %w", contact.Record.Error())
	}
	return
}
