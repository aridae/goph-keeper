package usersservice

// UserCredentials содержит учетные данные пользователя для входа в систему.
type UserCredentials struct {
	// Username имя пользователя.
	Username string
	// Password пароль пользователя.
	Password string
}
