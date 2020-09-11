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

// LoadBalancerProbesOperations contains the methods for the LoadBalancerProbes group.
type LoadBalancerProbesOperations interface {
	// Get - Gets load balancer probe.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string) (*ProbeResponse, error)
	// List - Gets all the load balancer probes.
	List(resourceGroupName string, loadBalancerName string) LoadBalancerProbeListResultPager
}

// LoadBalancerProbesClient implements the LoadBalancerProbesOperations interface.
// Don't use this type directly, use NewLoadBalancerProbesClient() instead.
type LoadBalancerProbesClient struct {
	*Client
	subscriptionID string
}

// NewLoadBalancerProbesClient creates a new instance of LoadBalancerProbesClient with the specified values.
func NewLoadBalancerProbesClient(c *Client, subscriptionID string) LoadBalancerProbesOperations {
	return &LoadBalancerProbesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *LoadBalancerProbesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Gets load balancer probe.
func (client *LoadBalancerProbesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string) (*ProbeResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, loadBalancerName, probeName)
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
func (client *LoadBalancerProbesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{probeName}", url.PathEscape(probeName))
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
func (client *LoadBalancerProbesClient) GetHandleResponse(resp *azcore.Response) (*ProbeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := ProbeResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Probe)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerProbesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the load balancer probes.
func (client *LoadBalancerProbesClient) List(resourceGroupName string, loadBalancerName string) LoadBalancerProbeListResultPager {
	return &loadBalancerProbeListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, loadBalancerName)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *LoadBalancerProbeListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerProbeListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerProbesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes"
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
func (client *LoadBalancerProbesClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerProbeListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := LoadBalancerProbeListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerProbeListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerProbesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
