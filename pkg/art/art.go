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
	if len(inputArgs) == 1 && inputArgs[0] == "-h" {
		responseData.Status = constants.HELP
		return responseData
	} else if len(inputArgs) == 2 && inputArgs[0] == "-d" {
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

func RunWebArtMultiLine(inputArgs []string) (singleLine string, MultiLine []string) {
	// TODO Modify function to return a st
	if len(inputArgs) == 1 && inputArgs[0] == "-h" {
		return constants.HELP, nil
	} else if len(inputArgs) == 1 {
		return utils.DecodeInput(inputArgs[0]), nil
	} else if len(inputArgs) == 2 && inputArgs[0] == "-e" {
		return utils.EncodeInput(inputArgs[1]), nil
	} else if len(inputArgs) == 3 && inputArgs[1] == "-m" {
		return "", utils.DecodeFile(inputArgs[2])
	} else if len(inputArgs) == 4 && inputArgs[1] == "-m" && inputArgs[2] == "-e" {
		return "", utils.DecodeFile(inputArgs[2])
	} else {
		return constants.HELP, nil
	}
}
