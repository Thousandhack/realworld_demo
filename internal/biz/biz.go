package biz

import "github.com/google/wire"

// ProviderSet is biz providers. 依赖注入的集合
var ProviderSet = wire.NewSet(NewSocialUsecase, NewUserUsecase)
