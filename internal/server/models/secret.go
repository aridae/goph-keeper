package models

type Secret struct {
	Accessor SecretAccessor
	Data     []byte
	Meta     map[string]string
}

type SecretAccessor struct {
	OwnerUsername string
	Key           string
}
