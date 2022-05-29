// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/hrittikhere/feedposter/platforms"
// 	"github.com/mmcdole/gofeed"
// )

// func main() {

// 	fp := gofeed.NewParser()

// 	feed, _ := fp.ParseURL("https://dev.to/feed/kcdchennai")

// 	for _, item := range feed.Items {

// 		NowTime := time.Now()
// 		ParsedNowTime := time.Unix(NowTime.Unix(), 0)

// 		PublishedTime := item.PublishedParsed
// 		ParsedPublishedTime := time.Unix(PublishedTime.Unix(), 0)

// 		if ParsedNowTime.Sub(ParsedPublishedTime).Hours() < 6 {

// 			PostTitle := item.Title
// 			PostLink := item.Link
// 			PostAuthor := item.Author.Name
// 			Category := item.Categories

// 			hashtags, _ := GetHastags(Category)

// 			Tweet := fmt.Sprintf("%s was published by %s  🎉🎉🎉 \n %s \n %s", PostTitle, PostAuthor, hashtags, PostLink)

// 			TweeetId, _ := cmd.PublishToTwitter(Tweet)

// 			fmt.Printf("%s with the TweetID: %s was posted! ✔️ \n ", PostTitle, TweeetId)

// 		}

// 	}

// }

// func GetHastags(Categories []string) (string, error) {

// 	var Category string

// 	for _, element := range Categories {
// 		Category += "#" + element + " "
// 	}

// 	return Category, nil
// }
