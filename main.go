package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	"github.com/mrtokyo33/PiTrainer/src/controllers"
	"github.com/mrtokyo33/PiTrainer/src/repositories/postgres"
	"github.com/mrtokyo33/PiTrainer/src/routes"
	"github.com/mrtokyo33/PiTrainer/src/usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database!")

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("error running migrations:", err)
	}

	log.Println("Migrations applied")

	userRepo := postgres.NewUserRepository(db)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepo)
	createUserHandler := controllers.NewCreateUserHandler(createUserUseCase)
	loginUseCase := usecases.NewLoginUserUseCase(userRepo)
	loginHandler := controllers.NewLoginUserHandler(loginUseCase)

	r := gin.Default()

	api := r.Group("/api")

	routes.InitHealthRoutes(api)
	routes.InitUserRoutes(api, createUserHandler.CreateUser)
	routes.InitAuthRoutes(api, loginHandler.Login)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
