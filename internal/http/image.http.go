package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type imagestruct struct {
	service iservice.ImageInterface
}

func NewImageHandler(service iservice.ImageInterface) ihttp.ImageInterface {
	return &imagestruct{
		service: service,
	}
}

func(image *imagestruct) PullImage(ctx *fiber.Ctx) error {
	var model dto.Image
	err := ctx.BodyParser(&model)
	if err != nil  {
		response := web.Failure{
			Message: "Tidak bisa membaca request",
			Code: 500,
			Error: err,
		}
		return ctx.Status(500).JSON(response)
	}

	errCreate := image.service.PullImage(ctx.UserContext(), model)
	if errCreate != nil {
		response := web.Failure{
			Message: "Tidak bisa pull image",
			Code: 500,
			Error: err,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Message: "Berhasil pull image",
		Code: 200,
		Data: nil,
	}

	return ctx.Status(200).JSON(response)
}

func(image *imagestruct) ListImage(ctx *fiber.Ctx) error {
	data, err := image.service.ListImage(ctx.UserContext())
	if err != nil {
		response := web.Failure{
			Code: 400,
			Message: "Image tidak ditemukan",
			Error: err,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Image ditemukan",
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func(image *imagestruct) DeleteImage(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := image.service.DeleteImage(id, ctx.UserContext())

	if err != nil {
		response := web.Failure{
			Code: 500,
			Message: "Image tidak berhasil dihapus",
			Error: err,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Image berhasil dihapus",
		Data: nil,
	}

	return ctx.Status(200).JSON(response)
}

func(image *imagestruct) InspectImage(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := image.service.InspectImage(id, ctx.UserContext())

	if err != nil {
		response := web.Failure{
			Code: 400,
			Message: "Image tidak ditemukan",
			Error: err,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Image berhasil ditemukan",
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func (image *imagestruct) BuildImage(ctx *fiber.Ctx) error {
	var model dto.BuildImage
	dockerfile, err := ctx.FormFile("file")
	if err != nil {
		response := web.Failure{
			Message: "Gagal binding form",
			Error: err.Error(),
			Code: 500,
		}

		return ctx.Status(500).JSON(response)
	}

	model.Dockerfile = dockerfile.Filename
	model.InputStream = ctx.Context().RequestBodyStream()
	model.OutputStream = ctx.Context().Response.BodyWriter()
	
	err = ctx.BodyParser(&model)
	if err != nil {
		response := web.Failure{
			Message: "Gagal binding body",
			Error: err.Error(),
			Code: 500,
		}

		return ctx.Status(500).JSON(response)
	}

	errBuild := image.service.BuildImage(model, ctx.UserContext())
	if errBuild != nil {
		response := web.Failure{
			Message: "Gagal build image",
			Error: errBuild.Error(),
			Code: 500,
		}

		return ctx.Status(500).JSON(response)
	}

	response := web.Success{
		Code: 200,
		Message: "Berhasil build image",
		Data: model.Name,
	}

	return ctx.Status(200).JSON(response)
}