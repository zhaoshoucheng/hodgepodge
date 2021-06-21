package file

import "github.com/360EntSecGroup-Skylar/excelize"

func GetResourceFromXlsx(path string,sheet string,opt func(row []string) (interface{}, error)) ([]interface{},error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	rows := f.GetRows(sheet)
	var res []interface{}
	for _, row := range rows {
		data, err := opt(row)
		if err != nil {
			continue
		}
		res = append(res, data)
	}
	return res, nil
}
