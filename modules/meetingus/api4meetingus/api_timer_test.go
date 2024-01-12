package api4meetingus

import (
	"context"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/facade4meetingus"
	"testing"
)

func TestToggleMeetingTimer(t *testing.T) {
	toggleTimer = func(ctx context.Context, userContext facade.User, params facade4meetingus.ToggleParams) (response facade4meetingus.ToggleTimerResponse, err error) {
		return
	}
	params := facade4meetingus.Params{
		RecordFactory: nil,
		BeforeSafe:    nil,
	}
	handler := ToggleMeetingTimer(params)
	if handler == nil {
		t.Fatal("handler = nil")
	}
	// TODO: implement
}

func TestToggleMemberTimer(t *testing.T) {
	toggleTimer = func(ctx context.Context, userContext facade.User, params facade4meetingus.ToggleParams) (response facade4meetingus.ToggleTimerResponse, err error) {
		return
	}
	params := facade4meetingus.Params{
		RecordFactory: nil,
		BeforeSafe:    nil,
	}
	handler := ToggleMemberTimer(params)
	if handler == nil {
		t.Fatal("handler = nil")
	}
	// TODO: implement
}
