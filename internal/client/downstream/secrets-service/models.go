package secretsservice

// Secret представляет информацию о секрете, хранящемся в сервисе SecretsService.
type Secret struct {
	// Key уникальный идентификатор секрета.
	Key string
	// Data содержимое секрета в байтах.
	Data []byte
	// Meta дополнительные метаданные секрета.
	Meta map[string]string
}
