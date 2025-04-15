package tests

import (
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/action"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/one_punch_handlers"
	"github/elliot9/class9/internal/role"
	"github/elliot9/class9/internal/rpg"
	"os"
	"strconv"
	"strings"
	"testing"
)

type RoleConfig struct {
	Name        string
	HP, MP, STR int
	Skills      []string
}

type TroopConfig struct {
	Roles []RoleConfig
}

func testRPG(t *testing.T, inputFile string, outputFile string) {
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

	troop1Config, troop2Config, actions := parseInputs(inputs)
	mockCLI := cli.NewMockCLI(actions)

	playerTroop := parseRoles(troop1Config, mockCLI, true)
	enemyTroop := parseRoles(troop2Config, mockCLI, false)

	rpg := rpg.NewRPG(playerTroop, enemyTroop, mockCLI)
	rpg.Battle.Start()
	rpg.Battle.PrintResult()

	// 驗證輸出
	actual := mockCLI.GetOutputs()
	for i, output := range outputs {
		if strings.TrimSpace(actual[i]) != strings.TrimSpace(output) {
			t.Errorf("index %d: expected %q, but got %q", i, output, actual[i])
		}
	}
}

func parseInputs(inputs []string) (troop1, troop2 TroopConfig, actions []string) {
	troopMap := map[string]*TroopConfig{
		"#軍隊-1-開始": &troop1,
		"#軍隊-2-開始": &troop2,
	}
	var currentTroop *TroopConfig

	for _, line := range inputs {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if troop, exists := troopMap[line]; exists {
			currentTroop = troop
			continue
		}

		if strings.HasSuffix(line, "-結束") {
			currentTroop = nil
			continue
		}

		if currentTroop != nil {
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				role := RoleConfig{
					Name:   parts[0],
					HP:     parseIntOrZero(parts[1]),
					MP:     parseIntOrZero(parts[2]),
					STR:    parseIntOrZero(parts[3]),
					Skills: parts[4:],
				}
				currentTroop.Roles = append(currentTroop.Roles, role)
			}
		} else {
			actions = append(actions, line)
		}
	}

	return troop1, troop2, actions
}

func parseIntOrZero(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func parseRoles(troopConfig TroopConfig, cli cli.CLI, isPlayerTroop bool) []interfaces.Role {
	roles := []interfaces.Role{}
	for i, roleConfig := range troopConfig.Roles {
		skills := []interfaces.Action{actions["普通攻擊"]}
		for _, skill := range roleConfig.Skills {
			skills = append(skills, actions[skill])
		}

		if i == 0 && isPlayerTroop {
			roles = append(roles, role.NewHero(roleConfig.Name, roleConfig.HP, roleConfig.MP, roleConfig.STR, cli, skills))
		} else {
			roles = append(roles, role.NewSeedAI(roleConfig.Name, roleConfig.HP, roleConfig.MP, roleConfig.STR, cli, skills))
		}
	}
	return roles
}

func TestPoison(t *testing.T) {
	testRPG(t, "data/poison.in", "data/poison.out")
}

func TestPetrochemical(t *testing.T) {
	testRPG(t, "data/petrochemical.in", "data/petrochemical.out")
}

func TestOnlyBasicAttack(t *testing.T) {
	testRPG(t, "data/only-basic-attack.in", "data/only-basic-attack.out")
}

func TestSelfHealing(t *testing.T) {
	testRPG(t, "data/self-healing.in", "data/self-healing.out")
}

func TestSummon(t *testing.T) {
	testRPG(t, "data/summon.in", "data/summon.out")
}

func TestWaterballAndFireball1v2(t *testing.T) {
	testRPG(t, "data/waterball-and-fireball-1v2.in", "data/waterball-and-fireball-1v2.out")
}

func TestSelfExplosion(t *testing.T) {
	testRPG(t, "data/self-explosion.in", "data/self-explosion.out")
}

func TestCheerUp(t *testing.T) {
	testRPG(t, "data/cheerup.in", "data/cheerup.out")
}

func TestCurse(t *testing.T) {
	testRPG(t, "data/curse.in", "data/curse.out")
}

func TestOnePunch(t *testing.T) {
	testRPG(t, "data/one-punch.in", "data/one-punch.out")
}

var HPOver500Handler interfaces.OnePunchHandler = one_punch_handlers.NewHPOver500()
var PoisonedOrPetrochemicalHandler interfaces.OnePunchHandler = one_punch_handlers.NewPoisonedOrPetrochemical()
var CheerUpHandler interfaces.OnePunchHandler = one_punch_handlers.NewCheerUp()
var NormalHandler interfaces.OnePunchHandler = one_punch_handlers.NewNormal()

func init() {
	HPOver500Handler.SetNext(PoisonedOrPetrochemicalHandler)
	PoisonedOrPetrochemicalHandler.SetNext(CheerUpHandler)
	CheerUpHandler.SetNext(NormalHandler)
}

var actions map[string]interfaces.Action = map[string]interfaces.Action{
	"普通攻擊": action.NewBasicAttack(),
	"水球":   action.NewWaterball(),
	"火球":   action.NewFireball(),
	"自我治療": action.NewSelfHealing(),
	"石化":   action.NewPetrochemical(),
	"下毒":   action.NewPoison(),
	"召喚":   action.NewSummon(),
	"自爆":   action.NewSelfExplosion(),
	"鼓舞":   action.NewCheerUp(),
	"詛咒":   action.NewCurse(),
	"一拳攻擊": action.NewOnePunch(HPOver500Handler),
}
