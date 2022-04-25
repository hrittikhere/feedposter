package main

import (
    "fmt"

    "github.com/mmcdole/gofeed"
)

func main() {


fp := gofeed.NewParser()

feed, _ := fp.ParseURL("https://dev.to/feed/kcdchennai")


for _, item := range feed.Items {

	fmt.Println(item.Title)
	fmt.Println(item.Link)
	fmt.Println(item.Published)
	fmt.Println(item.Author)

	}

}