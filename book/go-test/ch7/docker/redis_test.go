package docker

import (
	"context"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestRedis(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		t.Error(err)
	}
	defer redisC.Terminate(ctx)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err = rdb.Set("key", "value", 0).Err()
	assert.NoError(t, err)

	val, err := rdb.Get("key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value", val)
}
