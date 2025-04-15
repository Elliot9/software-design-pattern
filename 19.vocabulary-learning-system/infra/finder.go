package infra

import (
	"encoding/json"
	"github/elliot9/class19/word"
	"os"
)

type LocalJsonWordFinder struct {
	filePath string
	words    map[string]wordData
}

type wordData struct {
	Description string            `json:"description"`
	Definitions map[string]string `json:"definitions"`
}

func NewLocalJsonWordFinder(filePath string) *LocalJsonWordFinder {
	finder := &LocalJsonWordFinder{
		filePath: filePath,
		words:    make(map[string]wordData),
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(data, &finder.words)
	if err != nil {
		return nil
	}

	return finder
}

func (f *LocalJsonWordFinder) Find(name string) (word.Word, bool) {
	data, found := f.words[name]
	if !found {
		return word.Word{}, false
	}

	definitions := make([]word.Definition, 0)
	for pos, def := range data.Definitions {
		definitions = append(definitions, word.Definition{
			Pos:         word.PoS(pos),
			Explanation: def,
		})
	}

	return word.Word{
		Name:        name,
		Description: data.Description,
		Definitions: definitions,
	}, true
}

var _ word.Finder = &LocalJsonWordFinder{}
