// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type sparkBatchClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *sparkBatchClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CancelSparkBatchJob - Cancels a running spark batch job.
func (client *sparkBatchClient) CancelSparkBatchJob(ctx context.Context, batchId int32) (*http.Response, error) {
	req, err := client.CancelSparkBatchJobCreateRequest(ctx, batchId)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CancelSparkBatchJobHandleError(resp)
	}
	return resp.Response, nil
}

// CancelSparkBatchJobCreateRequest creates the CancelSparkBatchJob request.
func (client *sparkBatchClient) CancelSparkBatchJobCreateRequest(ctx context.Context, batchId int32) (*azcore.Request, error) {
	urlPath := "/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CancelSparkBatchJobHandleError handles the CancelSparkBatchJob error response.
func (client *sparkBatchClient) CancelSparkBatchJobHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// CreateSparkBatchJob - Create new spark batch job.
func (client *sparkBatchClient) CreateSparkBatchJob(ctx context.Context, sparkBatchJobOptions SparkBatchJobOptions, sparkBatchCreateSparkBatchJobOptions *SparkBatchCreateSparkBatchJobOptions) (*SparkBatchJobResponse, error) {
	req, err := client.CreateSparkBatchJobCreateRequest(ctx, sparkBatchJobOptions, sparkBatchCreateSparkBatchJobOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateSparkBatchJobHandleError(resp)
	}
	result, err := client.CreateSparkBatchJobHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSparkBatchJobCreateRequest creates the CreateSparkBatchJob request.
func (client *sparkBatchClient) CreateSparkBatchJobCreateRequest(ctx context.Context, sparkBatchJobOptions SparkBatchJobOptions, sparkBatchCreateSparkBatchJobOptions *SparkBatchCreateSparkBatchJobOptions) (*azcore.Request, error) {
	urlPath := "/batches"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if sparkBatchCreateSparkBatchJobOptions != nil && sparkBatchCreateSparkBatchJobOptions.Detailed != nil {
		query.Set("detailed", strconv.FormatBool(*sparkBatchCreateSparkBatchJobOptions.Detailed))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkBatchJobOptions)
}

// CreateSparkBatchJobHandleResponse handles the CreateSparkBatchJob response.
func (client *sparkBatchClient) CreateSparkBatchJobHandleResponse(resp *azcore.Response) (*SparkBatchJobResponse, error) {
	result := SparkBatchJobResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkBatchJob)
}

// CreateSparkBatchJobHandleError handles the CreateSparkBatchJob error response.
func (client *sparkBatchClient) CreateSparkBatchJobHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkBatchJob - Gets a single spark batch job.
func (client *sparkBatchClient) GetSparkBatchJob(ctx context.Context, batchId int32, sparkBatchGetSparkBatchJobOptions *SparkBatchGetSparkBatchJobOptions) (*SparkBatchJobResponse, error) {
	req, err := client.GetSparkBatchJobCreateRequest(ctx, batchId, sparkBatchGetSparkBatchJobOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetSparkBatchJobHandleError(resp)
	}
	result, err := client.GetSparkBatchJobHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSparkBatchJobCreateRequest creates the GetSparkBatchJob request.
func (client *sparkBatchClient) GetSparkBatchJobCreateRequest(ctx context.Context, batchId int32, sparkBatchGetSparkBatchJobOptions *SparkBatchGetSparkBatchJobOptions) (*azcore.Request, error) {
	urlPath := "/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if sparkBatchGetSparkBatchJobOptions != nil && sparkBatchGetSparkBatchJobOptions.Detailed != nil {
		query.Set("detailed", strconv.FormatBool(*sparkBatchGetSparkBatchJobOptions.Detailed))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSparkBatchJobHandleResponse handles the GetSparkBatchJob response.
func (client *sparkBatchClient) GetSparkBatchJobHandleResponse(resp *azcore.Response) (*SparkBatchJobResponse, error) {
	result := SparkBatchJobResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkBatchJob)
}

// GetSparkBatchJobHandleError handles the GetSparkBatchJob error response.
func (client *sparkBatchClient) GetSparkBatchJobHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkBatchJobs - List all spark batch jobs which are running under a particular spark pool.
func (client *sparkBatchClient) GetSparkBatchJobs(ctx context.Context, sparkBatchGetSparkBatchJobsOptions *SparkBatchGetSparkBatchJobsOptions) (*SparkBatchJobCollectionResponse, error) {
	req, err := client.GetSparkBatchJobsCreateRequest(ctx, sparkBatchGetSparkBatchJobsOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetSparkBatchJobsHandleError(resp)
	}
	result, err := client.GetSparkBatchJobsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSparkBatchJobsCreateRequest creates the GetSparkBatchJobs request.
func (client *sparkBatchClient) GetSparkBatchJobsCreateRequest(ctx context.Context, sparkBatchGetSparkBatchJobsOptions *SparkBatchGetSparkBatchJobsOptions) (*azcore.Request, error) {
	urlPath := "/batches"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if sparkBatchGetSparkBatchJobsOptions != nil && sparkBatchGetSparkBatchJobsOptions.From != nil {
		query.Set("from", strconv.FormatInt(int64(*sparkBatchGetSparkBatchJobsOptions.From), 10))
	}
	if sparkBatchGetSparkBatchJobsOptions != nil && sparkBatchGetSparkBatchJobsOptions.Size != nil {
		query.Set("size", strconv.FormatInt(int64(*sparkBatchGetSparkBatchJobsOptions.Size), 10))
	}
	if sparkBatchGetSparkBatchJobsOptions != nil && sparkBatchGetSparkBatchJobsOptions.Detailed != nil {
		query.Set("detailed", strconv.FormatBool(*sparkBatchGetSparkBatchJobsOptions.Detailed))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSparkBatchJobsHandleResponse handles the GetSparkBatchJobs response.
func (client *sparkBatchClient) GetSparkBatchJobsHandleResponse(resp *azcore.Response) (*SparkBatchJobCollectionResponse, error) {
	result := SparkBatchJobCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkBatchJobCollection)
}

// GetSparkBatchJobsHandleError handles the GetSparkBatchJobs error response.
func (client *sparkBatchClient) GetSparkBatchJobsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}