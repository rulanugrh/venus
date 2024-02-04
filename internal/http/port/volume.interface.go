package ihttp

import "github.com/gofiber/fiber/v2"

type VolumeInterface interface {
	CreateVolume(ctx *fiber.Ctx) error
	InspectVolume(ctx *fiber.Ctx) error
	ListVolume(ctx *fiber.Ctx) error
	DeleteVolume(ctx *fiber.Ctx) error
}