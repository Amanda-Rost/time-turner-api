package task

import (
	"time"
	"time-turner-api/internal/domain/user"
	"gorm.io/gorm"
)

type Status string
type Priority string

const (
	StatusPlanejada Status = "planejada"
	StatusConcluida Status = "concluida"
	StatusAdiada    Status = "adiada"
	StatusCancelada Status = "cancelada"

	PriorityBaixa Priority = "baixa"
	PriorityMedia Priority = "media"
	PriorityAlta  Priority = "alta"
)

type Task struct {
	ID        	string         `json:"id" gorm:"primaryKey;type:uuid"`
	UserID    	string         `json:"user_id" gorm:"type:uuid;not null"`          // Chave Estrangeira (FK)
	User      	*user.User          `json:"usuario,omitempty" gorm:"foreignKey:UserID"` // Relação no GORM (opcional na busca)
	Title     	string         `json:"titulo" gorm:"not null"`
	Date      	time.Time      `json:"data"`
	Alarm     	bool           `json:"alarme" gorm:"default:false"`
	Priority    Priority       `json:"prioridade" gorm:"type:varchar(20);default:'media'"`
	Status    	Status         `json:"status" gorm:"type:varchar(20);default:'planejada'"`
	Description string         `json:"descricao"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // Suporte nativo do GORM para Soft Delete
}
