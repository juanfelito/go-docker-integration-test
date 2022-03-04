package config

type Settings struct {
	DBName    string `envconfig:"DB_DATABASE"`
	DBHost    string `envconfig:"DB_HOST"`
	DBPass    string `envconfig:"DB_PASSWORD"`
	DBPort    int    `envconfig:"DB_PORT"`
	DBUser    string `envconfig:"DB_USERNAME"`
	DBSSLMode string `envconfig:"DB_SSL_MODE"`
	HTTPPort  int    `envconfig:"HTTP_PORT"`
}
