package mime2extension

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtToMime(t *testing.T) {
	var err error
	var mime string

	mime, err = Lookup(".mp4")
	makeAssertion(t, err, nil, mime, "video/mp4")

	mime, err = Lookup("/path/to/file.txt")
	makeAssertion(t, err, nil, mime, "text/plain")

	mime, err = Lookup("file.txt")
	makeAssertion(t, err, nil, mime, "text/plain")

	mime, err = Lookup(".TXT")
	makeAssertion(t, err, nil, mime, "text/plain")

	mime, err = Lookup("htm")
	makeAssertion(t, err, nil, mime, "text/html")

	mime, err = Lookup("foo")
	makeAssertion(t, err, errors.New("not found"), mime, "")

	mime, err = Lookup("")
	makeAssertion(t, err, errors.New("filePath is empty"), mime, "")
}

func TestMimeToExtSingle(t *testing.T) {
	var err error
	var ext string

	ext, err = Extension("video/mp4")
	makeAssertion(t, err, nil, ext, "mp4")

	ext, err = Extension("image/jpeg")
	makeAssertion(t, err, nil, ext, "jpeg")

	ext, err = Extension("foo")
	makeAssertion(t, err, errors.New("not found"), ext, "")
}

func TestMimeToExtList(t *testing.T) {
	var exts []string

	exts = Extensions("video/mp4")
	assertListsEqual(t, exts, []string{"mp4", "mp4v", "mpg4"})

	exts = Extensions("image/jpeg")
	assertListsEqual(t, exts, []string{"jpeg", "jpe", "jpg"})

	exts = Extensions("foo")
	assertListsEqual(t, exts, []string{""})
}

func makeAssertion(t *testing.T, actualErr, expectedErr error, actualOut, expectedOut string) {
	assert.Equal(t, actualErr, expectedErr)
	assert.Equal(t, actualOut, expectedOut)
}

func assertListsEqual(t *testing.T, l1, l2 []string) {
	for i := 0; i < len(l1); i++ {
		assert.Equal(t, l1[i], l2[i])
	}
}
