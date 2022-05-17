package main

import (
	"fmt"
	"time"

	"github.com/hrittikhere/feedposter/platforms"
	"github.com/mmcdole/gofeed"
)

func main() {

	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://dev.to/feed/kcdchennai")

	for _, item := range feed.Items {

		NowTime := time.Now()
		ParsedNowTime := time.Unix(NowTime.Unix(), 0)

		PublishedTime := item.PublishedParsed
		ParsedPublishedTime := time.Unix(PublishedTime.Unix(), 0)

		if ParsedNowTime.Sub(ParsedPublishedTime).Hours() < 24 {
			// 0 */4 * * *
			PostTitle := item.Title
			PostLink := item.Link
			PostAuthor := item.Author.Name
			// Category := item.Categories[0]

			Tweet := fmt.Sprintf("%s was published by %s 🎉🎉🎉 \n %s @hrittikhere TEST", PostTitle, PostAuthor, PostLink)

			TweeetId, _ := cmd.PublishToTwitter(Tweet)

			fmt.Println(TweeetId, "Posted")

		}

	}

}
