package handler

import (
	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/Van5sh/new-splitwise/internal/services"
	"github.com/gofiber/fiber/v2"
)

type GroupHandler struct {
	services *services.GroupServices
}

func NewGroupHandler(services *services.GroupServices) *GroupHandler {
	return &GroupHandler{services: services}
}

func (h *GroupHandler) GetGroups(c *fiber.Ctx) error {
	groups, err := h.services.GetGroups(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   groups,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) GetGroupById(c *fiber.Ctx) error {
	id := c.Params("id")
	valid_id, err := helpers.ValidateID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid group id",
		})
	}
	group, err := h.services.GetGroupById(c.Context(), valid_id.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   group,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) GetGroupByName(c *fiber.Ctx) error {
	groupName := c.Query("group_name")
	group, err := h.services.GetGroupByName(c.Context(), groupName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   group,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) CreateGroup(c *fiber.Ctx) error {
	var request struct {
		GroupName   string `json:"groupName"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	group, err := h.services.CreateGroup(c.Context(), request.GroupName, request.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   group,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) GetGroupsByUserId(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	groups, err := h.services.GetGroupsByUserId(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   groups,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) GetGroupMembers(c *fiber.Ctx) error {
	groupId := c.Params("group_id")
	members, err := h.services.GetGroupMembers(c.Context(), groupId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   members,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) GetGroupAdmins(c *fiber.Ctx) error {
	groupId := c.Params("group_id")
	admins, err := h.services.GetGroupAdmins(c.Context(), groupId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   admins,
			"code":   fiber.StatusOK,
		},
	)
}

func (h *GroupHandler) DeleteGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.services.DeleteGroup(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   "group deleted successfully",
			"code":   fiber.StatusOK,
		},
	)
}
func (h *GroupHandler) UpdateGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var request struct {
		GroupName   string `json:"groupName"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	group, err := h.services.UpdateGroup(c.Context(), id, request.GroupName, request.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "success",
			"data":   group,
			"code":   fiber.StatusOK,
		},
	)
}
