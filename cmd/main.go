package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rulanugrh/venus/config"
	handler "github.com/rulanugrh/venus/internal/http"
	"github.com/rulanugrh/venus/internal/routes"
	"github.com/rulanugrh/venus/internal/service"
)

func main() {
	conf := config.GetConfig()
	conn := config.GetClient()

	f := fiber.New()
	f.Use(cors.New(cors.Config{
		AllowOrigins: conf.Server.Origin,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodDelete,
			fiber.MethodPost,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	containerService := service.NewContainerService(conn)
	containerHandler := handler.NewContainerHandler(containerService)

	imageService := service.NewImageService(conn)
	imageHandler := handler.NewImageHandler(imageService)

	userService := service.NewUserService(*conf)
	userHandler := handler.NewUserHandler(userService)

	routes.RouteContainer(f, containerHandler)
	routes.RouteImage(f, imageHandler)
	routes.RouteUser(f, userHandler)

	server := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	log.Fatal(f.Listen(server))
}
