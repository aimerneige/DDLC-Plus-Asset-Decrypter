// (c) 2021 AimerNeige
// This code is licensed under MIT license (see LICENSE for details)

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

// CheckFileIsDir Check if file a directory
func CheckFileIsDir(f string) bool {
	fi, err := os.Stat(f)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return true
	} else {
		return false
	}
}
