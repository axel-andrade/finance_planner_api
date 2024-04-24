package main

import (
	"log"
	"os"

	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/server"
	mongo_database "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/mongo"
	redis_database "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/redis"
	"github.com/axel-andrade/finance_planner_api/internal/infra/bootstrap"
	"github.com/joho/godotenv"
)

/*
* The init function is called after all the variable declarations in the package have evaluated their initializers, and
* those are evaluated only after all the imported packages have been initialized.
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
