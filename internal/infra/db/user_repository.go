package db

import (
	"context"

	"gorm.io/gorm"

	"time-turner-api/internal/domain/user"
)

type userGormRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userGormRepository{db: db}
}

func (r *userGormRepository) Create(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userGormRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil // Retorna o ponteiro *user.User
}

func (r *userGormRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userGormRepository) Update(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Save(u).Error
}

func (r *userGormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&user.User{}, "id = ?", id).Error
}