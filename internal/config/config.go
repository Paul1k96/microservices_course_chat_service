package config

import "time"

// PGConfig represents configuration for PostgreSQL.
type PGConfig interface {
	GetDSN() string
}

// GRPCConfig represents configuration for gRPC.
type GRPCConfig interface {
	GetServerAddress() string
	GetAuthAddress() string
}

// RedisConfig represents configuration for Redis.
type RedisConfig interface {
	GetAddress() string
	GetConnectionTimeout() time.Duration
	GetMaxIdle() int
	GetIdleTimeout() time.Duration
}
