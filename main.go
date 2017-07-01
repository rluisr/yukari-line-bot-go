package main

import (
	"time"
	"runtime"

	cntrlr "github.com/rluisr/yukari-line-bot-go/controllers"
)

func main() {
	go cronCheckArticle()
	runtime.Gosched()

	cntrlr.StartServer()
}

func cronCheckArticle() {
	for {
		isHPDiff, isFCDiff := cntrlr.CheckDiff()
		if isHPDiff == true {
			cntrlr.PushMessage("HP")
		}

		if isFCDiff == true {
			cntrlr.PushMessage("FC")
		}

		time.Sleep(180 * time.Second) // 180秒
	}
}