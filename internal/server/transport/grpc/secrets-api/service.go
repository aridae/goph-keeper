package secretsapi

import (
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
)

type useCasesController interface {
}

type Implementation struct {
	desc.UnimplementedSecretsServiceServer

	useCasesController useCasesController
}

func New(useCasesController useCasesController) *Implementation {
	return &Implementation{useCasesController: useCasesController}
}
