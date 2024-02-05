package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/art/internal"
)

func DecodeInput(input string) {
	if internal.ValidateInput(input) {
		fmt.Println(internal.ValidateInput(input))
	}
}

// read in a file and Decode it
func DecodeFile(path string) {
	// get content
	// validate content
	// parse content
	// output content
}
