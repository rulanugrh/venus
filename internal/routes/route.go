package routes

import (
	"github.com/gofiber/fiber/v2"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
)

func RouteContainer(f fiber.Router, handler ihttp.ContainerInterface) {
	route := f.Group("/api/container")
	route.Post("/create", handler.CreateContainer)
	route.Get("/find", handler.ListContainer)
	route.Get("/find/:id", handler.InspectContainer)
	route.Get("/delete/:id", handler.DeleteContainer)
}

func RouteImage(f fiber.Router, handler ihttp.ImageInterface) {
	route := f.Group("/api/image")
	route.Post("/create", handler.PullImage)
	route.Get("/find", handler.ListImage)
	route.Get("/find/:id", handler.InspectImage)
	route.Get("/delete/:id", handler.DeleteImage)
}