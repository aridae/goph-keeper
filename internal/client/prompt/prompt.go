package prompt

import (
	"context"
	"fmt"
	createsecret "github.com/aridae/goph-keeper/internal/client/usecases/create-secret"
	"github.com/aridae/goph-keeper/internal/logger"
	"github.com/manifoldco/promptui"
)

type createSecretHandler interface {
	Handle(ctx context.Context, command createsecret.Request) error
}

type Service struct {
	createSecretHandler createSecretHandler
}

func NewService(
	createSecretHandler createSecretHandler,
) *Service {
	return &Service{
		createSecretHandler: createSecretHandler,
	}
}

var promptTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

type dialog struct {
	label   string
	prompts []prompt
}

func (d dialog) mustRun() {
	for _, p := range d.prompts {
		p.mustRun()
	}
}

type prompt struct {
	label               string
	entryText           string
	entryValidationFunc func(string) error
	acceptResultFunc    func(string)
}

func (p prompt) mustRun() {
	prompt := promptui.Prompt{
		Label:     p.entryText,
		Validate:  p.entryValidationFunc,
		Templates: promptTemplates,
	}

	result, err := prompt.Run()
	if err != nil {
		logger.Fatalf("prompt <label:%s> failed: %v", p.label, err)
	}

	p.acceptResultFunc(result)
}

func presentError(err error) {
	fmt.Printf("Command failed with error: %v", err)
}
