package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Credential ...
type Credential struct {
	Secret string
}

var (
	baseURL = "https://google.com/recaptcha/api/siteverify"
)

// Verify ...
func (cred *Credential) Verify(response, remoteip string) (res map[string]interface{}, err error) {
	postStr := url.Values{"secret": {cred.Secret}, "response": {response}, "remoteip": {remoteip}}
	responsePost, err := http.PostForm(baseURL, postStr)
	if err != nil {
		return res, err
	}

	defer responsePost.Body.Close()
	body, err := ioutil.ReadAll(responsePost.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
