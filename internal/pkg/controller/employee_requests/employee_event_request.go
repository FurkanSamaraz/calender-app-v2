package controller

import (
	"encoding/json"
	api_service "main/services/employee_requests"
	api_structure "main/structures/employee_requests"

	"github.com/gofiber/fiber/v2"
)

type EmployeeEventRequestController struct {
	Svc api_service.EmployeeEventRequestService
}

func (controller *EmployeeEventRequestController) GetEmployeeEventRequest(c *fiber.Ctx) error {
	var filter api_structure.EmployeeEventRequest
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// EmployeeEventRequestID := runtime_tools.ParseInterfaceToURID(c.Locals("EmployeeEventRequestID"))

	// filter.EmployeeEventRequestId = EmployeeEventRequestID

	result, rerr := controller.Svc.GetEmployeeEventRequest(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *EmployeeEventRequestController) UpdateEmployeeEventRequest(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.EmployeeEventRequest{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateEmployeeEventRequest(*id, editData)
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

func (controller *EmployeeEventRequestController) CreateEmployeeEventRequest(c *fiber.Ctx) error {
	data := api_structure.EmployeeEventRequest{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// EmployeeEventRequestID := runtime_tools.ParseInterfaceToURID(c.Locals("EmployeeEventRequestID"))

	// data.EmployeeEventRequestId = EmployeeEventRequestID

	result, rerr := controller.Svc.CreateEmployeeEventRequest(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *EmployeeEventRequestController) DeleteEmployeeEventRequest(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteEmployeeEventRequest(*id)

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
