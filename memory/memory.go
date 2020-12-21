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
	StatRead   uint64 = 0
	StatWrite  uint64 = 0
	StatDelete uint64 = 0
)
