package gogitlab

import (
	"encoding/json"
	"net/url"
)

const (
	repo_url_create_file = "/projects/:id/repository/files" // Create new file in repository
	repo_url_update_file = "/projects/:id/repository/files" // Update existing file in repository
)

type CommitResponse struct {
	File_Path   string
	Branch_Name string
}

/*
Create new file in repository.

    POST /projects/:id/repository/files

Parameters:

    file_path (required) - Full path to new file. Ex. lib/class.rb
    branch_name (required) - The name of branch
    encoding (optional) - 'text' or 'base64'. Text is default.
    content (required) - File content
    commit_message (required) - Commit message

*/
func (g *Gitlab) CreateFile(id, file_path, branch_name, encoding, content, commit_message string) (*CommitResponse, error) {

	path, opaque := g.ResourceUrlRaw(repo_url_create_file, map[string]string{":id": id})

	v := url.Values{}
	v.Set("file_path", file_path)
	v.Set("branch_name", branch_name)
	v.Set("encoding", encoding)
	v.Set("content", content)
	v.Set("commit_message", commit_message)

	body := v.Encode()
	var commitResponse *CommitResponse

	contents, err := g.buildAndExecRequestRaw("POST", path, opaque, []byte(body))

	if err == nil {
		err = json.Unmarshal(contents, &commitResponse)
	}

	return commitResponse, err
}

/*
Update existing file in repository.

    PUT /projects/:id/repository/files

Parameters:

    file_path (required) - Full path to new file. Ex. lib/class.rb
    branch_name (required) - The name of branch
    encoding (optional) - 'text' or 'base64'. Text is default.
    content (required) - File content
    commit_message (required) - Commit message

*/
func (g *Gitlab) UpdateFile(id, file_path, branch_name, encoding, content, commit_message string) (*CommitResponse, error) {

	path, opaque := g.ResourceUrlRaw(repo_url_create_file, map[string]string{":id": id})

	v := url.Values{}
	v.Set("file_path", file_path)
	v.Set("branch_name", branch_name)
	v.Set("encoding", encoding)
	v.Set("content", content)
	v.Set("commit_message", commit_message)

	body := v.Encode()
	var commitResponse *CommitResponse

	contents, err := g.buildAndExecRequestRaw("PUT", path, opaque, []byte(body))
	if err == nil {
		err = json.Unmarshal(contents, &commitResponse)
	}

	return commitResponse, err
}
