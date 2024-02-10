package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Server struct {
		Host   string
		Port   string
		Origin string
		Secret string
	}
	Admin struct {
		Email    string
		Password string
	}

	Opentelemetry struct {
		Name string
	}

	Docker struct {
		Email    string
		Password string
		Username string
	}
}

var app *App

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	conf := App{}
	if err := godotenv.Load(); err != nil {
		conf.Admin.Email = ""
		conf.Admin.Password = ""
		conf.Server.Host = "localhost"
		conf.Server.Port = "3000"

		conf.Opentelemetry.Name = ""
		return &conf
	}

	conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
	conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")

	conf.Server.Host = os.Getenv("SERVER_HOST")
	conf.Server.Port = os.Getenv("SERVER_PORT")
	conf.Server.Origin = os.Getenv("SERVER_ORIGIN")
	conf.Server.Secret = os.Getenv("SERVER_SECRET")

	conf.Opentelemetry.Name = os.Getenv("OTEL_SERVICE_NAME")
	conf.Docker.Email = os.Getenv("DOCKER_EMAIL")
	conf.Docker.Password = os.Getenv("DOCKER_PASSWORD")
	conf.Docker.Username = os.Getenv("DOCKER_USERNAME")

	return &conf
}
