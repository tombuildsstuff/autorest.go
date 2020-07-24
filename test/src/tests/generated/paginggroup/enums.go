package paginggroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Status enumerates the values for status.
type Status string

const (
	// Accepted ...
	Accepted Status = "Accepted"
	// Canceled ...
	Canceled Status = "canceled"
	// Created ...
	Created Status = "Created"
	// Creating ...
	Creating Status = "Creating"
	// Deleted ...
	Deleted Status = "Deleted"
	// Deleting ...
	Deleting Status = "Deleting"
	// Failed ...
	Failed Status = "Failed"
	// OK ...
	OK Status = "OK"
	// Succeeded ...
	Succeeded Status = "Succeeded"
	// Updated ...
	Updated Status = "Updated"
	// Updating ...
	Updating Status = "Updating"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, OK, Succeeded, Updated, Updating}
}
