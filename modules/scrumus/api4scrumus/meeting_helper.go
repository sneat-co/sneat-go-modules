package api4scrumus

import (
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/facade4meetingus"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/facade4scrumus"
)

// meetingParams holds records settings for MeetingID entity
var meetingParams = facade4meetingus.Params{
	RecordFactory: facade4scrumus.MeetingRecordFactory{},
	BeforeSafe:    beforeSafeScrum,
}

var beforeSafeScrum = func(ctx context.Context, tx dal.ReadwriteTransaction, params facade4meetingus.WorkerParams) error {
	if err := facade4scrumus.UpdateLastScrumIDIfNeeded(ctx, tx, params); err != nil {
		return fmt.Errorf("failed to update team with last scrum ContactID: %w", err)
	}
	return nil
}
