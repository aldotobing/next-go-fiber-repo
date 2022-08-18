package google

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// OauthGoogleURLAPI ...
const OauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// GetGoogleProfile ...
func GetGoogleProfile(token string) (res map[string]interface{}, err error) {
	response, err := http.Get(OauthGoogleURLAPI + token)
	if err != nil {
		return res, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func getUser(token string) (res []byte, err error) {
	response, err := http.Get(OauthGoogleURLAPI + token)
	if err != nil {
		return res, err
	}
	defer response.Body.Close()
	res, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
