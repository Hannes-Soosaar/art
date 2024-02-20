package art

import (
	"fmt"

	"gitea.kood.tech/hannessoosaar/art/constants"
	"gitea.kood.tech/hannessoosaar/art/internal"
	"gitea.kood.tech/hannessoosaar/art/pkg/models"
	"gitea.kood.tech/hannessoosaar/art/pkg/utils"
)

func InitializeAndRun(inputArgs []string) {
	if len(inputArgs) == 2 && inputArgs[1] == "-h" {
		fmt.Println(constants.HELP)
	} else if len(inputArgs) == 2 {
		internal.DisplayDecodedSingle(utils.DecodeInput(inputArgs[1]))
	} else if len(inputArgs) == 3 && inputArgs[1] == "-e" {
		internal.DisplayEncodedSingle(utils.EncodeInput(inputArgs[2]))
	} else if len(inputArgs) == 3 && inputArgs[1] == "-m" {
		internal.DisplayDecodedFile(utils.DecodeFile(inputArgs[2]))
	} else if len(inputArgs) == 4 && inputArgs[1] == "-m" && inputArgs[2] == "-e" {
		internal.DisplayEncodedFile(utils.EncodeFile(inputArgs[3]))
	} else {
		fmt.Println(constants.HELP)
	}
}

func RunWebArtSingleLine(inputArgs []string) (responseData models.ResponseBody) {
	if len(inputArgs) == 2 && inputArgs[0] == "-d" {
		responseData.DecodedText = utils.DecodeInput(inputArgs[1])
		if responseData.DecodedText != "" {
			responseData.Status = "OK"
		} else {
			responseData.Status = "NOK"
		}
		return responseData
	} else if len(inputArgs) == 2 && inputArgs[0] == "-e" {
		responseData.EncodedText = utils.EncodeInput(inputArgs[1])
		if responseData.EncodedText != "" {
			responseData.Status = "OK"
		} else {
			responseData.Status = "NOK"
		}
		return responseData
	}
	responseData.Status = "Wrong Case"
	return responseData
}

func RunWebArtMultiLine(option string, fileContent []string) (responseData models.ResponseBody) {
	fmt.Printf("Running router in : %s mode  and \n", option)
	if fileContent == nil {
		responseData.Status = "No file added"
		return responseData
	}
	if option == "" {
		responseData.Status = "No operation selected"
		return responseData
	}
	if option == "-d" {
		responseData.Status = "OK Decode"
		responseData.FileContent = utils.DecodeFileContent(fileContent)
	} else if option == "-e" {
		responseData.Status = "OK Encode"
		responseData.FileContent = utils.EncodeFileContent(fileContent)
	} else {
		responseData.Status = "NOK"
	}
	return responseData
}
