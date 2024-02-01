package config

import (
	"os"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/joho/godotenv"
)

type App struct {
	Server struct {
		Host string
		Port string
		Origin string
	}
	Admin struct {
		Email    string
		Password string
	}
}

var app *App

func GetClient() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil
	}

	return client
}

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

		return &conf
	}

	conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
	conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")
	conf.Server.Host = os.Getenv("SERVER_HOST")
	conf.Server.Port = os.Getenv("SERVER_PORT")

	return &conf
}