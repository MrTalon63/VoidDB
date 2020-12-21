package memory

import (
	"sync"
	"time"
)

var (
	Startup                  = time.Now()
	Data                     = make(map[string]interface{})
	Lock                     = sync.RWMutex{}
	StateChangedSaveRequired = false
)

var (
	StatRead   int64 = 0
	StatWrite  int64 = 0
	StatDelete int64 = 0
)
