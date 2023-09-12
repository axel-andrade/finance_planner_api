package main

import (
	"log"
	"os"

	"github.com/axel-andrade/finance_planner_api/internal/configuration/bootstrap"
	mongo_database "github.com/axel-andrade/finance_planner_api/internal/configuration/database/mongo"
	redis_database "github.com/axel-andrade/finance_planner_api/internal/configuration/database/redis"
	"github.com/axel-andrade/finance_planner_api/internal/configuration/http/server"
	"github.com/joho/godotenv"
)

/*
*
A função init por padrão é a primeira a ser executada pelo go.
Utilizada para configurar ou fazer um pré carregamento.
*
*/
func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	mongo_database.ConnectDB()
	redis_database.ConnectRedisDB()
}

func main() {
	dependecies := bootstrap.LoadDependencies()

	server := server.NewServer(os.Getenv("PORT"))
	server.AddRoutes(dependecies)
	server.Run()
}
