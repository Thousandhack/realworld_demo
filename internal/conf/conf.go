package conf

import "github.com/google/wire"

// ProviderSet is conf providers.
var ProviderSet = wire.NewSet(NewJWT)

// NewJWT 创建一个新的JWT配置
func NewJWT() *JWT {
	return &JWT{
		Secret: "realworld_demo_secret_key",
	}
}
