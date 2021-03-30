package logging

import (
	"log"
	"os"
)

// openLogFile 開啟log檔案
func openLogFile(prePath, filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir(prePath)
	case os.IsPermission(err):
		log.Println("Fail by Permission:", err.Error())
		panic(err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Fail by OpenFile:", err.Error())
		panic(err)
	}

	return handle
}

// mkDir 建立資料夾
func mkDir(path string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+path, os.ModePerm)
	if err != nil {
		log.Println("Fail by MkdirAll:", err.Error())
		panic(err)
	}
}
