package service

import (
	"example-middleware/model"
	"example-middleware/repository"
)

type IService interface {
	Get() []model.Data
	GetById(id int) model.Data
	Update(model.Data) error
	Delete(id int) error
	Create(model.Data) error
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) IService {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) Get() []model.Data {
	return s.Repository.Get()
}

func (s *Service) GetById(id int) model.Data {
	return s.Repository.GetById(id)
}

func (s *Service) Update(data model.Data) error {
	return s.Repository.Update(data)
}

func (s *Service) Delete(id int) error {
	return s.Repository.Delete(id)
}
func (s *Service) Create(data model.Data) error {
	return s.Repository.Create(data)
}
