package common

import "runtime"

func GetFileName(fileName string) string {
	var fullFileName string

	if runtime.GOOS == "windows" {
		fullFileName = "c:/" + fileName
	} else {
		fullFileName = fileName
	}

	return fullFileName
}
