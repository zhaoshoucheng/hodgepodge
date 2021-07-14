package validate

import (
	"testing"
)

type UserInfo struct {
	Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	RePasswd string `form:"repasswd" json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd"`
	Birthday string `form:"birthday" json:"birthday" validate:"required,date_format"`
	IDCard   string `form:"idcard" json:"idcard" validate:"required,id_format"`
}
func TestDefaultGetValidParams(t *testing.T) {
	InitValidate()
	u1 := &UserInfo{
		"1111111",
		"1111111",
		"1993-10-03",
		"123456",
	}
	err := ValidParams(u1)
	if err != nil {
		t.Log(err)
	} else {
		t.Log("OK!")
	}

	u2 := &UserInfo{
		"1111111",
		"1111111",
		"1993-1003",
		"123456",
	}
	err = ValidParams(u2)
	if err != nil {
		t.Log(err)
	} else {
		t.Log("OK!")
	}

	u3 := &UserInfo{
		"1111111",
		"1111111",
		"1993-10-03",
		"1234 56",
	}
	err = ValidParams(u3)
	if err != nil {
		t.Log(err)
	} else {
		t.Log("OK!")
	}

}
