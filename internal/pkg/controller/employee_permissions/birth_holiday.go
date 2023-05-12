package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/employee_permissions"
	api_structure "main/internal/pkg/structures/employee_permissions"

	"github.com/gofiber/fiber/v2"
)

type BirthHolidayController struct {
	Svc api_service.BirthHolidayService
}

func (controller *BirthHolidayController) GetBirthHoliday(c *fiber.Ctx) error {
	var filter api_structure.BirthHoliday
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// BirthHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("BirthHolidayID"))

	// filter.BirthHolidayId = BirthHolidayID

	result, rerr := controller.Svc.GetBirthHoliday(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *BirthHolidayController) UpdateBirthHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.BirthHoliday{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateBirthHoliday(*id, editData)
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

func (controller *BirthHolidayController) CreateBirthHoliday(c *fiber.Ctx) error {
	data := api_structure.BirthHoliday{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// BirthHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("BirthHolidayID"))

	// data.BirthHolidayId = BirthHolidayID

	result, rerr := controller.Svc.CreateBirthHoliday(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *BirthHolidayController) DeleteBirthHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteBirthHoliday(*id)

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
