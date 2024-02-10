package utils

import (
	"fmt"
	"strconv"

	"gitea.kood.tech/hannessoosaar/art/constants"
	"gitea.kood.tech/hannessoosaar/art/errors"
	"gitea.kood.tech/hannessoosaar/art/internal"
)

// perhaps it would be good to use a setter
func EncodeInput(inputString string) (encodedString string) {
	if !internal.ValidateEncodingInput(inputString) {
		errors.ErrEncodingCharacter()
		return constants.ERR_ENCODING_CHAR
	}
	var (
		lastChar   rune
		tempString string
	)
	charCount := 1
	for _, char := range inputString {
		if char == lastChar {
			// startIndex = i
			charCount++
		} else if charCount > 3 && char != lastChar {
			tempString += "[" + strconv.Itoa(charCount) + " " + string(lastChar) + "]"
			lastChar = char
			charCount = 1
		} else {
			tempString += string(lastChar) // writes the lastChar to
			lastChar = char
		}
	}
	if charCount == 1 {
		tempString += string(lastChar)
	} else {
		tempString += "[" + strconv.Itoa(charCount) + " " + string(lastChar) + "]"
	}
	fmt.Println(tempString)
	return tempString
}

func EncodeFile(filePath string) (encodedFile []string) {

	return encodedFile
}
