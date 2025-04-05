package prompt

import (
	"context"
	"errors"
	createsecret "github.com/aridae/goph-keeper/internal/client/usecases/create-secret"
)

func (s *Service) CreateSecret() {
	ctx := context.Background()

	type promptInput struct {
		key  string
		data []byte
	}

	input := promptInput{}

	dialog := dialog{
		label: "create-secret-dialog",
		prompts: []prompt{
			{
				label:     "input-secret-key",
				entryText: "Please input key of the secret to create:",
				entryValidationFunc: func(s string) error {
					if len(s) == 0 {
						return errors.New("empty key is unprocessable :(")
					}
					return nil
				},
				acceptResultFunc: func(s string) {
					input.key = s
				},
			},
			{
				label:     "input-secret-data",
				entryText: "Please input secret data:",
				entryValidationFunc: func(s string) error {
					if len(s) == 0 {
						return errors.New("empty data is unprocessable :(")
					}
					return nil
				},
				acceptResultFunc: func(s string) {
					input.data = []byte(s)
				},
			},
		},
	}

	dialog.mustRun()

	err := s.createSecretHandler.Handle(ctx, createsecret.Request{
		Key:  input.key,
		Data: input.data,
	})
	if err != nil {
		presentError(err)
	}
}
