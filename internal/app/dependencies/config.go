package dependencies

// EntrypointConfig конфиги для Entrypoint - вход в бизнес-лоигку прилоджения
type EntrypointConfig struct {
	LinkUcRandomCreateRetries int64
	RedirectHost              string

	VkGroupUrl string
}
