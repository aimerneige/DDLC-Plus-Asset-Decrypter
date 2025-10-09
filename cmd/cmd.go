// (c) 2021 AimerNeige
// This code is licensed under MIT license (see LICENSE for details)

package cmd

import (
	"fmt"
	"os"

	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/core"
	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/logs"
	"github.com/AimerNeige/DDLC-Plus-Asset-Decrypter/utils"
	"github.com/urfave/cli"
)

// Execute Cli execute
func Execute() {
	app := &cli.App{
		Name:      "ddlcpad",
		Usage:     "Doki Doki Literature Club Plus Asset Decrypter",
		UsageText: "ddlcpad [--encrypt] file/directory",
		Version:   "1.3.0",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "encrypt",
				Usage: "Encrypt the specified file or directory",
			},
		},
		Action: func(c *cli.Context) error {

			// if wrong args, exit
			if c.NArg() != 1 {
				logs.ErrorLog("Wrong args, please read the help info bellow.")
				cli.ShowAppHelp(c)
				return nil
			}

			// right args, continue
			logs.PrintAppInformation()

			// read args and output
			fileIn := c.Args().Get(0)
			logs.InfoLog(fmt.Sprintf("Input File: \"%s\".", fileIn))

			// check if file exist
			if !utils.CheckFileExist(fileIn) {
				logs.ErrorLog(fmt.Sprintf("File \"%s\" does not exist!", fileIn))
				logs.DecryptFailMsg()
				return nil
			}

			// check if a directory
			if utils.CheckFileIsDir(fileIn) {
				if c.Bool("encrypt") {
					logs.ErrorLog("Encrypting directories is not supported.")
					return nil
				}
				logs.InfoLog(fmt.Sprintf("Directory \"%s\" detect, try to decrypt all file in it.", fileIn))
				core.DecryptDirectly(fileIn)
				logs.DecryptDirectlySuccessfulMsg()
				return nil
			}
			// file exist and not a folder

			if c.Bool("encrypt") {
				// try to encrypt a single file

				logs.InfoLog(fmt.Sprintf("File \"%s\" detect, try to encrypt it.", fileIn))

				// try to create a *_encrypted.cy file
				fileOut := fmt.Sprintf("%s_encrypted.cy", fileIn)
				logs.InfoLog(fmt.Sprintf("Output File: \"%s\".", fileOut))

				// if a *_encrypted.cy exist, do nothing and exit (in case destroy some useful file)
				if utils.CheckFileExist(fileOut) {
					logs.ErrorLog(fmt.Sprintf("File (or folder) \"%s\" has exist.", fileOut))
					logs.EncryptFailMsg()
					return nil
				}

				// ready to decrypt file
				if core.Encrypt(fileIn, fileOut) {
					logs.EncryptSuccessfulMsg()
					return nil
				}

				// decrypt fail
				logs.EncryptFailMsg()
				return nil
			}

			// decrypt logic
			logs.InfoLog(fmt.Sprintf("File \"%s\" detected, trying to decrypt it.", fileIn))
			fileOut := fmt.Sprintf("%s.out", fileIn)
			logs.InfoLog(fmt.Sprintf("Output File: \"%s\".", fileOut))

			// if a *.out file exists, do nothing and exit
			if utils.CheckFileExist(fileOut) {
				logs.ErrorLog(fmt.Sprintf("File \"%s\" already exists.", fileOut))
				logs.DecryptFailMsg()
				return nil
			}

			// ready to decrypt file
			if core.Decrypt(fileIn, fileOut) {
				logs.DecryptSuccessfulMsg()
				return nil
			}

			// decrypt fail
			logs.DecryptFailMsg()
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logs.ErrorLog(err.Error())
	}
}
