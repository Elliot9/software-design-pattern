package tests

import (
	"context"
	"github/elliot9/class19/application"
	"github/elliot9/class19/infra"
	"os"
	"strings"
	"testing"
	"time"
)

func TestAddAndDelete(t *testing.T) {
	test(t, "add-and-delete.in", "add-and-delete.out")
}

func TestSearchWords(t *testing.T) {
	test(t, "search-words.in", "search-words.out")
}

// 測試案例過不了
func TestReviewWords(t *testing.T) {
	test(t, "review-words.in", "review-words.out")
}

func test(t *testing.T, inputFile string, outputFile string) {
	inputContent, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	inputs := strings.Split(string(inputContent), "\n")

	expectedOutput, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	outputs := strings.Split(string(expectedOutput), "\n")

	cli := infra.NewMockCLI(inputs)

	finder := infra.NewLocalJsonWordFinder("../data/words.json")
	repository := infra.NewMockWordRepository()
	randomizer := infra.NewMockRandomizer()
	app := application.NewVocabularyLearningSystem(finder, repository, randomizer, cli)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		app.Run()
		close(done)
	}()

	select {
	case <-done:
	case <-ctx.Done():
		// 驗證
		for i, output := range outputs[:len(outputs)-1] {
			if cli.GetOutputs()[i] != output {
				t.Fatalf("line %d: expected %s, but got %s", i, output, cli.GetOutputs()[i])
			}
		}
	}
}
