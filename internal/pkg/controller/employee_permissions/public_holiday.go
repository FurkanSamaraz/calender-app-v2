package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/employee_permissions"
	api_structure "main/internal/pkg/structures/employee_permissions"

	"github.com/gofiber/fiber/v2"
)

type PublicHolidayController struct {
	Svc api_service.PublicHolidayService
}

func (controller *PublicHolidayController) GetPublicHoliday(c *fiber.Ctx) error {
	var filter api_structure.PublicHoliday
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// PublicHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("PublicHolidayID"))

	// filter.PublicHolidayId = PublicHolidayID

	result, rerr := controller.Svc.GetPublicHoliday(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *PublicHolidayController) UpdatePublicHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.PublicHoliday{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdatePublicHoliday(*id, editData)
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

func (controller *PublicHolidayController) CreatePublicHoliday(c *fiber.Ctx) error {
	data := api_structure.PublicHoliday{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// PublicHolidayID := runtime_tools.ParseInterfaceToURID(c.Locals("PublicHolidayID"))

	// data.PublicHolidayId = PublicHolidayID

	result, rerr := controller.Svc.CreatePublicHoliday(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *PublicHolidayController) DeletePublicHoliday(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeletePublicHoliday(*id)

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
