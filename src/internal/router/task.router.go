package router

import "github.com/gofiber/fiber/v2"

func (r *FiberRouter) GetTask(path string, h func(c *fiber.Ctx)) {
	r.task.Get(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}

func (r *FiberRouter) PostTask(path string, h func(c *fiber.Ctx)) {
	r.task.Post(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}

func (r *FiberRouter) PatchTask(path string, h func(c *fiber.Ctx)) {
	r.task.Patch(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}

func (r *FiberRouter) DeleteTask(path string, h func(c *fiber.Ctx)) {
	r.task.Delete(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}
