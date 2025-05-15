package vcs

import (
	"fmt"

	"github.com/reactiveNeon/jvc/internal/utils"
)

func CheckoutJson(hash string) (any, error) {
	obj, err := utils.LoadObject(hash)
	if err != nil {
		return nil, err
	}

	switch obj["type"] {
	case "blob":
		return obj["value"], nil

	case "tree":
		structure := obj["structure"].(string)
		entries := obj["entries"].([]any)

		if structure == "array" {
			var result []any
			for _, e := range entries {
				entry := e.(map[string]any)
				// idx := int(entry["key"].(string))
				hash := entry["hash"].(string)

				value, err := CheckoutJson(hash)
				if err != nil {
					return nil, err
				}

				result = append(result, value)
			}
			return result, nil
		} else if structure == "object" {
			result := make(map[string]any)
			for _, e := range entries {
				entry := e.(map[string]any)
				key := entry["key"].(string)
				hash := entry["hash"].(string)
				value, err := CheckoutJson(hash)
				if err != nil {
					return nil, err
				}
				result[key] = value
			}
			return result, nil
		} else {
			return nil, fmt.Errorf("unsupported tree structure: %s", obj["structure"])
		}

	default:
		return nil, fmt.Errorf("unsupported object type: %s", obj["type"])
	}
}
