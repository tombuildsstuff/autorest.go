package optionalgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "tests/generated/optionalgroup"

// ArrayOptionalWrapper ...
type ArrayOptionalWrapper struct {
	Value *[]string `json:"value,omitempty"`
}

// ArrayWrapper ...
type ArrayWrapper struct {
	Value *[]string `json:"value,omitempty"`
}

// ClassOptionalWrapper ...
type ClassOptionalWrapper struct {
	Value *Product `json:"value,omitempty"`
}

// ClassWrapper ...
type ClassWrapper struct {
	Value *Product `json:"value,omitempty"`
}

// Error ...
type Error struct {
	autorest.Response `json:"-"`
	Status            *int32  `json:"status,omitempty"`
	Message           *string `json:"message,omitempty"`
}

// IntOptionalWrapper ...
type IntOptionalWrapper struct {
	Value *int32 `json:"value,omitempty"`
}

// IntWrapper ...
type IntWrapper struct {
	Value *int32 `json:"value,omitempty"`
}

// Product ...
type Product struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// StringOptionalWrapper ...
type StringOptionalWrapper struct {
	Value *string `json:"value,omitempty"`
}

// StringWrapper ...
type StringWrapper struct {
	Value *string `json:"value,omitempty"`
}
