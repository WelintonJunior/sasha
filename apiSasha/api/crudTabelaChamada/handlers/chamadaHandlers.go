package handlers

import (
	"strconv"

	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type ChamadaHandler struct {
	commons.GenericHandler[entity.Chamada]
}

func NewChamadaHandler(genericHandler commons.GenericHandler[entity.Chamada]) *ChamadaHandler {
	return &ChamadaHandler{
		genericHandler,
	}
}

func (h *ChamadaHandler) GetChamadaByID(ctx *fiber.Ctx, service services.ChamadaService) error {
	chaID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	result, err := service.GetChamadaByID(chaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (h *ChamadaHandler) GetAllDetChamadaFromChamadaId(ctx *fiber.Ctx, service services.ChamadaService) error {
	chaID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	result, err := service.GetAllDetChamadaFromChamadaId(chaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (h *ChamadaHandler) GetTokenFromChamadaId(ctx *fiber.Ctx, service services.ChamadaService) error {
	chaID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	result, err := service.GetTokenFromChamadaId(chaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (h *ChamadaHandler) UpdatePresencaAluno(ctx *fiber.Ctx, service services.ChamadaService) error {
	aluRA, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	_, err = service.UpdatePresencaAluno(aluRA)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("Erro")
	}

	return ctx.Status(fiber.StatusOK).JSON("Success")
}

func (h *ChamadaHandler) UpdateTokenInChamada(ctx *fiber.Ctx, service services.ChamadaService) error {
	token := ctx.Query("token")
	if token == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Token is required"})
	}

	chaID, err := strconv.Atoi(ctx.Query("chaID"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ChaID"})
	}

	resource, err := service.UpdateTokenInChamada(token, chaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Success", "resource": resource})
}

func (h *ChamadaHandler) UpdateChamadaTrue(ctx *fiber.Ctx, service services.ChamadaService) error {

	chaID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	resource, err := service.UpdateChamadaTrue(chaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Success", "resource": resource})
}

func (h *ChamadaHandler) UpdateChamadaFalse(ctx *fiber.Ctx, service services.ChamadaService) error {
	chaID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	resource, err := service.UpdateChamadaFalse(chaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Success", "resource": resource})
}
