package models

import (
	"github.com/PuerkitoBio/goquery"

	l "../libs"
)

var HPURL = "http://www.tamurayukari.com/"

func HPGetLatest() (string, string) {
	root := l.CreateRequest(HPURL)
	defer root.Close()

	doc, err := goquery.NewDocumentFromReader(root)
	if err != nil {
		panic(err)
	}

	latestTitle := doc.Find("tr > td > a").Eq(0).Text()
	latestUrl, _ := doc.Find("tr > td > a").Attr("href")

	return latestTitle, latestUrl
}
