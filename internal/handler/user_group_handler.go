package handler

import (
	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserGroupHandler struct {
	services *services.UserGroupService
}

func NewUserGroupHandler(services *services.UserGroupService) *UserGroupHandler {
	return &UserGroupHandler{
		services: services,
	}
}

func (h *UserGroupHandler) RemoveUserFromGroup(c *fiber.Ctx) error {
	userId := c.Params("id")
	groupId := c.Params("groupId")

	uid, err := helpers.ValidateId(userId)
	if err != nil {
		return err
	}
	gid, err := helpers.ValidateId(groupId)
	if err != nil {
		return err
	}
	err = h.services.RemoveUserFromGroup(uid, gid)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserGroupHandler) AddUserToGroup(c *fiber.Ctx) error {
	userId := c.Params("id")
	groupId := c.Params("groupId")
	uid, err := helpers.ValidateId(userId)
	if err != nil {
		return err
	}
	gid, err := helpers.ValidateId(groupId)
	if err != nil {
		return err
	}
	err = h.services.AddUserToGroup(uid, gid)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
