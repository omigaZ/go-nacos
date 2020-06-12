package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)
var SqlDB *xorm.Engine

func init()  {
	var err error
	SqlDB, err = xorm.NewEngine("mysql", "root:yunhuan29@/myadx?charset=utf8")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
