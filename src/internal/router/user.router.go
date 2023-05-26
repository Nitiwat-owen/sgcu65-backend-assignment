package router

import "github.com/gofiber/fiber/v2"

func (r *FiberRouter) GetUser(path string, h func(c *fiber.Ctx)) {
	r.user.Get(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}

func (r *FiberRouter) PostUser(path string, h func(c *fiber.Ctx)) {
	r.user.Post(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}

func (r *FiberRouter) PatchUser(path string, h func(c *fiber.Ctx)) {
	r.user.Patch(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}

func (r *FiberRouter) DeleteUser(path string, h func(c *fiber.Ctx)) {
	r.user.Delete(path, func(c *fiber.Ctx) error {
		h(c)
		return nil
	})
}
