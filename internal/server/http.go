package server

import (
	"context"
	v1 "realworld_demo/api/realworld/v1"
	"realworld_demo/internal/conf"
	auth "realworld_demo/internal/pkg/middleware"
	"realworld_demo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/realworld.v1.RealWorld/Login":        {},
		"/realworld.v1.RealWorld/Register":     {},
		"/realworld.v1.RealWorld/GetArticle":   {},
		"/realworld.v1.RealWorld/ListArticles": {},
		"/realworld.v1.RealWorld/GetComments":  {},
		"/realworld.v1.RealWorld/GetTags":      {},
		"/realworld.v1.RealWorld/GetProfile":   {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, s *service.RealWorldService, logger log.Logger) *http.Server {
	// 添加HTTP请求日志中间件
	httpLogger := log.NewHelper(logger)

	// 日志中间件
	logMiddleware := func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				httpLogger.Infof("收到HTTP请求: %s %s", tr.Kind(), tr.Operation())
			}
			return handler(ctx, req)
		}
	}

	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),

		// 添加自定义日志中间件
		http.Middleware(logMiddleware),

		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewSkipRoutersMatcher()).Build(),
			logging.Server(logger),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// 添加一个简单的调试路由
	srv.Route("/").GET("/debug/health", func(ctx http.Context) error {
		return ctx.JSON(200, map[string]interface{}{
			"status":  "ok",
			"message": "服务器正常运行",
		})
	})

	v1.RegisterRealWorldHTTPServer(srv, s)
	return srv
}
