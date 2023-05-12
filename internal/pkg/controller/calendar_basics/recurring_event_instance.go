package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/calendar_basics"
	api_structure "main/internal/pkg/structures/calendar_basics"

	"github.com/gofiber/fiber/v2"
)

type RecurringEventInstanceController struct {
	Svc api_service.RecurringEventInstanceService
}

func (controller *RecurringEventInstanceController) GetRecurringEventInstance(c *fiber.Ctx) error {
	var filter api_structure.RecurringEventInstance
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// RecurringEventInstanceID := runtime_tools.ParseInterfaceToURID(c.Locals("RecurringEventInstanceID"))

	// filter.RecurringEventInstanceId = RecurringEventInstanceID

	result, rerr := controller.Svc.GetRecurringEventInstance(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *RecurringEventInstanceController) UpdateRecurringEventInstance(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.RecurringEventInstance{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateRecurringEventInstance(*id, editData)
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

func (controller *RecurringEventInstanceController) CreateRecurringEventInstance(c *fiber.Ctx) error {
	data := api_structure.RecurringEventInstance{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// RecurringEventInstanceID := runtime_tools.ParseInterfaceToURID(c.Locals("RecurringEventInstanceID"))

	// data.RecurringEventInstanceId = RecurringEventInstanceID

	result, rerr := controller.Svc.CreateRecurringEventInstance(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *RecurringEventInstanceController) DeleteRecurringEventInstance(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteRecurringEventInstance(*id)

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
