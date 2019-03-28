package fs

import (
	"fmt"
	"math/rand"
)

const (
	permissionDir  = 0750
	permissionFile = 0750
)

var (
	muffinName = "muffin.jpg"
	muffinPath = fmt.Sprintf("%s/res/%s", projectPath(), muffinName)
	// testDir will contain the files generated for these tests
	testDir     = fmt.Sprintf("/tmp/sakaban-test-%d", rand.Intn(1e8))
	testFailDir = testDir + "/fail"
)
