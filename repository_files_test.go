package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFile(t *testing.T) {
	ts, gitlab := Stub("stubs/files/post.json")
	response, err := gitlab.CreateFile("1", "file", "branch", "text", "content", "message")

	assert.Equal(t, err, nil)
	assert.Equal(t, response.File_Name, "app/post.rb")
	defer ts.Close()
}

func TestUpdateFile(t *testing.T) {
	ts, gitlab := Stub("stubs/files/put.json")
	response, err := gitlab.UpdateFile("1", "file", "branch", "text", "content", "message")

	assert.Equal(t, err, nil)
	assert.Equal(t, response.File_Name, "app/put.rb")
	defer ts.Close()
}
