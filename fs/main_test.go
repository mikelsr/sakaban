package fs

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// prepare
	rand.Seed(time.Now().UTC().UnixNano())
	os.MkdirAll(testDir, permDir)
	os.Mkdir(testFailDir, 0000)
	// run tests
	m.Run()
	// cleanup
	os.RemoveAll(testDir)
}

func readGolden(t *testing.T, filename string) []byte {
	content, err := ioutil.ReadFile(filepath.Join(goldenDir, filename))
	if err != nil {
		t.Fatal(err)
	}
	return content
}
