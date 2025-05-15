package utils

import (
	"fmt"
	"os"

	"github.com/reactiveNeon/jvc/internal/constants"
)

func WriteHead(commitHash string) error {
	return os.WriteFile(fmt.Sprintf(".%v/HEAD", constants.JvcDirName), []byte(commitHash), 0644)
}

func ReadHead() (string, error) {
	data, err := os.ReadFile(fmt.Sprintf(".%v/HEAD", constants.JvcDirName))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

