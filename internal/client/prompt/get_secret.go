package prompt

import (
	"context"
	"fmt"
	promptio "github.com/aridae/goph-keeper/internal/client/prompt/dialog"
)

func (s *Service) RunGetSecretPrompt(ctx context.Context) {
	input := getSecretPromptInput{}
	dialog := buildGetSecretDialog(&input)
	dialog.MustRun()

	secret, err := s.getSecretHandler.Handle(ctx, input.key)
	if err != nil {
		dialog.PresentError(err, printableErrorMessage)
		return
	}

	dialog.PresentSuccess(fmt.Sprintf("Your secret: %s", string(secret.Data)))
}

type getSecretPromptInput struct {
	key string
}

func buildGetSecretDialog(input *getSecretPromptInput) promptio.Dialog {
	dialog := promptio.Dialog{
		Prompts: []promptio.Prompt{
			{
				EntryText: "What secret do you want to get?",
				AcceptResultFunc: func(s string) {
					input.key = s
				},
			},
		},
	}

	return dialog
}
