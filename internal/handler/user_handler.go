package handler

import (
	"go-user-api/internal/models"
	"go-user-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {

	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(
	c *fiber.Ctx,
) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Invalid request body",
			})
	}

	if err := h.service.CreateUser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"message": "User created successfully",
		})
}

func (h *UserHandler) GetUser(
	c *fiber.Ctx,
) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Invalid ID",
			})
	}

	user, err := h.service.GetUser(
		int32(id),
	)

	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"error": "User not found",
			})
	}

	return c.JSON(user)
}

func (h *UserHandler) ListUsers(
	c *fiber.Ctx,
) error {

	users, err := h.service.ListUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.JSON(users)
}

func (h *UserHandler) DeleteUser(
	c *fiber.Ctx,
) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Invalid ID",
			})
	}

	err = h.service.DeleteUser(
		int32(id),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.SendStatus(
		fiber.StatusNoContent,
	)
}

func (h *UserHandler) UpdateUser(
	c *fiber.Ctx,
) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Invalid ID",
			})
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Invalid request body",
			})
	}

	err = h.service.UpdateUser(
		int32(id),
		req,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}