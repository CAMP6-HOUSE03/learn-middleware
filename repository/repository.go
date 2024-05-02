package repository

import (
	"errors"
	"example-middleware/model"
)

// gunakan untuk mengolah data
type IRepository interface {
	Get() []model.Data
	GetById(id int) model.Data
	Update(model.Data) error
	Delete(id int) error
	Create(model.Data) error
}

// in memmory
type repository struct {
	data map[int]model.Data
}

func NewRepository() IRepository {
	return &repository{
		data: make(map[int]model.Data),
	}
}

func (r *repository) Get() []model.Data {
	var datas []model.Data

	for _, data := range r.data {
		datas = append(datas, data)
	}

	return datas
}

func (r *repository) GetById(id int) model.Data {
	return r.data[id]
}

func (r *repository) Update(newData model.Data) error {
	dataById := r.data[newData.ID]

	if dataById.ID == 0 {
		return errors.New("data not found")
	}

	dataById.Name = newData.Name
	dataById.Address = newData.Address
	dataById.PhoneNumber = newData.PhoneNumber
	dataById.Email = newData.Email

	r.data[newData.ID] = dataById
	return nil
}

func (r *repository) Create(data model.Data) error {
	if r.data[data.ID].ID != 0 {
		return errors.New("data already exist")
	}

	r.data[data.ID] = data
	return nil
}

func (r *repository) Delete(id int) error {
	delete(r.data, id)
	return nil
}
