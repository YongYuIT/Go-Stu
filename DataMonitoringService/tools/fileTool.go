package tools

import (
	"io/ioutil"
	"os"
)

type FileTool struct {
}

func (thiz *FileTool) ReadFileTxt(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(contents)
}
