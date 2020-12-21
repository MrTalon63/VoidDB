package endpoint

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
	"voiddb/memory"
)

func Stats(ctx *fiber.Ctx) error {

	var stats = `# current data entries count
voiddb_entries_total %d
# runtime in seconds
voiddb_runtime_seconds %d
# reads from the startup
voiddb_stat_read %d
# writes from the startup
voiddb_stat_write %d
# deletions from the startup
voiddb_stat_delete %d`

	var entriesTotal = len(memory.Data)
	var runtimeSeconds = time.Now().Unix() - memory.Startup.Unix()

	_, _ = fmt.Fprintf(ctx, stats,
		entriesTotal,
		runtimeSeconds,
		memory.StatRead,
		memory.StatWrite,
		memory.StatDelete)

	return nil
}
