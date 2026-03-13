package service

import (
	"real-holat/storage"
)

type ServiceI interface {
	User() UserServiceI
	InfrastructureType() InfrastructureTypeServiceI
	Verification() VerificationServiceI
}

type service struct {
	userSvc               UserServiceI
	infrastructureTypeSvc InfrastructureTypeServiceI
	verificationSvc       VerificationServiceI
}

func New(strg storage.StorageI) ServiceI {
	return &service{
		userSvc:               NewUserService(strg),
		infrastructureTypeSvc: NewInfrastructureTypeService(strg),
		verificationSvc:       NewVerificationService(strg),
	}
}

func (s *service) User() UserServiceI {
	return s.userSvc
}

func (s *service) InfrastructureType() InfrastructureTypeServiceI {
	return s.infrastructureTypeSvc
}

func (s *service) Verification() VerificationServiceI {
	return s.verificationSvc
}
