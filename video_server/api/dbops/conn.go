package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

// 只要调用dbops包，则该函数自动执行
func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}