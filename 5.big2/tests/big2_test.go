package tests

import (
	"github/elliot9/big2/common"
	"github/elliot9/big2/core"
	"github/elliot9/big2/entities/big2"
	"github/elliot9/big2/entities/deck"
	"github/elliot9/big2/infra/cli"
	"os"
	"regexp"
	"strings"
	"testing"
)

func testNormalBig2(t *testing.T, inputFile string, outputFile string) {
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

	// 設置遊戲
	mockCLI := cli.NewMockCLI(inputs[1:]) // 跳過第一行的牌組

	// 將第一行轉換為牌堆
	cardStrings := strings.Split(inputs[0], " ")
	cards := []core.Card{}
	for _, cardString := range cardStrings {
		// 用正則表達式取出花色和數字
		re := regexp.MustCompile(`([A-Z])\[([A-Z0-9]+)\]`)
		matches := re.FindStringSubmatch(cardString)
		suit := map[string]core.Suit{
			"C": core.Club,
			"D": core.Diamond,
			"H": core.Heart,
			"S": core.Spade,
		}[matches[1]]

		rank := map[string]core.Rank{
			"3":  core.Three,
			"4":  core.Four,
			"5":  core.Five,
			"6":  core.Six,
			"7":  core.Seven,
			"8":  core.Eight,
			"9":  core.Nine,
			"10": core.Ten,
			"J":  core.Jack,
			"Q":  core.Queen,
			"K":  core.King,
			"A":  core.Ace,
			"2":  core.Two,
		}[matches[2]]

		cards = append(cards, core.NewCard(rank, suit))
	}

	for i, j := 0, len(cards)-1; i < j; i, j = i+1, j-1 {
		cards[i], cards[j] = cards[j], cards[i]
	}

	deck := deck.NewMockDeck(cards)

	game := big2.NewNormalBig2(mockCLI, common.InitializePlayers(mockCLI), deck)
	game.SetCardPattern(common.InitializeCardPatterns())
	game.SetPlayValidator(common.InitializePlayValidator())
	game.Start()
	// 驗證輸出
	actual := mockCLI.GetOutputs()
	for i, output := range outputs {
		if strings.TrimSpace(actual[i]) != strings.TrimSpace(output) {
			t.Errorf("index %d: expected %q, but got %q", i, output, actual[i])
		}
	}
}

func TestAlwaysPlayFirstCard(t *testing.T) {
	testNormalBig2(t, "data/always-play-first-card.in", "data/always-play-first-card.out")
}

func TestFullHouse(t *testing.T) {
	testNormalBig2(t, "data/fullhouse.in", "data/fullhouse.out")
}

func TestIllegalActions(t *testing.T) {
	testNormalBig2(t, "data/illegal-actions.in", "data/illegal-actions.out")
}

func TestNormalNoErrorPlay1(t *testing.T) {
	testNormalBig2(t, "data/normal-no-error-play1.in", "data/normal-no-error-play1.out")
}

func TestNormalNoErrorPlay2(t *testing.T) {
	testNormalBig2(t, "data/normal-no-error-play2.in", "data/normal-no-error-play2.out")
}
