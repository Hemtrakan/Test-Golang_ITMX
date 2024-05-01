package handler

import (
	"Test-Golang-ITMX/access"
	"Test-Golang-ITMX/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func APICustomersHandler(c fiber.Router, db *gorm.DB) {
	customersAccess := access.CustomersAccess(db)
	customersService := service.CustomersService(customersAccess)
	customersHandler := CustomersHandler(customersService)

	c.Post("/customers", customersHandler.CreateCustomers)
	c.Put("/customers/:id", customersHandler.UpdateCustomers)
	c.Delete("/customers/:id", customersHandler.DeleteCustomers)
	c.Get("/customers", customersHandler.GetCustomers)
	c.Get("/customers/:id", customersHandler.GetCustomersById)
}

type CustomersHandlerInterface interface {
	CreateCustomers(c *fiber.Ctx) (Error error)
	UpdateCustomers(c *fiber.Ctx) (Error error)
	DeleteCustomers(c *fiber.Ctx) (Error error)
	GetCustomers(c *fiber.Ctx) (Error error)
	GetCustomersById(c *fiber.Ctx) (Error error)
}

type handler struct {
	service service.CustomersServiceInterface
}

func CustomersHandler(service service.CustomersServiceInterface) CustomersHandlerInterface {
	return handler{
		service: service,
	}
}
