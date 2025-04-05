package secretsservice

type Secret struct {
	Key  string
	Data []byte
	Meta map[string]string
}
