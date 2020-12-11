// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package custombaseurlgroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const telemetryInfo = "azsdk-go-custombaseurlgroup/<version>"

// ConnectionOptions contains configuration settings for the connection's pipeline.
type ConnectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging azcore.LogOptions
}

// DefaultConnectionOptions creates a ConnectionOptions type initialized with default values.
func DefaultConnectionOptions() ConnectionOptions {
	return ConnectionOptions{
		Retry:     azcore.DefaultRetryOptions(),
		Telemetry: azcore.DefaultTelemetryOptions(),
		Logging:   azcore.DefaultLogOptions(),
	}
}

func (c *ConnectionOptions) telemetryOptions() *azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return &to
}

// Connection - Test Infrastructure for AutoRest
type Connection struct {
	host string
	p    azcore.Pipeline
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
func NewConnection(host *string, options *ConnectionOptions) *Connection {
	if options == nil {
		o := DefaultConnectionOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewLogPolicy(&options.Logging))
	return NewConnectionWithPipeline(host, p)
}

// NewConnectionWithPipeline creates an instance of the Connection type with the specified endpoint and pipeline.
func NewConnectionWithPipeline(host *string, p azcore.Pipeline) *Connection {
	client := &Connection{
		p:    p,
		host: "host",
	}
	if host != nil {
		client.host = *host
	}
	return client
}

// Host returns part of the parameterized host.
func (c *Connection) Host() string {
	return c.host
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() azcore.Pipeline {
	return c.p
}
