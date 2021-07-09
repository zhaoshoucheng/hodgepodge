package file

import (
	"fmt"
	"testing"
)


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

func InTod(start1, end1, start2, end2 string) bool {
	if start2 == end2 {
		return true
	}
	if start1 == start2 {
		return true
	}
	if start1 <= end1 && start2 < end2 {
		if end1 <= start2 || start1 >= end2 {
			return false
		} else {
			return true
		}
	} else if start1 <= end1 && start2 > end2 {
		if end1 <= start2 && start1 >= end2 {
			return false
		} else {
			return true
		}
	} else if start1 >= end1 && start2 < end2 {
		if end1 <= start2 && start1 >= end2 {
			return false
		} else {
			return true
		}
	} else if start1 >= end1 && start2 > end2 {
		return true
	}
	// now way
	return false
}

func TestGetResourceFromJson2(t *testing.T) {
	start1 := "00:00:00"
	end1 := "24:00:00"
	start2 := "18:00:00"
	end2 := "00:00:00"
	t.Log(InTod(start1,end1,start2,end2))

	ddate := "2021-07-1"
	ttime := "17:00:01"
	t.Log(DateTimeVersion(ddate, ttime))

}

func DateTimeVersion(ddate, ttime string) string {
	var y, m, d, h, mm, s int
	fmt.Sscanf(ddate, "%d-%d-%d", &y, &m, &d)
	fmt.Sscanf(ttime, "%d:%d:%d", &h, &mm, &s)
	version := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", y, m, d, h, mm, s)
	return version
}
