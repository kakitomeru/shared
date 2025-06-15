package env

import (
	"errors"
	"os"
)

func returnError(key string) error {
	return errors.New(key + " is not set. Ensure you used Makefile's command to run the app. Also ensure you have .env file in the service's root path. See .env.example")
}

func LoadEnv(envs []string) error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return returnError("APP_ENV")
	}

	var err error
	for _, env := range envs {
		switch env {
		case "jwt":
			err = loadJwt()
		case "postgres":
			err = loadPostgres()
		case "otel":
			err = loadOtelCollector()
		case "auth":
			err = loadAuth()
		case "gateway":
			err = loadGateway()
		case "snippet":
			err = loadSnippet()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAppEnv() string {
	return os.Getenv("APP_ENV")
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetPostgresUser() string {
	return os.Getenv("POSTGRES_USER")
}

func GetPostgresPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}

func GetPostgresHost() string {
	return os.Getenv("POSTGRES_HOST")
}

func GetPostgresPort() string {
	return os.Getenv("POSTGRES_PORT")
}

func GetPostgresDB() string {
	return os.Getenv("POSTGRES_DB")
}

func GetOtelCollector() string {
	return os.Getenv("OTEL_COLLECTOR")
}

func GetAuthHost() string {
	return os.Getenv("AUTH_HOST")
}

func GetAuthPort() string {
	return os.Getenv("AUTH_PORT")
}

func GetGatewayHost() string {
	return os.Getenv("GATEWAY_HOST")
}

func GetGatewayPort() string {
	return os.Getenv("GATEWAY_PORT")
}

func GetSnippetHost() string {
	return os.Getenv("SNIPPET_HOST")
}

func GetSnippetPort() string {
	return os.Getenv("SNIPPET_PORT")
}

func loadJwt() error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return returnError("JWT_SECRET")
	}

	return nil
}

func loadPostgres() error {
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		return returnError("POSTGRES_USER")
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return returnError("POSTGRES_PASSWORD")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		return returnError("POSTGRES_HOST")
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		return returnError("POSTGRES_PORT")
	}

	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		return returnError("POSTGRES_DB")
	}

	return nil
}

func loadOtelCollector() error {
	otelCollector := os.Getenv("OTEL_COLLECTOR")
	if otelCollector == "" {
		return returnError("OTEL_COLLECTOR")
	}

	return nil
}

func loadAuth() error {
	authHost := os.Getenv("AUTH_HOST")
	if authHost == "" {
		return returnError("AUTH_HOST")
	}

	authPort := os.Getenv("AUTH_PORT")
	if authPort == "" {
		return returnError("AUTH_PORT")
	}

	return nil
}

func loadSnippet() error {
	snippetHost := os.Getenv("SNIPPET_HOST")
	if snippetHost == "" {
		return returnError("SNIPPET_HOST")
	}

	snippetPort := os.Getenv("SNIPPET_PORT")
	if snippetPort == "" {
		return returnError("SNIPPET_PORT")
	}

	return nil
}

func loadGateway() error {
	gatewayHost := os.Getenv("GATEWAY_HOST")
	if gatewayHost == "" {
		return returnError("GATEWAY_HOST")
	}

	gatewayPort := os.Getenv("GATEWAY_PORT")
	if gatewayPort == "" {
		return returnError("GATEWAY_PORT")
	}

	return nil
}
