package gorm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyOrm struct {
	Gdb *gorm.DB
}
func (myOrm *MyOrm) CreateInBatches(table string, insertData []interface{}) {
	myOrm.Gdb.Table(table).CreateInBatches(insertData, len(insertData))
}

func InitGormByDB(db *sql.DB) (gdb *gorm.DB, err error) {
	gdb, err = gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return
	}
	return
}


