package file

import (
	"bufio"
	"os"
)

func GetResourceFromTxt(fileName string) (resp []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		resp = append(resp, lineText)
	}
	return
}
