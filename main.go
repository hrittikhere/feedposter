package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hrittikhere/feedposter/extractor"
	"github.com/hrittikhere/feedposter/platforms"
	"github.com/mmcdole/gofeed"
)

func main() {

	file, err := os.ReadFile("feed.yaml")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
	}

	data := extractor.GetFeedURL(file)

	for _, FeedURL := range data {
		postParser(FeedURL)
	}

}

func postParser(FeedURL string) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL(FeedURL)

	for _, item := range feed.Items {

		NowTime := time.Now()
		ParsedNowTime := time.Unix(NowTime.Unix(), 0)

		PublishedTime := item.PublishedParsed
		ParsedPublishedTime := time.Unix(PublishedTime.Unix(), 0)

		if ParsedNowTime.Sub(ParsedPublishedTime).Hours() < 24 {

			PostTitle := item.Title
			PostLink := item.Link
			PostAuthor := item.Author.Name
			Category := item.Categories

			hashtags, _ := getHastags(Category)

			Tweet := fmt.Sprintf("%s was published by %s  🎉🎉🎉 \n %s \n %s", PostTitle, PostAuthor, hashtags, PostLink)

			TweeetId, _ := platform.PublishToTwitter(Tweet)

			fmt.Printf("%s with the TweetID: %s was posted! ✔️ \n ", PostTitle, TweeetId)

		}

	}
}

func getHastags(Categories []string) (string, error) {

	var Category string

	for _, element := range Categories {
		Category += "#" + element + " "
	}

	return Category, nil
}
