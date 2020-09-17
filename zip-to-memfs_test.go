package zipmemfs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	fs, err := New("example.zip")
	assert.Nil(t, err)

	// Check that we have:
	// a.txt
	// f/a.txt

	handle, err := fs.Open("example/a.txt")
	assert.Nil(t, err)

	data, err := ioutil.ReadAll(handle)
	assert.Nil(t, err)
	assert.Equal(t, "hello world\n", string(data))

	handle.Close()

	handle, err = fs.Open("example/b/a.txt")
	assert.Nil(t, err)

	data, err = ioutil.ReadAll(handle)
	assert.Nil(t, err)
	assert.Equal(t, "\"The important thing is not to stop questioning. Curiosity has its own reason for existing.\"", string(data))

	handle.Close()
}
