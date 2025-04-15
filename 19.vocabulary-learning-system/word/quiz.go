package word

import (
	"sort"
	"strings"
)

const BlankRate = 0.7

type Quiz struct {
	words        []Word
	currentIndex int
	point        int
	randomizer   Randomizer
}

func NewQuiz(words []Word, randomizer Randomizer) *Quiz {
	return &Quiz{
		words:        words,
		currentIndex: 0,
		point:        0,
		randomizer:   randomizer,
	}
}

func (q *Quiz) GetCurrentIndex() int {
	return q.currentIndex
}

func (q *Quiz) GetPoint() int {
	return q.point
}

func (q *Quiz) GetTotalWords() int {
	return len(q.words)
}

func (q *Quiz) GetCurrentQuestion() (wordBlank, definition string) {
	if q.currentIndex >= len(q.words) {
		return "", ""
	}

	word := q.words[q.currentIndex]
	posOrder := map[PoS]int{
		Verb:      1,
		Adjective: 2,
		Noun:      3,
		Adverb:    4,
	}

	sortedDefinitions := make([]Definition, len(word.Definitions))
	copy(sortedDefinitions, word.Definitions)

	sort.Slice(sortedDefinitions, func(i, j int) bool {
		return posOrder[sortedDefinitions[i].Pos] < posOrder[sortedDefinitions[j].Pos]
	})

	definition = sortedDefinitions[q.randomizer.IntN(len(sortedDefinitions))].Explanation

	wordBlank = q.createWordBlank(word)
	return wordBlank, definition
}

func (q *Quiz) CheckAnswer(answer string) (bool, string) {
	word := q.words[q.currentIndex]
	result := strings.EqualFold(answer, word.Name)

	if result {
		q.point++
	}

	q.currentIndex++
	return result, word.Name
}

// 需求敘述 + 測試案例不足
func (q *Quiz) createWordBlank(word Word) string {
	name := word.Name
	length := len(name)

	// 長度小於等於2不挖空
	if length <= 2 {
		return name
	}

	blankCount := int(float64(length) * BlankRate)
	if blankCount > length-2 {
		blankCount = length - 2 // 確保不挖頭尾
	}

	availablePositions := make([]int, 0, length-2)
	for i := 1; i < length-1; i++ {
		availablePositions = append(availablePositions, i)
	}

	q.randomizer.Shuffle(len(availablePositions), func(i, j int) {
		availablePositions[i], availablePositions[j] = availablePositions[j], availablePositions[i]
	})

	blanks := make(map[int]bool)
	for i := 0; i < blankCount; i++ {
		blanks[availablePositions[i]] = true
	}

	var result strings.Builder
	for i, char := range name {
		if blanks[i] {
			result.WriteRune('_')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
