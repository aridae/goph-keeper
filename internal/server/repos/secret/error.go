package secret

import "errors"

var (
	ErrOwnerUsernameKeyCombinationConstraintViolated = errors.New("unq_secrets__owner_username__key constraint violation")
)
