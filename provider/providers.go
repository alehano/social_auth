package provider

import (
	"github.com/alehano/social_auth/provider/facebook"
	"github.com/alehano/social_auth/provider/google"
	"github.com/alehano/social_auth/provider/mailru"
	"github.com/alehano/social_auth/provider/vk"
	"github.com/alehano/social_auth/provider/yandex"
	"github.com/alehano/social_auth/user"
	"golang.org/x/oauth2"
)

type Name string

const (
	VK       Name = "vk"
	Yandex   Name = "yandex"
	MailRU   Name = "mailru"
	Google   Name = "google"
	Facebook Name = "facebook"
)

type Credentials struct {
	ID     string
	Secret string
}

type param struct {
	Endpoint oauth2.Endpoint
	Scopes   []string
}

var params = map[Name]param{
	VK:       {Endpoint: vk.Endpoint, Scopes: []string{"email"}},
	Yandex:   {Endpoint: yandex.Endpoint},
	MailRU:   {Endpoint: mailru.Endpoint, Scopes: []string{"userinfo"}},
	Google:   {Endpoint: google.Endpoint, Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}},
	Facebook: {Endpoint: facebook.Endpoint, Scopes: []string{"email"}},
}

func GetUser(name Name, tok *oauth2.Token) (user.User, error) {
	switch name {
	case VK:
		return vk.GetUser(tok)
	case Yandex:
		return yandex.GetUser(tok)
	case MailRU:
		return mailru.GetUser(tok)
	case Google:
		return google.GetUser(tok)
	case Facebook:
		return facebook.GetUser(tok)
	}
	return user.User{}, nil
}

func GetConfig(name Name, data Credentials, redirectURL string) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     data.ID,
		ClientSecret: data.Secret,
		RedirectURL:  redirectURL,
		Scopes:       params[name].Scopes,
		Endpoint:     params[name].Endpoint,
	}
	return conf
}
