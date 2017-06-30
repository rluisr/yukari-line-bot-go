package libs

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectSql() (*sql.DB) {
	var DBFILE string = "./yukari.db"

	db, err := sql.Open("sqlite3", DBFILE)
	checkErr(err)

	return db
}

func GetOneBefore(tableName string) (string, string) {
	db := ConnectSql()

	rows, err := db.Query("SELECT title, url FROM " + tableName)
	checkErr(err)

	db.Close()

	var title string
	var url string

	// 本当は最後のレコードだけ欲しいけど、全部ループすることで結局は最後の値が返ってくるジレンマ
	for rows.Next() {
		err = rows.Scan(&title, &url)
		checkErr(err)
	}

	return title, url
}

func InsertToDB(tableName string, title string, url string) {
	db := ConnectSql()

	_, err := db.Exec(
		"INSERT INTO "+tableName+" (title, url) VALUES(?, ?)",
		title,
		url,
	)
	checkErr(err)

	db.Close()
}

func InsertUserToDB(tableName string, userId string) {
	db := ConnectSql()

	_, err := db.Exec(
		"INSERT INTO "+tableName+" (user_id) VALUES(?)",
		userId,
	)
	checkErr(err)

	db.Close()
}

func DeleteUserFromDB(userId string) {
	db := ConnectSql()

	stmt, err := db.Prepare("DELETE FROM Users WHERE user_id = ?")
	checkErr(err)

	res, err := stmt.Exec(userId)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	db.Close()
}

func GetUsers() ([]string){
	db := ConnectSql()

	rows, err := db.Query("SELECT user_id FROM Users")
	checkErr(err)

	db.Close()

	var userId string
	var arrayUserId []string

	for rows.Next() {
		err = rows.Scan(&userId)
		checkErr(err)
		arrayUserId = append(arrayUserId, userId)
	}

	return arrayUserId
}
