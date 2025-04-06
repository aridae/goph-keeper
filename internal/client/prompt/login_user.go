package prompt

import (
	"context"
	promptio "github.com/aridae/goph-keeper/internal/client/pkg/prompt-io"
	loginuser "github.com/aridae/goph-keeper/internal/client/usecases/login-user"
)

func (s *Service) RunLoginUserPrompt(ctx context.Context) {
	input := loginUserPromptInput{}
	dialog := buildLoginUserDialog(&input)
	dialog.MustRun()

	err := s.loginUserHandler.Handle(ctx, loginuser.Request{
		Login:    input.login,
		Password: input.password,
	})
	if err != nil {
		dialog.PresentError(err, printableErrorMessage)
		return
	}

	dialog.PresentSuccess("You are logged in!")
}

type loginUserPromptInput struct {
	login    string
	password string
}

func buildLoginUserDialog(input *loginUserPromptInput) promptio.Dialog {
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
