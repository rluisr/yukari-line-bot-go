package controllers

import (
	m "../models"
	l "../libs"
)

type HP struct {
	LatestTitle    string
	LatestURL      string
	OneBeforeTitle string
	OneBeforeURL   string
}

type FC struct {
	LatestTitle    string
	LatestURL      string
	OneBeforeTitle string
	OneBeforeURL   string
}

var hp = HP{}
var fc = FC{}

func HPGetLatest() {
	hp.LatestTitle, hp.LatestURL = m.HPGetLatest()
}

func HPGetOneBefore() {
	hp.OneBeforeTitle, hp.OneBeforeURL = l.GetOneBefore("hp")
}

func FCGetLatest() {
	fc.LatestTitle, fc.LatestURL = m.FCGetLatest()
}

func FCGetOneBefore() {
	fc.OneBeforeTitle, fc.OneBeforeURL = l.GetOneBefore("fc")
}

func InsertToDB(tableName string, title string, url string) {
	l.InsertToDB(tableName, title, url)
}

func CheckDiff() (bool, bool){
	HPGetLatest()
	HPGetOneBefore()
	FCGetLatest()
	FCGetOneBefore()

	var isHPDiff bool
	var isFCDiff bool

	if hp.LatestTitle != hp.OneBeforeTitle {
		InsertToDB("hp", hp.LatestTitle, hp.LatestURL)
		isHPDiff = true
	}

	if fc.LatestTitle != fc.OneBeforeTitle {
		InsertToDB("fc", fc.LatestTitle, fc.LatestURL)
		isFCDiff = true
	}

	return isHPDiff, isFCDiff
}
