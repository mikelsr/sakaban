package snap

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

const (
	testPermFile = 0755
	testPermDir  = 0755
)

var (
	testDir = filepath.Join(os.TempDir(), fmt.Sprintf("sakaban-test-%d",
		rand.Intn(1e5)))
	// unwriteable directory
	testFailDir = filepath.Join(testDir, "/fail")
)
