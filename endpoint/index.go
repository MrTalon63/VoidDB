package endpoint

import (
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	_, _ = ctx.WriteString("Running VoidDB 1.0")
	return nil
}
