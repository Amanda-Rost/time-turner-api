package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"time-turner-api/internal/domain/entities"
)

func main() {
	// 1. Configurar Conexão com o Banco de Dados
	dsn := "host=localhost user=postgres password=postgres dbname=timeturner port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar no banco de dados: %v", err)
	}

	// 2. Executar Auto-migração das Entidades (substitui o Prisma Migrate)
	db.AutoMigrate(&entities.Usuario{}, &entities.Tarefa{}, &entities.Compartilhamento{})

	// 3. Inicializar o Roteador HTTP (Gin)
	r := gin.Default()

	// 4. Rota Healthcheck de Exemplo
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API executando em Go!",
		})
	})

	// 5. Iniciar o Servidor na porta 8080
	log.Println("Servidor Go rodando na porta :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}