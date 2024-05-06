package handler

import (
	"Test-Golang-ITMX/repository"
	"Test-Golang-ITMX/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func APICustomersHandler(c fiber.Router, db *gorm.DB) {
	customersRepository := repository.NewCustomersRepository(db)
	customersService := service.NewCustomersService(customersRepository)
	customersHandler := NewCustomersHandler(customersService)

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

func NewCustomersHandler(service service.CustomersServiceInterface) CustomersHandlerInterface {
	return handler{
		service: service,
	}
}
