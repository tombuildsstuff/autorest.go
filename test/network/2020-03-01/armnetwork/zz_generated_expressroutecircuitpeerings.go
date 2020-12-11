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

// ExpressRouteCircuitPeeringsClient contains the methods for the ExpressRouteCircuitPeerings group.
// Don't use this type directly, use NewExpressRouteCircuitPeeringsClient() instead.
type ExpressRouteCircuitPeeringsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCircuitPeeringsClient creates a new instance of ExpressRouteCircuitPeeringsClient with the specified values.
func NewExpressRouteCircuitPeeringsClient(con *armcore.Connection, subscriptionID string) *ExpressRouteCircuitPeeringsClient {
	return &ExpressRouteCircuitPeeringsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified express route circuits.
func (client *ExpressRouteCircuitPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering, options *ExpressRouteCircuitPeeringsBeginCreateOrUpdateOptions) (ExpressRouteCircuitPeeringPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, circuitName, peeringName, peeringParameters, options)
	if err != nil {
		return ExpressRouteCircuitPeeringPollerResponse{}, err
	}
	result := ExpressRouteCircuitPeeringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCircuitPeeringPollerResponse{}, err
	}
	poller := &expressRouteCircuitPeeringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCircuitPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCircuitPeeringPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitPeeringPoller.ResumeToken().
func (client *ExpressRouteCircuitPeeringsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCircuitPeeringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitPeeringsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitPeeringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a peering in the specified express route circuits.
func (client *ExpressRouteCircuitPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering, options *ExpressRouteCircuitPeeringsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, peeringName, peeringParameters, options)
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
func (client *ExpressRouteCircuitPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering, options *ExpressRouteCircuitPeeringsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
	return req, req.MarshalAsJSON(peeringParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExpressRouteCircuitPeeringsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ExpressRouteCircuitPeeringResponse, error) {
	var val *ExpressRouteCircuitPeering
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCircuitPeeringResponse{}, err
	}
	return ExpressRouteCircuitPeeringResponse{RawResponse: resp.Response, ExpressRouteCircuitPeering: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCircuitPeeringsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified peering from the specified express route circuit.
func (client *ExpressRouteCircuitPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitPeeringsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *ExpressRouteCircuitPeeringsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitPeeringsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified peering from the specified express route circuit.
func (client *ExpressRouteCircuitPeeringsClient) delete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
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
func (client *ExpressRouteCircuitPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *ExpressRouteCircuitPeeringsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified peering for the express route circuit.
func (client *ExpressRouteCircuitPeeringsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsGetOptions) (ExpressRouteCircuitPeeringResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return ExpressRouteCircuitPeeringResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCircuitPeeringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteCircuitPeeringResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCircuitPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// getHandleResponse handles the Get response.
func (client *ExpressRouteCircuitPeeringsClient) getHandleResponse(resp *azcore.Response) (ExpressRouteCircuitPeeringResponse, error) {
	var val *ExpressRouteCircuitPeering
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCircuitPeeringResponse{}, err
	}
	return ExpressRouteCircuitPeeringResponse{RawResponse: resp.Response, ExpressRouteCircuitPeering: val}, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteCircuitPeeringsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all peerings in a specified express route circuit.
func (client *ExpressRouteCircuitPeeringsClient) List(resourceGroupName string, circuitName string, options *ExpressRouteCircuitPeeringsListOptions) ExpressRouteCircuitPeeringListResultPager {
	return &expressRouteCircuitPeeringListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, circuitName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ExpressRouteCircuitPeeringListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitPeeringListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCircuitPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitPeeringsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *ExpressRouteCircuitPeeringsClient) listHandleResponse(resp *azcore.Response) (ExpressRouteCircuitPeeringListResultResponse, error) {
	var val *ExpressRouteCircuitPeeringListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCircuitPeeringListResultResponse{}, err
	}
	return ExpressRouteCircuitPeeringListResultResponse{RawResponse: resp.Response, ExpressRouteCircuitPeeringListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteCircuitPeeringsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
