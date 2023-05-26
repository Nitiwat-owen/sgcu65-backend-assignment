package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"sgcu65-backend-assignment/src/config"
	"sgcu65-backend-assignment/src/database"
	"sgcu65-backend-assignment/src/internal/repository"
)

func main() {
	postgresConf, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatal(
			"failed to init postgres connection",
			zap.Error(err),
			zap.String("action", "init postgres connection"),
		)
	}

	postgresConn, err := database.InitPostgresDatabase(postgresConf)
	if err != nil {
		log.Fatal(
			"failed to init postgres connection",
			zap.Error(err),
			zap.String("action", "init postgres connection"),
		)
	}
	fmt.Println("Connected Database!!!")

	userRepo := repository.UserRepositoryImpl{DB: postgresConn}
	taskRepo := repository.TaskRepositoryImpl{DB: postgresConn}
	userTaskRepo := repository.UserTaskRepositoryImpl{DB: postgresConn}
}
