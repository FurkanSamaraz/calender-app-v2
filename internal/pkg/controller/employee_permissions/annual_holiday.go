package controller

import (
	"encoding/json"
	api_service "main/services/employee_permissions"
	api_structure "main/structures/employee_permissions"

	"github.com/gofiber/fiber/v2"
)

type AnnualHolidayController struct {
	Svc api_service.AnnualHolidayService
}

func (controller *AnnualHolidayController) GetAnnualHoliday(c *fiber.Ctx) error {
	var filter api_structure.AnnualHoliday
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// AnnualHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("AnnualHolidayID"))

	// filter.AnnualHolidayId = AnnualHolidayID

	result, rerr := controller.Svc.GetAnnualHoliday(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *AnnualHolidayController) UpdateAnnualHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.AnnualHoliday{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateAnnualHoliday(*id, editData)
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

func (controller *AnnualHolidayController) CreateAnnualHoliday(c *fiber.Ctx) error {
	data := api_structure.AnnualHoliday{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// AnnualHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("AnnualHolidayID"))

	// data.AnnualHolidayId = AnnualHolidayID

	result, rerr := controller.Svc.CreateAnnualHoliday(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *AnnualHolidayController) DeleteAnnualHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteAnnualHoliday(*id)

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
