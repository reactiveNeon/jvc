package model

// Blob represents the primitive data type, i.e, string, number (float64), bool, null (nil)
type Blob struct {
	Object
	Value any    `json:"value"`
}
