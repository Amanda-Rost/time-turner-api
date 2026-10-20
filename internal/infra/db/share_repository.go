package db

import (
	"context"

	"gorm.io/gorm"

	"time-turner-api/internal/domain/share"
)

type shareGormRepository struct {
	db *gorm.DB
}

func NewShareRepository(db *gorm.DB) share.ShareRepository {
	return &shareGormRepository{db: db}
}

func (r *shareGormRepository) Create(ctx context.Context, s *share.Share) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *shareGormRepository) Update(ctx context.Context, s *share.Share) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *shareGormRepository) FindAll(ctx context.Context) ([]*share.Share, error) {
	var shares []*share.Share
	// Preload carrega tanto o Usuario quanto a Tarefa associados ao compartilhamento
	err := r.db.WithContext(ctx).Preload("User").Preload("Task").Find(&shares).Error
	if err != nil {
		return nil, err
	}
	return shares, nil
}

func (r *shareGormRepository) FindByID(ctx context.Context, id string) (*share.Share, error) {
	var s share.Share
	err := r.db.WithContext(ctx).Preload("User").Preload("Task").First(&s, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}


func (r *shareGormRepository) FindByUserID(ctx context.Context, userID string) ([]*share.Share, error) {
	var shares []*share.Share
	err := r.db.WithContext(ctx).Preload("Task").Where("user_id = ?", userID).Find(&shares).Error
	if err != nil {
		return nil, err
	}
	return shares, nil
}

func (r *shareGormRepository) FindByTaskID(ctx context.Context, taskID string) ([]*share.Share, error) {
	var shares []*share.Share
	err := r.db.WithContext(ctx).Preload("User").Where("task_id = ?", taskID).Find(&shares).Error
	if err != nil {
		return nil, err
	}
	return shares, nil
}

func (r *shareGormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&share.Share{}, "id = ?", id).Error
}