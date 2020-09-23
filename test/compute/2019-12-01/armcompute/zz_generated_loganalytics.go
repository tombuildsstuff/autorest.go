// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// LogAnalyticsOperations contains the methods for the LogAnalytics group.
type LogAnalyticsOperations interface {
	// BeginExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window to show throttling activities.
	BeginExportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput) (*LogAnalyticsOperationResultPollerResponse, error)
	// ResumeExportRequestRateByInterval - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeExportRequestRateByInterval(token string) (LogAnalyticsOperationResultPoller, error)
	// BeginExportThrottledRequests - Export logs that show total throttled Api requests for this subscription in the given time window.
	BeginExportThrottledRequests(ctx context.Context, location string, parameters LogAnalyticsInputBase) (*LogAnalyticsOperationResultPollerResponse, error)
	// ResumeExportThrottledRequests - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeExportThrottledRequests(token string) (LogAnalyticsOperationResultPoller, error)
}

// LogAnalyticsClient implements the LogAnalyticsOperations interface.
// Don't use this type directly, use NewLogAnalyticsClient() instead.
type LogAnalyticsClient struct {
	*Client
	subscriptionID string
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient with the specified values.
func NewLogAnalyticsClient(c *Client, subscriptionID string) LogAnalyticsOperations {
	return &LogAnalyticsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *LogAnalyticsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *LogAnalyticsClient) BeginExportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput) (*LogAnalyticsOperationResultPollerResponse, error) {
	resp, err := client.ExportRequestRateByInterval(ctx, location, parameters)
	if err != nil {
		return nil, err
	}
	result := &LogAnalyticsOperationResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LogAnalyticsClient.ExportRequestRateByInterval", "azure-async-operation", resp, client.ExportRequestRateByIntervalHandleError)
	if err != nil {
		return nil, err
	}
	poller := &logAnalyticsOperationResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*LogAnalyticsOperationResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LogAnalyticsClient) ResumeExportRequestRateByInterval(token string) (LogAnalyticsOperationResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LogAnalyticsClient.ExportRequestRateByInterval", token, client.ExportRequestRateByIntervalHandleError)
	if err != nil {
		return nil, err
	}
	return &logAnalyticsOperationResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// ExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window to show throttling activities.
func (client *LogAnalyticsClient) ExportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput) (*azcore.Response, error) {
	req, err := client.ExportRequestRateByIntervalCreateRequest(ctx, location, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ExportRequestRateByIntervalHandleError(resp)
	}
	return resp, nil
}

// ExportRequestRateByIntervalCreateRequest creates the ExportRequestRateByInterval request.
func (client *LogAnalyticsClient) ExportRequestRateByIntervalCreateRequest(ctx context.Context, location string, parameters RequestRateByIntervalInput) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/logAnalytics/apiAccess/getRequestRateByInterval"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// ExportRequestRateByIntervalHandleResponse handles the ExportRequestRateByInterval response.
func (client *LogAnalyticsClient) ExportRequestRateByIntervalHandleResponse(resp *azcore.Response) (*LogAnalyticsOperationResultResponse, error) {
	result := LogAnalyticsOperationResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogAnalyticsOperationResult)
}

// ExportRequestRateByIntervalHandleError handles the ExportRequestRateByInterval error response.
func (client *LogAnalyticsClient) ExportRequestRateByIntervalHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

func (client *LogAnalyticsClient) BeginExportThrottledRequests(ctx context.Context, location string, parameters LogAnalyticsInputBase) (*LogAnalyticsOperationResultPollerResponse, error) {
	resp, err := client.ExportThrottledRequests(ctx, location, parameters)
	if err != nil {
		return nil, err
	}
	result := &LogAnalyticsOperationResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LogAnalyticsClient.ExportThrottledRequests", "azure-async-operation", resp, client.ExportThrottledRequestsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &logAnalyticsOperationResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*LogAnalyticsOperationResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LogAnalyticsClient) ResumeExportThrottledRequests(token string) (LogAnalyticsOperationResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LogAnalyticsClient.ExportThrottledRequests", token, client.ExportThrottledRequestsHandleError)
	if err != nil {
		return nil, err
	}
	return &logAnalyticsOperationResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// ExportThrottledRequests - Export logs that show total throttled Api requests for this subscription in the given time window.
func (client *LogAnalyticsClient) ExportThrottledRequests(ctx context.Context, location string, parameters LogAnalyticsInputBase) (*azcore.Response, error) {
	req, err := client.ExportThrottledRequestsCreateRequest(ctx, location, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ExportThrottledRequestsHandleError(resp)
	}
	return resp, nil
}

// ExportThrottledRequestsCreateRequest creates the ExportThrottledRequests request.
func (client *LogAnalyticsClient) ExportThrottledRequestsCreateRequest(ctx context.Context, location string, parameters LogAnalyticsInputBase) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/logAnalytics/apiAccess/getThrottledRequests"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// ExportThrottledRequestsHandleResponse handles the ExportThrottledRequests response.
func (client *LogAnalyticsClient) ExportThrottledRequestsHandleResponse(resp *azcore.Response) (*LogAnalyticsOperationResultResponse, error) {
	result := LogAnalyticsOperationResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogAnalyticsOperationResult)
}

// ExportThrottledRequestsHandleError handles the ExportThrottledRequests error response.
func (client *LogAnalyticsClient) ExportThrottledRequestsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}