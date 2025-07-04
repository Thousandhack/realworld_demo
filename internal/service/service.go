package service

import (
	"github.com/google/wire"

	v1 "realworld_demo/api/realworld/v1"
	"realworld_demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUsecase
	sc  *biz.SocialUsecase
	log *log.Helper
}

func NewRealWorldService(uc *biz.UserUsecase, sc *biz.SocialUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, sc: sc, log: log.NewHelper(logger)}
}
