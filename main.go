package main

import (
	"fmt"
	"time"

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

		if ParsedNowTime.Sub(ParsedPublishedTime).Hours() < 4 {
			// 0 */4 * * *
			fmt.Println(item.Title)
			fmt.Println(item.Link)
			fmt.Println(item.Author)
			fmt.Println("------------------------------------------------")
		}

	}

}
