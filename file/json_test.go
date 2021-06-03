package file

import "testing"


//复杂的json结构体
type Data struct {
	Name   string `json:"name"`
	Cities map[string]struct {
		Name string `json:"name"`
		Dis  map[string]struct {
			Name string `json:"name"`
		} `json:"districts"`
	} `json:"Cities"`
}

func TestGetResourceFromJson(t *testing.T) {
	filePath := "./test.json"
	data := make(map[string]*Data)
	err := GetResourceFromJson(filePath, &data)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(data["110000"])
	t.Log(data["120000"])
}
