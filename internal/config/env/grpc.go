package env

import (
	"net"
	"os"

	"github.com/Paul1k96/microservices_course_chat_service/internal/config"
)

const (
	grpcHost = "GRPC_HOST"
	grpcPort = "GRPC_PORT"

	grpcClientAuthHost = "GRPC_CLIENT_AUTH_HOST"
	grpcClientAuthPort = "GRPC_CLIENT_AUTH_PORT"
)

type grpcServerConfig struct {
	serverHost string
	serverPort string
	authHost   string
	authPort   string
}

// NewGRPCConfig returns a new config.GRPCConfig.
func NewGRPCConfig() config.GRPCConfig {
	var cfg grpcServerConfig

	cfg.serverHost = os.Getenv(grpcHost)
	cfg.serverPort = os.Getenv(grpcPort)

	cfg.authHost = os.Getenv(grpcClientAuthHost)
	cfg.authPort = os.Getenv(grpcClientAuthPort)

	return &cfg
}

// GetServerAddress returns the address of the gRPC server.
func (c *grpcServerConfig) GetServerAddress() string {
	return net.JoinHostPort(c.serverHost, c.serverPort)
}

// GetAuthAddress returns the address of the gRPC auth server.
func (c *grpcServerConfig) GetAuthAddress() string {
	return net.JoinHostPort(c.authHost, c.authPort)
}
