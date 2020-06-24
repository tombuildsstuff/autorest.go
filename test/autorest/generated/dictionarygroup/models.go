// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dictionarygroup

import (
	"fmt"
	"net/http"
	"time"
)

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// MapOfBoolResponse is the response envelope for operations that return a map[string]bool type.
type MapOfBoolResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": true, "1": false, "2": false, "3": true }
	Value *map[string]bool
}

// MapOfByteArrayResponse is the response envelope for operations that return a map[string][]byte type.
type MapOfByteArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": hex(FF FF FF FA), "1": hex(01 02 03), "2": hex (25, 29, 43)} with each elementencoded in base
	// 64
	Value *map[string][]byte
}

// MapOfDurationResponse is the response envelope for operations that return a map[string]time.Duration type.
type MapOfDurationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": "P123DT22H14M12.011S", "1": "P5DT1H0M0S"}
	Value *map[string]time.Duration
}

// MapOfFloat32Response is the response envelope for operations that return a map[string]float32 type.
type MapOfFloat32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
	Value *map[string]float32
}

// MapOfFloat64Response is the response envelope for operations that return a map[string]float64 type.
type MapOfFloat64Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
	Value *map[string]float64
}

// MapOfInt32Response is the response envelope for operations that return a map[string]int32 type.
type MapOfInt32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The null dictionary value
	Value *map[string]int32
}

// MapOfInt64Response is the response envelope for operations that return a map[string]int64 type.
type MapOfInt64Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
	Value *map[string]int64
}

// MapOfInterfaceResponse is the response envelope for operations that return a map[string]interface{} type.
type MapOfInterfaceResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// An dictionaries of dictionaries with value null
	Value *map[string]interface{}
}

// MapOfStringArrayResponse is the response envelope for operations that return a map[string][]string type.
type MapOfStringArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// a null array
	Value *map[string][]string
}

// MapOfStringResponse is the response envelope for operations that return a map[string]string type.
type MapOfStringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Dictionary of <string>
	Value *map[string]string
}

// MapOfTimeResponse is the response envelope for operations that return a map[string]time.Time type.
type MapOfTimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The dictionary value {"0": "2000-12-01", "1": "1980-01-02", "2": "1492-10-12"}
	Value *map[string]time.Time
}

// MapOfWidgetResponse is the response envelope for operations that return a map[string]Widget type.
type MapOfWidgetResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Dictionary of complex type with null value
	Value *map[string]Widget
}

type Widget struct {
	Integer *int32  `json:"integer,omitempty"`
	String  *string `json:"string,omitempty"`
}
