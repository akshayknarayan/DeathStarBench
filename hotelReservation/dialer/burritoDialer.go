// +build burrito

package dialer

import (
	"fmt"
	"os"
	"time"

	burrito "github.com/akshayknarayan/burrito/resolv-go"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	consul "github.com/hashicorp/consul/api"
	//lb "github.com/olivere/grpc/lb/consul"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var burrito_root = getBurritoRoot()

// DialOption allows optional config for dialer
type DialOption func(name string) (grpc.DialOption, error)

// WithTracer traces rpc calls
func WithTracer(tracer opentracing.Tracer) DialOption {
	return func(name string) (grpc.DialOption, error) {
		return grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)), nil
	}
}

// WithBalancer enables client side load balancing
func WithBalancer(registry *consul.Client) DialOption {
	// TODO with burrito, this doesn't really make sense?
	// disable for now. need to think about the load balancing story.
	return func(name string) (grpc.DialOption, error) {
		return grpc.EmptyDialOption{}, nil
	}

	//return func(name string) (grpc.DialOption, error) {
	//	r, err := lb.NewResolver(registry, name, "")
	//	if err != nil {
	//		return nil, err
	//	}
	//	return grpc.WithBalancer(grpc.RoundRobin(r)), nil
	//}
}

func WithBurritoDialer(burrito_root string) DialOption {
	return func(name string) (grpc.DialOption, error) {
		return grpc.WithDialer(
			burrito.BurritoDialer(burrito_root),
		), nil
	}
}

// Dial returns a load balanced grpc client conn with tracing interceptor
func Dial(name string, opts ...DialOption) (*grpc.ClientConn, error) {
	dialopts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Timeout:             120 * time.Second,
			PermitWithoutStream: true,
		}),
	}

	for _, fn := range opts {
		opt, err := fn(name)
		if err != nil {
			return nil, fmt.Errorf("config error: %v", err)
		}
		dialopts = append(dialopts, opt)
	}

	opt, err := WithBurritoDialer(burrito_root)(name)
	if err != nil {
		return nil, fmt.Errorf("config error: %v", err)
	}

	dialopts = append(dialopts, opt)

	conn, err := grpc.Dial(name, dialopts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s: %v", name, err)
	}

	return conn, nil
}

// Check:
// 1. BURRITO_ROOT environment variable
// 2. /burrito/controller
// If none found, panic.
func getBurritoRoot() string {
	env_root := os.Getenv("BURRITO_ROOT")
	if env_root != "" {
		return env_root
	}

	info, err := os.Stat("/burrito")
	if err != nil {
		panic(fmt.Errorf("Error opening /burrito: %v", err))
	}

	if !info.IsDir() {
		panic(fmt.Errorf("Opened /burrito, but is not a directory"))
	}

	return "/burrito"
}
