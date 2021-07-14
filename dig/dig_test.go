package dig

import "testing"

func TestGetNewStruct3(t *testing.T) {
	/*
	dig 用map管理所有对象，并发不安全，尽量在单线程(协程)中使用，使用时注意
	注册Provide 已经存在的对象会报错
	*/
	s3,err := GetNewStruct3()
	if err != nil{
		t.Log(err)
	} else {
		t.Log(s3.Name)
	}

	c, err := BuildContainer()
	if err != nil {
		t.Log("new BuildContainer err")
		return
	}
	t.Log(c.String())
}
