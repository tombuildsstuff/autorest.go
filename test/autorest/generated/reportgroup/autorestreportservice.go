// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package reportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// AutoRestReportServiceOperations contains the methods for the AutoRestReportService group.
type AutoRestReportServiceOperations interface {
	// GetOptionalReport - Get optional test coverage report
	GetOptionalReport(ctx context.Context, autoRestReportServiceGetOptionalReportOptions *AutoRestReportServiceGetOptionalReportOptions) (*MapOfInt32Response, error)
	// GetReport - Get test coverage report
	GetReport(ctx context.Context, autoRestReportServiceGetReportOptions *AutoRestReportServiceGetReportOptions) (*MapOfInt32Response, error)
}

// autoRestReportServiceOperations implements the AutoRestReportServiceOperations interface.
type autoRestReportServiceOperations struct {
	*Client
}

// GetOptionalReport - Get optional test coverage report
func (client *autoRestReportServiceOperations) GetOptionalReport(ctx context.Context, autoRestReportServiceGetOptionalReportOptions *AutoRestReportServiceGetOptionalReportOptions) (*MapOfInt32Response, error) {
	req, err := client.getOptionalReportCreateRequest(autoRestReportServiceGetOptionalReportOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getOptionalReportHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getOptionalReportCreateRequest creates the GetOptionalReport request.
func (client *autoRestReportServiceOperations) getOptionalReportCreateRequest(autoRestReportServiceGetOptionalReportOptions *AutoRestReportServiceGetOptionalReportOptions) (*azcore.Request, error) {
	urlPath := "/report/optional"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if autoRestReportServiceGetOptionalReportOptions != nil && autoRestReportServiceGetOptionalReportOptions.Qualifier != nil {
		query.Set("qualifier", *autoRestReportServiceGetOptionalReportOptions.Qualifier)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getOptionalReportHandleResponse handles the GetOptionalReport response.
func (client *autoRestReportServiceOperations) getOptionalReportHandleResponse(resp *azcore.Response) (*MapOfInt32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getOptionalReportHandleError(resp)
	}
	result := MapOfInt32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getOptionalReportHandleError handles the GetOptionalReport error response.
func (client *autoRestReportServiceOperations) getOptionalReportHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetReport - Get test coverage report
func (client *autoRestReportServiceOperations) GetReport(ctx context.Context, autoRestReportServiceGetReportOptions *AutoRestReportServiceGetReportOptions) (*MapOfInt32Response, error) {
	req, err := client.getReportCreateRequest(autoRestReportServiceGetReportOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getReportHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getReportCreateRequest creates the GetReport request.
func (client *autoRestReportServiceOperations) getReportCreateRequest(autoRestReportServiceGetReportOptions *AutoRestReportServiceGetReportOptions) (*azcore.Request, error) {
	urlPath := "/report"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if autoRestReportServiceGetReportOptions != nil && autoRestReportServiceGetReportOptions.Qualifier != nil {
		query.Set("qualifier", *autoRestReportServiceGetReportOptions.Qualifier)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *autoRestReportServiceOperations) getReportHandleResponse(resp *azcore.Response) (*MapOfInt32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getReportHandleError(resp)
	}
	result := MapOfInt32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getReportHandleError handles the GetReport error response.
func (client *autoRestReportServiceOperations) getReportHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}