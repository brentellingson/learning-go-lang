// Package main is a command line tool to log into an oauth2 server and retrieve a token.
// See: https://gist.github.com/marians/3b55318106df0e4e648158f1ffb43d38
package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

var (
	conf *oauth2.Config
	ctx  context.Context
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	// Use the authorization code that is pushed to the redirect
	// URL.
	code := queryParts["code"][0]
	log.Printf("code: %s\n", code)

	// Exchange will do the handshake to retrieve the initial access token.
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token: %s", tok.AccessToken)

	// show succes page
	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	if _, err := w.Write([]byte(msg)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	ctx = context.Background()
	conf = &oauth2.Config{
		ClientID:     "myclient",
		ClientSecret: "",
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:8081/realms/myrealm/protocol/openid-connect/auth",
			TokenURL: "http://localhost:8081/realms/myrealm/protocol/openid-connect/token",
		},
		// my own callback URL
		RedirectURL: "http://127.0.0.1:8082/oauth/callback",
	}

	// add transport for self-signed certificate to context
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	log.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)
	if err := open.Run(url); err != nil {
		log.Println(color.RedString("Failed to open browser to URL %s", url))
	}
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", url)

	http.HandleFunc("/oauth/callback", callbackHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
