package v1

import (
	"real-holat/config"
	"real-holat/internal/service"
)

type handlerV1 struct {
	cfg     *config.Config
	service service.ServiceI
}

type HandleV1 struct {
	Cfg     *config.Config
	Service service.ServiceI
}

func New(h *HandleV1) *handlerV1 {
	return &handlerV1{
		cfg:     h.Cfg,
		service: h.Service,
	}
}
