package facade4scrumus

import (
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/scrumus/models4scrumus"
)

// DeleteTask deletes task
func DeleteTask(ctx context.Context, userContext facade.User, request DeleteTaskRequest) (err error) {
	if err = request.Validate(); err != nil {
		return
	}

	return runTaskWorker(ctx, userContext, request,
		func(ctx context.Context, tx dal.ReadwriteTransaction, params taskWorkerParams) (err error) {
			if params.task == nil {
				//err = errors.New("task not found by ContactID: " + request.Task)
				return
			}
			tasks := make([]*models4scrumus.Task, 0, len(params.tasks))
			for _, task := range params.tasks {
				if task.ID != request.Task {
					tasks = append(tasks, task)
				}
			}
			if len(tasks) == len(params.tasks) {
				return nil
			}
			var updateValue interface{}
			if len(tasks) == 0 {
				updateValue = dal.DeleteField
			} else {
				if err = models4scrumus.ValidateTasks(tasks); err != nil {
					return err
				}
				updateValue = tasks
			}
			updates := []dal.Update{
				{
					Field: "v",
					Value: dal.Increment(1),
				},
				{
					Field: fmt.Sprintf("statuses.%v.byType.%v", request.ContactID, request.Type),
					Value: updateValue,
				},
			}
			if request.Type == "risk" {
				updates = append(updates, dal.Update{
					Field: "risksCount",
					Value: dal.Increment(-1),
				})
			}
			if request.Type == "qna" {
				updates = append(updates, dal.Update{
					Field: "questionsCount",
					Value: dal.Increment(-1),
				})
			}
			return tx.Update(ctx, params.Meeting.Key, updates)
		})
}
