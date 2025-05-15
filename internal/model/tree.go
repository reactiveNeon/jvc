package model

type TreeEntryStructure string

const (
	TreeEntryStructureArray  TreeEntryStructure = "array"
	TreeEntryStructureObject TreeEntryStructure = "object"
)

type TreeEntry struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

type Tree struct {
	Object
	Structure TreeEntryStructure `json:"structure"`
	Entries   []TreeEntry        `json:"entries"`
}
