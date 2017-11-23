package model

import (
	"os"
	"sync"
	// cloudgoLogger "util/logger"
)

var wg sync.WaitGroup
var fin = os.Stdin

// Load : load all resources for cloudgo.
func Load() {
}

// Save : Save all data for cloudgo.
func Save() error {
	return nil
}
