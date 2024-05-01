package handler

import (
	"Test-Golang-ITMX/model"
	"Test-Golang-ITMX/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h handler) CreateCustomers(c *fiber.Ctx) (Error error) {
	var httpResponse model.HttpResponse
	var request model.CustomersRequest
	if err := c.BodyParser(&request); err != nil {
		httpResponse.ErrorMsg = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	if err := utils.ValidateStruct(&request); err != nil {
		httpResponse.ErrorMsg = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	var data model.CustomersResponse
	data, Error = h.service.Create(request)
	if Error != nil {
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	httpResponse.Data = data
	return c.Status(fiber.StatusCreated).JSON(httpResponse)

}

func (h handler) UpdateCustomers(c *fiber.Ctx) (Error error) {
	var httpResponse model.HttpResponse
	var request model.CustomersRequest
	if err := c.BodyParser(&request); err != nil {
		httpResponse.ErrorMsg = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	if err := utils.ValidateStruct(&request); err != nil {
		httpResponse.ErrorMsg = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	strId := c.Params("id")
	if strId == "" {
		httpResponse.ErrorMsg = "Params invalid."
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	id, Error := strconv.Atoi(strId)
	if Error != nil {
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	var data model.CustomersResponse
	data, Error = h.service.Update(id, request)
	if Error != nil {
		var httpStatus = fiber.StatusBadRequest
		if Error.Error() == "record not found" {
			httpStatus = fiber.StatusNotFound
		}
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(httpStatus).JSON(httpResponse)
	}

	httpResponse.Data = data
	return c.Status(fiber.StatusCreated).JSON(httpResponse)
}

func (h handler) DeleteCustomers(c *fiber.Ctx) (Error error) {
	var httpResponse model.HttpResponse
	strId := c.Params("id")
	if strId == "" {
		httpResponse.ErrorMsg = "Params invalid."
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	id, Error := strconv.Atoi(strId)
	if Error != nil {
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}
	Error = h.service.Delete(id)
	if Error != nil {
		var httpStatus = fiber.StatusBadRequest
		if Error.Error() == "record not found" {
			httpStatus = fiber.StatusNotFound
		}
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(httpStatus).JSON(httpResponse)
	}

	httpResponse.Data = "Successful deletion"
	return c.Status(200).JSON(httpResponse)

}

func (h handler) GetCustomers(c *fiber.Ctx) (Error error) {
	var httpResponse model.HttpResponse
	var dataList []model.CustomersResponse
	dataList, Error = h.service.GetAll()
	if Error != nil {
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}
	httpResponse.Data = dataList
	return c.Status(200).JSON(httpResponse)
}

func (h handler) GetCustomersById(c *fiber.Ctx) (Error error) {
	var httpResponse model.HttpResponse
	strId := c.Params("id")
	if strId == "" {
		httpResponse.ErrorMsg = "Params invalid."
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}

	id, Error := strconv.Atoi(strId)
	if Error != nil {
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(fiber.StatusBadRequest).JSON(httpResponse)
	}
	var data model.CustomersResponse
	data, Error = h.service.GetById(id)
	if Error != nil {
		var httpStatus = fiber.StatusBadRequest
		if Error.Error() == "record not found" {
			httpStatus = fiber.StatusNotFound
		}
		httpResponse.ErrorMsg = Error.Error()
		return c.Status(httpStatus).JSON(httpResponse)
	}

	httpResponse.Data = data
	return c.Status(200).JSON(httpResponse)
}
