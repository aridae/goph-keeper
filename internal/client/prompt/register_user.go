package prompt

import (
	"context"
	promptio "github.com/aridae/goph-keeper/internal/client/prompt/dialog"
	registeruser "github.com/aridae/goph-keeper/internal/client/usecases/register-user"
)

func (s *Service) RunRegisterUserPrompt(ctx context.Context) {
	input := registerUserPromptInput{}
	dialog := buildRegisterUserDialog(&input)
	dialog.MustRun()

	err := s.registerUserHandler.Handle(ctx, registeruser.Request{
		Login:    input.login,
		Password: input.password,
	})
	if err != nil {
		dialog.PresentError(err, printableErrorMessage)
		return
	}

	dialog.PresentSuccess("Your account has been successfully created!")
}

type registerUserPromptInput struct {
	login    string
	password string
}

func buildRegisterUserDialog(input *registerUserPromptInput) promptio.Dialog {
	dialog := promptio.Dialog{
		Prompts: []promptio.Prompt{
			{
				EntryText: "Login:",
				AcceptResultFunc: func(s string) {
					input.login = s
				},
			},
			{
				EntryText: "Password:",
				AcceptResultFunc: func(s string) {
					input.password = s
				},
			},
		},
	}

	return dialog
}
