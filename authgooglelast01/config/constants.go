package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "http://localhost:3000/google/callback",
		Scopes: []string{
			"SCOPE1",
			"SCOPE2",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
