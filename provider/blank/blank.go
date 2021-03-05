package blank

import (
	"encoding/json"
	"fmt"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
	"net/http"
)

type UserData struct {
	ID        string `json:"id"`
	ClientID  string `json:"client_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
	}
	return u, nil
}

func GetUserData(token string) (UserData, error) {
	url := fmt.Sprintf("https://example.com/?access_token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return UserData{}, err
	}
	defer resp.Body.Close()

	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s", b)

	dec := json.NewDecoder(resp.Body)
	var ud UserData
	err = dec.Decode(&ud)
	if err != nil {
		return UserData{}, err
	}
	return ud, nil
}
