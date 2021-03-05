/*
https://oauth.yandex.ru/
*/

package yandex

import (
	"encoding/json"
	"fmt"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
	"net/http"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://oauth.yandex.ru/authorize",
	TokenURL: "https://oauth.yandex.ru/token",
}

type UserData struct {
	ID              string   `json:"id"`
	Login           string   `json:"login"`
	ClientID        string   `json:"client_id"`
	DisplayName     string   `json:"display_name"`
	RealName        string   `json:"real_name"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	Sex             string   `json:"sex"`
	DefaultEmail    string   `json:"default_email"`
	Emails          []string `json:"emails"`
	DefaultAvatarID string   `json:"default_avatar_id"`
	IsAvatarEmpty   bool     `json:"is_avatar_empty"`
}

func GetUser(tok *oauth2.Token) (user.User, error) {
	ud, err := GetUserData(tok.AccessToken)
	if err != nil {
		return user.User{}, err
	}

	u := user.User{
		ID:        ud.ID,
		Email:     ud.DefaultEmail,
		FirstName: ud.FirstName,
		LastName:  ud.LastName,
	}

	// Photo
	if !ud.IsAvatarEmpty {
		u.Photo = fmt.Sprintf("https://avatars.yandex.net/get-yapic/%s/islands-200", ud.DefaultAvatarID)
	}

	return u, nil
}

// https://yandex.ru/dev/passport/doc/dg/reference/response.html
func GetUserData(token string) (UserData, error) {
	url := fmt.Sprintf("https://login.yandex.ru/info?format=json&oauth_token=%s", token)
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
