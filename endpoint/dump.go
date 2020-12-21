package endpoint

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"voiddb/memory"
)

func Dump(ctx *fiber.Ctx) error {

	memory.StatRead++
	output, err := json.Marshal(memory.Data)

	if err != nil {
		_, _ = ctx.WriteString(fmt.Sprint(err))
		return nil
	}

	_, _ = ctx.WriteString(string(output))
	return nil
}

// mxj does not escape map values
//func Dump(ctx *fasthttp.RequestCtx) {
//
//	mv := mxj.Map(mem.Data)
//	xmlValue, err := mv.XmlIndent("", "  ", "data")
//
//	if err != nil {
//		_, _ = ctx.WriteString(fmt.Sprint(err))
//		return
//	}
//
//	_, _ = ctx.WriteString(string(xmlValue))
//}
