package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rulanugrh/venus/config"
	handler "github.com/rulanugrh/venus/internal/http"
	"github.com/rulanugrh/venus/internal/routes"
	"github.com/rulanugrh/venus/internal/service"
	"github.com/rulanugrh/venus/util"
)

func main() {
	conf := config.GetConfig()
	conn := util.GetClient()
	tracer := util.GetTracer()

	opentele := util.InitTracer()
	defer func(){
		if err := opentele.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down: %v", err)
		}
	}()

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

	file, _ := os.OpenFile("./log/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone: "Asia/Jakarta",
		Output: file,
	}))

	f.Use(otelfiber.Middleware())
	containerService := service.NewContainerService(conn, tracer)
	containerHandler := handler.NewContainerHandler(containerService)

	imageService := service.NewImageService(conn, tracer)
	imageHandler := handler.NewImageHandler(imageService)

	userService := service.NewUserService(*conf)
	userHandler := handler.NewUserHandler(userService)
	
	networkService := service.NewNetworkService(conn, tracer)
	networkHandler := handler.NewNetworkHandler(networkService)

	volumeService := service.NewVolumeService(conn, tracer)
	volumeHandler := handler.NewVolumeHandler(volumeService)

	routes.RouteContainer(f, containerHandler)
	routes.RouteImage(f, imageHandler)
	routes.RouteUser(f, userHandler)
	routes.RouteNetwork(f, networkHandler)
	routes.RouteVolume(f, volumeHandler)

	server := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	log.Fatal(f.Listen(server))
}
