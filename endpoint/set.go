package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"voiddb/memory"
)

func Set(ctx *fiber.Ctx) error {
	memory.StatWrite++
	memory.StateChangedSaveRequired = true
	memory.Lock.Lock()
	memory.Data[ctx.Params("key")] = ctx.Body()
	memory.Lock.Unlock()
	return nil
}
