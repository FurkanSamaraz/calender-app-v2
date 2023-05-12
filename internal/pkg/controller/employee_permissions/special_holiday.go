package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/employee_permissions"
	api_structure "main/internal/pkg/structures/employee_permissions"

	"github.com/gofiber/fiber/v2"
)

type SpecialHolidayController struct {
	Svc api_service.SpecialHolidayService
}

func (controller *SpecialHolidayController) GetSpecialHoliday(c *fiber.Ctx) error {
	var filter api_structure.SpecialHoliday
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// SpecialHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("SpecialHolidayID"))

	// filter.SpecialHolidayId = SpecialHolidayID

	result, rerr := controller.Svc.GetSpecialHoliday(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *SpecialHolidayController) UpdateSpecialHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.SpecialHoliday{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateSpecialHoliday(*id, editData)
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

func (controller *SpecialHolidayController) CreateSpecialHoliday(c *fiber.Ctx) error {
	data := api_structure.SpecialHoliday{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// SpecialHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("SpecialHolidayID"))

	// data.SpecialHolidayId = SpecialHolidayID

	result, rerr := controller.Svc.CreateSpecialHoliday(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *SpecialHolidayController) DeleteSpecialHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteSpecialHoliday(*id)

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
