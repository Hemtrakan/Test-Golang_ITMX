package service

import "Test-Golang-ITMX/model"

func (s Service) GetById(id int) (res model.CustomersResponse, Error error) {
	var data model.Customers
	data, Error = s.repository.GetById(id)
	if Error != nil {
		return
	}
	res = s.MappingDataGetBy(data)
	return
}

func (s Service) MappingDataGetBy(data model.Customers) (res model.CustomersResponse) {
	res = model.CustomersResponse{
		Id:   data.Id,
		Name: data.Name,
		Age:  data.Age,
	}
	return
}

func (s Service) GetAll() (res []model.CustomersResponse, Error error) {
	var dataList []model.Customers
	dataList, Error = s.repository.GetAll()
	if Error != nil {
		return
	}
	res = s.MappingDataList(dataList)
	return
}

func (s Service) MappingDataList(dataList []model.Customers) (res []model.CustomersResponse) {
	for _, data := range dataList {
		res = append(res, model.CustomersResponse{
			Id:   data.Id,
			Name: data.Name,
			Age:  data.Age,
		})
	}
	return
}
func (s Service) Create(req model.CustomersRequest) (res model.CustomersResponse, Error error) {
	var data model.Customers
	data, Error = s.repository.Create(req)
	if Error != nil {
		return
	}
	res = s.MappingDataGetBy(data)
	return
}

func (s Service) Update(id int, req model.CustomersRequest) (res model.CustomersResponse, Error error) {
	var data model.Customers

	data, Error = s.repository.GetById(id)
	if Error != nil {
		return
	}

	data, Error = s.repository.Update(id, req)
	if Error != nil {
		return
	}
	res = s.MappingDataGetBy(data)
	return
}

func (s Service) Delete(id int) (Error error) {
	Error = s.repository.Delete(id)
	if Error != nil {
		return
	}
	return
}
