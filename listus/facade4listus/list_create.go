package facade4listus

import (
	"context"
	"errors"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	dbmodels2 "github.com/sneat-co/sneat-go-core/models/dbmodels"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/listus/const4listus"
	"github.com/sneat-co/sneat-go-modules/listus/dal4listus"
	"github.com/sneat-co/sneat-go-modules/listus/models4listus"
	"github.com/strongo/random"
	"strings"
	"time"
)

// CreateList creates a new list
func CreateList(ctx context.Context, user facade.User, request CreateListRequest) (response CreateListResponse, err error) {
	request.Title = strings.TrimSpace(request.Title)
	if err = request.Validate(); err != nil {
		return
	}
	err = dal4teamus.CreateTeamItem(ctx, user, "", request.TeamRequest, const4listus.ModuleID,
		func(ctx context.Context, tx dal.ReadwriteTransaction, params *dal4teamus.ModuleTeamWorkerParams[*models4listus.ListusTeamDto]) (err error) {
			var listGroup *models4listus.ListGroup
			for _, lg := range params.TeamModuleEntry.Data.ListGroups {
				if lg.Type == request.Type {
					if listGroup != nil {
						return fmt.Errorf("team record has at least 2 list groups with the same type: %v", lg.Type)
					}
					listGroup = lg
				}
			}
			isNewListGroup := listGroup == nil
			if isNewListGroup {
				listGroup = &models4listus.ListGroup{Type: request.Type, Title: request.Type}
				params.TeamModuleEntry.Data.ListGroups = append(params.TeamModuleEntry.Data.ListGroups, listGroup)
			} else if err = listGroup.Validate(); err != nil {
				return fmt.Errorf("list group received from team record is not valid: %w", err)
			}

			for i, brief := range listGroup.Lists {
				if brief.Title == request.Title {
					return fmt.Errorf("an attempt to create a new list with duplicate title [%v] equal to list at index %v {id=%v, type=%v}", request.Title, i, brief.ID, brief.Type)
				}
			}
			id := request.Type
			idTryNumber := 0
		checkId:
			if idTryNumber++; idTryNumber == 10 {
				return errors.New("too many attempts to generate random list ContactID")
			}
			for _, brief := range listGroup.Lists {
				if brief.ID == id {
					id = random.ID(2)
					goto checkId
				}
			}
			modified := dbmodels2.Modified{
				By: user.GetID(),
				At: time.Now(),
			}
			list := models4listus.ListDto{
				WithModified: dbmodels2.WithModified{
					WithCreated: dbmodels2.WithCreated{
						CreatedAt: modified.At,
						CreatedBy: modified.By,
					},
					WithUpdated: dbmodels2.WithUpdated{
						UpdatedAt: modified.At,
						UpdatedBy: modified.By,
					},
				},
				WithTeamIDs: dbmodels2.WithTeamIDs{
					TeamIDs: []string{request.TeamID},
				},
				ListBase: models4listus.ListBase{
					Type:  request.Type,
					Title: request.Title,
				},
			}
			if err = list.Validate(); err != nil {
				return fmt.Errorf("formed list DTO struct is not valid: %w", err)
			}
			listKey := dal4listus.NewTeamListKey(request.TeamID, id)
			listRecord := dal.NewRecordWithData(listKey, &list)
			if err = tx.Insert(ctx, listRecord); err != nil {
				return fmt.Errorf("failed to insert list record")
			}
			listGroup.Lists = append(listGroup.Lists, &models4listus.ListBrief{
				ID: id,
				ListBase: models4listus.ListBase{
					Type:  request.Type,
					Title: request.Type,
				},
			})
			if err = listGroup.Validate(); err != nil {
				return fmt.Errorf("list group is not valid after adding a new list: %w", err)
			}
			params.TeamUpdates = append(params.TeamUpdates, dal.Update{
				Field: "listGroups." + request.Type,
				Value: listGroup,
			})
			if err != nil {
				return fmt.Errorf("failed to generate new list ContactID: %w", err)
			}
			return nil
		},
	)
	return
}
