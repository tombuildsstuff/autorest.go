package complexgroupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"tests/generated/complexgroup"
)

// BasicClientAPI contains the set of methods on the BasicClient type.
type BasicClientAPI interface {
	GetEmpty(ctx context.Context) (result complexgroup.Basic, err error)
	GetInvalid(ctx context.Context) (result complexgroup.Basic, err error)
	GetNotProvided(ctx context.Context) (result complexgroup.Basic, err error)
	GetNull(ctx context.Context) (result complexgroup.Basic, err error)
	GetValid(ctx context.Context) (result complexgroup.Basic, err error)
	PutValid(ctx context.Context, complexBody complexgroup.Basic) (result autorest.Response, err error)
}

var _ BasicClientAPI = (*complexgroup.BasicClient)(nil)

// PrimitiveClientAPI contains the set of methods on the PrimitiveClient type.
type PrimitiveClientAPI interface {
	GetBool(ctx context.Context) (result complexgroup.BooleanWrapper, err error)
	GetByte(ctx context.Context) (result complexgroup.ByteWrapper, err error)
	GetDate(ctx context.Context) (result complexgroup.DateWrapper, err error)
	GetDateTime(ctx context.Context) (result complexgroup.DatetimeWrapper, err error)
	GetDateTimeRfc1123(ctx context.Context) (result complexgroup.Datetimerfc1123Wrapper, err error)
	GetDouble(ctx context.Context) (result complexgroup.DoubleWrapper, err error)
	GetDuration(ctx context.Context) (result complexgroup.DurationWrapper, err error)
	GetFloat(ctx context.Context) (result complexgroup.FloatWrapper, err error)
	GetInt(ctx context.Context) (result complexgroup.IntWrapper, err error)
	GetLong(ctx context.Context) (result complexgroup.LongWrapper, err error)
	GetString(ctx context.Context) (result complexgroup.StringWrapper, err error)
	PutBool(ctx context.Context, complexBody complexgroup.BooleanWrapper) (result autorest.Response, err error)
	PutByte(ctx context.Context, complexBody complexgroup.ByteWrapper) (result autorest.Response, err error)
	PutDate(ctx context.Context, complexBody complexgroup.DateWrapper) (result autorest.Response, err error)
	PutDateTime(ctx context.Context, complexBody complexgroup.DatetimeWrapper) (result autorest.Response, err error)
	PutDateTimeRfc1123(ctx context.Context, complexBody complexgroup.Datetimerfc1123Wrapper) (result autorest.Response, err error)
	PutDouble(ctx context.Context, complexBody complexgroup.DoubleWrapper) (result autorest.Response, err error)
	PutDuration(ctx context.Context, complexBody complexgroup.DurationWrapper) (result autorest.Response, err error)
	PutFloat(ctx context.Context, complexBody complexgroup.FloatWrapper) (result autorest.Response, err error)
	PutInt(ctx context.Context, complexBody complexgroup.IntWrapper) (result autorest.Response, err error)
	PutLong(ctx context.Context, complexBody complexgroup.LongWrapper) (result autorest.Response, err error)
	PutString(ctx context.Context, complexBody complexgroup.StringWrapper) (result autorest.Response, err error)
}

var _ PrimitiveClientAPI = (*complexgroup.PrimitiveClient)(nil)

// ArrayClientAPI contains the set of methods on the ArrayClient type.
type ArrayClientAPI interface {
	GetEmpty(ctx context.Context) (result complexgroup.ArrayWrapper, err error)
	GetNotProvided(ctx context.Context) (result complexgroup.ArrayWrapper, err error)
	GetValid(ctx context.Context) (result complexgroup.ArrayWrapper, err error)
	PutEmpty(ctx context.Context, complexBody complexgroup.ArrayWrapper) (result autorest.Response, err error)
	PutValid(ctx context.Context, complexBody complexgroup.ArrayWrapper) (result autorest.Response, err error)
}

var _ ArrayClientAPI = (*complexgroup.ArrayClient)(nil)

// DictionaryClientAPI contains the set of methods on the DictionaryClient type.
type DictionaryClientAPI interface {
	GetEmpty(ctx context.Context) (result complexgroup.DictionaryWrapper, err error)
	GetNotProvided(ctx context.Context) (result complexgroup.DictionaryWrapper, err error)
	GetNull(ctx context.Context) (result complexgroup.DictionaryWrapper, err error)
	GetValid(ctx context.Context) (result complexgroup.DictionaryWrapper, err error)
	PutEmpty(ctx context.Context, complexBody complexgroup.DictionaryWrapper) (result autorest.Response, err error)
	PutValid(ctx context.Context, complexBody complexgroup.DictionaryWrapper) (result autorest.Response, err error)
}

var _ DictionaryClientAPI = (*complexgroup.DictionaryClient)(nil)

// InheritanceClientAPI contains the set of methods on the InheritanceClient type.
type InheritanceClientAPI interface {
	GetValid(ctx context.Context) (result complexgroup.Siamese, err error)
	PutValid(ctx context.Context, complexBody complexgroup.Siamese) (result autorest.Response, err error)
}

var _ InheritanceClientAPI = (*complexgroup.InheritanceClient)(nil)

// PolymorphismClientAPI contains the set of methods on the PolymorphismClient type.
type PolymorphismClientAPI interface {
	GetComplicated(ctx context.Context) (result complexgroup.SalmonModel, err error)
	GetValid(ctx context.Context) (result complexgroup.FishModel, err error)
	PutComplicated(ctx context.Context, complexBody complexgroup.BasicSalmon) (result autorest.Response, err error)
	PutMissingDiscriminator(ctx context.Context, complexBody complexgroup.BasicSalmon) (result complexgroup.SalmonModel, err error)
	PutValid(ctx context.Context, complexBody complexgroup.BasicFish) (result autorest.Response, err error)
	PutValidMissingRequired(ctx context.Context, complexBody complexgroup.BasicFish) (result autorest.Response, err error)
}

var _ PolymorphismClientAPI = (*complexgroup.PolymorphismClient)(nil)

// PolymorphicrecursiveClientAPI contains the set of methods on the PolymorphicrecursiveClient type.
type PolymorphicrecursiveClientAPI interface {
	GetValid(ctx context.Context) (result complexgroup.FishModel, err error)
	PutValid(ctx context.Context, complexBody complexgroup.BasicFish) (result autorest.Response, err error)
}

var _ PolymorphicrecursiveClientAPI = (*complexgroup.PolymorphicrecursiveClient)(nil)

// ReadonlypropertyClientAPI contains the set of methods on the ReadonlypropertyClient type.
type ReadonlypropertyClientAPI interface {
	GetValid(ctx context.Context) (result complexgroup.ReadonlyObj, err error)
	PutValid(ctx context.Context, complexBody complexgroup.ReadonlyObj) (result autorest.Response, err error)
}

var _ ReadonlypropertyClientAPI = (*complexgroup.ReadonlypropertyClient)(nil)

// FlattencomplexClientAPI contains the set of methods on the FlattencomplexClient type.
type FlattencomplexClientAPI interface {
	GetValid(ctx context.Context) (result complexgroup.MyBaseTypeModel, err error)
}

var _ FlattencomplexClientAPI = (*complexgroup.FlattencomplexClient)(nil)
