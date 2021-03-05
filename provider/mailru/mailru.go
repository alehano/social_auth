/*
https://o2.mail.ru/app/
*/

package mailru

import (
	"encoding/json"
	"fmt"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/mailru"
	"net/http"
)

var Endpoint = mailru.Endpoint

type UserData struct {
	ID        string `json:"id"`
	ClientID  string `json:"client_id"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Locale    string `json:"locale"`
	Image     string `json:"image"`
}

func GetUser(tok *oauth2.Token) (user.User, error) {
	ud, err := GetUserData(tok.AccessToken)
	if err != nil {
		return user.User{}, err
	}
	u := user.User{
		ID:        ud.ID,
		Email:     ud.Email,
		FirstName: ud.FirstName,
		LastName:  ud.LastName,
		Photo:     ud.Image,
	}
	return u, nil
}

func GetUserData(token string) (UserData, error) {
	url := fmt.Sprintf("https://oauth.mail.ru/userinfo?access_token=%s", token)
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
