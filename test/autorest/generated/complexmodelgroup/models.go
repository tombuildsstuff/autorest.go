// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexmodelgroup

import (
	"fmt"
	"net/http"
)

type CatalogArray struct {
	// Array of products
	ProductArray *[]Product `json:"productArray,omitempty"`
}

type CatalogArrayOfDictionary struct {
	// Array of dictionary of products
	ProductArrayOfDictionary *[]map[string]Product `json:"productArrayOfDictionary,omitempty"`
}

// CatalogArrayResponse is the response envelope for operations that return a CatalogArray type.
type CatalogArrayResponse struct {
	CatalogArray *CatalogArray

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type CatalogDictionary struct {
	// Dictionary of products
	ProductDictionary *map[string]Product `json:"productDictionary,omitempty"`
}

type CatalogDictionaryOfArray struct {
	// Dictionary of Array of product
	ProductDictionaryOfArray *map[string][]Product `json:"productDictionaryOfArray,omitempty"`
}

// CatalogDictionaryResponse is the response envelope for operations that return a CatalogDictionary type.
type CatalogDictionaryResponse struct {
	CatalogDictionary *CatalogDictionary

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

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

// The product documentation.
type Product struct {
	// Capacity of product. For example, 4 people.
	Capacity *string `json:"capacity,omitempty"`

	// Description of product.
	Description *string `json:"description,omitempty"`

	// Display name of product.
	DisplayName *string `json:"display_name,omitempty"`

	// Image URL representing the product.
	Image *string `json:"image,omitempty"`

	// Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco
	// will have a different product_id than uberX in Los Angeles.
	ProductID *string `json:"product_id,omitempty"`
}
