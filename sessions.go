package gogitlab

import (
	"encoding/json"
	"net/url"
)

const (
	repo_url_new_session = "/session" // Create new file in repository
)

type Session struct {
	Id                 int
	Username           string
	Email              string
	Name               string
	Private_Token      string
	Blocked            bool
	Created_At         string
	Bio                string
	Skype              string
	Linkedin           string
	Twitter            string
	Website_Url        string
	Dark_Scheme        bool
	Theme_Id           int
	Is_Admin           bool
	Can_Create_Group   bool
	Can_Create_Team    bool
	Can_Create_Project bool
}

/*
Login to get private token

    POST /session

Parameters:

    login (required) - The login of user
    email (required if login missing) - The email of user
    password (required) - Valid password

*/
func (g *Gitlab) NewSession(login, email, password string) (*Session, error) {

	path, opaque := g.TokenlessResourceUrlRaw(repo_url_new_session, nil)

	v := url.Values{}
	v.Set("login", login)
	v.Set("email", email)
	v.Set("password", password)

	body := v.Encode()
	var session *Session

	contents, err := g.buildAndExecRequestRaw("POST", path, opaque, []byte(body))

	if err == nil {
		err = json.Unmarshal(contents, &session)
	}

	return session, err
}
