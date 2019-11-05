package mime2extension

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtToMime(t *testing.T) {
	var err error
	var mime string

	err, mime = Lookup(".mp4")
	makeAssertion(t, err, nil, mime, "video/mp4")

	err, mime = Lookup("/path/to/file.txt")
	makeAssertion(t, err, nil, mime, "text/plain")

	err, mime = Lookup("file.txt")
	makeAssertion(t, err, nil, mime, "text/plain")

	err, mime = Lookup(".TXT")
	makeAssertion(t, err, nil, mime, "text/plain")

	err, mime = Lookup("htm")
	makeAssertion(t, err, nil, mime, "text/html")

	err, mime = Lookup("foo")
	makeAssertion(t, err, errors.New("not found"), mime, "")

	err, mime = Lookup("")
	makeAssertion(t, err, errors.New("file_path is empty"), mime, "")
}

func TestMimeToExtSingle(t *testing.T) {
	var err error
	var ext string

	err, ext = Extension("video/mp4")
	makeAssertion(t, err, nil, ext, "mp4")

	err, ext = Extension("image/jpeg")
	makeAssertion(t, err, nil, ext, "jpeg")

	err, ext = Extension("foo")
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
