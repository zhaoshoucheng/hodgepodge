package file

import (
	"fmt"
	"testing"
)

func TestGetResourceFromTxt(t *testing.T) {
	file := "/Users/zhaoshoucheng/data/src/mysrc/hodgepodge/tools/cp_file/cp_file.txt"
	lins, err := GetResourceFromTxt(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range lins {
		fmt.Println(value)
	}
}
