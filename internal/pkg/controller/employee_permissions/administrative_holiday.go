package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/employee_permissions"
	api_structure "main/internal/pkg/structures/employee_permissions"

	"github.com/gofiber/fiber/v2"
)

type AdministrativeHolidayController struct {
	Svc api_service.AdministrativeHolidayService
}

func (controller *AdministrativeHolidayController) GetAdministrativeHoliday(c *fiber.Ctx) error {
	var filter api_structure.AdministrativeHoliday
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// AdministrativeHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("AdministrativeHolidayID"))

	// filter.AdministrativeHolidayId = AdministrativeHolidayID

	result, rerr := controller.Svc.GetAdministrativeHoliday(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *AdministrativeHolidayController) UpdateAdministrativeHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.AdministrativeHoliday{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateAdministrativeHoliday(*id, editData)
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

func (controller *AdministrativeHolidayController) CreateAdministrativeHoliday(c *fiber.Ctx) error {
	data := api_structure.AdministrativeHoliday{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// AdministrativeHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("AdministrativeHolidayID"))

	// data.AdministrativeHolidayId = AdministrativeHolidayID

	result, rerr := controller.Svc.CreateAdministrativeHoliday(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *AdministrativeHolidayController) DeleteAdministrativeHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteAdministrativeHoliday(*id)

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
