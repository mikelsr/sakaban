package fs

import (
	"fmt"
	"os"
)

// ProjectPath returns the directory this project is supposed to be at
func ProjectPath() string {
	return fmt.Sprintf("%s/src/bitbucket.org/mikelsr/sakaban", os.Getenv("GOPATH"))
}
