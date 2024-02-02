package ihttp

import "github.com/gofiber/fiber/v2"

type ImageInterface interface {
	PullImage(ctx *fiber.Ctx) error
	ListImage(ctx *fiber.Ctx) error
	DeleteImage(ctx *fiber.Ctx) error
	InspectImage(ctx *fiber.Ctx) error
	BuildImage(ctx *fiber.Ctx) error
}