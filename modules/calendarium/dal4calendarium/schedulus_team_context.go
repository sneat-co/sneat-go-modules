package dal4calendarium

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/const4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/models4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/teamus/dal4teamus"
)

type CalendariumTeamContext = record.DataWithID[string, *models4calendarium.CalendariumTeamDto]

func NewCalendariumTeamKey(teamID string) *dal.Key {
	return dal4teamus.NewTeamModuleKey(teamID, const4calendarium.ModuleID)
}

func NewCalendariumTeamContext(teamID string) CalendariumTeamContext {
	key := NewCalendariumTeamKey(teamID)
	return record.NewDataWithID(teamID, key, new(models4calendarium.CalendariumTeamDto))
}

func GetCalendariumTeam(ctx context.Context, tx dal.ReadwriteTransaction, teamID string) (CalendariumTeamContext, error) {
	calendariumTeam := NewCalendariumTeamContext(teamID)
	return calendariumTeam, tx.Get(ctx, calendariumTeam.Record)
}
