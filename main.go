package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Scopes:       []string{"SCOPE"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://127.0.0.1:8080/server-authorize",
			TokenURL: "http://127.0.0.1:8081/server-token",
		},
	}

	url := conf.AuthCodeURL("your_state")
	println(fmt.Sprintf("Visit the URL for the auth dialog: %v", url))

	oauth2Token, err := conf.Exchange(ctx, "YOUR_AUTH_CODE")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, oauth2Token)
	resp, err := client.Get("http://127.0.0.1:8081/server-resource")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	str, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	println(fmt.Sprintf("Response: %s", resp.Status))
	println(fmt.Sprintf("Response: %s", string(str)))
}
