package go_gorm

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	conf "github.com/zhaoshoucheng/hodgepodge/conf"
)

type MyOrm struct {
	Gdb *gorm.DB
}

func (orm *MyOrm)CreatTable(models interface{},) error {
	if orm.Gdb.HasTable(models) {
		return errors.New("table have exists")
	}
	return orm.Gdb.CreateTable(models).Error
}
func NewMyOrm(sqlConf *conf.MySQLConf) (myOrm *MyOrm, err error) {
	gdb, err := gorm.Open(sqlConf.DriverName, sqlConf.DataSourceName)
	if err != nil {
		return
	}
	gdb.DB().SetMaxIdleConns(sqlConf.MaxIdleConn)
	gdb.DB().SetMaxOpenConns(sqlConf.MaxOpenConn)
	return &MyOrm{Gdb: gdb}, nil
}



