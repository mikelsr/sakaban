package fs

import (
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

// TestMain will create and delete the testing directory
// before and after running tests, respectively
func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.MkdirAll(testDir, permRWX)
	os.Mkdir(testFailDir, permRWX)
	ioutil.WriteFile(testFailFile, nil, permNoRWX)
	os.Chmod(testFailDir, permNoRWX)
	m.Run()
	// cleanup
	os.RemoveAll(testDir)
}
