package mime2extension

//build Lookup and Extension and Extensions

import (
	"errors"
	"path/filepath"
	"strings"
)

// Lookup looks up a mimetype from a file extension or a file path
func Lookup(filePath string) (string, error) {

	if filePath == "" {
		return "", errors.New("filePath is empty")
	}

	filePath = strings.ToLower(filePath)
	extension := filepath.Ext(filePath)

	if extension == "" {
		extension = filePath
	} else {
		extension = extension[1:]
	}

	if val, ok := bdMap.extToMime[extension]; ok {
		return val, nil
	}

	return "", errors.New("not found")
}

// Extension returns an extension for the given mime type
func Extension(mime string) (string, error) {

	mime = strings.ToLower(mime)

	if val, ok := bdMap.mimeToExt[mime]; ok {
		return val[0], nil
	}
	return "", errors.New("not found")
}

// Extensions returns a list of possible extensions for the given mime type
func Extensions(mime string) []string {

	mime = strings.ToLower(mime)

	if val, ok := bdMap.mimeToExt[mime]; ok {
		return val
	}

	return []string{}
}
