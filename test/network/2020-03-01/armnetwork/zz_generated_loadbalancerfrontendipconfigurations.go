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

// LoadBalancerFrontendIPConfigurationsOperations contains the methods for the LoadBalancerFrontendIPConfigurations group.
type LoadBalancerFrontendIPConfigurationsOperations interface {
	// Get - Gets load balancer frontend IP configuration.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string) (*FrontendIPConfigurationResponse, error)
	// List - Gets all the load balancer frontend IP configurations.
	List(resourceGroupName string, loadBalancerName string) LoadBalancerFrontendIPConfigurationListResultPager
}

// LoadBalancerFrontendIPConfigurationsClient implements the LoadBalancerFrontendIPConfigurationsOperations interface.
// Don't use this type directly, use NewLoadBalancerFrontendIPConfigurationsClient() instead.
type LoadBalancerFrontendIPConfigurationsClient struct {
	*Client
	subscriptionID string
}

// NewLoadBalancerFrontendIPConfigurationsClient creates a new instance of LoadBalancerFrontendIPConfigurationsClient with the specified values.
func NewLoadBalancerFrontendIPConfigurationsClient(c *Client, subscriptionID string) LoadBalancerFrontendIPConfigurationsOperations {
	return &LoadBalancerFrontendIPConfigurationsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *LoadBalancerFrontendIPConfigurationsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Gets load balancer frontend IP configuration.
func (client *LoadBalancerFrontendIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string) (*FrontendIPConfigurationResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, loadBalancerName, frontendIPConfigurationName)
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
func (client *LoadBalancerFrontendIPConfigurationsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{frontendIPConfigurationName}", url.PathEscape(frontendIPConfigurationName))
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
func (client *LoadBalancerFrontendIPConfigurationsClient) GetHandleResponse(resp *azcore.Response) (*FrontendIPConfigurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := FrontendIPConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FrontendIPConfiguration)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerFrontendIPConfigurationsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the load balancer frontend IP configurations.
func (client *LoadBalancerFrontendIPConfigurationsClient) List(resourceGroupName string, loadBalancerName string) LoadBalancerFrontendIPConfigurationListResultPager {
	return &loadBalancerFrontendIPConfigurationListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, loadBalancerName)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *LoadBalancerFrontendIPConfigurationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerFrontendIPConfigurationListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerFrontendIPConfigurationsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations"
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
func (client *LoadBalancerFrontendIPConfigurationsClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerFrontendIPConfigurationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := LoadBalancerFrontendIPConfigurationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerFrontendIPConfigurationListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerFrontendIPConfigurationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
