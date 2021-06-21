package file

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestGetResourceFromXlsx(t *testing.T) {
	data := "./city.xlsx"
	type City struct {
		ID int
		Area string
		CityName string
	}
	resArr, err := GetResourceFromXlsx(data,"Sheet1", func(row []string) (interface{}, error) {
		if len(row) < 3 {
			return nil, errors.New("row len lack")
		}
		id, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, errors.New(fmt.Sprintln("atoi err", err))
		}
		return &City{ID: id, Area: row[0],CityName: row[1]}, nil
	})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resArr[0])
}
