package user

import (
	"github.com/gofiber/fiber/v2"
	"sgcu65-backend-assignment/src/internal/repository"
)

type Handler struct {
	userRepo repository.IUserRepository
}

func NewHandler(userRepo repository.IUserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

func (h *Handler) FindAllUsers(c *fiber.Ctx) {

}

func (h *Handler) CreateUser(c *fiber.Ctx) {

}

func (h *Handler) FindUserById(c *fiber.Ctx) {

}

func (h *Handler) UpdateUser(c *fiber.Ctx) {

}

func (h *Handler) DeleteUser(c *fiber.Ctx) {

}
