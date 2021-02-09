package lrogroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LRORetrysClient is the long-running Operation for AutoRest
type LRORetrysClient struct {
	BaseClient
}

// NewLRORetrysClient creates an instance of the LRORetrysClient client.
func NewLRORetrysClient() LRORetrysClient {
	return NewLRORetrysClientWithBaseURI(DefaultBaseURI)
}

// NewLRORetrysClientWithBaseURI creates an instance of the LRORetrysClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLRORetrysClientWithBaseURI(baseURI string) LRORetrysClient {
	return LRORetrysClient{NewWithBaseURI(baseURI)}
}

// Delete202Retry200 long running delete request, service returns a 500, then a 202 to the initial request. Polls
// return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LRORetrysClient) Delete202Retry200(ctx context.Context) (result LRORetrysDelete202Retry200Future, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.Delete202Retry200")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Delete202Retry200Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Delete202Retry200", nil, "Failure preparing request")
		return
	}

	result, err = client.Delete202Retry200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Delete202Retry200", nil, "Failure sending request")
		return
	}

	return
}

// Delete202Retry200Preparer prepares the Delete202Retry200 request.
func (client LRORetrysClient) Delete202Retry200Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/delete/202/retry/200"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Delete202Retry200Sender sends the Delete202Retry200 request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) Delete202Retry200Sender(req *http.Request) (future LRORetrysDelete202Retry200Future, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysDelete202Retry200Future", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysDelete202Retry200Future")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// Delete202Retry200Responder handles the response to the Delete202Retry200 request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) Delete202Retry200Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAsyncRelativeRetrySucceeded long running delete request, service returns a 500, then a 202 to the initial
// request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceeded(ctx context.Context) (result LRORetrysDeleteAsyncRelativeRetrySucceededFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.DeleteAsyncRelativeRetrySucceeded")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteAsyncRelativeRetrySucceededPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteAsyncRelativeRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteAsyncRelativeRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteAsyncRelativeRetrySucceeded", nil, "Failure sending request")
		return
	}

	return
}

// DeleteAsyncRelativeRetrySucceededPreparer prepares the DeleteAsyncRelativeRetrySucceeded request.
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceededPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/deleteasync/retry/succeeded"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAsyncRelativeRetrySucceededSender sends the DeleteAsyncRelativeRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceededSender(req *http.Request) (future LRORetrysDeleteAsyncRelativeRetrySucceededFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysDeleteAsyncRelativeRetrySucceededFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysDeleteAsyncRelativeRetrySucceededFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteAsyncRelativeRetrySucceededResponder handles the response to the DeleteAsyncRelativeRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceededResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteProvisioning202Accepted200Succeeded long running delete request, service returns a 500, then a  202 to the
// initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last
// poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LRORetrysClient) DeleteProvisioning202Accepted200Succeeded(ctx context.Context) (result LRORetrysDeleteProvisioning202Accepted200SucceededFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.DeleteProvisioning202Accepted200Succeeded")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteProvisioning202Accepted200SucceededPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteProvisioning202Accepted200Succeeded", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteProvisioning202Accepted200SucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteProvisioning202Accepted200Succeeded", nil, "Failure sending request")
		return
	}

	return
}

// DeleteProvisioning202Accepted200SucceededPreparer prepares the DeleteProvisioning202Accepted200Succeeded request.
func (client LRORetrysClient) DeleteProvisioning202Accepted200SucceededPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteProvisioning202Accepted200SucceededSender sends the DeleteProvisioning202Accepted200Succeeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) DeleteProvisioning202Accepted200SucceededSender(req *http.Request) (future LRORetrysDeleteProvisioning202Accepted200SucceededFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (p Product, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysDeleteProvisioning202Accepted200SucceededFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysDeleteProvisioning202Accepted200SucceededFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysDeleteProvisioning202Accepted200SucceededFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.DeleteProvisioning202Accepted200SucceededResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysDeleteProvisioning202Accepted200SucceededFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// DeleteProvisioning202Accepted200SucceededResponder handles the response to the DeleteProvisioning202Accepted200Succeeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) DeleteProvisioning202Accepted200SucceededResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Post202Retry200 long running post request, service returns a 500, then a 202 to the initial request, with 'Location'
// and 'Retry-After' headers, Polls return a 200 with a response body after success
// Parameters:
// product - product to put
func (client LRORetrysClient) Post202Retry200(ctx context.Context, product *Product) (result LRORetrysPost202Retry200Future, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.Post202Retry200")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Post202Retry200Preparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Post202Retry200", nil, "Failure preparing request")
		return
	}

	result, err = client.Post202Retry200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Post202Retry200", nil, "Failure sending request")
		return
	}

	return
}

// Post202Retry200Preparer prepares the Post202Retry200 request.
func (client LRORetrysClient) Post202Retry200Preparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/post/202/retry/200"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Post202Retry200Sender sends the Post202Retry200 request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) Post202Retry200Sender(req *http.Request) (future LRORetrysPost202Retry200Future, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPost202Retry200Future", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysPost202Retry200Future")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// Post202Retry200Responder handles the response to the Post202Retry200 request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) Post202Retry200Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PostAsyncRelativeRetrySucceeded long running post request, service returns a 500, then a 202 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// Parameters:
// product - product to put
func (client LRORetrysClient) PostAsyncRelativeRetrySucceeded(ctx context.Context, product *Product) (result LRORetrysPostAsyncRelativeRetrySucceededFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.PostAsyncRelativeRetrySucceeded")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PostAsyncRelativeRetrySucceededPreparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PostAsyncRelativeRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.PostAsyncRelativeRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PostAsyncRelativeRetrySucceeded", nil, "Failure sending request")
		return
	}

	return
}

// PostAsyncRelativeRetrySucceededPreparer prepares the PostAsyncRelativeRetrySucceeded request.
func (client LRORetrysClient) PostAsyncRelativeRetrySucceededPreparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/postasync/retry/succeeded"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostAsyncRelativeRetrySucceededSender sends the PostAsyncRelativeRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) PostAsyncRelativeRetrySucceededSender(req *http.Request) (future LRORetrysPostAsyncRelativeRetrySucceededFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPostAsyncRelativeRetrySucceededFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysPostAsyncRelativeRetrySucceededFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// PostAsyncRelativeRetrySucceededResponder handles the response to the PostAsyncRelativeRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) PostAsyncRelativeRetrySucceededResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Put201CreatingSucceeded200 long running put request, service returns a 500, then a 201 to the initial request, with
// an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’
// with ProvisioningState=’Succeeded’
// Parameters:
// product - product to put
func (client LRORetrysClient) Put201CreatingSucceeded200(ctx context.Context, product *Product) (result LRORetrysPut201CreatingSucceeded200Future, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.Put201CreatingSucceeded200")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Put201CreatingSucceeded200Preparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Put201CreatingSucceeded200", nil, "Failure preparing request")
		return
	}

	result, err = client.Put201CreatingSucceeded200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Put201CreatingSucceeded200", nil, "Failure sending request")
		return
	}

	return
}

// Put201CreatingSucceeded200Preparer prepares the Put201CreatingSucceeded200 request.
func (client LRORetrysClient) Put201CreatingSucceeded200Preparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/put/201/creating/succeeded/200"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Put201CreatingSucceeded200Sender sends the Put201CreatingSucceeded200 request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) Put201CreatingSucceeded200Sender(req *http.Request) (future LRORetrysPut201CreatingSucceeded200Future, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (p Product, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPut201CreatingSucceeded200Future", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysPut201CreatingSucceeded200Future")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPut201CreatingSucceeded200Future", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.Put201CreatingSucceeded200Responder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPut201CreatingSucceeded200Future", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// Put201CreatingSucceeded200Responder handles the response to the Put201CreatingSucceeded200 request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) Put201CreatingSucceeded200Responder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutAsyncRelativeRetrySucceeded long running put request, service returns a 500, then a 200 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// Parameters:
// product - product to put
func (client LRORetrysClient) PutAsyncRelativeRetrySucceeded(ctx context.Context, product *Product) (result LRORetrysPutAsyncRelativeRetrySucceededFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LRORetrysClient.PutAsyncRelativeRetrySucceeded")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutAsyncRelativeRetrySucceededPreparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PutAsyncRelativeRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.PutAsyncRelativeRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PutAsyncRelativeRetrySucceeded", nil, "Failure sending request")
		return
	}

	return
}

// PutAsyncRelativeRetrySucceededPreparer prepares the PutAsyncRelativeRetrySucceeded request.
func (client LRORetrysClient) PutAsyncRelativeRetrySucceededPreparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/putasync/retry/succeeded"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutAsyncRelativeRetrySucceededSender sends the PutAsyncRelativeRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) PutAsyncRelativeRetrySucceededSender(req *http.Request) (future LRORetrysPutAsyncRelativeRetrySucceededFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LRORetrysClient) (p Product, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPutAsyncRelativeRetrySucceededFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("lrogroup.LRORetrysPutAsyncRelativeRetrySucceededFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPutAsyncRelativeRetrySucceededFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.PutAsyncRelativeRetrySucceededResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysPutAsyncRelativeRetrySucceededFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// PutAsyncRelativeRetrySucceededResponder handles the response to the PutAsyncRelativeRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) PutAsyncRelativeRetrySucceededResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
