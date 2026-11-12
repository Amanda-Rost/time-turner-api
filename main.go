package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"time-turner-api/internal/infra/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, utilizando variáveis de ambiente do sistema")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=timeturner port=5432 sslmode=disable"
	}

	database, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Falha ao inicializar o banco de dados: %v", err)
	}

	// instancia os repositórios injetando a conexão
	userRepo := db.NewUserRepository(database)
	taskRepo := db.NewTaskRepository(database)
	shareRepo := db.NewShareRepository(database)

	_ = userRepo
	_ = taskRepo
	_ = shareRepo

	// cria a instância do servidor e ativa automaticamente dois middlewares
	// Logger e Recovery
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API executando em Go!",
		})
	})

	// inicia servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor Go rodando na porta :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}