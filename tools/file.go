package tools

import (
	"errors"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"io"
	"os"
)

var (
	UnknownFileType = errors.New("Unknown file type")
)

func getFileKind(file io.Reader) (kind types.Type, err error) {
	b := make([]byte, 262)
	if _, err = file.Read(b); err != nil {
		return
	}
	return filetype.Match(b)
}

func GetExt(file io.Reader) (ext string, err error) {
	kind, err := getFileKind(file)
	if err != nil {
		return
	}
	if kind == filetype.Unknown {
		return "", UnknownFileType
	}
	return kind.Extension, nil
}

func GetMimeType(file io.Reader) (mime string, err error) {
	kind, err := getFileKind(file)
	if err != nil {
		return
	}
	if kind == filetype.Unknown {
		return "", UnknownFileType
	}
	return kind.MIME.Value, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
