package fs

import (
	"fmt"
	"math/rand"
)

const (
	permRWX   = 0755
	permNoRWX = 0000
)

var (
	muffinHash        = "QmNPMibWJrRkXmkDUatzkLXKnvPmUBDNyHXX7h52c15Ehy"
	muffinBlockHashes = []string{
		"Qmergz5cpuEGWeUFPL4JgEpvQeFZkqFUYFARebK61yEc1F",
		"QmU4j8pB5m2xhxJN9QmfHVmLZxvknDUMMJ1nDwubs8pkZe",
		"QmatFE9ZMwSRJbg9zzXwcd6uNnmkS7bZ9ENpciXdEJZ6ZM",
	}
	muffinPath = fmt.Sprintf("%s/res/muffin.jpg", ProjectPath())
	// testDir will contain the files generated for this tests
	testDir      = fmt.Sprintf("/tmp/sakaban-test-%d", rand.Intn(1e8))
	testFile     = testDir + "/file"
	testFailDir  = testDir + "/fail"
	testFailFile = testFailDir + "/file"
)
