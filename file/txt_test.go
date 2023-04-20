package file

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetResourceFromTxt(t *testing.T) {
	file := "./write.txt"
	lins, err := GetResourceFromTxt(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range lins {
		fmt.Println(value)
	}
}
func TestWriteToTxt(t *testing.T) {
	for i := 0; i < 5; i++ {
		err := WriteToTxt("./write.txt", strconv.Itoa(i)+"\n")
		if err != nil {
			panic(err)
		}
	}
}
