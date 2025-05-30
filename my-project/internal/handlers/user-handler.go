package handlers

import (
	"github.com/my-project/internal/services/home"
	"github.com/my-project/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *userservice.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService *userservice.UserService, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validate,
	}
}

func (h *UserHandler) GetListUsers(ctx *fiber.Ctx) error {
	result, err := h.userService.GetListUsers(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(utils.NewResponse(result, "succes get data", fiber.StatusOK))
}
