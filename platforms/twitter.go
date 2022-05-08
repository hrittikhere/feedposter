package cmd

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func AuthenticateTwitter() (*twitter.Client, error) {

	const ConsumerKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const ConsumerSecret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const AccessToken = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const AccessSecret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	Config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	Token := oauth1.NewToken(AccessToken, AccessSecret)

	httpClient := Config.Client(oauth1.NoContext, Token)

	Client := twitter.NewClient(httpClient)

	return Client, nil
}

func PublishToTwitter(Tweet string) (string, error) {

	Client, _ := AuthenticateTwitter()

	Response, _, _ := Client.Statuses.Update(Tweet, nil)

	return Response.IDStr, nil
}
