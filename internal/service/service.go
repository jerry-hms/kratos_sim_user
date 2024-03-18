package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos_sim/api/user/service/v1"
	"kratos_sim/app/user/service/internal/biz"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UserServiceServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(log.With(logger, "module", "server/server"))}
}
