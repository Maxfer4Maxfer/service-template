package grpcclient

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	pb "github.com/maxfer4maxfer/service-template/internal/entrypoint/grpc/pb"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
)

var (
	// errGrpcServerUnavailable throws when a gRPC server is unavailable.
	errGrpcServerUnavailable = errors.New("gRPC server is unavailable")
	// errGrpcServerStandardError throws when a gRPC server returns
	// one of the standard errors.
	errGrpcServerStandardError = errors.New("gRPC server returned a standard error")
	// errGrpcServerUnknowError throws when a gRPC server returns
	// one of an unknow error.
	errGrpcServerUnknowError = errors.New("gRPC server returned an unknow error")
)

// Client represents a wrapper for the "errors-supplier-account" service.
type Client struct {
	logger *zerolog.Logger
	addr   string
}

// New returns a new Client instance.
func New(logger *zerolog.Logger, addr string) *Client {
	return &Client{
		logger: logger,
		addr:   addr,
	}
}

func (c *Client) client() (pb.CalculatorClient, *grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure())
	if err != nil {
		c.logger.Error().Err(err).
			Msg("Calculator.grpcClient: cannot connect to gRPC server")

		return nil, nil, errGrpcServerUnavailable.Wrap(err).Method("Multiply")
	}

	return pb.NewCalculatorClient(conn), conn, nil
}

func (c *Client) handleError(err error, method string) error {
	_, ok := status.FromError(err)
	if ok {
		return errGrpcServerStandardError.Wrap(err).Method(method)
	}

	return errGrpcServerUnknowError.Wrap(err).Method(method)
}

// Subtract subtracts one number for other.
func (c *Client) Subtract(
	ctx context.Context, a int, b int,
) (
	sub int, err error,
) {
	client, conn, err := c.client()
	if err != nil {
		return sub, err
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r, err := client.Subtract(ctx, &pb.SubtractRequest{A: int32(a), B: int32(b)})
	if err != nil {
		return 0, c.handleError(err, "client.Subtract")
	}

	c.logger.Debug().
		Str("method", "Subtract").
		Interface("resp", r).
		Msg("Calculator.grpcClient: call to a remote endpoint")

	sub = int(r.GetSub())

	return sub, nil
}

// Multiply multiplies too numbers.
func (c *Client) Multiply(
	ctx context.Context, a int, b int,
) (
	mult int, err error,
) {
	client, conn, err := c.client()
	if err != nil {
		return mult, err
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r, err := client.Multiply(ctx, &pb.MultiplyRequest{A: int32(a), B: int32(b)})
	if err != nil {
		return 0, c.handleError(err, "client.Multiply")
	}

	c.logger.Debug().
		Str("method", "Multiply").
		Interface("resp", r).
		Msg("Calculator.grpcClient: call to a remote endpoint")

	mult = int(r.GetMult())

	return mult, nil
}

// Pi returns pi number with a given length.
func (c *Client) Pi(ctx context.Context, count int) (pi string, err error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure())
	if err != nil {
		c.logger.Error().Err(err).
			Msg("Calculator.grpcClient: Pi: cannot connect to gRPC server")

		return "", errGrpcServerUnavailable.Wrap(err).Method("Pi")
	}

	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r, err := client.Pi(ctx, &pb.PiRequest{Count: int32(count)})
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			return "", errGrpcServerStandardError.Wrap(err).Method("client.Pi")
		}

		return "", errGrpcServerUnknowError.Wrap(err).Method("client.Pi")
	}

	c.logger.Debug().
		Str("method", "Pi").
		Interface("resp", r).
		Msg("Calculator.grpcClient: Pi: call to a remote endpoint")

	pi = r.GetPi()

	return pi, nil
}
