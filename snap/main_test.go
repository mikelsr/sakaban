package snap

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// prepare
	rand.Seed(time.Now().UTC().UnixNano())
	os.MkdirAll(testDir, testPermDir)
	os.Mkdir(testFailDir, 0000)
	// run tests
	m.Run()
	// cleanup
	os.RemoveAll(testDir)
}
