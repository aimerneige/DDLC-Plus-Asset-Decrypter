package utils

import (
	"os"
)

// GetFileSize Get file size without open file
func GetFileSize(f string) (ret int64) {
	fi, err := os.Stat(f)
	if err != nil {
		ret = -1
	} else {
		ret = fi.Size()
	}
	return
}

// CheckFileExist Check if file exist
func CheckFileExist(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	} else {
		return false
	}
}
