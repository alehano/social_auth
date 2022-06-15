/*
https://developers.facebook.com/apps/
*/

package facebook

import (
	"encoding/json"
	"fmt"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"net/http"
)

var Endpoint = facebook.Endpoint

type UserData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
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
		Photo:     ud.Picture.Data.URL,
	}
	return u, nil
}

func GetUserData(token string) (UserData, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v9.0/me?fields=id,first_name,last_name,email,picture.type(large)&access_token=%s", token)
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
