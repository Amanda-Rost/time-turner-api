package share

import "context"

type ShareRepository interface {
    Create(ctx context.Context, user *Share) error
    Update(ctx context.Context, task *Share) error
    FindAll(ctx context.Context) ([]*Share, error)
	FindById(ctx context.Context, id string) (*Share, error)
	FindByUserId(ctx context.Context, userId string) ([]*Share, error)
    FindByTaskId(ctx context.Context, taskId string) ([]*Share, error)
    Delete(ctx context.Context, id string) error
}
