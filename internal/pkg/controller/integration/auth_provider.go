package controller

import (
	"encoding/json"
	api_service "main/internal/pkg/services/integration"
	api_structure "main/internal/pkg/structures/integration"

	"github.com/gofiber/fiber/v2"
)

type AuthProviderController struct {
	Svc api_service.AuthProviderService
}

func (controller *AuthProviderController) GetAuthProvider(c *fiber.Ctx) error {
	var filter api_structure.AuthProvider
	// if err := runtime_tools.ParseQuery(c, &filter); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"type":    "queryParser",
	// 		"message": err.Error(),
	// 	})
	// }
	// //PRE Lines Registered
	// AuthProviderID := runtime_tools.ParseInterfaceToURID(c.Locals("AuthProviderID"))

	// filter.AuthProviderId = AuthProviderID

	result, rerr := controller.Svc.GetAuthProvider(filter)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *AuthProviderController) UpdateAuthProvider(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	editData := api_structure.AuthProvider{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateAuthProvider(*id, editData)
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

func (controller *AuthProviderController) CreateAuthProvider(c *fiber.Ctx) error {
	data := api_structure.AuthProvider{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	// AuthProviderID := runtime_tools.ParseInterfaceToURID(c.Locals("AuthProviderID"))

	// data.AuthProviderId = AuthProviderID

	result, rerr := controller.Svc.CreateAuthProvider(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (controller *AuthProviderController) DeleteAuthProvider(c *fiber.Ctx) error {

	var id *int
	json.Unmarshal([]byte(c.Params("id")), id)

	deleteErr := controller.Svc.DeleteAuthProvider(*id)

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
