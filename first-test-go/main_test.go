package main

import (
	"context"
	"io"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:6.2.6-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer redisC.Terminate(ctx)

	if err != nil {
		t.Error(err)
	}

	logs, err := redisC.Logs(ctx)
	if err != nil {
		t.Error(err)
	}

	bytes, err := io.ReadAll(logs)
	if err != nil {
		t.Error(err)
	}

	t.Logf(string(bytes))
}
