package persistence

import (
	"backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var dbx *sqlx.DB

func openx() (*sqlx.DB, error) {
	cfg := config.GetConfig()
	return sqlx.Open(cfg.Db.Driver, cfg.Db.URL)
}

func Close() {
	if dbx != nil {
		err := dbx.Close()
		CheckErr(err)
	}
}

func ConnectX() {
	_dbx, err := openx()
	CheckErr(err)
	log.Println("Db connected")
	_dbx.SetMaxIdleConns(10)
	_dbx.SetMaxOpenConns(40)
	dbx = _dbx
}

func CheckErr(err error) {
	if err != nil {
		Close()
		log.Fatalln(err.Error())
	}
}

func LogError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func GetDbx() *sqlx.DB {
	return dbx
}
