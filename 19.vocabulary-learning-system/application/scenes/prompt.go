package scenes

import (
	"fmt"
	"github/elliot9/class19/word"
)

type Prompt interface {
	Render() []string
}

type SimplePrompt struct {
	text string
}

func NewSimplePrompt(text string) *SimplePrompt {
	return &SimplePrompt{text: text}
}

func (p *SimplePrompt) Render() []string {
	return []string{p.text}
}

type QuestionPrompt struct {
	wordService *word.WordService
}

func NewQuestionPrompt(wordService *word.WordService) *QuestionPrompt {
	return &QuestionPrompt{wordService: wordService}
}

func (p *QuestionPrompt) Render() []string {
	index := p.wordService.GetCurrentQuiz().GetCurrentIndex()
	wordBlank, definition := p.wordService.GetCurrentQuestion()

	return []string{
		fmt.Sprintf("%d. Question: %s: %s", index+1, wordBlank, definition),
		"Answer: ",
	}
}
