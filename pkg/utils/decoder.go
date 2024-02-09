package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gitea.kood.tech/hannessoosaar/art/internal"
)

func DecodeInput(inputToDecode string) (decodedOutput string) {
	var (
		omitCounter    int
		encodedSection string
		cleanSection   string
		decodedSection string
	)
	if internal.ValidateInput(inputToDecode) {
		for i, char := range inputToDecode {
			if char == '[' {
				encodedSection = internal.ExtractEncodedSection(inputToDecode[i:])
				omitCounter = len(encodedSection)
				cleanSection = internal.RemoveFirstAndLastChar(encodedSection)
				decodedSection = decodeSection(cleanSection)
				decodedOutput += decodedSection
			}
			if omitCounter > 0 {
				omitCounter--
			} else {
				decodedOutput += string(inputToDecode[i])
			}
		}
	}
	fmt.Printf("FINAL DECODED %s \n", decodedOutput)
	return decodedOutput
}

func DecodeFile(path string) {
	// TODO add functionality
}

func decodeSection(cleanSection string) (decodedSection string) {
	sections := strings.SplitN(cleanSection, " ", 2)
	multiplier, _ := strconv.Atoi(string(sections[0]))
	compressedSection := string(sections[1])
	for i := 0; i < multiplier; i++ {
		decodedSection += compressedSection
	}
	return decodedSection
}
