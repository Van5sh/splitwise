package handler

import (
	//	"github.com/Van5sh/new-splitwise/internal/helpers/response"

	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	services *services.UserServices
}

func NewUserHandler(services *services.UserServices) *UserHandler {
	return &UserHandler{services: services}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.services.GetUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   users,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *UserHandler) GetUserId(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.services.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   user,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	userName := c.FormValue("user_name")
	FirebaseID := c.FormValue("firebase_id")
	role := c.FormValue("role")
	email := c.FormValue("email")
	res, err := h.services.CreateUser(c.Context(), userName, FirebaseID, role, email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   res,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *UserHandler) GetUserByFirebaseID(c *fiber.Ctx) error {
	firebaseID := c.Params("firebase_id")
	if firebaseID == "" {
		return c.Status(fiber.ErrNotFound.Code).JSON(
			fiber.Map{
				"error": "Firebase ID is required",
			},
		)
	}

	user, err := h.services.GetUserByFirebaseID(c.Context(), firebaseID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   user,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *UserHandler) UpdateUserRole(c *fiber.Ctx) error {
	userId := c.Params("id")
	role := c.Params("role")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}
	res, err := h.services.UpdateUserRole(c.Context(), userId, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   res,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == " " {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}
	err := h.services.DeleteUser(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"code":   fiber.StatusNoContent,
		},
	)
}
