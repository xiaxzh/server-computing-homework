package entities

import (
	_ "github.com/mxk/go-sqlite/sqlite3" // as database driver of this programe

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// ORM engine
var xormEngine *xorm.Engine

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error
	// create engine
	xormEngine, err = xorm.NewEngine("sqlite3", "./agenda.db")
	checkErr(err)

	xormEngine.SetMapper(core.GonicMapper{})

	// sync the struct changes to database
	checkErr(xormEngine.Sync2(new(User)))
}
