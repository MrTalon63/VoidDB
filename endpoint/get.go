package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"voiddb/memory"
)

func Get(ctx *fiber.Ctx) error {
	memory.StatRead++
	raw := memory.Data[ctx.Params("key")]
	if raw == nil {
		_, _ = ctx.WriteString("")
		return nil
	}
	data := raw.([]byte)
	_, _ = ctx.Write(data)
	return nil
}
