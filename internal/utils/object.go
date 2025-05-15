package utils

import (
	"compress/zlib"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/reactiveNeon/jvc/internal/constants"
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
	dir := filepath.Join(fmt.Sprintf(".%v/objects", constants.JvcDirName), hash[:2])
	file := hash[2:]

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(dir, file))
	if err != nil {
		return err
	}
	defer f.Close()

	w := zlib.NewWriter(f)
	defer w.Close()

	_, err = w.Write(data)

	return err
}

func LoadObject(hash string) (map[string]any, error) {
	dir := filepath.Join(fmt.Sprintf(".%v/objects", constants.JvcDirName), hash[:2])
	file := hash[2:]

	f, err := os.Open(filepath.Join(dir, file))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := zlib.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}
