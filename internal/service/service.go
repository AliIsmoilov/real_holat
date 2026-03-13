package service

import (
	"real-holat/storage"
)

type ServiceI interface {
	User() UserServiceI
	InfrastructureType() InfrastructureTypeServiceI
}

type service struct {
	userSvc               UserServiceI
	infrastructureTypeSvc InfrastructureTypeServiceI
}

func New(strg storage.StorageI) ServiceI {
	return &service{
		userSvc:               NewUserService(strg),
		infrastructureTypeSvc: NewInfrastructureTypeService(strg),
	}
}

func (s *service) User() UserServiceI {
	return s.userSvc
}

func (s *service) InfrastructureType() InfrastructureTypeServiceI {
	return s.infrastructureTypeSvc
}
