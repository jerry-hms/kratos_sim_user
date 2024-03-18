package server

import (
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"go.opentelemetry.io/otel/sdk/trace"
	"kratos_sim/api/user/service/v1"
	"kratos_sim/app/user/service/internal/conf"
	"kratos_sim/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, u *service.UserService, ac *conf.Auth, tp *trace.TracerProvider, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.Key), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServiceServer(srv, u)
	return srv
}
