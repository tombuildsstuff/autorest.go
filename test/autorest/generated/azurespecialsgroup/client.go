// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const telemetryInfo = "azsdk-go-azurespecialsgroup/<version>"

// ClientOptions contains configuration settings for the default client's pipeline.
type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return to
}

// Client - Test Infrastructure for AutoRest
type Client struct {
	u string
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.
func NewDefaultClient(options *ClientOptions) *Client {
	return NewClient(DefaultEndpoint, options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, options *ClientOptions) *Client {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) *Client {
	return &Client{u: endpoint, p: p}
}

// XMSClientRequestIDOperations returns the XMSClientRequestIDOperations associated with this client.
func (client *Client) XMSClientRequestIDOperations() XMSClientRequestIDOperations {
	return &xmsClientRequestIdoperations{Client: client}
}

// SubscriptionInCredentialsOperations returns the SubscriptionInCredentialsOperations associated with this client.
func (client *Client) SubscriptionInCredentialsOperations(subscriptionID string) SubscriptionInCredentialsOperations {
	return &subscriptionInCredentialsOperations{Client: client, subscriptionID: subscriptionID}
}

// SubscriptionInMethodOperations returns the SubscriptionInMethodOperations associated with this client.
func (client *Client) SubscriptionInMethodOperations() SubscriptionInMethodOperations {
	return &subscriptionInMethodOperations{Client: client}
}

// APIVersionDefaultOperations returns the APIVersionDefaultOperations associated with this client.
func (client *Client) APIVersionDefaultOperations() APIVersionDefaultOperations {
	return &apiVersionDefaultOperations{Client: client}
}

// APIVersionLocalOperations returns the APIVersionLocalOperations associated with this client.
func (client *Client) APIVersionLocalOperations() APIVersionLocalOperations {
	return &apiVersionLocalOperations{Client: client}
}

// SkipURLEncodingOperations returns the SkipURLEncodingOperations associated with this client.
func (client *Client) SkipURLEncodingOperations() SkipURLEncodingOperations {
	return &skipUrlEncodingOperations{Client: client}
}

// OdataOperations returns the OdataOperations associated with this client.
func (client *Client) OdataOperations() OdataOperations {
	return &odataOperations{Client: client}
}

// HeaderOperations returns the HeaderOperations associated with this client.
func (client *Client) HeaderOperations() HeaderOperations {
	return &headerOperations{Client: client}
}