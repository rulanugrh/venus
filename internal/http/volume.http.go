package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type volumestruct struct {
	service iservice.VolumeInterface
}

func NewVolumeHandler(service iservice.VolumeInterface) ihttp.VolumeInterface {
	return &volumestruct{
		service: service,
	}
}

func(volume *volumestruct) CreateVolume(ctx *fiber.Ctx) error {
	var models dto.Volume

	err := ctx.BodyParser(&models)
	if err != nil {
		response := web.Failure{
			Message: "Gagal binding data",
			Error: err.Error(),
			Code: 500,
		}

		return ctx.Status(500).JSON(response)
	}

	data, err := volume.service.CreateVolume(models, ctx.UserContext())
	if err != nil {
		response := web.Failure{
			Message: "Gagal create volume",
			Error: err.Error(),
			Code: 400,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "Success create volume",
		Data: data,
		Code: 200,
	}

	return ctx.Status(200).JSON(response)
}

func(volume *volumestruct) InspectVolume(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	data, err := volume.service.InspectVolume(name, ctx.UserContext())
	if err != nil {
		response := web.Failure{
			Message: "Gagal inspect volume",
			Error: err.Error(),
			Code: 400,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "Data ditemukan",
		Code: 200,
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func(volume *volumestruct) ListVolume(ctx *fiber.Ctx) error {
	data, err := volume.service.ListVolume(ctx.UserContext())
	if err != nil {
		response := web.Failure{
			Message: "volume tidak ditemukan",
			Error: err.Error(),
			Code: 400,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "volume ditemukan",
		Code: 200,
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func(volume *volumestruct) DeleteVolume(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	err := volume.service.DeleteVolume(name, ctx.UserContext())
	
	if err != nil {
		response := web.Failure{
			Message: "gagal delete volume",
			Error: err.Error(),
			Code: 400,
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "volume telah dihapus",
		Code: 200,
		Data: nil,
	}

	return ctx.Status(200).JSON(response)
}
