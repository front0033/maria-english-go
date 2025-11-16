package config

// Config содержит конфигурационные данные приложения
type Config struct {
	Token string
}

// LoadConfig загружает конфигурацию (в данном случае, только токен)
func LoadConfig() Config {
	token := "8275702411:AAF61gObZMJOuB5gSnqRhG896f-dcigZ5O0"
	
	return Config{
		Token: token,
	}
}
