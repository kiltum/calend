package main

// go get github.com/mmcdole/gofeed

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	time.LoadLocation("Europe/Moscow")
	feed, _ := fp.ParseURL("http://www.calend.ru/img/export/today-holidays.rss")
	//fmt.Println(er)

	today := strconv.Itoa(time.Now().Day())
	text := ""
	// fmt.Println(feed)
	for _, element := range feed.Items {
		if strings.HasPrefix(element.Title, today) {
			if len(text) == 0 {
				t := strings.Split(element.Title, " ")
				text = t[0] + " " + t[1] + ": "
			}
			text = text + element.Title[strings.Index(element.Title, "- ")+2:] + "; "
		}
	}
	fmt.Println(text[:len(text)-2])
}
