package fs

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

const (
	// TODO: why doesn't it work with 0644?
	permDir  = 0755
	permFile = 0755
)

var (
	// directory containing golden files
	goldenDir = "testdata"
	// directory to run tests on
	testDir = filepath.Join(os.TempDir(), fmt.Sprintf("sakaban-test-%d",
		rand.Intn(1e5)))
	// unwriteable directory
	testFailDir = filepath.Join(testDir, "/fail")
)
