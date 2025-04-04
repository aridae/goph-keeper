package models

import "time"

type Secret struct {
	Accessor  SecretAccessor
	Data      []byte
	Meta      map[string]string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SecretAccessor struct {
	OwnerUsername string
	Key           string
}
