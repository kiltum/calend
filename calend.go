package main

// go get github.com/mmcdole/gofeed

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"strconv"
	"strings"
	"time"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://www.calend.ru/img/export/today-holidays.rss")
	time.LoadLocation("Europe/Moscow")
	today := strconv.Itoa(time.Now().Day())
	text := ""
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
