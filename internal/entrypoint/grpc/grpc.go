package grpc

import (
	"context"
	"net"

	goerrors "errors"

	// needed for interact with gRPC service trought Insomnia or Postman
	_ "github.com/jnewmano/grpc-json-proxy/codec"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	pb "github.com/maxfer4maxfer/service-template/internal/entrypoint/grpc/pb"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

var (
	errServiceUnknownError = errors.New(
		"unexpected error while calling the service")
)

// Config holds configuration for a GRPC server.
type Config struct {
	Address string
}

// Service declares interactions with the errors service.
type Service interface {
	Subtract(ctx context.Context, a int, b int) (sub int, err error)
	Multiply(ctx context.Context, a int, b int) (mult int, err error)
	Pi(ctx context.Context, count int) (string, error)
}

// ServerGRPC is a wraper of http.Server.
type ServerGRPC struct {
	logger  *zerolog.Logger
	cfg     Config
	service Service
	server  *grpc.Server
}

// New returns a GRPC server.
func New(logger *zerolog.Logger, svc Service, cfg Config) *ServerGRPC {
	s := &ServerGRPC{
		logger:  logger,
		cfg:     cfg,
		service: svc,
	}

	return s
}

// Start starts GRPC Server.
func (s *ServerGRPC) Start() chan error {
	serverErrors := make(chan error, 1)

	lis, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		s.logger.Fatal().Err(err).Msgf("failed to listen: %v", s.cfg.Address)
	}

	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			s.UnaryServerInterceptorCorrelationID(),
			s.UnaryServerInterceptorLogging(),
		)),
	)

	pb.RegisterCalculatorServer(s.server, s)

	go func() {
		s.logger.Info().Msgf("Start GRPC Listening %s", s.cfg.Address)
		serverErrors <- s.server.Serve(lis)
	}()

	return serverErrors
}

// Shutdown stops GRPC Server.
func (s *ServerGRPC) Shutdown(ctx context.Context) {
	s.server.GracefulStop()
}

// Multiply multiplies too numbers.
func (s *ServerGRPC) Multiply(
	ctx context.Context, in *pb.MultiplyRequest,
) (
	out *pb.MultiplyReply, pbErr error,
) {
	out = &pb.MultiplyReply{}

	handleError := func(
		err error, out *pb.MultiplyReply,
	) (
		*pb.MultiplyReply, error,
	) {
		var serr errors.ServiceError

		if !goerrors.As(err, &serr) {
			serr = errServiceUnknownError.Wrap(err)
		}

		out.Code = int32(serr.Code())
		out.Error = serr.Error()

		return out, pbErr
	}

	mult, err := s.service.Multiply(ctx, int(in.GetA()), int(in.GetB()))
	if err != nil {
		return handleError(err, out)
	}

	res := &pb.MultiplyReply{
		Mult: int32(mult),
	}

	return res, err
}

// Subtract multiplies too numbers.
func (s *ServerGRPC) Subtract(
	ctx context.Context, in *pb.SubtractRequest,
) (
	out *pb.SubtractReply, pbErr error,
) {
	out = &pb.SubtractReply{}

	handleError := func(
		err error, out *pb.SubtractReply,
	) (
		*pb.SubtractReply, error,
	) {
		var serr errors.ServiceError

		if !goerrors.As(err, &serr) {
			serr = errServiceUnknownError.Wrap(err)
		}

		out.Code = int32(serr.Code())
		out.Error = serr.Error()

		return out, pbErr
	}

	sub, err := s.service.Subtract(ctx, int(in.GetA()), int(in.GetB()))
	if err != nil {
		return handleError(err, out)
	}

	res := &pb.SubtractReply{
		Sub: int32(sub),
	}

	return res, err
}

// Pi execute a pi method.
func (s *ServerGRPC) Pi(
	ctx context.Context, in *pb.PiRequest,
) (
	out *pb.PiReply, pbErr error,
) {
	out = &pb.PiReply{}

	handleError := func(
		err error, out *pb.PiReply,
	) (
		*pb.PiReply, error,
	) {
		var serr errors.ServiceError

		if !goerrors.As(err, &serr) {
			serr = errServiceUnknownError.Wrap(err)
		}

		out.Code = int32(serr.Code())
		out.Error = serr.Error()

		return out, pbErr
	}

	pi, err := s.service.Pi(ctx, int(in.GetCount()))
	if err != nil {
		return handleError(err, out)
	}

	res := &pb.PiReply{
		Pi: pi,
	}

	return res, err
}
