package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"sgcu65-backend-assignment/src/config"
	"sgcu65-backend-assignment/src/database"
	"sgcu65-backend-assignment/src/internal/handler/task"
	"sgcu65-backend-assignment/src/internal/handler/user"
	"sgcu65-backend-assignment/src/internal/repository"
	"sgcu65-backend-assignment/src/internal/router"
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

	userHandler := user.NewHandler(&userRepo)
	taskHandler := task.NewHandler(&taskRepo, &userTaskRepo)

	r := router.NewFiberRouter()

	r.GetUser("/", userHandler.FindAllUsers)
	r.GetUser("/:userId", userHandler.FindUserById)
	r.PostUser("/", userHandler.CreateUser)
	r.PatchUser("/:userId", userHandler.UpdateUser)
	r.DeleteUser("/:userId", userHandler.DeleteUser)

	r.GetTask("/", taskHandler.FindAllTask)
	r.GetTask("/:taskId", taskHandler.FindTaskById)
	r.PostTask("/", taskHandler.CreateTask)
	r.PostTask("/assign", taskHandler.AssignUser)
	r.PatchTask("/:taskId", taskHandler.UpdateTask)
	r.DeleteTask("/:taskId", taskHandler.DeleteTask)
	r.DeleteTask("/:taskId/remove/userId", taskHandler.RemoveUser)
}
