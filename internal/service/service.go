package service

import (
	"real-holat/storage"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ServiceI interface {
	User() UserServiceI
	InfrastructureType() InfrastructureTypeServiceI
	Infrastructure() InfrastructureServiceI
	Verification() VerificationServiceI
	Report() ReportServiceI
	R2() R2ServiceI
}

type service struct {
	userSvc               UserServiceI
	infrastructureTypeSvc InfrastructureTypeServiceI
	infrastructureSvc     InfrastructureServiceI
	verificationSvc       VerificationServiceI
	reportSvc             ReportServiceI
	r2                    R2ServiceI
}

func New(strg storage.StorageI, r2client *s3.Client) ServiceI {
	var r2svc R2ServiceI
	if r2client != nil {
		r2svc = NewR2Service(r2client)
	} else {
		panic("R2 client not initialized")
	}

	return &service{
		userSvc:               NewUserService(strg),
		infrastructureTypeSvc: NewInfrastructureTypeService(strg),
		infrastructureSvc:     NewInfrastructureService(strg),
		verificationSvc:       NewVerificationService(strg),
		reportSvc:             NewReportService(strg),
		r2:                    r2svc,
	}
}

func (s *service) User() UserServiceI {
	return s.userSvc
}

func (s *service) InfrastructureType() InfrastructureTypeServiceI {
	return s.infrastructureTypeSvc
}

func (s *service) Infrastructure() InfrastructureServiceI {
	return s.infrastructureSvc
}

func (s *service) Verification() VerificationServiceI {
	return s.verificationSvc
}

func (s *service) Report() ReportServiceI {
	return s.reportSvc
}

func (s *service) R2() R2ServiceI {
	return s.r2
}
