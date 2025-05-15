package utils

import (
	"crypto/sha256"
	"os"
	"path/filepath"
	"encoding/hex"
	"encoding/json"
)

func HashObject(obj any) (string, []byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", nil, err
	}

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), data, nil
}

func WriteObject(hash string, data []byte) error {
	dir := ".jvc/objects"
	os.MkdirAll(dir, 0755)

	return os.WriteFile(filepath.Join(dir, hash), data, 0644)
}

func LoadObject(hash string) (map[string]any, error) {
	data, err := os.ReadFile(filepath.Join(".jvc/objects", hash))
	if err != nil {
		return nil, err
	}

	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}
