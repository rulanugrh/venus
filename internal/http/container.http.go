package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type containerstruct struct {
	service iservice.ContainerInterface
}

func NewContainerHandler(service iservice.ContainerInterface) ihttp.ContainerInterface {
	return &containerstruct{
		service: service,
	}
}

func(container *containerstruct) CreateContainer(ctx *fiber.Ctx) error {
	var model dto.Container
	err := ctx.BodyParser(&model)
	if err != nil {
		response := web.Failure{
			Code: 500,
			Message: "Tidak bisa binding",
			Error: err,
		}

		return ctx.Status(500).JSON(response)
	}

	data, errCreate := container.service.Create(model, ctx.Context())
	if errCreate != nil {
		response := web.Failure{
			Code: 500,
			Message: "Tidak bisa request create container",
			Error: errCreate,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Success create container",
		Data: data,
	}

	return ctx.Status(200).JSON(response)
	
}

func(container *containerstruct) ListContainer(ctx *fiber.Ctx) error {
	data, err := container.service.ListContainer(ctx.UserContext())
	if err != nil {
		response := web.Failure{
			Code: 400,
			Message: "Container tidak ditemukan",
			Error: err,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Container ditemukan",
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func(container *containerstruct) DeleteContainer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := container.service.DeleteContaienr(id, ctx.UserContext())
	if err != nil {
		response := web.Failure{
			Code: 400,
			Message: "Gagal delete container",
			Error: err,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "Container berhasil dihapus",
		Code: 200,
		Data: nil,
	}

	return ctx.Status(200).JSON(response)
}

func(container *containerstruct) InspectContainer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := container.service.InspectContainer(ctx.UserContext(), id)

	if err != nil {
		response := web.Failure{
			Code: 400,
			Message: "Gagal inspect container dengan id ini",
			Error: err,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Container ditemukan",
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func(container *containerstruct) ExecContainer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := container.service.ExecContainer(id, ctx.Context().RequestBodyStream(), ctx.Context().Request.BodyWriter(), ctx.UserContext())

	if err != nil {
		response := web.Failure{
			Code: 500,
			Message: "Gagal exec container dengan id ini",
			Error: err,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Message: "Success exec container",
		Code: 200,
		Data: "Berhasil exec container",
	}

	return ctx.Status(200).JSON(response)
}

func(container *containerstruct) Logger(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	err := container.service.Logger(name, ctx.UserContext(), ctx.Context().Request.BodyWriter())
	if err != nil {
		response := web.Failure{
			Code: 500,
			Message: "Gagal melihat log container dengan id ini",
			Error: err,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Message: "Success melihat log container ini",
		Code: 200,
		Data: "Berhasil",
	}

	return ctx.Status(200).JSON(response)
}