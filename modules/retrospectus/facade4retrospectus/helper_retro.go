package facade4retrospectus

import (
	"context"
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/facade4meetingus"
)

var runRetroWorker = func(ctx context.Context, userID string, request facade4meetingus.Request, worker facade4meetingus.Worker) error {
	return facade4meetingus.RunMeetingWorker(ctx, userID, request, MeetingRecordFactory{}, worker)
}
