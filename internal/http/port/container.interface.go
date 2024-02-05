package ihttp

import "github.com/gofiber/fiber/v2"

type ContainerInterface interface {
	CreateContainer(ctx *fiber.Ctx) error
	ListContainer(ctx *fiber.Ctx) error
	DeleteContainer(ctx *fiber.Ctx) error
	InspectContainer(ctx *fiber.Ctx) error
	ExecContainer(ctx *fiber.Ctx) error
	Logger(ctx *fiber.Ctx) error
}