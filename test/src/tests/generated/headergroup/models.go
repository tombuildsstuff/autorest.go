package headergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// The package's fully qualified name.
const fqdn = "tests/generated/headergroup"

// GreyscaleColors enumerates the values for greyscale colors.
type GreyscaleColors string

const (
	// Black ...
	Black GreyscaleColors = "black"
	// GREY ...
	GREY GreyscaleColors = "GREY"
	// White ...
	White GreyscaleColors = "White"
)

// PossibleGreyscaleColorsValues returns an array of possible values for the GreyscaleColors const type.
func PossibleGreyscaleColorsValues() []GreyscaleColors {
	return []GreyscaleColors{Black, GREY, White}
}

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}