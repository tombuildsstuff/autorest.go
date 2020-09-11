// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPRetryOperations contains the methods for the HTTPRetry group.
type HTTPRetryOperations interface {
	// Delete503 - Return 503 status code, then 200 after retry
	Delete503(ctx context.Context) (*http.Response, error)
	// Get502 - Return 502 status code, then 200 after retry
	Get502(ctx context.Context) (*http.Response, error)
	// Head408 - Return 408 status code, then 200 after retry
	Head408(ctx context.Context) (*http.Response, error)
	// Options502 - Return 502 status code, then 200 after retry
	Options502(ctx context.Context) (*BoolResponse, error)
	// Patch500 - Return 500 status code, then 200 after retry
	Patch500(ctx context.Context) (*http.Response, error)
	// Patch504 - Return 504 status code, then 200 after retry
	Patch504(ctx context.Context) (*http.Response, error)
	// Post503 - Return 503 status code, then 200 after retry
	Post503(ctx context.Context) (*http.Response, error)
	// Put500 - Return 500 status code, then 200 after retry
	Put500(ctx context.Context) (*http.Response, error)
	// Put504 - Return 504 status code, then 200 after retry
	Put504(ctx context.Context) (*http.Response, error)
}

// HTTPRetryClient implements the HTTPRetryOperations interface.
// Don't use this type directly, use NewHTTPRetryClient() instead.
type HTTPRetryClient struct {
	*Client
}

// NewHTTPRetryClient creates a new instance of HTTPRetryClient with the specified values.
func NewHTTPRetryClient(c *Client) HTTPRetryOperations {
	return &HTTPRetryClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *HTTPRetryClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Delete503 - Return 503 status code, then 200 after retry
func (client *HTTPRetryClient) Delete503(ctx context.Context) (*http.Response, error) {
	req, err := client.Delete503CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Delete503HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete503CreateRequest creates the Delete503 request.
func (client *HTTPRetryClient) Delete503CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/503"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(true)
}

// Delete503HandleResponse handles the Delete503 response.
func (client *HTTPRetryClient) Delete503HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Delete503HandleError(resp)
	}
	return resp.Response, nil
}

// Delete503HandleError handles the Delete503 error response.
func (client *HTTPRetryClient) Delete503HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get502 - Return 502 status code, then 200 after retry
func (client *HTTPRetryClient) Get502(ctx context.Context) (*http.Response, error) {
	req, err := client.Get502CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Get502HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get502CreateRequest creates the Get502 request.
func (client *HTTPRetryClient) Get502CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/502"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Get502HandleResponse handles the Get502 response.
func (client *HTTPRetryClient) Get502HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Get502HandleError(resp)
	}
	return resp.Response, nil
}

// Get502HandleError handles the Get502 error response.
func (client *HTTPRetryClient) Get502HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head408 - Return 408 status code, then 200 after retry
func (client *HTTPRetryClient) Head408(ctx context.Context) (*http.Response, error) {
	req, err := client.Head408CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Head408HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Head408CreateRequest creates the Head408 request.
func (client *HTTPRetryClient) Head408CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/408"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Head408HandleResponse handles the Head408 response.
func (client *HTTPRetryClient) Head408HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Head408HandleError(resp)
	}
	return resp.Response, nil
}

// Head408HandleError handles the Head408 error response.
func (client *HTTPRetryClient) Head408HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Options502 - Return 502 status code, then 200 after retry
func (client *HTTPRetryClient) Options502(ctx context.Context) (*BoolResponse, error) {
	req, err := client.Options502CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Options502HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Options502CreateRequest creates the Options502 request.
func (client *HTTPRetryClient) Options502CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/502"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Options502HandleResponse handles the Options502 response.
func (client *HTTPRetryClient) Options502HandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Options502HandleError(resp)
	}
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// Options502HandleError handles the Options502 error response.
func (client *HTTPRetryClient) Options502HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch500 - Return 500 status code, then 200 after retry
func (client *HTTPRetryClient) Patch500(ctx context.Context) (*http.Response, error) {
	req, err := client.Patch500CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Patch500HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Patch500CreateRequest creates the Patch500 request.
func (client *HTTPRetryClient) Patch500CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/500"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(true)
}

// Patch500HandleResponse handles the Patch500 response.
func (client *HTTPRetryClient) Patch500HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Patch500HandleError(resp)
	}
	return resp.Response, nil
}

// Patch500HandleError handles the Patch500 error response.
func (client *HTTPRetryClient) Patch500HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch504 - Return 504 status code, then 200 after retry
func (client *HTTPRetryClient) Patch504(ctx context.Context) (*http.Response, error) {
	req, err := client.Patch504CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Patch504HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Patch504CreateRequest creates the Patch504 request.
func (client *HTTPRetryClient) Patch504CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/504"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(true)
}

// Patch504HandleResponse handles the Patch504 response.
func (client *HTTPRetryClient) Patch504HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Patch504HandleError(resp)
	}
	return resp.Response, nil
}

// Patch504HandleError handles the Patch504 error response.
func (client *HTTPRetryClient) Patch504HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post503 - Return 503 status code, then 200 after retry
func (client *HTTPRetryClient) Post503(ctx context.Context) (*http.Response, error) {
	req, err := client.Post503CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Post503HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Post503CreateRequest creates the Post503 request.
func (client *HTTPRetryClient) Post503CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/503"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(true)
}

// Post503HandleResponse handles the Post503 response.
func (client *HTTPRetryClient) Post503HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Post503HandleError(resp)
	}
	return resp.Response, nil
}

// Post503HandleError handles the Post503 error response.
func (client *HTTPRetryClient) Post503HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put500 - Return 500 status code, then 200 after retry
func (client *HTTPRetryClient) Put500(ctx context.Context) (*http.Response, error) {
	req, err := client.Put500CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Put500HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put500CreateRequest creates the Put500 request.
func (client *HTTPRetryClient) Put500CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/500"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(true)
}

// Put500HandleResponse handles the Put500 response.
func (client *HTTPRetryClient) Put500HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Put500HandleError(resp)
	}
	return resp.Response, nil
}

// Put500HandleError handles the Put500 error response.
func (client *HTTPRetryClient) Put500HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put504 - Return 504 status code, then 200 after retry
func (client *HTTPRetryClient) Put504(ctx context.Context) (*http.Response, error) {
	req, err := client.Put504CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Put504HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put504CreateRequest creates the Put504 request.
func (client *HTTPRetryClient) Put504CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/retry/504"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.MarshalAsJSON(true)
}

// Put504HandleResponse handles the Put504 response.
func (client *HTTPRetryClient) Put504HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Put504HandleError(resp)
	}
	return resp.Response, nil
}

// Put504HandleError handles the Put504 error response.
func (client *HTTPRetryClient) Put504HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
