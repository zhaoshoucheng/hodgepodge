package gorm

import "testing"

func TestMyOrm_CreateInBatches(t *testing.T) {
	db, err := InitMysqlDB()
	if err != nil {
		t.Log(err)
		return
	}
	/*
	err = ExecSql(db, CreatCity)
	if err != nil {
		t.Log(err)
		return
	}
	 */

	gorm, err := InitGormByDB(db)
	if err != nil {
		t.Log(err)
		return
	}
	//myOrm := MyOrm{Gdb: gorm}
	mockList := Mock()
	gorm.Table("base_city").CreateInBatches(mockList, len(mockList))
	//myOrm.CreateInBatches("base_city",mockList)
}

func Mock() []interface{} {
	var list []interface{}
	for i := 0; i < 5; i++ {
		list = append(list, &BaseCity{
			CityID: i,
			ProvinceID: i,
			Name:"test",
		})
	}
	return list
}
