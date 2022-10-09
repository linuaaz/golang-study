package gnomock

import (
	"testing"

	redisclient "github.com/go-redis/redis/v7"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
	"github.com/stretchr/testify/require"
)

func TestRedis(t *testing.T) {
	vs := make(map[string]interface{})

	vs["a"] = "foo"
	vs["b"] = 42
	vs["c"] = true

	p := redis.Preset(redis.WithValues(vs))
	container, _ := gnomock.Start(p,
		gnomock.WithDebugMode(),
		gnomock.WithUseLocalImagesFirst(),
	)

	defer func() { _ = gnomock.Stop(container) }()

	addr := container.DefaultAddress()
	client := redisclient.NewClient(&redisclient.Options{Addr: addr})

	v1, err := client.Get("a").Result()
	require.Equal(t, "foo", v1)
	require.NoError(t, err)

	var number int

	err = client.Get("b").Scan(&number)
	require.Equal(t, 42, number)
	require.NoError(t, err)

	var flag bool

	err = client.Get("c").Scan(&flag)
	require.Equal(t, true, flag)
	require.NoError(t, err)
}
