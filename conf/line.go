package conf

import "os"

var CHANNEL_SECRET string = os.Getenv("LINE_CHANNEL_SECRET")
var CHANNEL_TOKEN string = os.Getenv("LINE_CHANNEL_TOKEN")
var FOLLOW_MSG string = "登録ありがとうございます 􀄃􀆰3 hearts􏿿\n\n・田村ゆかり公式サイト\n・ファンクラブサイト\nの更新をお知らせ致します􀐂􀄝light bulb􏿿\n\nなお、当アカウントをブロックすることで利用の停止ができます 􀄃􀆐content􏿿"

// arg: HP or FC {String}
func CreateUpdateMsg(arg string, title string, url string) (string) {
	var UPDATE_MSG string

	if arg == "HP" {
		UPDATE_MSG = "【BOTよりお知らせ】\r\n公式サイトが更新されました！\r\n\r\n【タイトル】\r\n" + title + "\r\n\r\n詳しくはこちらへ！\r\n" + url
	}
	if arg == "FC" {
		UPDATE_MSG = "【BOTよりお知らせ】\r\nFCサイトが更新されました！\r\n\r\n【タイトル】\r\n" + title + "\r\n\r\n詳しくはこちらへ！\r\n" + url
	}

	return UPDATE_MSG
}
