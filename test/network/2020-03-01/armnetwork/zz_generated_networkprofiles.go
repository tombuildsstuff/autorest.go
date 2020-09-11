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

// NetworkProfilesOperations contains the methods for the NetworkProfiles group.
type NetworkProfilesOperations interface {
	// CreateOrUpdate - Creates or updates a network profile.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkProfileName string, parameters NetworkProfile) (*NetworkProfileResponse, error)
	// BeginDelete - Deletes the specified network profile.
	BeginDelete(ctx context.Context, resourceGroupName string, networkProfileName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified network profile in a specified resource group.
	Get(ctx context.Context, resourceGroupName string, networkProfileName string, networkProfilesGetOptions *NetworkProfilesGetOptions) (*NetworkProfileResponse, error)
	// List - Gets all network profiles in a resource group.
	List(resourceGroupName string) NetworkProfileListResultPager
	// ListAll - Gets all the network profiles in a subscription.
	ListAll() NetworkProfileListResultPager
	// UpdateTags - Updates network profile tags.
	UpdateTags(ctx context.Context, resourceGroupName string, networkProfileName string, parameters TagsObject) (*NetworkProfileResponse, error)
}

// NetworkProfilesClient implements the NetworkProfilesOperations interface.
// Don't use this type directly, use NewNetworkProfilesClient() instead.
type NetworkProfilesClient struct {
	*Client
	subscriptionID string
}

// NewNetworkProfilesClient creates a new instance of NetworkProfilesClient with the specified values.
func NewNetworkProfilesClient(c *Client, subscriptionID string) NetworkProfilesOperations {
	return &NetworkProfilesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *NetworkProfilesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates a network profile.
func (client *NetworkProfilesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkProfileName string, parameters NetworkProfile) (*NetworkProfileResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, networkProfileName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NetworkProfilesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, parameters NetworkProfile) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NetworkProfilesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*NetworkProfileResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result := NetworkProfileResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkProfile)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NetworkProfilesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified network profile.
func (client *NetworkProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkProfileName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, networkProfileName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("NetworkProfilesClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkProfilesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkProfilesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *NetworkProfilesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *NetworkProfilesClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *NetworkProfilesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified network profile in a specified resource group.
func (client *NetworkProfilesClient) Get(ctx context.Context, resourceGroupName string, networkProfileName string, networkProfilesGetOptions *NetworkProfilesGetOptions) (*NetworkProfileResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkProfileName, networkProfilesGetOptions)
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
func (client *NetworkProfilesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, networkProfilesGetOptions *NetworkProfilesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if networkProfilesGetOptions != nil && networkProfilesGetOptions.Expand != nil {
		query.Set("$expand", *networkProfilesGetOptions.Expand)
	}
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *NetworkProfilesClient) GetHandleResponse(resp *azcore.Response) (*NetworkProfileResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := NetworkProfileResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkProfile)
}

// GetHandleError handles the Get error response.
func (client *NetworkProfilesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all network profiles in a resource group.
func (client *NetworkProfilesClient) List(resourceGroupName string) NetworkProfileListResultPager {
	return &networkProfileListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *NetworkProfileListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkProfileListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *NetworkProfilesClient) ListCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *NetworkProfilesClient) ListHandleResponse(resp *azcore.Response) (*NetworkProfileListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := NetworkProfileListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkProfileListResult)
}

// ListHandleError handles the List error response.
func (client *NetworkProfilesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the network profiles in a subscription.
func (client *NetworkProfilesClient) ListAll() NetworkProfileListResultPager {
	return &networkProfileListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx)
		},
		responder: client.ListAllHandleResponse,
		advancer: func(ctx context.Context, resp *NetworkProfileListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkProfileListResult.NextLink)
		},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *NetworkProfilesClient) ListAllCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkProfiles"
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

// ListAllHandleResponse handles the ListAll response.
func (client *NetworkProfilesClient) ListAllHandleResponse(resp *azcore.Response) (*NetworkProfileListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListAllHandleError(resp)
	}
	result := NetworkProfileListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkProfileListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *NetworkProfilesClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates network profile tags.
func (client *NetworkProfilesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkProfileName string, parameters TagsObject) (*NetworkProfileResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, networkProfileName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *NetworkProfilesClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *NetworkProfilesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*NetworkProfileResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result := NetworkProfileResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkProfile)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *NetworkProfilesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
