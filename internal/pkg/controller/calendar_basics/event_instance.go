package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/calendar_basics"
	api_structure "main/internal/pkg/structures/calendar_basics"

	"github.com/gofiber/fiber/v2"
)

type EventInstanceController struct {
	Svc api_service.EventInstanceService
}

func (controller *EventInstanceController) GetEventInstance(c *fiber.Ctx) error {
	var filter api_structure.EventInstance
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// EventInstanceID := runtime_tools.ParseInterfaceToURID(c.Locals("EventInstanceID"))

	// filter.EventInstanceId = EventInstanceID

	result, rerr := controller.Svc.GetEventInstance(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *EventInstanceController) UpdateEventInstance(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.EventInstance{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateEventInstance(*id, editData)
	if uerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Update Data",
			"message": uerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"succes":  true,
		"message": "Updated Successfully",
		"type":    "Update Data",
	})
}

func (controller *EventInstanceController) CreateEventInstance(c *fiber.Ctx) error {
	data := api_structure.EventInstance{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// EventInstanceID := runtime_tools.ParseInterfaceToURID(c.Locals("EventInstanceID"))

	// data.EventInstanceId = EventInstanceID

	result, rerr := controller.Svc.CreateEventInstance(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *EventInstanceController) DeleteEventInstance(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteEventInstance(*id)

	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Delete Data",
			"message": deleteErr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted Successfully",
		"type":    "Delete Data",
		"success": true,
	})
}
