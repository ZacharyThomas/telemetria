package datadog

import (
	"context"

	"github.com/DataDog/datadog-go/statsd"
)

type statsdContextKey string

const statsdClient = statsdContextKey("statsd.client")

// Set the statsd client to the context
func SetClient(ctx context.Context, client *statsd.Client) context.Context {
	return context.WithValue(ctx, statsdClient, client)
}

// Create new statsd client and assign to context
func New(ctx context.Context, namespace string, options ...statsd.Option) context.Context {
	options = append(options, statsd.WithNamespace(namespace))

	client, err := statsd.New("", options...)
	if err != nil {
		panic(err)
	}

	return SetClient(ctx, client)
}

func ClientFromContext(ctx context.Context) *statsd.Client {
	value := ctx.Value(statsdClient)
	if value == nil {
		panic("No statsd client found in context")
	}

	return value.(*statsd.Client)
}