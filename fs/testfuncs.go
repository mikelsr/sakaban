package fs

import (
	"fmt"
	"os"
)

// projectPath returns the directory this project is supposed to be at
func projectPath() string {
	return fmt.Sprintf("%s/src/bitbucket.org/mikelsr/sakaban", os.Getenv("GOPATH"))
}
