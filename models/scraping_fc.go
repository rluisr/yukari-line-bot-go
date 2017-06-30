package models

import (
	"github.com/PuerkitoBio/goquery"

	l "../libs"
)

var FCURL = "https://www.mellowpretty.com"

func FCGetLatest() (string, string) {
	root := l.CreateRequest(FCURL)
	defer root.Close()

	doc, err := goquery.NewDocumentFromReader(root)
	if err != nil {
		panic(err)
	}

	latestTitle := doc.Find("div.post > h3 > a").Eq(0).Text()

	return latestTitle, FCURL
}
