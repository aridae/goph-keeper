package prompt

import (
	"context"
	promptio "github.com/aridae/goph-keeper/internal/client/pkg/prompt-io"
	createsecret "github.com/aridae/goph-keeper/internal/client/usecases/create-secret"
)

func (s *Service) RunCreateSecretPrompt(ctx context.Context) {
	input := createSecretPromptInput{}
	dialog := buildCreateSecretDialog(&input)
	dialog.MustRun()

	err := s.createSecretHandler.Handle(ctx, createsecret.Request{
		Key:  input.key,
		Data: input.data,
	})
	if err != nil {
		dialog.PresentError(err)
		return
	}

	dialog.PresentSuccess("Your secret is safe with me!")
}

type createSecretPromptInput struct {
	key  string
	data []byte
}

func buildCreateSecretDialog(input *createSecretPromptInput) promptio.Dialog {
	dialog := promptio.Dialog{
		Prompts: []promptio.Prompt{
			{
				EntryText: "Please input key of the secret to create:",
				AcceptResultFunc: func(s string) {
					input.key = s
				},
			},
			{
				EntryText: "Please input secret data:",
				AcceptResultFunc: func(s string) {
					input.data = []byte(s)
				},
			},
		},
	}

	return dialog
}
