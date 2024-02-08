package utils

import (
	"fmt"
	"strconv"

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
				// fmt.Printf("SECTION TO BE ADDED: %s  FULL TEXT : %s \n", decodedSection, decodedOutput)
				decodedOutput += decodedSection
			}
		}
		if omitCounter > 0 {
			omitCounter--
		} else {
			decodedOutput += inputToDecode
			fmt.Println(decodedOutput) // should grow
		}

	}
	fmt.Printf("FINAL DECODED %s \n" , decodedOutput)
	return decodedOutput
}

func DecodeFile(path string) {

}

func decodeSection(cleanSection string) (decodedSection string) {
	multiplier, _ := strconv.Atoi(string(cleanSection[0]))
	compressedSection := string(cleanSection[1:])
	for i := 0; i < multiplier; i++ {
		decodedSection += compressedSection
	}
	return decodedSection
}
