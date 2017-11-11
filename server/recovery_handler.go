package server

import (
	"fmt"
	"runtime/debug"
)

// RecoveryHandler handles grpc recovery.
func RecoveryHandler(p interface{}) (err error) {
	stack := string(debug.Stack())
	err = fmt.Errorf("GRPC recovery handler error: %s. Stack trace: %s", p, stack)
	return
}
