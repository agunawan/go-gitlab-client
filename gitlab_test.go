package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceUrl(t *testing.T) {
	gitlab := NewGitlab("http://base_url/", "api_path", "token")

	assert.Equal(t, gitlab.ResourceUrl(projects_url, nil), "http://base_url/api_path/projects?private_token=token")
	assert.Equal(t, gitlab.ResourceUrl(project_url, map[string]string{":id": "123"}), "http://base_url/api_path/projects/123?private_token=token")
	tokenless_url, _ := gitlab.TokenlessResourceUrlRaw(projects_url, nil)
	assert.Equal(t, tokenless_url, "http://base_url/api_path/projects")
}
