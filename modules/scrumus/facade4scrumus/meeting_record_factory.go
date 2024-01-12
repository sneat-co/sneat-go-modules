package facade4scrumus

import (
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/models4meetingus"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/models4scrumus"
)

// MeetingRecordFactory factory
type MeetingRecordFactory struct {
}

// Collection "api4meetingus"
func (MeetingRecordFactory) Collection() string {
	return "api4meetingus"
}

// NewRecordData create new api4meetingus record
func (MeetingRecordFactory) NewRecordData() models4meetingus.MeetingInstance {
	return &models4scrumus.Scrum{}
}
