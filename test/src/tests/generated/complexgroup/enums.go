package complexgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CMYKColors enumerates the values for cmyk colors.
type CMYKColors string

const (
	// BlacK ...
	BlacK CMYKColors = "blacK"
	// Cyan ...
	Cyan CMYKColors = "cyan"
	// Magenta ...
	Magenta CMYKColors = "Magenta"
	// YELLOW ...
	YELLOW CMYKColors = "YELLOW"
)

// PossibleCMYKColorsValues returns an array of possible values for the CMYKColors const type.
func PossibleCMYKColorsValues() []CMYKColors {
	return []CMYKColors{BlacK, Cyan, Magenta, YELLOW}
}

// FishType enumerates the values for fish type.
type FishType string

const (
	// FishTypeDotFish ...
	FishTypeDotFish FishType = "DotFish"
	// FishTypeDotSalmon ...
	FishTypeDotSalmon FishType = "DotSalmon"
)

// PossibleFishTypeValues returns an array of possible values for the FishType const type.
func PossibleFishTypeValues() []FishType {
	return []FishType{FishTypeDotFish, FishTypeDotSalmon}
}

// FishtypeBasicFish enumerates the values for fishtype basic fish.
type FishtypeBasicFish string

const (
	// FishtypeCookiecuttershark ...
	FishtypeCookiecuttershark FishtypeBasicFish = "cookiecuttershark"
	// FishtypeFish ...
	FishtypeFish FishtypeBasicFish = "Fish"
	// FishtypeGoblin ...
	FishtypeGoblin FishtypeBasicFish = "goblin"
	// FishtypeSalmon ...
	FishtypeSalmon FishtypeBasicFish = "salmon"
	// FishtypeSawshark ...
	FishtypeSawshark FishtypeBasicFish = "sawshark"
	// FishtypeShark ...
	FishtypeShark FishtypeBasicFish = "shark"
	// FishtypeSmartSalmon ...
	FishtypeSmartSalmon FishtypeBasicFish = "smart_salmon"
)

// PossibleFishtypeBasicFishValues returns an array of possible values for the FishtypeBasicFish const type.
func PossibleFishtypeBasicFishValues() []FishtypeBasicFish {
	return []FishtypeBasicFish{FishtypeCookiecuttershark, FishtypeFish, FishtypeGoblin, FishtypeSalmon, FishtypeSawshark, FishtypeShark, FishtypeSmartSalmon}
}

// GoblinSharkColor enumerates the values for goblin shark color.
type GoblinSharkColor string

const (
	// Brown ...
	Brown GoblinSharkColor = "brown"
	// Gray ...
	Gray GoblinSharkColor = "gray"
	// Pink ...
	Pink GoblinSharkColor = "pink"
)

// PossibleGoblinSharkColorValues returns an array of possible values for the GoblinSharkColor const type.
func PossibleGoblinSharkColorValues() []GoblinSharkColor {
	return []GoblinSharkColor{Brown, Gray, Pink}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindKind1 ...
	KindKind1 Kind = "Kind1"
	// KindMyBaseType ...
	KindMyBaseType Kind = "MyBaseType"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindKind1, KindMyBaseType}
}

// MyKind enumerates the values for my kind.
type MyKind string

const (
	// Kind1 ...
	Kind1 MyKind = "Kind1"
)

// PossibleMyKindValues returns an array of possible values for the MyKind const type.
func PossibleMyKindValues() []MyKind {
	return []MyKind{Kind1}
}
