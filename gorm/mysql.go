package gorm

import (
	"database/sql"
	"fmt"
	"github.com/zhaoshoucheng/hodgepodge/conf"
	"time"
)

func InitMysqlDB() (db *sql.DB, err error) {
	dbConf := conf.InitMysqlConf()
	db, err = sql.Open(dbConf.DriverName, dbConf.DataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbConf.MaxOpenConn)
	db.SetMaxIdleConns(dbConf.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(dbConf.MaxConnLifeTime) * time.Second)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ExecSql(db *sql.DB,creatSQl string) (err error) {
	stmtCity, err := db.Prepare(creatSQl)
	if err != nil {
		fmt.Println("db.Prepare err", err)
		return err
	}
	_, err = stmtCity.Exec()
	if err != nil {
		fmt.Println("stmt exec err", err)
		return err
	}
	return nil
}

