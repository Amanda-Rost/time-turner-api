package share

import (
	"time"
	"time-turner-api/internal/domain/task"
	"time-turner-api/internal/domain/user"

	"gorm.io/gorm"
)

type Permission string

const (
	PermissionLeitor       Permission = "leitor"
	PermissionParticipante Permission = "participante"
	PermissionResponsavel  Permission = "responsavel"
)

type Share struct {
	ID         string         `json:"id" gorm:"primaryKey;type:uuid"`
	TaskID     string         `json:"tarefa_id" gorm:"type:uuid;not null"`         // Chave Estrangeira (FK)
	Task       *task.Task     `json:"tarefa,omitempty" gorm:"foreignKey:TaskID"` // Relação no GORM (opcional na busca)
	UserID     string         `json:"usuario_id" gorm:"type:uuid;not null"`
	User       *user.User     `json:"usuario,omitempty" gorm:"foreignKey:UserID"`
	Permission Permission     `json:"permissao" gorm:"type:varchar(20);default:'leitor'"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // Suporte nativo do GORM para Soft Delete
}
