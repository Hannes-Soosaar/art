package errors

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/art/constants"
)

func ErrNotNum() {
	fmt.Println(constants.ERR_ARG1_NOT_NUM)
}
func ErrNoSecondArg() {
	fmt.Println(constants.ERR_ARG_NO_SECOND_ARG)
}
func ErrNoSpace() {
	fmt.Println(constants.ERR_ARG_NO_SPACE)
}
func ErrUnbalanced() {
	fmt.Println(constants.ERR_ARG_UNBALANCED)
}

func ErrReadingFile() {
	fmt.Printf(constants.ERR_READING_FILE)
}

func ErrWritingFile() {
	fmt.Println(constants.ERR_WRITING_FILE)
}
func ErrInvalidInput() {
	fmt.Println(constants.HELP)
}

func ErrOpeningFile() {
	fmt.Println(constants.ERR_OPENING_FILE)
}
func ErrEncodingCharacter() {
	fmt.Println(constants.ERR_ENCODING_CHAR)
}