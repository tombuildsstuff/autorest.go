// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// IntClient contains the methods for the Int group.
// Don't use this type directly, use NewIntClient() instead.
type IntClient struct {
	con *Connection
}

// NewIntClient creates a new instance of IntClient with the specified values.
func NewIntClient(con *Connection) *IntClient {
	return &IntClient{con: con}
}

// Get - Get an int enum
func (client *IntClient) Get(ctx context.Context, options *IntGetOptions) (IntEnumResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return IntEnumResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IntEnumResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IntEnumResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntClient) getCreateRequest(ctx context.Context, options *IntGetOptions) (*azcore.Request, error) {
	urlPath := "/nonStringEnums/int/get"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntClient) getHandleResponse(resp *azcore.Response) (IntEnumResponse, error) {
	var val *IntEnum
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IntEnumResponse{}, err
	}
	return IntEnumResponse{RawResponse: resp.Response, Value: val}, nil
}

// getHandleError handles the Get error response.
func (client *IntClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Put - Put an int enum
func (client *IntClient) Put(ctx context.Context, options *IntPutOptions) (StringResponse, error) {
	req, err := client.putCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putHandleError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *IntClient) putCreateRequest(ctx context.Context, options *IntPutOptions) (*azcore.Request, error) {
	urlPath := "/nonStringEnums/int/put"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Input)
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *IntClient) putHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// putHandleError handles the Put error response.
func (client *IntClient) putHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
