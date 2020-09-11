// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PolymorphicrecursiveOperations contains the methods for the Polymorphicrecursive group.
type PolymorphicrecursiveOperations interface {
	// GetValid - Get complex types that are polymorphic and have recursive references
	GetValid(ctx context.Context) (*FishResponse, error)
	// PutValid - Put complex types that are polymorphic and have recursive references
	PutValid(ctx context.Context, complexBody FishClassification) (*http.Response, error)
}

// PolymorphicrecursiveClient implements the PolymorphicrecursiveOperations interface.
// Don't use this type directly, use NewPolymorphicrecursiveClient() instead.
type PolymorphicrecursiveClient struct {
	*Client
}

// NewPolymorphicrecursiveClient creates a new instance of PolymorphicrecursiveClient with the specified values.
func NewPolymorphicrecursiveClient(c *Client) PolymorphicrecursiveOperations {
	return &PolymorphicrecursiveClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PolymorphicrecursiveClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetValid - Get complex types that are polymorphic and have recursive references
func (client *PolymorphicrecursiveClient) GetValid(ctx context.Context) (*FishResponse, error) {
	req, err := client.GetValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetValidCreateRequest creates the GetValid request.
func (client *PolymorphicrecursiveClient) GetValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *PolymorphicrecursiveClient) GetValidHandleResponse(resp *azcore.Response) (*FishResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetValidHandleError(resp)
	}
	result := FishResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// GetValidHandleError handles the GetValid error response.
func (client *PolymorphicrecursiveClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutValid - Put complex types that are polymorphic and have recursive references
func (client *PolymorphicrecursiveClient) PutValid(ctx context.Context, complexBody FishClassification) (*http.Response, error) {
	req, err := client.PutValidCreateRequest(ctx, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutValidCreateRequest creates the PutValid request.
func (client *PolymorphicrecursiveClient) PutValidCreateRequest(ctx context.Context, complexBody FishClassification) (*azcore.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleResponse handles the PutValid response.
func (client *PolymorphicrecursiveClient) PutValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutValidHandleError(resp)
	}
	return resp.Response, nil
}

// PutValidHandleError handles the PutValid error response.
func (client *PolymorphicrecursiveClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
