package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/employee_requests"
	api_structure "main/internal/pkg/structures/employee_requests"

	"github.com/gofiber/fiber/v2"
)

type EmployeeDevelopmentController struct {
	Svc api_service.EmployeeDevelopmentService
}

func (controller *EmployeeDevelopmentController) GetEmployeeDevelopment(c *fiber.Ctx) error {
	var filter api_structure.EmployeeDevelopment
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// EmployeeDevelopmentID := runtime_tools.ParseInterfaceToURID(c.Locals("EmployeeDevelopmentID"))

	// filter.EmployeeDevelopmentId = EmployeeDevelopmentID

	result, rerr := controller.Svc.GetEmployeeDevelopment(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *EmployeeDevelopmentController) UpdateEmployeeDevelopment(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.EmployeeDevelopment{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateEmployeeDevelopment(*id, editData)
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

func (controller *EmployeeDevelopmentController) CreateEmployeeDevelopment(c *fiber.Ctx) error {
	data := api_structure.EmployeeDevelopment{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// EmployeeDevelopmentID := runtime_tools.ParseInterfaceToURID(c.Locals("EmployeeDevelopmentID"))

	// data.EmployeeDevelopmentId = EmployeeDevelopmentID

	result, rerr := controller.Svc.CreateEmployeeDevelopment(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *EmployeeDevelopmentController) DeleteEmployeeDevelopment(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteEmployeeDevelopment(*id)

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
