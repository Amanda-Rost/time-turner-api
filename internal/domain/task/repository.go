package task

import "context"

type TaskRepository interface {
    Create(ctx context.Context, user *Task) error
    Update(ctx context.Context, task *Task) error
    FindAll(ctx context.Context) ([]*Task, error)
	FindById(ctx context.Context, id string) (*Task, error)
	FindByUserId(ctx context.Context, userId string) ([]*Task, error)
    Delete(ctx context.Context, id string) error
}
