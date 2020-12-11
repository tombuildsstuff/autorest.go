// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HeaderClient contains the methods for the Header group.
// Don't use this type directly, use NewHeaderClient() instead.
type HeaderClient struct {
	con *Connection
}

// NewHeaderClient creates a new instance of HeaderClient with the specified values.
func NewHeaderClient(con *Connection) *HeaderClient {
	return &HeaderClient{con: con}
}

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func (client *HeaderClient) CustomNamedRequestID(ctx context.Context, fooClientRequestId string, options *HeaderCustomNamedRequestIDOptions) (HeaderCustomNamedRequestIDResponse, error) {
	req, err := client.customNamedRequestIdCreateRequest(ctx, fooClientRequestId, options)
	if err != nil {
		return HeaderCustomNamedRequestIDResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HeaderCustomNamedRequestIDResponse{}, client.customNamedRequestIdHandleError(resp)
	}
	return client.customNamedRequestIdHandleResponse(resp)
}

// customNamedRequestIdCreateRequest creates the CustomNamedRequestID request.
func (client *HeaderClient) customNamedRequestIdCreateRequest(ctx context.Context, fooClientRequestId string, options *HeaderCustomNamedRequestIDOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestId"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("foo-client-request-id", fooClientRequestId)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIdHandleResponse handles the CustomNamedRequestID response.
func (client *HeaderClient) customNamedRequestIdHandleResponse(resp *azcore.Response) (HeaderCustomNamedRequestIDResponse, error) {
	result := HeaderCustomNamedRequestIDResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIdHandleError handles the CustomNamedRequestID error response.
func (client *HeaderClient) customNamedRequestIdHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func (client *HeaderClient) CustomNamedRequestIDHead(ctx context.Context, fooClientRequestId string, options *HeaderCustomNamedRequestIDHeadOptions) (HeaderCustomNamedRequestIDHeadResponse, error) {
	req, err := client.customNamedRequestIdHeadCreateRequest(ctx, fooClientRequestId, options)
	if err != nil {
		return HeaderCustomNamedRequestIDHeadResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDHeadResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotFound) {
		return HeaderCustomNamedRequestIDHeadResponse{}, client.customNamedRequestIdHeadHandleError(resp)
	}
	return client.customNamedRequestIdHeadHandleResponse(resp)
}

// customNamedRequestIdHeadCreateRequest creates the CustomNamedRequestIDHead request.
func (client *HeaderClient) customNamedRequestIdHeadCreateRequest(ctx context.Context, fooClientRequestId string, options *HeaderCustomNamedRequestIDHeadOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdHead"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("foo-client-request-id", fooClientRequestId)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIdHeadHandleResponse handles the CustomNamedRequestIDHead response.
func (client *HeaderClient) customNamedRequestIdHeadHandleResponse(resp *azcore.Response) (HeaderCustomNamedRequestIDHeadResponse, error) {
	result := HeaderCustomNamedRequestIDHeadResponse{RawResponse: resp.Response}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIdHeadHandleError handles the CustomNamedRequestIDHead error response.
func (client *HeaderClient) customNamedRequestIdHeadHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request, via a parameter group
func (client *HeaderClient) CustomNamedRequestIDParamGrouping(ctx context.Context, headerCustomNamedRequestIdParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	req, err := client.customNamedRequestIdParamGroupingCreateRequest(ctx, headerCustomNamedRequestIdParamGroupingParameters)
	if err != nil {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, client.customNamedRequestIdParamGroupingHandleError(resp)
	}
	return client.customNamedRequestIdParamGroupingHandleResponse(resp)
}

// customNamedRequestIdParamGroupingCreateRequest creates the CustomNamedRequestIDParamGrouping request.
func (client *HeaderClient) customNamedRequestIdParamGroupingCreateRequest(ctx context.Context, headerCustomNamedRequestIdParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdParamGrouping"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("foo-client-request-id", headerCustomNamedRequestIdParamGroupingParameters.FooClientRequestId)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIdParamGroupingHandleResponse handles the CustomNamedRequestIDParamGrouping response.
func (client *HeaderClient) customNamedRequestIdParamGroupingHandleResponse(resp *azcore.Response) (HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	result := HeaderCustomNamedRequestIDParamGroupingResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIdParamGroupingHandleError handles the CustomNamedRequestIDParamGrouping error response.
func (client *HeaderClient) customNamedRequestIdParamGroupingHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
