package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/reactiveNeon/jvc/internal/constants"
)

func getAllHashesByType(objType string) ([]string, error) {
	var hashes []string

	err := filepath.Walk(fmt.Sprintf(".%v/objects", constants.JvcDirName), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		parts := strings.Split(path, string(os.PathSeparator))
		dir := parts[len(parts)-2]
		file := parts[len(parts)-1]

		hash := fmt.Sprintf("%s%s", dir, file)

		obj, err := LoadObject(hash)
		if err != nil {
			return nil
		}

		if obj["type"] == objType {
			hashes = append(hashes, hash)
		}

		return nil
	})

	return hashes, err
}

func GetTreeHashFromCommitHash(commitHash string) (string, error) {
	obj, err := LoadObject(commitHash)
	if err != nil {
		return "", err
	}

	treeHash, ok := obj["tree"].(string)
	if !ok {
		return "", fmt.Errorf("tree hash not found in commit object")
	}

	return treeHash, nil
}

func GetAllBlobHashes() ([]string, error) {
	return getAllHashesByType("blob")
}

func GetAllTreeHashes() ([]string, error) {
	return getAllHashesByType("tree")
}

func GetAllCommitHashes() ([]string, error) {
	return getAllHashesByType("commit")
}
