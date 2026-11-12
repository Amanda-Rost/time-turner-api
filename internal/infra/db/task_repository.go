package db

import (
	"context"

	"gorm.io/gorm"

	"time-turner-api/internal/domain/task"
)

type taskGormRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) task.TaskRepository {
	return &taskGormRepository{db: db}
}

func (r *taskGormRepository) Create(ctx context.Context, t *task.Task) error {
	return r.db.WithContext(ctx).Create(t).Error
}

func (r *taskGormRepository) Update(ctx context.Context, t *task.Task) error {
	return r.db.WithContext(ctx).Save(t).Error
}

func (r *taskGormRepository) FindAll(ctx context.Context) ([]*task.Task, error) {
	var tasks []*task.Task
	err := r.db.WithContext(ctx).Preload("User").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskGormRepository) FindById(ctx context.Context, id string) (*task.Task, error) {
	var t task.Task // Struct simples em memória
	err := r.db.WithContext(ctx).Preload("User").First(&t, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil 
}


func (r *taskGormRepository) FindByUserId(ctx context.Context, userId string) ([]*task.Task, error) {
	var tasks []*task.Task
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskGormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&task.Task{}, "id = ?", id).Error
}