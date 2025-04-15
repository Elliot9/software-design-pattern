package scenes

import (
	"fmt"
	"github/elliot9/class19/word"
	"sort"
	"strconv"
	"strings"
)

type Messenger interface {
	Render() []string
}

type SimpleMessenger struct {
	message string
}

func NewSimpleMessenger(message string) *SimpleMessenger {
	return &SimpleMessenger{message: message}
}

func (m *SimpleMessenger) Render() []string {
	return []string{m.message}
}

type CurrentWordsMessenger struct {
	wordService *word.WordService
}

func NewCurrentWordsMessenger(wordService *word.WordService) *CurrentWordsMessenger {
	return &CurrentWordsMessenger{wordService: wordService}
}

func (m *CurrentWordsMessenger) Render() []string {
	words := m.wordService.GetAllWords()
	if len(words) == 0 {
		return []string{"You don't have any words."}
	}

	wordsString := []string{}
	for _, word := range words {
		wordsString = append(wordsString, word.Name)
	}
	return []string{fmt.Sprintf("Current Words: %s", strings.Join(wordsString, ", "))}
}

type GreetingMessenger struct {
	wordService *word.WordService
}

func NewGreetingMessenger(wordService *word.WordService) *GreetingMessenger {
	return &GreetingMessenger{wordService: wordService}
}

func (m *GreetingMessenger) Render() []string {
	words := m.wordService.GetAllWords()

	if len(words) == 0 {
		return []string{"You don't have any words."}
	}

	score, total := 0, 0

	lastQuiz := m.wordService.GetLastQuiz()
	if lastQuiz != nil {
		score, total = lastQuiz.GetPoint(), lastQuiz.GetTotalWords()
	}

	return []string{
		"Are you ready for the review test?",
		"Latest point: " + strconv.Itoa(score) + "/" + strconv.Itoa(total),
	}
}

type SearchWordsMessenger struct {
	word *word.Word
}

func NewSearchWordsMessenger(word *word.Word) *SearchWordsMessenger {
	return &SearchWordsMessenger{word: word}
}

func (m *SearchWordsMessenger) Render() []string {
	output := []string{
		"Word: " + m.word.Name,
		"Description: " + m.word.Description,
	}

	// for pass the testing
	posOrder := map[word.PoS]int{
		word.Verb:      1,
		word.Adjective: 2,
		word.Noun:      3,
		word.Adverb:    4,
	}

	sortedDefinitions := make([]word.Definition, len(m.word.Definitions))
	copy(sortedDefinitions, m.word.Definitions)

	sort.Slice(sortedDefinitions, func(i, j int) bool {
		return posOrder[sortedDefinitions[i].Pos] < posOrder[sortedDefinitions[j].Pos]
	})

	for _, definition := range sortedDefinitions {
		output = append(output, fmt.Sprintf("%s - %s", definition.Pos.String(), definition.Explanation))
	}

	return output
}

type AddNewWordsMessenger struct {
	currentWords  []word.Word
	wordsAdded    []word.Word
	wordsNotFound []string
}

func NewAddNewWordsMessenger(currentWords []word.Word, wordsAdded []word.Word, wordsNotFound []string) *AddNewWordsMessenger {
	return &AddNewWordsMessenger{currentWords: currentWords, wordsAdded: wordsAdded, wordsNotFound: wordsNotFound}
}

func (m *AddNewWordsMessenger) Render() []string {
	currentWordsString := []string{}
	for _, word := range m.currentWords {
		currentWordsString = append(currentWordsString, word.Name)
	}

	wordsAddedString := []string{}
	for _, word := range m.wordsAdded {
		wordsAddedString = append(wordsAddedString, word.Name)
	}

	output := []string{
		"Current Words: " + strings.Join(currentWordsString, ", "),
	}

	if len(wordsAddedString) > 0 {
		output = append(output, "Words successfully added: "+strings.Join(wordsAddedString, ", "))
	}

	if len(m.wordsNotFound) > 0 {
		output = append(output, "Words not found: "+strings.Join(m.wordsNotFound, ", "))
	}

	return output
}

type DeleteWordsMessenger struct {
	currentWords  []word.Word
	wordsDeleted  []word.Word
	wordsNotFound []string
}

func NewDeleteWordsMessenger(currentWords []word.Word, wordsDeleted []word.Word, wordsNotFound []string) *DeleteWordsMessenger {
	return &DeleteWordsMessenger{currentWords: currentWords, wordsDeleted: wordsDeleted, wordsNotFound: wordsNotFound}
}

func (m *DeleteWordsMessenger) Render() []string {
	currentWordsString := []string{}
	for _, word := range m.currentWords {
		currentWordsString = append(currentWordsString, word.Name)
	}

	output := []string{}

	if len(currentWordsString) > 0 {
		output = append(output, "Current Words: "+strings.Join(currentWordsString, ", "))
	} else {
		output = append(output, "You don't have any words.")
	}

	if len(m.wordsDeleted) > 0 {
		wordsDeletedString := []string{}
		for _, word := range m.wordsDeleted {
			wordsDeletedString = append(wordsDeletedString, word.Name)
		}
		output = append(output, "Words successfully deleted: "+strings.Join(wordsDeletedString, ", "))
	}

	if len(m.wordsNotFound) > 0 {
		wordsNotFoundString := []string{}
		for _, word := range m.wordsNotFound {
			wordsNotFoundString = append(wordsNotFoundString, word)
		}
		output = append(output, "Words not found: "+strings.Join(wordsNotFoundString, ", "))
	}

	return output
}

type QuestionMessenger struct {
	wordService *word.WordService
}

func NewQuestionMessenger(wordService *word.WordService) *QuestionMessenger {
	return &QuestionMessenger{wordService: wordService}
}

func (m *QuestionMessenger) Render() []string {
	return []string{
		"Point: " + strconv.Itoa(m.wordService.GetCurrentQuiz().GetPoint()),
		"Remaining: " + strconv.Itoa(max(m.wordService.GetCurrentQuiz().GetTotalWords()-m.wordService.GetCurrentQuiz().GetCurrentIndex(), 0)),
	}
}

type ResultMessenger struct {
	wordService *word.WordService
}

func NewResultMessenger(wordService *word.WordService) *ResultMessenger {
	return &ResultMessenger{wordService: wordService}
}

func (m *ResultMessenger) Render() []string {
	return []string{
		"The exam is over!",
		"You got " + strconv.Itoa(m.wordService.GetLastQuiz().GetPoint()) + "/" + strconv.Itoa(m.wordService.GetLastQuiz().GetTotalWords()) + " point.",
	}
}
