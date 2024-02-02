package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
	"github.com/rulanugrh/venus/internal/middleware"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type userstruct struct {
	service iservice.UserInterface
}

func NewUserHandler(service iservice.UserInterface) ihttp.UserInterface{
	return &userstruct{
		service: service,
	}
}

func(user *userstruct) Login(ctx *fiber.Ctx) error {
	var model dto.User
	err := ctx.BodyParser(&model)
	if err != nil {
		response := web.Failure{
			Code: 500,
			Message: "Gagal binding request",
			Error: err.Error(),
		}

		return ctx.Status(500).JSON(response)
	}

	errLogin := user.service.Login(model)
	if errLogin != nil {
		response := web.WebValidationError{
			Message: "Gagal login",
			Errors: errLogin,
		}

		return ctx.Status(400).JSON(response)
	}

	token, errToken := middleware.CreateToken(model)
	if errToken != nil {
		response := web.Failure{
			Code: 500,
			Message: "Gagal generate token",
			Error: errToken,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Success login",
		Data: token,
	}

	return ctx.Status(200).JSON(response)
}