package services

import (
	"main/repository"
)

type Services struct{
	Repo *repository.Repository
}

func NewServices() *Services {
	return &Services{
		Repo : repository.NewRepository() ,
	}
}
