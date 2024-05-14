package internal

import "sync"

var (
	globalApp     Application
	globalAppOnce sync.Once
)
