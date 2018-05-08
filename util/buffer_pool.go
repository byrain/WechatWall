package util

import (
	"bytes"
	"sync"
)

var TextBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4KB
	},
}

var MediaBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 10<<20)) // 10MB
	},
}
