package model

type Commit struct {
	Object
	Tree      string `json:"tree"`
	Parent    string `json:"parent"`
	Message   string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}
