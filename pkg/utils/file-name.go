package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MostajeranMohammad/static-file-server/internal/contracts/consts"
	"github.com/google/uuid"
)

func GenerateFileName(bucketName string, fileFormat int) string {
	extName := ""
	switch fileFormat {
	case consts.JPG:
		extName = "jpeg"

	case consts.PNG:
		extName = "png"

	case consts.WEBP:
	default:
		extName = "webp"

	}
	return fmt.Sprintf("%s-%s.%s", bucketName, uuid.New().String(), extName)
}

func GetBucketNameFromFileName(fileName string) (string, error) {
	bucketName := strings.Split(fileName, "-")[0]
	if bucketName == "" {
		return "", errors.New("wrong fileName format")
	}
	return bucketName, nil
}
