package controller

import (
	"encoding/json"
	api_service "main/services/employee_permissions"
	api_structure "main/structures/employee_permissions"

	"github.com/gofiber/fiber/v2"
)

type ExcuseHolidayController struct {
	Svc api_service.ExcuseHolidayService
}

func (controller *ExcuseHolidayController) GetExcuseHoliday(c *fiber.Ctx) error {
	var filter api_structure.ExcuseHoliday
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// ExcuseHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("ExcuseHolidayID"))

	// filter.ExcuseHolidayId = ExcuseHolidayID

	result, rerr := controller.Svc.GetExcuseHoliday(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *ExcuseHolidayController) UpdateExcuseHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.ExcuseHoliday{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateExcuseHoliday(*id, editData)
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

func (controller *ExcuseHolidayController) CreateExcuseHoliday(c *fiber.Ctx) error {
	data := api_structure.ExcuseHoliday{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// ExcuseHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("ExcuseHolidayID"))

	// data.ExcuseHolidayId = ExcuseHolidayID

	result, rerr := controller.Svc.CreateExcuseHoliday(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *ExcuseHolidayController) DeleteExcuseHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteExcuseHoliday(*id)

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
