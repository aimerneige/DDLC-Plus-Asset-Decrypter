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
		UsageText: "ddlcpad file",
		Version:   "1.0.0",
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 && c.NArg() < 2 {
				logs.PrintAppInformation()
				fileIn := c.Args().Get(0)
				fileOut := fmt.Sprintf("%s.out", fileIn)
				logs.InfoLog(fmt.Sprintf("Input File: \"%s\".", fileIn))
				logs.InfoLog(fmt.Sprintf("Output File: \"%s\".", fileOut))
				if utils.CheckFileExist(fileOut) {
					logs.ErrorLog(fmt.Sprintf("File \"%s\" has exist, program exit!", fileOut))
				} else {
					if core.Decrypt(fileIn, fileOut) {
						logs.DecryptSuccessfulMsg()
					} else {
						logs.DecryptFailMsg()
					}
				}
			} else {
				cli.ShowAppHelp(c)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logs.ErrorLog(err.Error())
	}
}
