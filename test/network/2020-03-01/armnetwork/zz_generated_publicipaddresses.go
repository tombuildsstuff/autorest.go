// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PublicIPAddressesClient contains the methods for the PublicIPAddresses group.
// Don't use this type directly, use NewPublicIPAddressesClient() instead.
type PublicIPAddressesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPublicIPAddressesClient creates a new instance of PublicIPAddressesClient with the specified values.
func NewPublicIPAddressesClient(con *armcore.Connection, subscriptionID string) *PublicIPAddressesClient {
	return &PublicIPAddressesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP address.
func (client *PublicIPAddressesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesBeginCreateOrUpdateOptions) (PublicIPAddressPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return PublicIPAddressPollerResponse{}, err
	}
	result := PublicIPAddressPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPAddressesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return PublicIPAddressPollerResponse{}, err
	}
	poller := &publicIPAddressPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PublicIPAddressResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new PublicIPAddressPoller from the specified resume token.
// token - The value must come from a previous call to PublicIPAddressPoller.ResumeToken().
func (client *PublicIPAddressesClient) ResumeCreateOrUpdate(token string) (PublicIPAddressPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPAddressesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &publicIPAddressPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP address.
func (client *PublicIPAddressesClient) createOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PublicIPAddressesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PublicIPAddressesClient) createOrUpdateHandleResponse(resp *azcore.Response) (PublicIPAddressResponse, error) {
	var val *PublicIPAddress
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressResponse{}, err
	}
	return PublicIPAddressResponse{RawResponse: resp.Response, PublicIPAddress: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PublicIPAddressesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified public IP address.
func (client *PublicIPAddressesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, publicIPAddressName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPAddressesClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *PublicIPAddressesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPAddressesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified public IP address.
func (client *PublicIPAddressesClient) delete(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publicIPAddressName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PublicIPAddressesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PublicIPAddressesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified public IP address in a specified resource group.
func (client *PublicIPAddressesClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesGetOptions) (PublicIPAddressResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, publicIPAddressName, options)
	if err != nil {
		return PublicIPAddressResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PublicIPAddressResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PublicIPAddressResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PublicIPAddressesClient) getCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublicIPAddressesClient) getHandleResponse(resp *azcore.Response) (PublicIPAddressResponse, error) {
	var val *PublicIPAddress
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressResponse{}, err
	}
	return PublicIPAddressResponse{RawResponse: resp.Response, PublicIPAddress: val}, nil
}

// getHandleError handles the Get error response.
func (client *PublicIPAddressesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetVirtualMachineScaleSetPublicIPAddress - Get the specified public IP address in a virtual machine scale set.
func (client *PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *PublicIPAddressesGetVirtualMachineScaleSetPublicIPAddressOptions) (PublicIPAddressResponse, error) {
	req, err := client.getVirtualMachineScaleSetPublicIPAddressCreateRequest(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, ipConfigurationName, publicIPAddressName, options)
	if err != nil {
		return PublicIPAddressResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PublicIPAddressResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PublicIPAddressResponse{}, client.getVirtualMachineScaleSetPublicIPAddressHandleError(resp)
	}
	return client.getVirtualMachineScaleSetPublicIPAddressHandleResponse(resp)
}

// getVirtualMachineScaleSetPublicIPAddressCreateRequest creates the GetVirtualMachineScaleSetPublicIPAddress request.
func (client *PublicIPAddressesClient) getVirtualMachineScaleSetPublicIPAddressCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *PublicIPAddressesGetVirtualMachineScaleSetPublicIPAddressOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineScaleSetName}", url.PathEscape(virtualMachineScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualmachineIndex}", url.PathEscape(virtualmachineIndex))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2018-10-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getVirtualMachineScaleSetPublicIPAddressHandleResponse handles the GetVirtualMachineScaleSetPublicIPAddress response.
func (client *PublicIPAddressesClient) getVirtualMachineScaleSetPublicIPAddressHandleResponse(resp *azcore.Response) (PublicIPAddressResponse, error) {
	var val *PublicIPAddress
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressResponse{}, err
	}
	return PublicIPAddressResponse{RawResponse: resp.Response, PublicIPAddress: val}, nil
}

// getVirtualMachineScaleSetPublicIPAddressHandleError handles the GetVirtualMachineScaleSetPublicIPAddress error response.
func (client *PublicIPAddressesClient) getVirtualMachineScaleSetPublicIPAddressHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all public IP addresses in a resource group.
func (client *PublicIPAddressesClient) List(resourceGroupName string, options *PublicIPAddressesListOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *PublicIPAddressesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPAddressesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublicIPAddressesClient) listHandleResponse(resp *azcore.Response) (PublicIPAddressListResultResponse, error) {
	var val *PublicIPAddressListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressListResultResponse{}, err
	}
	return PublicIPAddressListResultResponse{RawResponse: resp.Response, PublicIPAddressListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *PublicIPAddressesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the public IP addresses in a subscription.
func (client *PublicIPAddressesClient) ListAll(options *PublicIPAddressesListAllOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *PublicIPAddressesClient) listAllCreateRequest(ctx context.Context, options *PublicIPAddressesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *PublicIPAddressesClient) listAllHandleResponse(resp *azcore.Response) (PublicIPAddressListResultResponse, error) {
	var val *PublicIPAddressListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressListResultResponse{}, err
	}
	return PublicIPAddressListResultResponse{RawResponse: resp.Response, PublicIPAddressListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *PublicIPAddressesClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListVirtualMachineScaleSetPublicIPAddresses - Gets information about all public IP addresses on a virtual machine scale set level.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddresses(resourceGroupName string, virtualMachineScaleSetName string, options *PublicIPAddressesListVirtualMachineScaleSetPublicIPAddressesOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listVirtualMachineScaleSetPublicIPAddressesCreateRequest(ctx, resourceGroupName, virtualMachineScaleSetName, options)
		},
		responder: client.listVirtualMachineScaleSetPublicIPAddressesHandleResponse,
		errorer:   client.listVirtualMachineScaleSetPublicIPAddressesHandleError,
		advancer: func(ctx context.Context, resp PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listVirtualMachineScaleSetPublicIPAddressesCreateRequest creates the ListVirtualMachineScaleSetPublicIPAddresses request.
func (client *PublicIPAddressesClient) listVirtualMachineScaleSetPublicIPAddressesCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, options *PublicIPAddressesListVirtualMachineScaleSetPublicIPAddressesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineScaleSetName}", url.PathEscape(virtualMachineScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2018-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listVirtualMachineScaleSetPublicIPAddressesHandleResponse handles the ListVirtualMachineScaleSetPublicIPAddresses response.
func (client *PublicIPAddressesClient) listVirtualMachineScaleSetPublicIPAddressesHandleResponse(resp *azcore.Response) (PublicIPAddressListResultResponse, error) {
	var val *PublicIPAddressListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressListResultResponse{}, err
	}
	return PublicIPAddressListResultResponse{RawResponse: resp.Response, PublicIPAddressListResult: val}, nil
}

// listVirtualMachineScaleSetPublicIPAddressesHandleError handles the ListVirtualMachineScaleSetPublicIPAddresses error response.
func (client *PublicIPAddressesClient) listVirtualMachineScaleSetPublicIPAddressesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListVirtualMachineScaleSetVMPublicIPaddresses - Gets information about all public IP addresses in a virtual machine IP configuration in a virtual machine
// scale set.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPaddresses(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, options *PublicIPAddressesListVirtualMachineScaleSetVMPublicIPaddressesOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listVirtualMachineScaleSetVMPublicIpaddressesCreateRequest(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, ipConfigurationName, options)
		},
		responder: client.listVirtualMachineScaleSetVMPublicIpaddressesHandleResponse,
		errorer:   client.listVirtualMachineScaleSetVMPublicIpaddressesHandleError,
		advancer: func(ctx context.Context, resp PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listVirtualMachineScaleSetVMPublicIpaddressesCreateRequest creates the ListVirtualMachineScaleSetVMPublicIPaddresses request.
func (client *PublicIPAddressesClient) listVirtualMachineScaleSetVMPublicIpaddressesCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, options *PublicIPAddressesListVirtualMachineScaleSetVMPublicIPaddressesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineScaleSetName}", url.PathEscape(virtualMachineScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualmachineIndex}", url.PathEscape(virtualmachineIndex))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2018-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listVirtualMachineScaleSetVMPublicIpaddressesHandleResponse handles the ListVirtualMachineScaleSetVMPublicIPaddresses response.
func (client *PublicIPAddressesClient) listVirtualMachineScaleSetVMPublicIpaddressesHandleResponse(resp *azcore.Response) (PublicIPAddressListResultResponse, error) {
	var val *PublicIPAddressListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressListResultResponse{}, err
	}
	return PublicIPAddressListResultResponse{RawResponse: resp.Response, PublicIPAddressListResult: val}, nil
}

// listVirtualMachineScaleSetVMPublicIpaddressesHandleError handles the ListVirtualMachineScaleSetVMPublicIPaddresses error response.
func (client *PublicIPAddressesClient) listVirtualMachineScaleSetVMPublicIpaddressesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates public IP address tags.
func (client *PublicIPAddressesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesUpdateTagsOptions) (PublicIPAddressResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return PublicIPAddressResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PublicIPAddressResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PublicIPAddressResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PublicIPAddressesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PublicIPAddressesClient) updateTagsHandleResponse(resp *azcore.Response) (PublicIPAddressResponse, error) {
	var val *PublicIPAddress
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPAddressResponse{}, err
	}
	return PublicIPAddressResponse{RawResponse: resp.Response, PublicIPAddress: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *PublicIPAddressesClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
