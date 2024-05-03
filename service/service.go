package service

import (
	"example-middleware/model"
	"example-middleware/repository"
)

type IService interface {
	Get(role string) []model.Data
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

func (s *Service) Get(role string) []model.Data {
	// business logic

	// get data dulu, dari repository
	data := s.Repository.Get()

	// logic
	var finalData []model.Data

	for _, dat := range data {
		if role == "admin" {
			finalData = append(finalData, dat)
		} else {
			// logic untuk mengubah
			// 0813213123132
			// 081xxxxxxxxxx

			var securePhoneNumber string

			for i := 0; i < len(dat.PhoneNumber); i++ {
				if i <= 2 {
					securePhoneNumber += string(dat.PhoneNumber[i])
				} else {
					securePhoneNumber += "x"
				}
			}

			dat.PhoneNumber = securePhoneNumber
			finalData = append(finalData, dat)
		}
	}

	return finalData
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
