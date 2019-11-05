package mime2extension

//build Lookup and Extension and Extensions

import (
	"errors"
	"path/filepath"
	"strings"
)

/**
 * Given file path and it will return mime type
 * err if not found mime
 */
func Lookup(file_path string) (error, string) {

	if file_path == "" {
		return errors.New("file_path is empty"), ""
	}

	file_path = strings.ToLower(file_path)
	extension := filepath.Ext(file_path)

	if extension == "" {
		extension = file_path
	} else {
		extension = extension[1:]
	}

	if val, ok := bdMap.extToMime[extension]; ok {
		return nil, val
	}

	return errors.New("Not found"), ""
}

/**
 * Given file extension when input is mimetype
 * err if not found extension
 */
func Extension(mime string) (error, string) {

	mime = strings.ToLower(mime)

	if val, ok := bdMap.mimeToExt[mime]; ok {
		return nil, val[0]
	}
	return errors.New("Not found"), ""
}

/**
 * Given list of files extension when input is mimetype
 */
func Extensions(mime string) []string {

	mime = strings.ToLower(mime)

	if val, ok := bdMap.mimeToExt[mime]; ok {
		return val
	}

	return []string{}
}
