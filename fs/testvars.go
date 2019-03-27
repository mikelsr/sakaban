package fs

import "fmt"

var (
	muffinName = "muffin.jpg"
	muffinPath = fmt.Sprintf("%s/res/%s", projectPath(), muffinName)
)
