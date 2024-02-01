package main

import (
	"fmt"
	"os"
	"gitea.kood.tech/hannessoosaar/art/pkg/utils"
	"gitea.kood.tech/hannessoosaar/art/constants"
)

func main() {
	inputArgs := os.Args
	initializeAndRun(inputArgs)
}

func initializeAndRun(inputArgs []string) {
	if len(inputArgs) == 2 && inputArgs[1] == "-h" {
		fmt.Println(constants.HELP)
	} else if len(inputArgs) == 2 {
		utils.decodeInput(inputArgs[2])
	} else if len(inputArgs) == 3 && inputArgs[2] == "-encode" {
//		utils.encodeInput(inputArgs[3])
	}
}
