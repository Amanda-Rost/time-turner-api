package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"time-turner-api/internal/domain/share"
	"time-turner-api/internal/domain/task"
	"time-turner-api/internal/domain/user"
)

func ConnectPostgres(dsn string) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no banco de dados: %w", err)
	}

	// Executa as migrações automáticas das tabelas
	err = database.AutoMigrate(
		&user.User{},
		&task.Task{},
		&share.Share{},
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao executar auto-migrate: %w", err)
	}

	log.Println("Conexão com PostgreSQL e migrações concluídas!")
	return database, nil
}