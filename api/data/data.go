package data

import (
	"github.com/HouzuoGuo/tiedot/db"
)

var Db *db.DB

func InitDb(dir string) (err error) {
	// (Create if not exist) open a database
	if Db, err = db.OpenDB(dir); err != nil {
		return
	}
	return
}
