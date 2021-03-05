/*
https://vk.com/apps?act=manage
*/

package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"net/http"
)

var Endpoint = vk.Endpoint

type ResponseData struct {
	Response []UserData `json:"response"`
}

type UserData struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhotoURL  string `json:"photo_200"`
}

func GetUser(tok *oauth2.Token) (user.User, error) {
	userID := 0
	if uID, ok := tok.Extra("user_id").(float64); ok {
		userID = int(uID)
	}
	if userID == 0 {
		return user.User{}, errors.New("zero user ID")
	}

	ud, err := GetUserData(userID, tok)
	if err != nil {
		return user.User{}, err
	}

	u := user.User{
		FirstName: ud.FirstName,
		LastName:  ud.LastName,
		Photo:     ud.PhotoURL,
	}

	u.ID = fmt.Sprintf("%d", userID)
	u.Email = fmt.Sprintf("%s", tok.Extra("email"))

	return u, nil
}

// https://vk.com/dev/users.get
func GetUserData(id int, tok *oauth2.Token) (UserData, error) {
	version := "5.126"
	url := fmt.Sprintf("https://api.vk.com/method/users.get?user_ids=%d&fields=photo_200&access_token=%s&v=%s",
		id, tok.AccessToken, version)
	resp, err := http.Get(url)
	if err != nil {
		return UserData{}, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var res ResponseData
	err = dec.Decode(&res)
	if err != nil {
		return UserData{}, err
	}
	if len(res.Response) == 0 {
		return UserData{}, nil
	}
	return res.Response[0], nil
}
