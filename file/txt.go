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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		resp = append(resp, lineText)
	}
	return
}

func WriteToTxt(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 写入文件
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
