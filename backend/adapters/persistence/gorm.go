package persistence

import (
	"backend/config"
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dbgorm *gorm.DB
var db *sql.DB

func openGorm() (*gorm.DB, error) {
	cfg := config.GetConfig()
	_db := postgres.Open(cfg.Db.URL)
	return gorm.Open(_db, &gorm.Config{})
}

func CloseGorm() {
	if dbgorm != nil {
		err := db.Close()
		CheckErr(err)
	}
}

func ConnectGorm() {
	_dbgorm, err := openGorm()
	CheckErr(err)
	log.Println("Db connected")
	dbgorm = _dbgorm
	db, err = dbgorm.DB()

	if err != nil {
		panic("failed to get database")
	}
}

func GetDbGorm() *gorm.DB {
	return dbgorm
}
