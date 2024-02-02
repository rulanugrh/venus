package routes

import (
	"github.com/gofiber/fiber/v2"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
	"github.com/rulanugrh/venus/internal/middleware"
)

func RouteContainer(f fiber.Router, handler ihttp.ContainerInterface) {
	route := f.Group("/api/container", middleware.JWTVerify())
	route.Post("/create", handler.CreateContainer)
	route.Get("/find", handler.ListContainer)
	route.Get("/find/:id", handler.InspectContainer)
	route.Get("/delete/:id", handler.DeleteContainer)
}

func RouteImage(f fiber.Router, handler ihttp.ImageInterface) {
	route := f.Group("/api/image", middleware.JWTVerify())
	route.Post("/create", handler.PullImage)
	route.Get("/find", handler.ListImage)
	route.Get("/find/:id", handler.InspectImage)
	route.Get("/delete/:id", handler.DeleteImage)
}

func RouteUser(f fiber.Router, handler ihttp.UserInterface) {
	route := f.Group("/api/user")
	route.Post("/login", handler.Login)
}