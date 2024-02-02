package main

import (
	"fmt"
	"os"

	"gitea.kood.tech/hannessoosaar/art/constants"
	"gitea.kood.tech/hannessoosaar/art/pkg/utils"
)

func main() {
	inputArgs := os.Args
	initializeAndRun(inputArgs)
}

func initializeAndRun(inputArgs []string) {
	if len(inputArgs) == 2 && inputArgs[1] == "-h" {
		fmt.Println(constants.HELP)
	} else if len(inputArgs) == 2 {
		fmt.Println("decode")
		utils.DecodeInput(inputArgs[1])
	} else if len(inputArgs) == 3 && inputArgs[1] == "-encode" {
		fmt.Println("encode")
		utils.EncodeInput(inputArgs[2])
	} else if len(inputArgs) == 3 && inputArgs[1] == "-multi" {
		//TODO add functionality to decode a file
		fmt.Println("decode a file ")
	} else if len(inputArgs) == 4 && inputArgs[1] == "-multi" && inputArgs[2] == "-encode" {
		//TODO add functionality to encode a file
		fmt.Println("encode a file")
	} else {

		//make a internal parser for this 
		inputString := ""
		for i, args := range inputArgs {
			if i > 0 {
				inputString += args + " "
			}
		}
		utils.DecodeInput(inputString)
		fmt.Println(inputString)
	}
}
