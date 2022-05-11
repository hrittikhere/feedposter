package cmd

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

func AuthenticateTwitter() (*twitter.Client) {

	const ConsumerKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const ConsumerSecret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const AccessToken = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const AccessSecret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	Config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	Token := oauth1.NewToken(AccessToken, AccessSecret)

	httpClient := Config.Client(oauth1.NoContext, Token)

	Client := twitter.NewClient(httpClient)

	return Client
}

func PublishToTwitter(Tweet string) (string, error) {

	Client := AuthenticateTwitter()


	Response, _, err := Client.Statuses.Update(Tweet, nil)
	if err != nil {
		log.Fatal("Failed Tweet ",err)
	}

	TweetId := Response.IDStr

	return TweetId, nil
}
