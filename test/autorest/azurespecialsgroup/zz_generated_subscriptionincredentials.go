// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionInCredentialsOperations contains the methods for the SubscriptionInCredentials group.
type SubscriptionInCredentialsOperations interface {
	// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error)
	// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to null, and client-side validation should prevent you from making this call
	PostMethodGlobalNull(ctx context.Context) (*http.Response, error)
	// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostMethodGlobalValid(ctx context.Context) (*http.Response, error)
	// PostPathGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostPathGlobalValid(ctx context.Context) (*http.Response, error)
	// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostSwaggerGlobalValid(ctx context.Context) (*http.Response, error)
}

// SubscriptionInCredentialsClient implements the SubscriptionInCredentialsOperations interface.
// Don't use this type directly, use NewSubscriptionInCredentialsClient() instead.
type SubscriptionInCredentialsClient struct {
	*Client
	subscriptionID string
}

// NewSubscriptionInCredentialsClient creates a new instance of SubscriptionInCredentialsClient with the specified values.
func NewSubscriptionInCredentialsClient(c *Client, subscriptionID string) SubscriptionInCredentialsOperations {
	return &SubscriptionInCredentialsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *SubscriptionInCredentialsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error) {
	req, err := client.PostMethodGlobalNotProvidedValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PostMethodGlobalNotProvidedValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostMethodGlobalNotProvidedValidCreateRequest creates the PostMethodGlobalNotProvidedValid request.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/globalNotProvided/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2015-07-01-preview")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// PostMethodGlobalNotProvidedValidHandleResponse handles the PostMethodGlobalNotProvidedValid response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodGlobalNotProvidedValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodGlobalNotProvidedValidHandleError handles the PostMethodGlobalNotProvidedValid error response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to null, and client-side validation should prevent you from making this call
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNull(ctx context.Context) (*http.Response, error) {
	req, err := client.PostMethodGlobalNullCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PostMethodGlobalNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostMethodGlobalNullCreateRequest creates the PostMethodGlobalNull request.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNullCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/null/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PostMethodGlobalNullHandleResponse handles the PostMethodGlobalNull response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNullHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodGlobalNullHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodGlobalNullHandleError handles the PostMethodGlobalNull error response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.PostMethodGlobalValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PostMethodGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostMethodGlobalValidCreateRequest creates the PostMethodGlobalValid request.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PostMethodGlobalValidHandleResponse handles the PostMethodGlobalValid response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodGlobalValidHandleError handles the PostMethodGlobalValid error response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostPathGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostPathGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.PostPathGlobalValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PostPathGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostPathGlobalValidCreateRequest creates the PostPathGlobalValid request.
func (client *SubscriptionInCredentialsClient) PostPathGlobalValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PostPathGlobalValidHandleResponse handles the PostPathGlobalValid response.
func (client *SubscriptionInCredentialsClient) PostPathGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostPathGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostPathGlobalValidHandleError handles the PostPathGlobalValid error response.
func (client *SubscriptionInCredentialsClient) PostPathGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.PostSwaggerGlobalValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PostSwaggerGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostSwaggerGlobalValidCreateRequest creates the PostSwaggerGlobalValid request.
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PostSwaggerGlobalValidHandleResponse handles the PostSwaggerGlobalValid response.
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostSwaggerGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostSwaggerGlobalValidHandleError handles the PostSwaggerGlobalValid error response.
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
