// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"time"
)

type CMYKColors string

const (
	CMYKColorsBlack CMYKColors = "blacK"
	CMYKColorsCyan CMYKColors = "cyan"
	CMYKColorsMagenta CMYKColors = "Magenta"
	CMYKColorsYellow CMYKColors = "YELLOW"
)

func PossibleCMYKColorsValues() []CMYKColors {
	return []CMYKColors{CMYKColorsBlack, CMYKColorsCyan, CMYKColorsMagenta, CMYKColorsYellow}
}

// GoblinSharkColor - Colors possible
type GoblinSharkColor string

const (
	GoblinSharkColorBrown GoblinSharkColor = "brown"
	GoblinSharkColorGray GoblinSharkColor = "gray"
	// GoblinSharkColorLowerred - Lowercase RED
	GoblinSharkColorLowerred GoblinSharkColor = "red"
	GoblinSharkColorPink GoblinSharkColor = "pink"
	// GoblinSharkColorUpperred - Uppercase RED
	GoblinSharkColorUpperred GoblinSharkColor = "RED"
)

func PossibleGoblinSharkColorValues() []GoblinSharkColor {
	return []GoblinSharkColor{GoblinSharkColorBrown, GoblinSharkColorGray, GoblinSharkColorLowerred, GoblinSharkColorPink, GoblinSharkColorUpperred}
}

// ArrayGetEmptyResponse contains the response from method Array.GetEmpty.
type ArrayGetEmptyResponse struct {
	ArrayWrapper *ArrayWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// ArrayGetNotProvidedResponse contains the response from method Array.GetNotProvided.
type ArrayGetNotProvidedResponse struct {
	ArrayWrapper *ArrayWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// ArrayGetValidResponse contains the response from method Array.GetValid.
type ArrayGetValidResponse struct {
	ArrayWrapper *ArrayWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// ArrayPutEmptyResponse contains the response from method Array.PutEmpty.
type ArrayPutEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// ArrayPutValidResponse contains the response from method Array.PutValid.
type ArrayPutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type ArrayWrapper struct {
	Array []string `json:"array,omitempty"`
}

type Basic struct {
	Color *CMYKColors `json:"color,omitempty"`
	// Basic Id
	ID *int32 `json:"id,omitempty"`
	// Name property with a very long description that does not fit on a single line and a line break.
	Name *string `json:"name,omitempty"`
}

// BasicGetEmptyResponse contains the response from method Basic.GetEmpty.
type BasicGetEmptyResponse struct {
	Basic *Basic
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// BasicGetInvalidResponse contains the response from method Basic.GetInvalid.
type BasicGetInvalidResponse struct {
	Basic *Basic
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// BasicGetNotProvidedResponse contains the response from method Basic.GetNotProvided.
type BasicGetNotProvidedResponse struct {
	Basic *Basic
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// BasicGetNullResponse contains the response from method Basic.GetNull.
type BasicGetNullResponse struct {
	Basic *Basic
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// BasicGetValidResponse contains the response from method Basic.GetValid.
type BasicGetValidResponse struct {
	Basic *Basic
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// BasicPutValidResponse contains the response from method Basic.PutValid.
type BasicPutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type BooleanWrapper struct {
	FieldFalse *bool `json:"field_false,omitempty"`
	FieldTrue *bool `json:"field_true,omitempty"`
}

type ByteWrapper struct {
	Field []byte `json:"field,omitempty"`
}

type Cat struct {
	Color *string `json:"color,omitempty"`
	Hates []Dog `json:"hates,omitempty"`
}

type Cookiecuttershark struct {
}

type DateWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Leap *time.Time `json:"leap,omitempty"`
}

type DatetimeWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now *time.Time `json:"now,omitempty"`
}

type Datetimerfc1123Wrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now *time.Time `json:"now,omitempty"`
}

// DictionaryGetEmptyResponse contains the response from method Dictionary.GetEmpty.
type DictionaryGetEmptyResponse struct {
	DictionaryWrapper *DictionaryWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// DictionaryGetNotProvidedResponse contains the response from method Dictionary.GetNotProvided.
type DictionaryGetNotProvidedResponse struct {
	DictionaryWrapper *DictionaryWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// DictionaryGetNullResponse contains the response from method Dictionary.GetNull.
type DictionaryGetNullResponse struct {
	DictionaryWrapper *DictionaryWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// DictionaryGetValidResponse contains the response from method Dictionary.GetValid.
type DictionaryGetValidResponse struct {
	DictionaryWrapper *DictionaryWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// DictionaryPutEmptyResponse contains the response from method Dictionary.PutEmpty.
type DictionaryPutEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// DictionaryPutValidResponse contains the response from method Dictionary.PutValid.
type DictionaryPutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type DictionaryWrapper struct {
	// Dictionary of <components·schemas·dictionary_wrapper·properties·defaultprogram·additionalproperties>
	DefaultProgram map[string]*string `json:"defaultProgram,omitempty"`
}

type Dog struct {
	Food *string `json:"food,omitempty"`
}

type DotFish struct {
	FishType *string `json:"fish.type,omitempty"`
	Species *string `json:"species,omitempty"`
}

type DotFishMarket struct {
	Fishes []DotFish `json:"fishes,omitempty"`
	Salmons []DotSalmon `json:"salmons,omitempty"`
	SampleFish *DotFish `json:"sampleFish,omitempty"`
	SampleSalmon *DotSalmon `json:"sampleSalmon,omitempty"`
}

type DotSalmon struct {
	Iswild *bool `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

type DoubleWrapper struct {
	Field1 *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

type DurationWrapper struct {
	Field *time.Duration `json:"field,omitempty"`
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

func newError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

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

type Fish struct {
	Fishtype *string `json:"fishtype,omitempty"`
	Length *float32 `json:"length,omitempty"`
	Siblings []Fish `json:"siblings,omitempty"`
	Species *string `json:"species,omitempty"`
}

// FlattencomplexGetValidResponse contains the response from method Flattencomplex.GetValid.
type FlattencomplexGetValidResponse struct {
	MyBaseType *MyBaseType
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type FloatWrapper struct {
	Field1 *float32 `json:"field1,omitempty"`
	Field2 *float32 `json:"field2,omitempty"`
}

type Goblinshark struct {
	// Colors possible
	Color *GoblinSharkColor `json:"color,omitempty"`
	Jawsize *int32 `json:"jawsize,omitempty"`
}

// InheritanceGetValidResponse contains the response from method Inheritance.GetValid.
type InheritanceGetValidResponse struct {
	Siamese *Siamese
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// InheritancePutValidResponse contains the response from method Inheritance.PutValid.
type InheritancePutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type IntWrapper struct {
	Field1 *int32 `json:"field1,omitempty"`
	Field2 *int32 `json:"field2,omitempty"`
}

type LongWrapper struct {
	Field1 *int64 `json:"field1,omitempty"`
	Field2 *int64 `json:"field2,omitempty"`
}

type MyBaseHelperType struct {
	PropBh1 *string `json:"propBH1,omitempty"`
}

type MyBaseType struct {
	Helper *MyBaseHelperType `json:"helper,omitempty"`
	Kind *string `json:"kind,omitempty"`
	PropB1 *string `json:"propB1,omitempty"`
}

type MyDerivedType struct {
	PropD1 *string `json:"propD1,omitempty"`
}

type Pet struct {
	ID *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PolymorphicrecursiveGetValidResponse contains the response from method Polymorphicrecursive.GetValid.
type PolymorphicrecursiveGetValidResponse struct {
	Fish *Fish
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphicrecursivePutValidResponse contains the response from method Polymorphicrecursive.PutValid.
type PolymorphicrecursivePutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismGetComplicatedResponse contains the response from method Polymorphism.GetComplicated.
type PolymorphismGetComplicatedResponse struct {
	Salmon *Salmon
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismGetComposedWithDiscriminatorResponse contains the response from method Polymorphism.GetComposedWithDiscriminator.
type PolymorphismGetComposedWithDiscriminatorResponse struct {
	DotFishMarket *DotFishMarket
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismGetComposedWithoutDiscriminatorResponse contains the response from method Polymorphism.GetComposedWithoutDiscriminator.
type PolymorphismGetComposedWithoutDiscriminatorResponse struct {
	DotFishMarket *DotFishMarket
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismGetDotSyntaxResponse contains the response from method Polymorphism.GetDotSyntax.
type PolymorphismGetDotSyntaxResponse struct {
	DotFish *DotFish
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismGetValidResponse contains the response from method Polymorphism.GetValid.
type PolymorphismGetValidResponse struct {
	Fish *Fish
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismPutComplicatedResponse contains the response from method Polymorphism.PutComplicated.
type PolymorphismPutComplicatedResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismPutMissingDiscriminatorResponse contains the response from method Polymorphism.PutMissingDiscriminator.
type PolymorphismPutMissingDiscriminatorResponse struct {
	Salmon *Salmon
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismPutValidMissingRequiredResponse contains the response from method Polymorphism.PutValidMissingRequired.
type PolymorphismPutValidMissingRequiredResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PolymorphismPutValidResponse contains the response from method Polymorphism.PutValid.
type PolymorphismPutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetBoolResponse contains the response from method Primitive.GetBool.
type PrimitiveGetBoolResponse struct {
	BooleanWrapper *BooleanWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetByteResponse contains the response from method Primitive.GetByte.
type PrimitiveGetByteResponse struct {
	ByteWrapper *ByteWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetDateResponse contains the response from method Primitive.GetDate.
type PrimitiveGetDateResponse struct {
	DateWrapper *DateWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetDateTimeRFC1123Response contains the response from method Primitive.GetDateTimeRFC1123.
type PrimitiveGetDateTimeRFC1123Response struct {
	Datetimerfc1123Wrapper *Datetimerfc1123Wrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetDateTimeResponse contains the response from method Primitive.GetDateTime.
type PrimitiveGetDateTimeResponse struct {
	DatetimeWrapper *DatetimeWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetDoubleResponse contains the response from method Primitive.GetDouble.
type PrimitiveGetDoubleResponse struct {
	DoubleWrapper *DoubleWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetDurationResponse contains the response from method Primitive.GetDuration.
type PrimitiveGetDurationResponse struct {
	DurationWrapper *DurationWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetFloatResponse contains the response from method Primitive.GetFloat.
type PrimitiveGetFloatResponse struct {
	FloatWrapper *FloatWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetIntResponse contains the response from method Primitive.GetInt.
type PrimitiveGetIntResponse struct {
	IntWrapper *IntWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetLongResponse contains the response from method Primitive.GetLong.
type PrimitiveGetLongResponse struct {
	LongWrapper *LongWrapper
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitiveGetStringResponse contains the response from method Primitive.GetString.
type PrimitiveGetStringResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
	StringWrapper *StringWrapper
}

// PrimitivePutBoolResponse contains the response from method Primitive.PutBool.
type PrimitivePutBoolResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutByteResponse contains the response from method Primitive.PutByte.
type PrimitivePutByteResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutDateResponse contains the response from method Primitive.PutDate.
type PrimitivePutDateResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutDateTimeRFC1123Response contains the response from method Primitive.PutDateTimeRFC1123.
type PrimitivePutDateTimeRFC1123Response struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutDateTimeResponse contains the response from method Primitive.PutDateTime.
type PrimitivePutDateTimeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutDoubleResponse contains the response from method Primitive.PutDouble.
type PrimitivePutDoubleResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutDurationResponse contains the response from method Primitive.PutDuration.
type PrimitivePutDurationResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutFloatResponse contains the response from method Primitive.PutFloat.
type PrimitivePutFloatResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutIntResponse contains the response from method Primitive.PutInt.
type PrimitivePutIntResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutLongResponse contains the response from method Primitive.PutLong.
type PrimitivePutLongResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PrimitivePutStringResponse contains the response from method Primitive.PutString.
type PrimitivePutStringResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type ReadonlyObj struct {
	ID *string `json:"id,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// ReadonlypropertyGetValidResponse contains the response from method Readonlyproperty.GetValid.
type ReadonlypropertyGetValidResponse struct {
	ReadonlyObj *ReadonlyObj
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// ReadonlypropertyPutValidResponse contains the response from method Readonlyproperty.PutValid.
type ReadonlypropertyPutValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type Salmon struct {
	Iswild *bool `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

type Sawshark struct {
	Picture []byte `json:"picture,omitempty"`
}

type Shark struct {
	Age *int32 `json:"age,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}

type Siamese struct {
	Breed *string `json:"breed,omitempty"`
}

type SmartSalmon struct {
	CollegeDegree *string `json:"college_degree,omitempty"`
}

type StringWrapper struct {
	Empty *string `json:"empty,omitempty"`
	Field *string `json:"field,omitempty"`
	Null *string `json:"null,omitempty"`
}

