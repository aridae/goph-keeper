package promptio

import (
	"errors"
	"github.com/aridae/goph-keeper/internal/logger"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var promptTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

type Dialog struct {
	Prompts []Prompt
}

func (d Dialog) MustRun() {
	for _, p := range d.Prompts {
		p.MustRun()
	}
}

func (d Dialog) PresentSuccess(message string) {
	color.Blue(message)
}

func (d Dialog) PresentError(err error) {
	color.Red("Error occurred: %v", err)
}

type Prompt struct {
	EntryText        string
	AcceptResultFunc func(string)
}

func (p Prompt) MustRun() {
	prompt := promptui.Prompt{
		Label: p.EntryText,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New("empty prompt entry")
			}
			return nil
		},
		Templates: promptTemplates,
	}

	result, err := prompt.Run()
	if err != nil {
		logger.Fatalf("failed to run prompt: %v", err)
	}

	p.AcceptResultFunc(result)
}
