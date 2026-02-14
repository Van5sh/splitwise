package handler

import (
	"errors"

	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserDetailsHandler struct {
	services *services.UserDetailsService
}

func NewUserDetailsHandler(services *services.UserDetailsService) *UserDetailsHandler {
	return &UserDetailsHandler{
		services: services,
	}
}

func (h *UserDetailsHandler) GetUserDetails(c *fiber.Ctx) error {
	userId := c.Params("id")
	uid, err := helpers.ValidateID(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			"Invalid user id",
		)
	}
	user, err := h.services.GetUserDetailsByUserID(c.Context(), uid.String())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":   user,
		"status": "success",
	})
}

func (h *UserDetailsHandler) GetUserDetailsByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return helpers.ValidationError(
			"Email is required",
			nil,
			errors.New("email is required"),
		)
	}
	user, err := h.services.GetUserDetailsByEmail(c.Context(), email)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":   user,
		"status": "success",
	})
}

func (h *UserDetailsHandler) UpdateUserDetails(c *fiber.Ctx) error {
	userId := c.FormValue("id")
	email := c.FormValue("email")
	uid, err := helpers.ValidateID(userId)
	if err != nil {
		return helpers.ValidationError("Wrong Input", nil, &fiber.Error{})
	}
	user, err := h.services.UpdateUserDetails(c.Context(), uid.String(), email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			"Error updating user details",
		)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":   user,
		"status": "success",
	})
}

func (h *UserDetailsHandler) CheckUserInGroup(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	groupId := c.Params("group_id")
	uid, err := helpers.ValidateID(userId)
	if err != nil {
		return helpers.ValidationError("Wrong Input", nil, &fiber.Error{})
	}
	gid, err := helpers.ValidateID(groupId)
	if err != nil {
		return helpers.ValidationError("Wrong Input", nil, &fiber.Error{})
	}
	res, err := h.services.CheckUserInGroupParams(c.Context(), uid.String(), gid.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			"Error checking user in group",
		)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"is_member": res,
		"status":    "success",
	})
}
