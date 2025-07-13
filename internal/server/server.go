package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers. 依赖注入的集合
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
