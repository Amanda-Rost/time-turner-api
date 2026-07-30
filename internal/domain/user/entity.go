package user

import (
	"time"
)

// User representa a entidade de domínio do usuário.
type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"nome" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Pashword     string    `json:"-" gorm:"not null"` // "-" oculta a senha no JSON de saída
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}