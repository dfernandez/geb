package domain

import "golang.org/x/oauth2"

type Token struct {
	OAuthToken *oauth2.Token
	Platform string
	ProfileUrl string
}
