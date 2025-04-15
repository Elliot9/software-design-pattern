package infra

import "github/elliot9/class19/word"

type MockWordRepository struct {
	words []word.Word
	index map[string]int
}

func NewMockWordRepository() *MockWordRepository {
	return &MockWordRepository{words: make([]word.Word, 0), index: make(map[string]int)}
}

func (r *MockWordRepository) Add(word word.Word) {
	if _, ok := r.index[word.Name]; ok {
		return
	}

	r.words = append(r.words, word)
	r.index[word.Name] = len(r.words) - 1
}

func (r *MockWordRepository) Delete(word word.Word) bool {
	if index, ok := r.index[word.Name]; ok {
		r.words = append(r.words[:index], r.words[index+1:]...)

		delete(r.index, word.Name)

		for i := index; i < len(r.words); i++ {
			r.index[r.words[i].Name] = i
		}

		return true
	}
	return false
}

func (r *MockWordRepository) Get(name string) (word.Word, bool) {
	if index, ok := r.index[name]; ok {
		return r.words[index], true
	}
	return word.Word{}, false
}

func (r *MockWordRepository) GetAll() []word.Word {
	return r.words
}

var _ word.Repository = &MockWordRepository{}
