// (c) 2021 AimerNeige
// This code is licensed under MIT license (see LICENSE for details)

package core

import (
	"fmt"
	"io/ioutil"

	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/logs"
	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/utils"
)

const DECRYPT_KEY = 40 // decrypt key BY MlgmXyysd

// DecryptDirectly Decrypt data and create a new file.
func Decrypt(in, out string) bool {
	// check if file exist
	if utils.CheckFileExist(in) {
		logs.InfoLog(fmt.Sprintf("Start read file \"%s\". Size: %d", in, utils.GetFileSize(in)))
	} else {
		logs.ErrorLog("File \"%s\" does not exist!")
		return false
	}
	// try to read the whole file
	data, err := ioutil.ReadFile(in)
	if err != nil {
		logs.ErrorLog(logs.FailToReadFile(in))
		return false
	}
	logs.InfoLog(fmt.Sprintf("File \"%s\" read successful!", in))
	// start decrypt
	logs.InfoLog("Decrypt Start")
	data = xorByte(data)
	logs.InfoLog("Decrypt End")
	// write output file
	logs.InfoLog(fmt.Sprintf("Start writing file \"%s\".", out))
	err = ioutil.WriteFile(out, data, 0644)
	if err != nil {
		logs.ErrorLog(logs.FailToWriteFile(out))
		return false
	}
	logs.InfoLog(fmt.Sprintf("Write file \"%s\" successful.", out))
	return true
}

func xorByte(str []byte) (ret []byte) {
	for i := 0; i < len(str); i++ {
		ret = append(ret, str[i]^DECRYPT_KEY)
	}
	return
}
