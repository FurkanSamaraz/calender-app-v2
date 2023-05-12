package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/notifications"
	api_structure "main/internal/pkg/structures/notifications"

	"github.com/gofiber/fiber/v2"
)

type Send_EmailsController struct {
	Svc api_service.Send_EmailsService
}

func (controller *Send_EmailsController) GetSend_Emails(c *fiber.Ctx) error {
	var filter api_structure.Send_Emails
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// Send_EmailsID := runtime_tools.ParseInterfaceToURID(c.Locals("Send_EmailsID"))

	// filter.Send_EmailsId = Send_EmailsID

	result, rerr := controller.Svc.GetSend_Emails(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *Send_EmailsController) UpdateSend_Emails(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.Send_Emails{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateSend_Emails(*id, editData)
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

func (controller *Send_EmailsController) CreateSend_Emails(c *fiber.Ctx) error {
	data := api_structure.Send_Emails{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// Send_EmailsID := runtime_tools.ParseInterfaceToURID(c.Locals("Send_EmailsID"))

	// data.Send_EmailsId = Send_EmailsID

	result, rerr := controller.Svc.CreateSend_Emails(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *Send_EmailsController) DeleteSend_Emails(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteSend_Emails(*id)

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

func (controller *Send_EmailsController) Send_Email(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.Send_SendEmail(*id)
	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Send Email",
			"message": deleteErr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(deleteErr)
}
