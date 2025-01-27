package router

import "github.com/gofiber/fiber/v2"

type FiberRouter struct {
	*fiber.App
	user fiber.Router
	task fiber.Router
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	user := r.Group("/user")
	task := r.Group("/task")

	return &FiberRouter{r, user, task}
}
