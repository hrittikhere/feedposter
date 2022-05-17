package cmd

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func AuthenticateTwitter() *twitter.Client {

	ConsumerKey := os.Getenv("CONSUMER_KEY")
	ConsumerSecret := os.Getenv("CONSUMER_SECRET")
	AccessToken := os.Getenv("ACCESS_TOKEN")
	AccessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

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
		log.Fatal("Failed Tweet ", err)
	}

	TweetId := Response.IDStr

	return TweetId, nil
}
