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

// NetworkInterfaceIPConfigurationsOperations contains the methods for the NetworkInterfaceIPConfigurations group.
type NetworkInterfaceIPConfigurationsOperations interface {
	// Get - Gets the specified network interface ip configuration.
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string) (*NetworkInterfaceIPConfigurationResponse, error)
	// List - Get all ip configurations in a network interface.
	List(resourceGroupName string, networkInterfaceName string) NetworkInterfaceIPConfigurationListResultPager
}

// NetworkInterfaceIPConfigurationsClient implements the NetworkInterfaceIPConfigurationsOperations interface.
// Don't use this type directly, use NewNetworkInterfaceIPConfigurationsClient() instead.
type NetworkInterfaceIPConfigurationsClient struct {
	*Client
	subscriptionID string
}

// NewNetworkInterfaceIPConfigurationsClient creates a new instance of NetworkInterfaceIPConfigurationsClient with the specified values.
func NewNetworkInterfaceIPConfigurationsClient(c *Client, subscriptionID string) NetworkInterfaceIPConfigurationsOperations {
	return &NetworkInterfaceIPConfigurationsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *NetworkInterfaceIPConfigurationsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Gets the specified network interface ip configuration.
func (client *NetworkInterfaceIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string) (*NetworkInterfaceIPConfigurationResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkInterfaceName, ipConfigurationName)
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
func (client *NetworkInterfaceIPConfigurationsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
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
func (client *NetworkInterfaceIPConfigurationsClient) GetHandleResponse(resp *azcore.Response) (*NetworkInterfaceIPConfigurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := NetworkInterfaceIPConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceIPConfiguration)
}

// GetHandleError handles the Get error response.
func (client *NetworkInterfaceIPConfigurationsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Get all ip configurations in a network interface.
func (client *NetworkInterfaceIPConfigurationsClient) List(resourceGroupName string, networkInterfaceName string) NetworkInterfaceIPConfigurationListResultPager {
	return &networkInterfaceIPConfigurationListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, networkInterfaceName)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *NetworkInterfaceIPConfigurationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceIPConfigurationListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *NetworkInterfaceIPConfigurationsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
func (client *NetworkInterfaceIPConfigurationsClient) ListHandleResponse(resp *azcore.Response) (*NetworkInterfaceIPConfigurationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := NetworkInterfaceIPConfigurationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceIPConfigurationListResult)
}

// ListHandleError handles the List error response.
func (client *NetworkInterfaceIPConfigurationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
