// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import "fmt"

type ArrayOptionalWrapper struct {
	Value *[]string `json:"value,omitempty"`
}

type ArrayWrapper struct {
	Value *[]string `json:"value,omitempty"`
}

type ClassOptionalWrapper struct {
	Value *Product `json:"value,omitempty"`
}

type ClassWrapper struct {
	Value *Product `json:"value,omitempty"`
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

// ExplicitPostOptionalArrayHeaderOptions contains the optional parameters for the Explicit.PostOptionalArrayHeader method.
type ExplicitPostOptionalArrayHeaderOptions struct {
	HeaderParameter *[]string
}

// ExplicitPostOptionalArrayParameterOptions contains the optional parameters for the Explicit.PostOptionalArrayParameter
// method.
type ExplicitPostOptionalArrayParameterOptions struct {
	BodyParameter *[]string
}

// ExplicitPostOptionalArrayPropertyOptions contains the optional parameters for the Explicit.PostOptionalArrayProperty method.
type ExplicitPostOptionalArrayPropertyOptions struct {
	BodyParameter *ArrayOptionalWrapper
}

// ExplicitPostOptionalClassParameterOptions contains the optional parameters for the Explicit.PostOptionalClassParameter
// method.
type ExplicitPostOptionalClassParameterOptions struct {
	BodyParameter *Product
}

// ExplicitPostOptionalClassPropertyOptions contains the optional parameters for the Explicit.PostOptionalClassProperty method.
type ExplicitPostOptionalClassPropertyOptions struct {
	BodyParameter *ClassOptionalWrapper
}

// ExplicitPostOptionalIntegerHeaderOptions contains the optional parameters for the Explicit.PostOptionalIntegerHeader method.
type ExplicitPostOptionalIntegerHeaderOptions struct {
	HeaderParameter *int32
}

// ExplicitPostOptionalIntegerParameterOptions contains the optional parameters for the Explicit.PostOptionalIntegerParameter
// method.
type ExplicitPostOptionalIntegerParameterOptions struct {
	BodyParameter *int32
}

// ExplicitPostOptionalIntegerPropertyOptions contains the optional parameters for the Explicit.PostOptionalIntegerProperty
// method.
type ExplicitPostOptionalIntegerPropertyOptions struct {
	BodyParameter *IntOptionalWrapper
}

// ExplicitPostOptionalStringHeaderOptions contains the optional parameters for the Explicit.PostOptionalStringHeader method.
type ExplicitPostOptionalStringHeaderOptions struct {
	BodyParameter *string
}

// ExplicitPostOptionalStringParameterOptions contains the optional parameters for the Explicit.PostOptionalStringParameter
// method.
type ExplicitPostOptionalStringParameterOptions struct {
	BodyParameter *string
}

// ExplicitPostOptionalStringPropertyOptions contains the optional parameters for the Explicit.PostOptionalStringProperty
// method.
type ExplicitPostOptionalStringPropertyOptions struct {
	BodyParameter *StringOptionalWrapper
}

// ImplicitGetOptionalGlobalQueryOptions contains the optional parameters for the Implicit.GetOptionalGlobalQuery method.
type ImplicitGetOptionalGlobalQueryOptions struct {
}

// ImplicitPutOptionalBodyOptions contains the optional parameters for the Implicit.PutOptionalBody method.
type ImplicitPutOptionalBodyOptions struct {
	BodyParameter *string
}

// ImplicitPutOptionalHeaderOptions contains the optional parameters for the Implicit.PutOptionalHeader method.
type ImplicitPutOptionalHeaderOptions struct {
	QueryParameter *string
}

// ImplicitPutOptionalQueryOptions contains the optional parameters for the Implicit.PutOptionalQuery method.
type ImplicitPutOptionalQueryOptions struct {
	QueryParameter *string
}

type IntOptionalWrapper struct {
	Value *int32 `json:"value,omitempty"`
}

type IntWrapper struct {
	Value *int32 `json:"value,omitempty"`
}

type Product struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type StringOptionalWrapper struct {
	Value *string `json:"value,omitempty"`
}

type StringWrapper struct {
	Value *string `json:"value,omitempty"`
}
