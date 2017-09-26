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

// HTTPServerFailureClient is the test Infrastructure for AutoRest
type HTTPServerFailureClient struct {
	ManagementClient
}

// NewHTTPServerFailureClient creates an instance of the HTTPServerFailureClient client.
func NewHTTPServerFailureClient() HTTPServerFailureClient {
	return NewHTTPServerFailureClientWithBaseURI(DefaultBaseURI)
}

// NewHTTPServerFailureClientWithBaseURI creates an instance of the HTTPServerFailureClient client.
func NewHTTPServerFailureClientWithBaseURI(baseURI string) HTTPServerFailureClient {
	return HTTPServerFailureClient{NewWithBaseURI(baseURI)}
}

// Delete505 return 505 status code - should be represented in the client as an error
//
// booleanValue is simple boolean value true
func (client HTTPServerFailureClient) Delete505(booleanValue *bool) (result Error, err error) {
	req, err := client.Delete505Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Delete505", nil, "Failure preparing request")
		return
	}

	resp, err := client.Delete505Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Delete505", resp, "Failure sending request")
		return
	}

	result, err = client.Delete505Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Delete505", resp, "Failure responding to request")
	}

	return
}

// Delete505Preparer prepares the Delete505 request.
func (client HTTPServerFailureClient) Delete505Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/505"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Delete505Sender sends the Delete505 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Delete505Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Delete505Responder handles the response to the Delete505 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Delete505Responder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get501 return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Get501() (result Error, err error) {
	req, err := client.Get501Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Get501", nil, "Failure preparing request")
		return
	}

	resp, err := client.Get501Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Get501", resp, "Failure sending request")
		return
	}

	result, err = client.Get501Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Get501", resp, "Failure responding to request")
	}

	return
}

// Get501Preparer prepares the Get501 request.
func (client HTTPServerFailureClient) Get501Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/501"))
	return preparer.Prepare(&http.Request{})
}

// Get501Sender sends the Get501 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Get501Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Get501Responder handles the response to the Get501 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Get501Responder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Head501 return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Head501() (result Error, err error) {
	req, err := client.Head501Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Head501", nil, "Failure preparing request")
		return
	}

	resp, err := client.Head501Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Head501", resp, "Failure sending request")
		return
	}

	result, err = client.Head501Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Head501", resp, "Failure responding to request")
	}

	return
}

// Head501Preparer prepares the Head501 request.
func (client HTTPServerFailureClient) Head501Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/501"))
	return preparer.Prepare(&http.Request{})
}

// Head501Sender sends the Head501 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Head501Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Head501Responder handles the response to the Head501 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Head501Responder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Post505 return 505 status code - should be represented in the client as an error
//
// booleanValue is simple boolean value true
func (client HTTPServerFailureClient) Post505(booleanValue *bool) (result Error, err error) {
	req, err := client.Post505Preparer(booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Post505", nil, "Failure preparing request")
		return
	}

	resp, err := client.Post505Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Post505", resp, "Failure sending request")
		return
	}

	result, err = client.Post505Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Post505", resp, "Failure responding to request")
	}

	return
}

// Post505Preparer prepares the Post505 request.
func (client HTTPServerFailureClient) Post505Preparer(booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/505"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare(&http.Request{})
}

// Post505Sender sends the Post505 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Post505Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// Post505Responder handles the response to the Post505 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Post505Responder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
