package go_gorm

import (
	"github.com/zhaoshoucheng/hodgepodge/conf"
	"testing"
)

func TestNewMyOrm(t *testing.T) {
	myOrm, err := NewMyOrm(conf.InitMysqlConf())
	if err != nil {
		t.Log(err)
		return
	}
	type NewTable struct {
		ID 		 int `json:"id" gorm:"column(id)`
		Ip       string `json:"ip" gorm:"column(ip)`
		CID      int    `json:"cid" gorm:"column(cid)`

	}
	table := &NewTable{}
	err = myOrm.CreatTable(table)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("Done !")
}
