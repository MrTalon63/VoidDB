package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"voiddb/memory"
)

func Del(ctx *fiber.Ctx) error {
	memory.StatDelete++
	memory.StateChangedSaveRequired = true
	memory.Lock.Lock()
	delete(memory.Data, ctx.Params("key"))
	memory.Lock.Unlock()
	return nil
}
