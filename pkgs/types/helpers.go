package types

import (
	"os"
)

func ReadYamlFile(filePath string, data DataParser) error {
	// check if file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return data.ParseData(file)
}
