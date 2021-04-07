package service

import (
	"golang.org/x/oauth2"
	"os"
)

type LoginService struct{
}

func (ls LoginService) GetLoginUrl() (url string){
	scopes := []string{"repo"}
	github := oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SERVER_HOST"),
		Scopes:       scopes,
		Endpoint:     oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	loginUrl := github.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return loginUrl
}

func (ls LoginService) GetToken(code string) (token string, err error){
	scopes := []string{"repo"}
	github := oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SERVER_HOST"),
		Scopes:       scopes,
		Endpoint:     oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	tokenObj, err := github.Exchange(oauth2.NoContext, code)
	return tokenObj.AccessToken, err
}