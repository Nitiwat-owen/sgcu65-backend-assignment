package task

import (
	"github.com/gofiber/fiber/v2"
	"sgcu65-backend-assignment/src/internal/repository"
)

type Handler struct {
	taskRepo     repository.ITaskRepository
	userTaskRepo repository.IUserTaskRepository
}

func NewHandler(taskRepo repository.ITaskRepository, userTaskRepo repository.IUserTaskRepository) *Handler {
	return &Handler{taskRepo: taskRepo, userTaskRepo: userTaskRepo}
}

func (h *Handler) CreateTask(c *fiber.Ctx) {

}

func (h *Handler) FindAllTask(c *fiber.Ctx) {

}

func (h *Handler) FindTaskById(c *fiber.Ctx) {

}

func (h *Handler) UpdateTask(c *fiber.Ctx) {

}

func (h *Handler) DeleteTask(c *fiber.Ctx) {

}

func (h *Handler) AssignUser(c *fiber.Ctx) {
	
}

func (h *Handler) RemoveUser(c *fiber.Ctx) {

}
