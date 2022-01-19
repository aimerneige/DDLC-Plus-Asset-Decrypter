// (c) 2021 AimerNeige
// This code is licensed under MIT license (see LICENSE for details)

package logs

import (
	"fmt"
	"time"

	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/utils"
)

// DecryptFailMsg Print app and author info.
func PrintAppInformation() {
	fmt.Print("\033[32m***********************************************************\033[0m\n")
	fmt.Print("\033[32m*      Doki Doki Literature Club Plus Asset Decrypter     *\033[0m\n")
	fmt.Print("\033[32m* Author: AimerNeige                                      *\033[0m\n")
	fmt.Print("\033[32m* Github: github.com/AimerNeige/DDLC-Plus-Asset-Decrypter *\033[0m\n")
	fmt.Print("\033[32m* Refer: github.com/MlgmXyysd/DDLC-Plus-Asset-Decrypter   *\033[0m\n")
	fmt.Print("\033[32m* Email: aimer.neige.soft@gmail.com                       *\033[0m\n")
	fmt.Print("\033[32m***********************************************************\033[0m\n")
}

// DecryptFailMsg Print help message when decrypt success.
func DecryptSuccessfulMsg() {
	fmt.Print("\033[32m***********************************************************\033[0m\n")
	fmt.Print("\033[32m*                  Decrypted Successful!                  *\033[0m\n")
	fmt.Print("\033[32m* Congratulations! You have successfully decrypted the cy *\033[0m\n")
	fmt.Print("\033[32m* file! Now you can use AssetStudio to get all the assert *\033[0m\n")
	fmt.Print("\033[32m* files. Here to download AssertStudio:                   *\033[0m\n")
	fmt.Print("\033[32m* <https://github.com/Perfare/AssetStudio>                *\033[0m\n")
	fmt.Print("\033[32m***********************************************************\033[0m\n")
}

// DecryptDirectlySuccessfulMsg Print help message when decrypt directly success.
func DecryptDirectlySuccessfulMsg() {
	fmt.Print("\033[32m*****************************************************\033[0m\n")
	fmt.Print("\033[32m*               Decrypted Successful!               *\033[0m\n")
	fmt.Print("\033[32m* Decryption has been completed.                    *\033[0m\n")
	fmt.Print("\033[32m* There may be something wrong, all error has been  *\033[0m\n")
	fmt.Print("\033[32m* skiped, you can check the log for more detail.    *\033[0m\n")
	fmt.Print("\033[32m* Now you can use AssetStudio to get all the assert *\033[0m\n")
	fmt.Print("\033[32m* files. Here to download AssertStudio:             *\033[0m\n")
	fmt.Print("\033[32m* <https://github.com/Perfare/AssetStudio>          *\033[0m\n")
	fmt.Print("\033[32m*****************************************************\033[0m\n")
}

// DecryptFailMsg Print help message when decrypt fail.
func DecryptFailMsg() {
	fmt.Print("\033[31m***********************************************************\033[0m\n")
	fmt.Print("\033[31m*                     Decrypted Fail!                     *\033[0m\n")
	fmt.Print("\033[31m* Opus, there is something wrong. You can check the error *\033[0m\n")
	fmt.Print("\033[31m* log and try again.                                      *\033[0m\n")
	fmt.Print("\033[31m* If you still have problem, post a issue here:           *\033[0m\n")
	fmt.Print("\033[31m* github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/issue   *\033[0m\n")
	fmt.Print("\033[31m***********************************************************\033[0m\n")
}

// ErrorLog Print error log with color and time.
func ErrorLog(msg string) {
	fmt.Print("\033[31m[E]\033[0m")
	fmt.Printf("\033[34m %s \033[0m", getFormattedTime())
	fmt.Printf("\033[35m%s\n\033[0m", msg)
}

// InfoLog Print info log with color and time.
func InfoLog(msg string) {
	fmt.Print("\033[32m[I]\033[0m")
	fmt.Printf("\033[34m %s \033[0m", getFormattedTime())
	fmt.Printf("\033[35m%s\n\033[0m", msg)
}

// FailToReadFile Print fail to read file info.
func IsReadingFile(fileName string) string {
	return fmt.Sprintf("ddlcpad is reading file \"%s\". Size: %d.", fileName, utils.GetFileSize(fileName))
}

// FailToWriteFile Print fail to write file info.
func IsWritingFile(fileName string) string {
	return fmt.Sprintf("ddlcpad is writing file \"%s\".", fileName)
}

// FailToReadFile Print fail to read file info.
func FailToReadFile(fileName string) string {
	return fmt.Sprintf("Fail to read file %s.", fileName)
}

// FailToWriteFile Print fail to write file info.
func FailToWriteFile(fileName string) string {
	return fmt.Sprintf("Fail to write file %s.", fileName)
}

// getFormattedTime Get Formatted time stirng.
func getFormattedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
