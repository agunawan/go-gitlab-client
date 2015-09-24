package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {
	ts, gitlab := Stub("stubs/sessions/post.json")
	response, err := gitlab.NewSession("user", "email", "password")

	assert.Equal(t, err, nil)
	assert.Equal(t, response.Private_Token, "dd34asd13as")
	defer ts.Close()
}
