package controller

import (
	"encoding/json"
	api_service "main/services/calendar_basics"
	api_structure "main/structures/calendar_basics"

	"github.com/gofiber/fiber/v2"
)

type RecurringEventController struct {
	Svc api_service.RecurringEventService
}

func (controller *RecurringEventController) GetRecurringEvent(c *fiber.Ctx) error {
	var filter api_structure.RecurringEvent
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// RecurringEventID := runtime_tools.ParseInterfaceToURID(c.Locals("RecurringEventID"))

	// filter.RecurringEventId = RecurringEventID

	result, rerr := controller.Svc.GetRecurringEvent(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *RecurringEventController) UpdateRecurringEvent(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.RecurringEvent{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateRecurringEvent(*id, editData)
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

func (controller *RecurringEventController) CreateRecurringEvent(c *fiber.Ctx) error {
	data := api_structure.RecurringEvent{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// RecurringEventID := runtime_tools.ParseInterfaceToURID(c.Locals("RecurringEventID"))

	// data.RecurringEventId = RecurringEventID

	result, rerr := controller.Svc.CreateRecurringEvent(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *RecurringEventController) DeleteRecurringEvent(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteRecurringEvent(*id)

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
