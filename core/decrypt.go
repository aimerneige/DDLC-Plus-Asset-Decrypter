// (c) 2021 AimerNeige
// This code is licensed under MIT license (see LICENSE for details)

package core

import (
	"fmt"
	"io/ioutil"

	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/logs"
	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/utils"
)

const decrypt_key = 40 // decrypt key BY MlgmXyysd

// Decrypt Decrypt data and create a new file.
func Decrypt(in, out string) bool {
	logs.InfoLog(fmt.Sprintf("Start read file \"%s\". Size: %d bytes.", in, utils.GetFileSize(in)))

	// try to read the whole file
	data, err := ioutil.ReadFile(in)
	if err != nil {
		logs.ErrorLog(logs.FailToReadFile(in))
		return false
	}
	logs.InfoLog(fmt.Sprintf("File \"%s\" read successful!", in))

	// check cy file
	if !cyFileCheck(data[:7]) {
		logs.ErrorLog(fmt.Sprintf("File \"%s\" does not a cy file.", in))
		return false
	}

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

// xorByte XOR all bytes input with decrypt key
func xorByte(str []byte) (ret []byte) {
	for i := 0; i < len(str); i++ {
		ret = append(ret, str[i]^decrypt_key)
	}
	return
}

// cyFileCheck Check file header and makesure the file input is a *.cy file
func cyFileCheck(header []byte) bool {
	if len(header) != 7 {
		return false
	}
	right_header := [7]byte{0x7D, 0x46, 0x41, 0x5C, 0x51, 0x6E, 0x7B}
	for i, b := range right_header {
		if b != header[i] {
			return false
		}
	}
	return true
}
