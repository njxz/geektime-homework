package server

import (
	dv1 "geektime/api/mydemo/v1"
	"geektime/internal/conf"
	"geektime/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.MyDemoServer, mydemo *service.Mydemo) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(),
			validate.Validator(),
		),
	}
	if c.Grpc != "" {
		opts = append(opts, grpc.Address(c.Grpc))
	}
	srv := grpc.NewServer(opts...)
	dv1.RegisterHelloServer(srv, mydemo)
	return srv
}
