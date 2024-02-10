package main

import (
	"fmt"
	"os"

	"gitea.kood.tech/hannessoosaar/art/constants"
	"gitea.kood.tech/hannessoosaar/art/internal"
	"gitea.kood.tech/hannessoosaar/art/pkg/utils"
)

func main() {
	initializeAndRun(os.Args)
}

func initializeAndRun(inputArgs []string) {
	if len(inputArgs) == 2 && inputArgs[1] == "-h" {
		fmt.Println(constants.HELP)
	} else if len(inputArgs) == 2 {
		internal.DisplayDecodedSingle(utils.DecodeInput(inputArgs[1]))
	} else if len(inputArgs) == 3 && inputArgs[1] == "-e" {
		internal.DisplayEncodedSingle(utils.EncodeInput(inputArgs[2]))
	} else if len(inputArgs) == 3 && inputArgs[1] == "-m" {
		internal.DisplayDecodedFile(utils.DecodeFile(inputArgs[2]))
	} else if len(inputArgs) == 4 && inputArgs[1] == "-m" && inputArgs[2] == "-e" {
		internal.DisplayEncodedFile(utils.EncodeFile(inputArgs[3]))
	} else {
		//! why do I need this?
		inputString := ""
		for i, args := range inputArgs {
			if i > 0 {
				inputString += args + " "
			}
		}
		utils.DecodeInput(inputString)
	}
}
