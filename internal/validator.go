package internal

import "unicode"

// Takes in a string and analyzes it to to see if square [] are balanced
func BracketBalance(inputToAnalyze string) (isBalanced bool) {
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
	}else {
		return false 
	}
}
// checks if the number after the "[" is a unicode number
func FirstArgIsNumber(inputToAnalyze string) (firstArgIsNumber bool){
	for i,char:=range inputToAnalyze {
		if char == '['  && i+1 < len(inputToAnalyze) {
			nextChar := rune(inputToAnalyze[i+1])
			if !(unicode.IsDigit(nextChar)){
				return false
			}
		}
	}
	return true
}
