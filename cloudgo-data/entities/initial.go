package entities

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var mydb *sql.DB

// use godbc if define db
func init() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	mydb = db
}

var myOrm *xorm.Engine
// use xorm if define engine
// func init() {
// 	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
// 	checkErr(err)
// 	myOrm = engine
// }

type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DaoSource struct {
	SQLExecer
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}