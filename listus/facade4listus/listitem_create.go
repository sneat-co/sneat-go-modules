package facade4listus

import (
	"context"
	"errors"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-core/modules/teamus/facade4teamus"
	"github.com/sneat-co/sneat-go-modules/listus"
	"github.com/sneat-co/sneat-go-modules/listus/dal4listus"
	"github.com/sneat-co/sneat-go-modules/listus/models4listus"
	"github.com/strongo/random"
	"github.com/strongo/slice"
	"github.com/strongo/validation"
)

// CreateListItems creates list items
func CreateListItems(ctx context.Context, userContext facade.User, request CreateListItemsRequest) (response CreateListItemResponse, err error) {
	if err = request.Validate(); err != nil {
		return
	}
	uid := userContext.GetID()

	err = dal4teamus.RunModuleTeamWorker(ctx, userContext, request.TeamRequest, listus.ModuleID,
		func(ctx context.Context, tx dal.ReadwriteTransaction, params *dal4teamus.ModuleTeamWorkerParams[*models4listus.ListusTeamDto]) error {
			listID := models4listus.GetFullListID(request.ListType, request.ListID)
			listKey := dal4listus.NewTeamListKey(request.TeamID, listID)
			var listDto models4listus.ListDto
			var listRecord = dal.NewRecordWithData(listKey, &listDto)
			var isNewList bool
			if err := tx.Get(ctx, listRecord); err != nil {
				if !dal.IsNotFound(err) {
					return fmt.Errorf("failed to get list record: %w", err)
				}
				isNewList = true

				isOkToAutoCreateList :=
					request.ListType == models4listus.ListTypeToBuy && request.ListID == "groceries" ||
						request.ListType == models4listus.ListTypeToWatch && request.ListID == "movies"

				team, err := facade4teamus.GetTeamByID(ctx, tx, request.TeamID)
				if err != nil {
					return fmt.Errorf("failed to get team record: %w", err)
				}
				if slice.Index(team.Data.UserIDs, uid) < 0 {
					return fmt.Errorf("user have no access to this team")
				}
				var isExistingListGroup bool
				var listGroup *models4listus.ListGroup
				for _, lg := range params.TeamModuleEntry.Data.ListGroups {
					if lg.Type == request.ListType {
						listGroup = lg
						isExistingListGroup = true
						break
					}
				}
				if !isExistingListGroup {
					listGroup = &models4listus.ListGroup{
						Type:  request.ListType,
						Title: request.ListType,
					}
					params.TeamModuleEntry.Data.ListGroups = append(params.TeamModuleEntry.Data.ListGroups, listGroup)
				}
				var isExistingList bool
				for _, l := range listGroup.Lists {
					if l.ID == request.ListID {
						isExistingList = true
						break
					}
				}
				if !isExistingList {
					if !isOkToAutoCreateList {
						return validation.NewErrBadRequestFieldValue("listID", "unknown list")
					}
					listBrief := models4listus.ListBrief{
						ID: request.ListID,
						ListBase: models4listus.ListBase{
							Type:  request.ListType,
							Title: request.ListID,
						},
					}
					if listBrief.Type == models4listus.ListTypeToBuy && listBrief.ID == "groceries" {
						listBrief.Emoji = "🛒"
					}
					listGroup.Lists = append(listGroup.Lists, &listBrief)
					if err := tx.Update(ctx, team.Key, []dal.Update{
						{
							Field: "lists." + request.ListType,
							Value: listGroup,
						},
					}); err != nil {
						return fmt.Errorf("failed to update team record: %w", err)
					}
				}

				listDto.TeamIDs = []string{request.TeamID}
				listDto.UserIDs = []string{uid}
				listDto.Type = request.ListType
				listDto.Title = request.ListID
				if request.ListType == "to-buy" && request.ListID == "groceries" {
					listDto.Emoji = "🛒"
				}
			}
			for i, item := range request.Items {
				id, err := generateRandomListItemID(listDto.Items, item.ID)
				if err != nil {
					return fmt.Errorf("failed to generate random id for item #%v: %w", i, err)
				}
				listItem := models4listus.ListItemBrief{
					ID:           id,
					ListItemBase: item.ListItemBase,
				}
				listDto.Items = append(listDto.Items, &listItem)
			}
			listDto.Count = len(listDto.Items)
			if err := listDto.Validate(); err != nil {
				return fmt.Errorf("list record is not valid: %w", err)
			}
			if isNewList {
				if err := tx.Insert(ctx, listRecord); err != nil {
					return fmt.Errorf("failed to insert list record: %w", err)
				}
			} else {
				if slice.Index(listDto.UserIDs, uid) < 0 {
					return errors.New("current user does not have access to the list: userID=" + uid)
				}
				if err := tx.Update(ctx, listKey, []dal.Update{
					{
						Field: "items",
						Value: listDto.Items,
					},
					{
						Field: "count",
						Value: len(listDto.Items),
					},
				}); err != nil {
					return fmt.Errorf("failed to update list record: %w", err)
				}
			}
			return nil
		})
	return
}

func generateRandomListItemID(items []*models4listus.ListItemBrief, initialID string) (id string, err error) {
	isDuplicateID := func() bool {
		for _, item := range items {
			if item.ID == id {
				return true
			}
		}
		return false
	}
	id = initialID
	if !isDuplicateID() {
		return
	}
next:
	for i := 0; i <= 100; i++ {
		id = random.ID(3)
		if isDuplicateID() {
			continue next
		}
		return
	}
	return "", errors.New("too many attempts to generate random ContactID")
}
