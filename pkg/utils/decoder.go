package utils

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/art/internal"
)

func DecodeInput(input string) (decodedOutput string) {
	decodedOutput = ""
	if internal.ValidateInput(input) {
		for i, char := range input {
			if char == '[' {
				multiplier := int(input[i+1])
				character := input[i+3]
				for j := 0; j < multiplier; j++ {
					decodedOutput += string(character)
				}
			}
			decodedOutput += string(char)
		}
		fmt.Printf("The decoded picture: %s \n", decodedOutput)
		return decodedOutput
	}
	return ""
}

// read in a file and Decode it
func DecodeFile(path string) {
	// get content
	// validate content
	// parse content
	// output content
}
