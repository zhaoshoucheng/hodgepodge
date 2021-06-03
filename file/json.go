package file

import (
	"encoding/json"
	"os"
)
func GetResourceFromJson(filePath string, data interface{}) error {
	filePtr, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer filePtr.Close()

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&data)
	return err
}
