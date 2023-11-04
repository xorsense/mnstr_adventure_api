package integration

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/xorsense/mnstr_adventure_api/adventurer"
)

type registrationState struct {
	Adventurer *adventurer.Adventurer
}

var state = &registrationState{}

func adventurerWithName(adverb, name string) error {
	var exists bool
	switch adverb {
	case "no", "not an":
		exists = false
	default:
		exists = true
	}

	a, err := adventurer.FindWithName(name)

	switch exists {
	case true:
		if a == nil {
			return fmt.Errorf(`
expected: adventurer with name %s was found
     got: adventurer with name %s was NOT found
`, name, name)
		}
		state.Adventurer = a
	default:
		if err == nil {
			return fmt.Errorf(`
expected: adventurer with name %s was NOT found
     got: adventurer with name %s was found
`, name, name)
		}
	}
	return nil
}

func viewFor(screen string) error {
	return godog.ErrPending
}

func enterInput(field, value string) error {
	return godog.ErrPending
}

func selectAction(label string) error {
	return godog.ErrPending
}

func seeItem(exists, item string) error {
	return godog.ErrPending
}

func stating(item, statment string) error {
	return godog.ErrPending
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^[Tt]here is (\w+)? adventurer with the name of (.+)`, adventurerWithName)
	ctx.Step(`^I am at the (\w+) screen`, viewFor)
	ctx.Step(`^I enter the (.+) of (.+)`, enterInput)
	ctx.Step(`^I should (not)? see a (\w+)`, seeItem)
	ctx.Step(`(\w+) stating "(.+)"`, stating)
}
