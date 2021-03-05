/*
https://console.developers.google.com/apis/credentials
*/

package google

import (
	"encoding/json"
	"fmt"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

var Endpoint = google.Endpoint

type UserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GetUser(tok *oauth2.Token) (user.User, error) {
	ud, err := GetUserData(tok.AccessToken)
	if err != nil {
		return user.User{}, err
	}
	u := user.User{
		ID:        ud.ID,
		Email:     ud.Email,
		FirstName: ud.GivenName,
		LastName:  ud.FamilyName,
		Photo:     ud.Picture,
	}
	if u.LastName == u.FirstName {
		u.LastName = ""
	}
	return u, nil
}

func GetUserData(token string) (UserData, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return UserData{}, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var ud UserData
	err = dec.Decode(&ud)
	if err != nil {
		return UserData{}, err
	}
	return ud, nil
}
