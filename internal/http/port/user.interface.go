package ihttp

import "github.com/gofiber/fiber/v2"

type UserInterface interface {
	Login(ctx *fiber.Ctx) error
}