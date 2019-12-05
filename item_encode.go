package base64Captcha

import (
	"fmt"
	"os"
	"path/filepath"
)

func ItemWriteToFile(cap Item, outputDir, fileName, fileExt string) error {
	filePath := filepath.Join(outputDir, fileName+"."+fileExt)
	if !pathExists(outputDir) {
		_ = os.MkdirAll(outputDir, os.ModePerm)
	}
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("%s is invalid path.error:%v", filePath, err)
		return err
	}
	defer file.Close()
	_, err = cap.WriteTo(file)
	return err
}
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
