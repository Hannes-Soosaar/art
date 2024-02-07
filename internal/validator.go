package internal

import (
	"fmt"
	"unicode"

	"gitea.kood.tech/hannessoosaar/art/errors"
)

// Takes in a string and analyzes it to to see if square [] are balanced
func bracketBalance(inputToAnalyze string) (isBalanced bool) {
	var (
		rightBracketCount int = 0
		leftBracketCount  int = 0
	)
	for _, char := range inputToAnalyze {
		if char == '[' {
			rightBracketCount++
		}
		if char == ']' {
			leftBracketCount++
		}
		if rightBracketCount < leftBracketCount {
			return false
		}
	}
	if rightBracketCount == leftBracketCount {
		return true
	} else {
		return false
	}
}

// checks if the number after the "[" is a unicode number
func firstArgIsNumber(encodedSection string) (firstArgIsNumber bool) {
	return unicode.IsDigit(rune(encodedSection[1]))
}

// Checks to see if the encoded section designated for data that is compressed is not blank
func hasSecondArg(encodedSection string) (hasSecondArg bool) {
	thirdArg := rune(encodedSection[3])
	fmt.Println(thirdArg)
	return thirdArg != ' '
}

// Checks to see if the second char is a space
func separatedBySpace(encodedSection string) (separatedBySpace bool) {
	if encodedSection[2] == ' ' {
		return true
	} else {
		return false
	}
}
func ValidateInput(inputToAnalyze string) (validInput bool) {
	var (
		startIndex int = 0
		endIndex   int = 0
	)
	if bracketBalance(inputToAnalyze) { // make sure there is an exit here
		for i, char := range inputToAnalyze {
			if char == '[' {
				startIndex = i
				endIndex = 0
			}
			if char == ']' {
				endIndex = i
			}
			if endIndex > 0 {
				encodedSection := getEncodedSection(inputToAnalyze[startIndex : endIndex+1])
				// if len(encodedSection) < 5 {
				// 	errors.ErrInvalidInput() //TODO make a dedicated error
				// 	return false
				// } else
				if !firstArgIsNumber(encodedSection) {
					errors.ErrNotNum()
					return false
				} else if !separatedBySpace(encodedSection) {
					errors.ErrNoSpace()
					return false
				} else if !hasSecondArg(encodedSection) {
					errors.ErrNoSecondArg()
					return false
				}
			}
		}
	} else {
		errors.ErrUnbalanced()
		return false
	}
	return true
}

// to make things more readable will return only a section of the string.
func getEncodedSection(sectionOfInputToAnalyze string) (encodedSection string) {
	encodedSection = sectionOfInputToAnalyze
	return encodedSection
}

// this branch must be pushed to origin