// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/url"
)

const scope = "https://storage.azure.com/.default"

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

type Client struct {
	u *url.URL
	p azcore.Pipeline
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, cred azcore.Credential, options *ClientOptions) (*Client, error) {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)
	}
	return &Client{u: u, p: p}, nil
}

// ServiceOperations returns the ServiceOperations associated with this client.
func (client *Client) ServiceOperations() ServiceOperations {
	return &serviceOperations{Client: client}
}

// ContainerOperations returns the ContainerOperations associated with this client.
func (client *Client) ContainerOperations() ContainerOperations {
	return &containerOperations{Client: client}
}

// DirectoryOperations returns the DirectoryOperations associated with this client.
func (client *Client) DirectoryOperations(pathRenameMode *PathRenameMode) DirectoryOperations {
	return &directoryOperations{Client: client, pathRenameMode: pathRenameMode}
}

// BlobOperations returns the BlobOperations associated with this client.
func (client *Client) BlobOperations(pathRenameMode *PathRenameMode) BlobOperations {
	return &blobOperations{Client: client, pathRenameMode: pathRenameMode}
}

// PageBlobOperations returns the PageBlobOperations associated with this client.
func (client *Client) PageBlobOperations() PageBlobOperations {
	return &pageBlobOperations{Client: client}
}

// AppendBlobOperations returns the AppendBlobOperations associated with this client.
func (client *Client) AppendBlobOperations() AppendBlobOperations {
	return &appendBlobOperations{Client: client}
}

// BlockBlobOperations returns the BlockBlobOperations associated with this client.
func (client *Client) BlockBlobOperations() BlockBlobOperations {
	return &blockBlobOperations{Client: client}
}
