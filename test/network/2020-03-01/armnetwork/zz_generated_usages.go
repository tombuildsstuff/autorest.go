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

// UsagesOperations contains the methods for the Usages group.
type UsagesOperations interface {
	// List - List network usages for a subscription.
	List(location string) UsagesListResultPager
}

// UsagesClient implements the UsagesOperations interface.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	*Client
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
func NewUsagesClient(c *Client, subscriptionID string) UsagesOperations {
	return &UsagesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *UsagesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - List network usages for a subscription.
func (client *UsagesClient) List(location string) UsagesListResultPager {
	return &usagesListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, location)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *UsagesListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.UsagesListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *UsagesClient) ListCreateRequest(ctx context.Context, location string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *UsagesClient) ListHandleResponse(resp *azcore.Response) (*UsagesListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := UsagesListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.UsagesListResult)
}

// ListHandleError handles the List error response.
func (client *UsagesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
