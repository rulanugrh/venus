package ihttp

import "github.com/gofiber/fiber/v2"

type NetworkInterface interface {
	CreateNetwork(ctx *fiber.Ctx) error
	InspectNetwork(ctx *fiber.Ctx) error
	ListNetworks(ctx *fiber.Ctx) error
}
