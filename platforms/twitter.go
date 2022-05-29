package platforms

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func authenticateTwitter() *twitter.Client {

	ConsumerKey, ConsumerKeyError := os.LookupEnv("CONSUMER_KEY")
	if !ConsumerKeyError {
		log.Fatal("CONSUMER_KEY Not Available ❌ Add it to GitHub Secrets ")
	}
	ConsumerSecret, ConsumerSecretError := os.LookupEnv("CONSUMER_SECRET")
	if !ConsumerSecretError {
		log.Fatal("CONSUMER_SECRET Not Available ❌ Add it to GitHub Secrets ")
	}
	AccessToken, AccessTokenError := os.LookupEnv("ACCESS_TOKEN")
	if !AccessTokenError {
		log.Fatal("ACCESS_TOKEN Not Available ❌ Add it to GitHub Secrets ")
	}
	AccessSecret, AccessSecretError := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !AccessSecretError {
		log.Fatal("ACCESS_TOKEN_SECRET Not Available ❌ Add it to GitHub Secrets ")
	}

	Config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	Token := oauth1.NewToken(AccessToken, AccessSecret)

	httpClient := Config.Client(oauth1.NoContext, Token)

	Client := twitter.NewClient(httpClient)

	return Client
}

func PublishToTwitter(Tweet string) (string, error) {

	Client := authenticateTwitter()

	Response, _, err := Client.Statuses.Update(Tweet, nil)
	if err != nil {
		log.Fatal("Failed Tweet ", err)
	}

	TweetId := Response.IDStr

	return TweetId, nil
}
