package facade4calendarium

import (
	"context"
	"errors"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/dto4calendarium"
	"github.com/sneat-co/sneat-go-modules/modules/calendarium/models4calendarium"
	"github.com/strongo/slice"
)

func AdjustSlot(ctx context.Context, user facade.User, request dto4calendarium.HappeningSlotDateRequest) (err error) {
	if err = request.Validate(); err != nil {
		return
	}

	var worker = func(ctx context.Context, tx dal.ReadwriteTransaction, params *happeningWorkerParams) (err error) {
		switch params.Happening.Dto.Type {
		case "single":
			return errors.New("only recurring happenings can be adjusted, single happenings should be updated")
		case "recurring":
			if err = adjustRecurringSlot(ctx, tx, params.Happening, request); err != nil {
				return fmt.Errorf("failed to adjust recurring happening: %w", err)
			}
			return err
		}
		return
	}

	if err = modifyHappening(ctx, user, request.HappeningRequest, worker); err != nil {
		return err
	}
	return nil
}

func adjustRecurringSlot(ctx context.Context, tx dal.ReadwriteTransaction, happening models4calendarium.HappeningContext, request dto4calendarium.HappeningSlotDateRequest) (err error) {
	//for _, teamID := range happening.Dto.TeamIDs { // TODO: run in parallel in go routine if > 1
	if err := adjustSlotInCalendarDay(ctx, tx, request.TeamID, happening.ID, request); err != nil {
		return fmt.Errorf("failed to adjust slot in calendar day record for teamID=%v: %w", request.TeamID, err)
	}
	//}
	return nil
}

func adjustSlotInCalendarDay(ctx context.Context, tx dal.ReadwriteTransaction, teamID, happeningID string, request dto4calendarium.HappeningSlotDateRequest) error {
	calendarDay := models4calendarium.NewCalendarDayContext(teamID, request.Date)
	if err := tx.Get(ctx, calendarDay.Record); err != nil {
		if !dal.IsNotFound(err) {
			return fmt.Errorf("failed to get calendar day record: %w", err)
		}
	}
	_, adjustment := calendarDay.Dto.GetAdjustment(happeningID, request.Slot.ID)
	if adjustment == nil {
		adjustment = &models4calendarium.HappeningAdjustment{
			HappeningID: happeningID,
		}
		calendarDay.Dto.HappeningAdjustments = append(calendarDay.Dto.HappeningAdjustments, adjustment)
	}
	adjustment.Slot = request.Slot
	var happeningIDsChanged bool
	if happeningIDsChanged = slice.Index(calendarDay.Dto.HappeningIDs, happeningID) < 0; happeningIDsChanged {
		calendarDay.Dto.HappeningIDs = append(calendarDay.Dto.HappeningIDs, happeningID)
	}

	if err := calendarDay.Dto.Validate(); err != nil {
		return fmt.Errorf("calednar day record is not valid: %w", err)
	}

	if calendarDay.Record.Exists() {
		updates := []dal.Update{
			{Field: "happeningAdjustments", Value: calendarDay.Dto.HappeningAdjustments},
		}
		if happeningIDsChanged {
			updates = append(updates, dal.Update{
				Field: "happeningIDs", Value: calendarDay.Dto.HappeningIDs,
			})
		}
		if err := tx.Update(ctx, calendarDay.Key, updates); err != nil {
			return fmt.Errorf("failed to update calendar day record with happening adjustment: %w", err)
		}
	} else {
		if err := tx.Insert(ctx, calendarDay.Record); err != nil {
			return fmt.Errorf("failed to insert calendar day record with happening adjustment: %w", err)
		}
	}
	return nil
}
