package main

import (
	"github.com/mmcdole/gofeed"
	"fmt"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://www.acquired.fm/episodes?format=rss")
	fmt.Println(feed.Title)
	fmt.Println(feed.Link)
	fmt.Println(feed.Updated)
	fmt.Println(feed.Items)

}
