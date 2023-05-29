package utils

import (
	"errors"
	"strconv"
)

func ValidateAllParamsAreIntString(input []string) error {
	Err := ""
	for _, v := range input {
		_, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			Err += err.Error() + "/n"
		}
	}
	if Err != "" {
		return errors.New(Err)
	} else {
		return nil
	}
}

func ConvertStringArrayToIntArray(input []string) []int32 {
	output := []int32{}
	for _, v := range input {
		convertedV, _ := strconv.ParseInt(v, 10, 32)
		output = append(output, int32(convertedV))
	}
	return output
}
