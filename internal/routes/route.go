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
	route.Get("/execute/:id", handler.ExecContainer)
	route.Delete("/delete/:id", handler.DeleteContainer)
}

func RouteImage(f fiber.Router, handler ihttp.ImageInterface) {
	route := f.Group("/api/image", middleware.JWTVerify())
	route.Post("/create", handler.PullImage)
	route.Post("/build", handler.BuildImage)
	route.Get("/find", handler.ListImage)
	route.Get("/find/:id", handler.InspectImage)
	route.Delete("/delete/:id", handler.DeleteImage)
}

func RouteNetwork(f fiber.Router, handler ihttp.NetworkInterface) {
	route := f.Group("/api/network", middleware.JWTVerify())
	route.Post("/create", handler.CreateNetwork)
	route.Get("/find/:id", handler.InspectNetwork)
	route.Get("/find", handler.ListNetworks)
	route.Delete("/delete/:id", handler.DeleteNetwork)
}

func RouteVolume(f fiber.Router, handler ihttp.VolumeInterface) {
	route := f.Group("/api/volume", middleware.JWTVerify())
	route.Post("/create", handler.CreateVolume)
	route.Get("/find", handler.ListVolume)
	route.Get("/find/:id", handler.InspectVolume)
	route.Delete("/delete/:id", handler.DeleteVolume)
}

func RouteUser(f fiber.Router, handler ihttp.UserInterface) {
	route := f.Group("/api/user")
	route.Post("/login", handler.Login)
}
