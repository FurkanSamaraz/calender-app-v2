package controller

import (
	"encoding/json"
	api_service "main/services/calendar_basics"
	api_structure "main/structures/calendar_basics"

	"github.com/gofiber/fiber/v2"
)

type ImportantDayController struct {
	Svc api_service.ImportantDayService
}

func (controller *ImportantDayController) GetImportantDay(c *fiber.Ctx) error {
	var filter api_structure.ImportantDay
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// ImportantDayID := runtime_tools.ParseInterfaceToURID(c.Locals("ImportantDayID"))

	// filter.ImportantDayId = ImportantDayID

	result, rerr := controller.Svc.GetImportantDay(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *ImportantDayController) UpdateImportantDay(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.ImportantDay{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateImportantDay(*id, editData)
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

func (controller *ImportantDayController) CreateImportantDay(c *fiber.Ctx) error {
	data := api_structure.ImportantDay{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// ImportantDayID := runtime_tools.ParseInterfaceToURID(c.Locals("ImportantDayID"))

	// data.ImportantDayId = ImportantDayID

	result, rerr := controller.Svc.CreateImportantDay(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *ImportantDayController) DeleteImportantDay(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteImportantDay(*id)

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
