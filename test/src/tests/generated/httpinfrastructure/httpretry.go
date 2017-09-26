package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// HTTPRetryClient is the test Infrastructure for AutoRest
type HTTPRetryClient struct {
	ManagementClient
}

// NewHTTPRetryClient creates an instance of the HTTPRetryClient client.
func NewHTTPRetryClient() HTTPRetryClient {
	return NewHTTPRetryClientWithBaseURI(DefaultBaseURI)
}

// NewHTTPRetryClientWithBaseURI creates an instance of the HTTPRetryClient client.
func NewHTTPRetryClientWithBaseURI(baseURI string) HTTPRetryClient {
	return HTTPRetryClient{NewWithBaseURI(baseURI)}
}

// Delete503 return 503 status code, then 200 after retry
//
// booleanValue is simple boolean value true
func (client HTTPRetryClient) Delete503(booleanValue *bool) (result autorest.Response, err error) {
	req, err := client.Delete503Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Delete503", nil, "Failure preparing request")
		return
	}

	resp, err := client.Delete503Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Delete503", resp, "Failure sending request")
		return
	}

	result, err = client.Delete503Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Delete503", resp, "Failure responding to request")
	}

	return
}

// Delete503Preparer prepares the Delete503 request.
func (client HTTPRetryClient) Delete503Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/503"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Delete503Sender sends the Delete503 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Delete503Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Delete503Responder handles the response to the Delete503 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Delete503Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get502 return 502 status code, then 200 after retry
func (client HTTPRetryClient) Get502() (result autorest.Response, err error) {
	req, err := client.Get502Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Get502", nil, "Failure preparing request")
		return
	}

	resp, err := client.Get502Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Get502", resp, "Failure sending request")
		return
	}

	result, err = client.Get502Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Get502", resp, "Failure responding to request")
	}

	return
}

// Get502Preparer prepares the Get502 request.
func (client HTTPRetryClient) Get502Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/502"))
	return preparer.Prepare(&http.Request{})
}

// Get502Sender sends the Get502 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Get502Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Get502Responder handles the response to the Get502 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Get502Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Head408 return 408 status code, then 200 after retry
func (client HTTPRetryClient) Head408() (result autorest.Response, err error) {
	req, err := client.Head408Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Head408", nil, "Failure preparing request")
		return
	}

	resp, err := client.Head408Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Head408", resp, "Failure sending request")
		return
	}

	result, err = client.Head408Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Head408", resp, "Failure responding to request")
	}

	return
}

// Head408Preparer prepares the Head408 request.
func (client HTTPRetryClient) Head408Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/408"))
	return preparer.Prepare(&http.Request{})
}

// Head408Sender sends the Head408 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Head408Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Head408Responder handles the response to the Head408 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Head408Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Patch500 return 500 status code, then 200 after retry
//
// booleanValue is simple boolean value true
func (client HTTPRetryClient) Patch500(booleanValue *bool) (result autorest.Response, err error) {
	req, err := client.Patch500Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Patch500", nil, "Failure preparing request")
		return
	}

	resp, err := client.Patch500Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Patch500", resp, "Failure sending request")
		return
	}

	result, err = client.Patch500Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Patch500", resp, "Failure responding to request")
	}

	return
}

// Patch500Preparer prepares the Patch500 request.
func (client HTTPRetryClient) Patch500Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/500"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Patch500Sender sends the Patch500 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Patch500Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Patch500Responder handles the response to the Patch500 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Patch500Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Patch504 return 504 status code, then 200 after retry
//
// booleanValue is simple boolean value true
func (client HTTPRetryClient) Patch504(booleanValue *bool) (result autorest.Response, err error) {
	req, err := client.Patch504Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Patch504", nil, "Failure preparing request")
		return
	}

	resp, err := client.Patch504Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Patch504", resp, "Failure sending request")
		return
	}

	result, err = client.Patch504Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Patch504", resp, "Failure responding to request")
	}

	return
}

// Patch504Preparer prepares the Patch504 request.
func (client HTTPRetryClient) Patch504Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/504"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Patch504Sender sends the Patch504 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Patch504Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Patch504Responder handles the response to the Patch504 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Patch504Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Post503 return 503 status code, then 200 after retry
//
// booleanValue is simple boolean value true
func (client HTTPRetryClient) Post503(booleanValue *bool) (result autorest.Response, err error) {
	req, err := client.Post503Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Post503", nil, "Failure preparing request")
		return
	}

	resp, err := client.Post503Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Post503", resp, "Failure sending request")
		return
	}

	result, err = client.Post503Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Post503", resp, "Failure responding to request")
	}

	return
}

// Post503Preparer prepares the Post503 request.
func (client HTTPRetryClient) Post503Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/503"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Post503Sender sends the Post503 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Post503Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Post503Responder handles the response to the Post503 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Post503Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Put500 return 500 status code, then 200 after retry
//
// booleanValue is simple boolean value true
func (client HTTPRetryClient) Put500(booleanValue *bool) (result autorest.Response, err error) {
	req, err := client.Put500Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Put500", nil, "Failure preparing request")
		return
	}

	resp, err := client.Put500Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Put500", resp, "Failure sending request")
		return
	}

	result, err = client.Put500Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Put500", resp, "Failure responding to request")
	}

	return
}

// Put500Preparer prepares the Put500 request.
func (client HTTPRetryClient) Put500Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/500"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Put500Sender sends the Put500 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Put500Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Put500Responder handles the response to the Put500 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Put500Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Put504 return 504 status code, then 200 after retry
//
// booleanValue is simple boolean value true
func (client HTTPRetryClient) Put504(booleanValue *bool) (result autorest.Response, err error) {
	req, err := client.Put504Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Put504", nil, "Failure preparing request")
		return
	}

	resp, err := client.Put504Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Put504", resp, "Failure sending request")
		return
	}

	result, err = client.Put504Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPRetryClient", "Put504", resp, "Failure responding to request")
	}

	return
}

// Put504Preparer prepares the Put504 request.
func (client HTTPRetryClient) Put504Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/retry/504"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Put504Sender sends the Put504 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPRetryClient) Put504Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Put504Responder handles the response to the Put504 request. The method always
// closes the http.Response Body.
func (client HTTPRetryClient) Put504Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
