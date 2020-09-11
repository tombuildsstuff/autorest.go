// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancerBackendAddressPoolsOperations contains the methods for the LoadBalancerBackendAddressPools group.
type LoadBalancerBackendAddressPoolsOperations interface {
	// Get - Gets load balancer backend address pool.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*BackendAddressPoolResponse, error)
	// List - Gets all the load balancer backed address pools.
	List(resourceGroupName string, loadBalancerName string) LoadBalancerBackendAddressPoolListResultPager
}

// LoadBalancerBackendAddressPoolsClient implements the LoadBalancerBackendAddressPoolsOperations interface.
// Don't use this type directly, use NewLoadBalancerBackendAddressPoolsClient() instead.
type LoadBalancerBackendAddressPoolsClient struct {
	*Client
	subscriptionID string
}

// NewLoadBalancerBackendAddressPoolsClient creates a new instance of LoadBalancerBackendAddressPoolsClient with the specified values.
func NewLoadBalancerBackendAddressPoolsClient(c *Client, subscriptionID string) LoadBalancerBackendAddressPoolsOperations {
	return &LoadBalancerBackendAddressPoolsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *LoadBalancerBackendAddressPoolsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Gets load balancer backend address pool.
func (client *LoadBalancerBackendAddressPoolsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*BackendAddressPoolResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *LoadBalancerBackendAddressPoolsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{backendAddressPoolName}", url.PathEscape(backendAddressPoolName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *LoadBalancerBackendAddressPoolsClient) GetHandleResponse(resp *azcore.Response) (*BackendAddressPoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := BackendAddressPoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BackendAddressPool)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerBackendAddressPoolsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the load balancer backed address pools.
func (client *LoadBalancerBackendAddressPoolsClient) List(resourceGroupName string, loadBalancerName string) LoadBalancerBackendAddressPoolListResultPager {
	return &loadBalancerBackendAddressPoolListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, loadBalancerName)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *LoadBalancerBackendAddressPoolListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerBackendAddressPoolListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerBackendAddressPoolsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *LoadBalancerBackendAddressPoolsClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerBackendAddressPoolListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := LoadBalancerBackendAddressPoolListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerBackendAddressPoolListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerBackendAddressPoolsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
