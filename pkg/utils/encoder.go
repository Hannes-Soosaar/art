package utils

import (
	"strconv"

	"gitea.kood.tech/hannessoosaar/art/constants"
	"gitea.kood.tech/hannessoosaar/art/internal"
)

func EncodeInput(inputString string) string {
	if internal.ValidateEncodingInput(inputString) {
		numOfLet := 1 //num of chars
		encodedStr := inputString
		for numOfLet <= 3 {
			encodedStr = SingleLineEncoder(encodedStr, numOfLet)
			numOfLet++
		}
		return encodedStr
	}
	return ""
}

func SingleLineEncoder(inputString string, numOfLet int) string {
	finalStr := ""
	var checkLastInd int
	for i := 0; i < len(inputString)-numOfLet; i++ {
		subStr := inputString[i : i+numOfLet]
		lastInd, numOfReps, pattern := GetRepPatterns(inputString, subStr, i)
		if numOfReps == 1 {
			finalStr += string(inputString[i])
			lastInd = i
		} else {
			finalStr += "[" + strconv.Itoa(numOfReps) + " " + string(pattern) + "]"
			i = lastInd - 1
		}
		checkLastInd = lastInd
	}
	if len(inputString)-checkLastInd == 1 {
		finalStr += string(inputString[checkLastInd])
	} else if checkLastInd < len(inputString) {
		finalStr += inputString[checkLastInd+1:]
	}
	return finalStr
}

// get repetitive series of substring in a string
func GetRepPatterns(inputString, subStr string, startInd int) (int, int, string) {
	pattern := subStr
	numOfReps := 0
	for len(pattern)+startInd <= len(inputString) && inputString[startInd:startInd+len(pattern)] == pattern {
		numOfReps++
		pattern += subStr
	}
	if numOfReps > 1 {
		startInd = startInd + (len(subStr) * numOfReps)
	}
	return startInd, numOfReps, subStr
}

// get repetitive series of substring in a string
func EncodeFile(filePath string) (encodedFile []string) {
	fileContent := internal.GetFileContent(internal.OpenFile(filePath))
	for _, content := range fileContent {
		encodedOutput := EncodeInput(content)
		encodedFile = append(encodedFile, encodedOutput+"\n")
	}
	if constants.WRITE_TO_FILE {
	internal.CreateEncodedFile(encodedFile, filePath)
	}
	return encodedFile
}
