// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type linkedServiceClient struct {
	con *connection
}

// CreateOrUpdateLinkedService - Creates or updates a linked service.
func (client *linkedServiceClient) createOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceBeginCreateOrUpdateLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateLinkedServiceCreateRequest(ctx, linkedServiceName, linkedService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateLinkedServiceCreateRequest creates the CreateOrUpdateLinkedService request.
func (client *linkedServiceClient) createOrUpdateLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceBeginCreateOrUpdateLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(linkedService)
}

// createOrUpdateLinkedServiceHandleResponse handles the CreateOrUpdateLinkedService response.
func (client *linkedServiceClient) createOrUpdateLinkedServiceHandleResponse(resp *azcore.Response) (LinkedServiceResourceResponse, error) {
	var val *LinkedServiceResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	return LinkedServiceResourceResponse{RawResponse: resp.Response, LinkedServiceResource: val}, nil
}

// createOrUpdateLinkedServiceHandleError handles the CreateOrUpdateLinkedService error response.
func (client *linkedServiceClient) createOrUpdateLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteLinkedService - Deletes a linked service.
func (client *linkedServiceClient) deleteLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceBeginDeleteLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.deleteLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// deleteLinkedServiceCreateRequest creates the DeleteLinkedService request.
func (client *linkedServiceClient) deleteLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *LinkedServiceBeginDeleteLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteLinkedServiceHandleError handles the DeleteLinkedService error response.
func (client *linkedServiceClient) deleteLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLinkedService - Gets a linked service.
func (client *linkedServiceClient) GetLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceGetLinkedServiceOptions) (LinkedServiceResourceResponse, error) {
	req, err := client.getLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return LinkedServiceResourceResponse{}, client.getLinkedServiceHandleError(resp)
	}
	return client.getLinkedServiceHandleResponse(resp)
}

// getLinkedServiceCreateRequest creates the GetLinkedService request.
func (client *linkedServiceClient) getLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *LinkedServiceGetLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getLinkedServiceHandleResponse handles the GetLinkedService response.
func (client *linkedServiceClient) getLinkedServiceHandleResponse(resp *azcore.Response) (LinkedServiceResourceResponse, error) {
	var val *LinkedServiceResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	return LinkedServiceResourceResponse{RawResponse: resp.Response, LinkedServiceResource: val}, nil
}

// getLinkedServiceHandleError handles the GetLinkedService error response.
func (client *linkedServiceClient) getLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLinkedServicesByWorkspace - Lists linked services.
func (client *linkedServiceClient) GetLinkedServicesByWorkspace(options *LinkedServiceGetLinkedServicesByWorkspaceOptions) LinkedServiceListResponsePager {
	return &linkedServiceListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getLinkedServicesByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getLinkedServicesByWorkspaceHandleResponse,
		errorer:   client.getLinkedServicesByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp LinkedServiceListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LinkedServiceListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getLinkedServicesByWorkspaceCreateRequest creates the GetLinkedServicesByWorkspace request.
func (client *linkedServiceClient) getLinkedServicesByWorkspaceCreateRequest(ctx context.Context, options *LinkedServiceGetLinkedServicesByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getLinkedServicesByWorkspaceHandleResponse handles the GetLinkedServicesByWorkspace response.
func (client *linkedServiceClient) getLinkedServicesByWorkspaceHandleResponse(resp *azcore.Response) (LinkedServiceListResponseResponse, error) {
	var val *LinkedServiceListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LinkedServiceListResponseResponse{}, err
	}
	return LinkedServiceListResponseResponse{RawResponse: resp.Response, LinkedServiceListResponse: val}, nil
}

// getLinkedServicesByWorkspaceHandleError handles the GetLinkedServicesByWorkspace error response.
func (client *linkedServiceClient) getLinkedServicesByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameLinkedService - Renames a linked service.
func (client *linkedServiceClient) renameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceBeginRenameLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.renameLinkedServiceCreateRequest(ctx, linkedServiceName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// renameLinkedServiceCreateRequest creates the RenameLinkedService request.
func (client *linkedServiceClient) renameLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceBeginRenameLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameLinkedServiceHandleError handles the RenameLinkedService error response.
func (client *linkedServiceClient) renameLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
