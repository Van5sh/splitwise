package handler

import (
	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserGroupHandler struct {
	service *services.UserGroupService
}

func NewUserGroupHandler(service *services.UserGroupService) *UserGroupHandler {
	return &UserGroupHandler{service: service}
}

func (h *UserGroupHandler) RemoveUserFromGroup(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	groupID := c.Params("group_id")

	uid, err := helpers.ValidateID(userID)
	if err != nil {
		appErr := helpers.ValidationError("invalid user id", nil, err)
		return c.Status(appErr.Status).JSON(appErr)
	}
	gid, err := helpers.ValidateID(groupID)
	if err != nil {
		appErr := helpers.ValidationError("invalid group id", nil, err)
		return c.Status(appErr.Status).JSON(appErr)
	}

	if err := h.service.RemoveUserFromGroup(
		c.Context(),
		uid.String(),
		gid.String(),
	); err != nil {
		appErr := helpers.InternalServerError(
			"failed to remove user from group",
			nil,
			err,
		)
		return c.Status(appErr.Status).JSON(appErr)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "user removed from group successfully",
		"userId":  userID,
		"groupId": groupID,
	})
}

func (h *UserGroupHandler) AddUserToGroup(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	groupID := c.Params("group_id")

	uid, err := helpers.ValidateID(userID)
	if err != nil {
		appErr := helpers.ValidationError("invalid user id", nil, err)
		return c.Status(appErr.Status).JSON(appErr)
	}

	gid, err := helpers.ValidateID(groupID)
	if err != nil {
		appErr := helpers.ValidationError("invalid group id", nil, err)
		return c.Status(appErr.Status).JSON(appErr)
	}

	if err := h.service.AddUserToGroup(
		c.Context(),
		uid.String(),
		gid.String(),
	); err != nil {
		appErr := helpers.InternalServerError(
			"failed to add user to group",
			nil,
			err,
		)
		return c.Status(appErr.Status).JSON(appErr)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "user added to group successfully",
		"userId":  userID,
		"groupId": groupID,
	})
}
