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

// ServiceAssociationLinksOperations contains the methods for the ServiceAssociationLinks group.
type ServiceAssociationLinksOperations interface {
	// List - Gets a list of service association links for a subnet.
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*ServiceAssociationLinksListResultResponse, error)
}

// ServiceAssociationLinksClient implements the ServiceAssociationLinksOperations interface.
// Don't use this type directly, use NewServiceAssociationLinksClient() instead.
type ServiceAssociationLinksClient struct {
	*Client
	subscriptionID string
}

// NewServiceAssociationLinksClient creates a new instance of ServiceAssociationLinksClient with the specified values.
func NewServiceAssociationLinksClient(c *Client, subscriptionID string) ServiceAssociationLinksOperations {
	return &ServiceAssociationLinksClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ServiceAssociationLinksClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Gets a list of service association links for a subnet.
func (client *ServiceAssociationLinksClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*ServiceAssociationLinksListResultResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *ServiceAssociationLinksClient) ListCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
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
func (client *ServiceAssociationLinksClient) ListHandleResponse(resp *azcore.Response) (*ServiceAssociationLinksListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := ServiceAssociationLinksListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceAssociationLinksListResult)
}

// ListHandleError handles the List error response.
func (client *ServiceAssociationLinksClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
