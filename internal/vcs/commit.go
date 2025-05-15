package vcs

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/reactiveNeon/jvc/internal/model"
	"github.com/reactiveNeon/jvc/internal/utils"
)

func StoreJson(value any) (string, error) {
	switch v := value.(type) {
	case string, float64, bool, nil:
		val, err := json.Marshal(v)
		if err != nil {
			return "", err
		}

		blob := model.Blob{
			Object: model.Object{
				Type: "blob",
				Size: len(val), // get the byte size of the string
			},
			Value: v,
		}

		hash, data, err := utils.HashObject(blob)
		if err != nil {
			return "", err
		}

		if err := utils.WriteObject(hash, data); err != nil {
			return "", err
		}

		return hash, nil

	case map[string]any:
		var entries []model.TreeEntry
		keys := make([]string, 0, len(v))

		for key := range v {
			keys = append(keys, key)
		}

		sort.Strings(keys) // sort keys for consistent ordering/hashing

		for _, key := range keys {
			hash, err := StoreJson(v[key])
			if err != nil {
				return "", err
			}
			entries = append(entries, model.TreeEntry{
				Key:  key,
				Hash: hash,
			})
		}

		tree := model.Tree{
			Object: model.Object{
				Type: "tree",
				Size: len(entries),
			},
			Structure: model.TreeEntryStructureObject,
			Entries:   entries,
		}

		hash, data, err := utils.HashObject(tree)
		if err != nil {
			return "", err
		}

		if err := utils.WriteObject(hash, data); err != nil {
			return "", err
		}

		return hash, nil

	case []any:
		var entries []model.TreeEntry
		for i, item := range v {
			hash, err := StoreJson(item)
			if err != nil {
				return "", err
			}
			entries = append(entries, model.TreeEntry{
				Key:  strconv.Itoa(i),
				Hash: hash,
			})
		}

		tree := model.Tree{
			Object: model.Object{
				Type: "tree",
				Size: len(entries),
			},
			Structure: model.TreeEntryStructureArray,
			Entries:   entries,
		}

		hash, data, err := utils.HashObject(tree)
		if err != nil {
			return "", err
		}

		if err := utils.WriteObject(hash, data); err != nil {
			return "", err
		}

		return hash, nil

	default:
		return "", fmt.Errorf("unsupported type: %T", v)
	}
}

func StoreCommit(treeHash, parentHash, message string) (string, error) {
	timestamp := utils.GetCurrentTimestamp()

	commit := model.Commit{
		Object: model.Object{
			Type: "commit",
			Size: len(message),
		},
		Tree:      treeHash,
		Parent:    parentHash,
		Message:   message,
		Timestamp: timestamp,
	}

	hash, data, err := utils.HashObject(commit)
	if err != nil {
		return "", err
	}

	if err := utils.WriteObject(hash, data); err != nil {
		return "", err
	}

	return hash, nil
}
