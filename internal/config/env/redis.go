package env

import (
	"os"
	"strconv"
	"time"

	"github.com/Paul1k96/microservices_course_chat_service/internal/config"
)

const (
	redisHost        = "REDIS_HOST"
	redisPort        = "REDIS_PORT"
	redisConnTimeout = "REDIS_CONN_TIMEOUT"
	redisMaxIdle     = "REDIS_MAX_IDLE"
	redisIdleTimeout = "REDIS_IDLE_TIMEOUT"
)

type redisConfig struct {
	host        string
	port        string
	connTimeout time.Duration
	maxIdle     int
	idleTimeout time.Duration
}

// NewRedisConfig returns a new config.RedisConfig.
func NewRedisConfig() (config.RedisConfig, error) {
	var cfg redisConfig

	cfg.host = os.Getenv(redisHost)
	cfg.port = os.Getenv(redisPort)

	connTimeout, err := time.ParseDuration(os.Getenv(redisConnTimeout))
	if err != nil {
		return nil, err
	}
	cfg.connTimeout = connTimeout

	maxIdle, err := strconv.Atoi(os.Getenv(redisMaxIdle))
	if err != nil {
		return nil, err
	}
	cfg.maxIdle = maxIdle

	idleTimeout, err := time.ParseDuration(os.Getenv(redisIdleTimeout))
	if err != nil {
		return nil, err
	}
	cfg.idleTimeout = idleTimeout

	return &cfg, nil
}

// GetAddress returns the address.
func (c *redisConfig) GetAddress() string {
	return c.host + ":" + c.port
}

// GetConnectionTimeout returns the connection timeout.
func (c *redisConfig) GetConnectionTimeout() time.Duration {
	return c.connTimeout
}

// GetMaxIdle returns the max idle.
func (c *redisConfig) GetMaxIdle() int {
	return c.maxIdle
}

// GetIdleTimeout returns the idle timeout.
func (c *redisConfig) GetIdleTimeout() time.Duration {
	return c.idleTimeout
}
