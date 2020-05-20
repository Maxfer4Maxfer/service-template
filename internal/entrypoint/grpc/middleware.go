package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
	"google.golang.org/grpc"
)

var (
	errNoCorrelationID = errors.New(
		"incomming request's ctx should be with CorrelationID")
)

// UnaryServerInterceptorCorrelationID attachs a new Correlation ID
// to a gRPC request.
func (s *ServerGRPC) UnaryServerInterceptorCorrelationID() func(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		cID := correlationid.Extract(ctx)
		if cID == "" {
			ctx, _ = correlationid.Assign(ctx)

			s.logger.Warn().Err(errNoCorrelationID).Msg("")
		}

		resp, err := handler(ctx, req)

		return resp, err
	}
}

// UnaryServerInterceptorLogging log start and end of a gRPC session.
func (s *ServerGRPC) UnaryServerInterceptorLogging() func(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		cID := correlationid.Extract(ctx)

		s.logger.Info().
			Str("correlationID", cID).
			Msg("incomming http request")

		resp, err := handler(ctx, req)

		s.logger.Info().
			Str("correlationID", cID).
			Str("duration", time.Since(start).String()).
			Msg("finishing handle http request")

		return resp, err
	}
}
