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
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *linkedServiceClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdateLinkedService - Creates or updates a linked service.
func (client *linkedServiceClient) CreateOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, linkedServiceCreateOrUpdateLinkedServiceOptions *LinkedServiceCreateOrUpdateLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateLinkedServiceCreateRequest(ctx, linkedServiceName, linkedService, linkedServiceCreateOrUpdateLinkedServiceOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdateLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateLinkedServiceCreateRequest creates the CreateOrUpdateLinkedService request.
func (client *linkedServiceClient) CreateOrUpdateLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, linkedServiceCreateOrUpdateLinkedServiceOptions *LinkedServiceCreateOrUpdateLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if linkedServiceCreateOrUpdateLinkedServiceOptions != nil && linkedServiceCreateOrUpdateLinkedServiceOptions.IfMatch != nil {
		req.Header.Set("If-Match", *linkedServiceCreateOrUpdateLinkedServiceOptions.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(linkedService)
}

// CreateOrUpdateLinkedServiceHandleResponse handles the CreateOrUpdateLinkedService response.
func (client *linkedServiceClient) CreateOrUpdateLinkedServiceHandleResponse(resp *azcore.Response) (*LinkedServiceResourceResponse, error) {
	result := LinkedServiceResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LinkedServiceResource)
}

// CreateOrUpdateLinkedServiceHandleError handles the CreateOrUpdateLinkedService error response.
func (client *linkedServiceClient) CreateOrUpdateLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteLinkedService - Deletes a linked service.
func (client *linkedServiceClient) DeleteLinkedService(ctx context.Context, linkedServiceName string) (*azcore.Response, error) {
	req, err := client.DeleteLinkedServiceCreateRequest(ctx, linkedServiceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// DeleteLinkedServiceCreateRequest creates the DeleteLinkedService request.
func (client *linkedServiceClient) DeleteLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteLinkedServiceHandleError handles the DeleteLinkedService error response.
func (client *linkedServiceClient) DeleteLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLinkedService - Gets a linked service.
func (client *linkedServiceClient) GetLinkedService(ctx context.Context, linkedServiceName string, linkedServiceGetLinkedServiceOptions *LinkedServiceGetLinkedServiceOptions) (*LinkedServiceResourceResponse, error) {
	req, err := client.GetLinkedServiceCreateRequest(ctx, linkedServiceName, linkedServiceGetLinkedServiceOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetLinkedServiceHandleError(resp)
	}
	result, err := client.GetLinkedServiceHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLinkedServiceCreateRequest creates the GetLinkedService request.
func (client *linkedServiceClient) GetLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, linkedServiceGetLinkedServiceOptions *LinkedServiceGetLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if linkedServiceGetLinkedServiceOptions != nil && linkedServiceGetLinkedServiceOptions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *linkedServiceGetLinkedServiceOptions.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetLinkedServiceHandleResponse handles the GetLinkedService response.
func (client *linkedServiceClient) GetLinkedServiceHandleResponse(resp *azcore.Response) (*LinkedServiceResourceResponse, error) {
	result := LinkedServiceResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LinkedServiceResource)
}

// GetLinkedServiceHandleError handles the GetLinkedService error response.
func (client *linkedServiceClient) GetLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLinkedServicesByWorkspace - Lists linked services.
func (client *linkedServiceClient) GetLinkedServicesByWorkspace() LinkedServiceListResponsePager {
	return &linkedServiceListResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetLinkedServicesByWorkspaceCreateRequest(ctx)
		},
		responder: client.GetLinkedServicesByWorkspaceHandleResponse,
		errorer:   client.GetLinkedServicesByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *LinkedServiceListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LinkedServiceListResponse.NextLink)
		},
	}
}

// GetLinkedServicesByWorkspaceCreateRequest creates the GetLinkedServicesByWorkspace request.
func (client *linkedServiceClient) GetLinkedServicesByWorkspaceCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/linkedservices"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetLinkedServicesByWorkspaceHandleResponse handles the GetLinkedServicesByWorkspace response.
func (client *linkedServiceClient) GetLinkedServicesByWorkspaceHandleResponse(resp *azcore.Response) (*LinkedServiceListResponseResponse, error) {
	result := LinkedServiceListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LinkedServiceListResponse)
}

// GetLinkedServicesByWorkspaceHandleError handles the GetLinkedServicesByWorkspace error response.
func (client *linkedServiceClient) GetLinkedServicesByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}