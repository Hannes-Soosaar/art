package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/art/internal"
)

func DecodeInput(input string) {
	var (
		isBalanced bool
		firstArgIsNumber bool
	)
	isBalanced = internal.BracketBalance(input)
	firstArgIsNumber = internal.FirstArgIsNumber(input)
	fmt.Println(isBalanced)
	fmt.Println(firstArgIsNumber)

	fmt.Println(input)
}

func DecodeFile(path string) {
	// get content
	// validate content
	// parse content
	// output content
}
