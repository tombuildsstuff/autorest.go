// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlrouptest

import (
	"context"
	"generatortests/autorest/generated/xmlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func toTimePtr(layout string, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

func TestGetACLs(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetACLs(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]xmlgroup.SignedIDentifier{
		{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &xmlgroup.AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.SignedIdentifiers, expected)
}

func TestGetComplexTypeRefNoMeta(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetComplexTypeRefNoMeta(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.RootWithRefAndNoMeta{
		RefToModel: &xmlgroup.ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	helpers.DeepEqualOrFatal(t, result.RootWithRefAndNoMeta, expected)
}

func TestGetComplexTypeRefWithMeta(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetComplexTypeRefWithMeta(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected := &xmlgroup.RootWithRefAndMeta{
		RefToModel: &xmlgroup.ComplexTypeWithMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	helpers.DeepEqualOrFatal(t, result.RootWithRefAndMeta, expected)
}

func TestGetEmptyChildElement(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetEmptyChildElement(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	}
	helpers.DeepEqualOrFatal(t, result.Banana, expected)
}

func TestGetEmptyList(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetEmptyList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.Slideshow{}
	helpers.DeepEqualOrFatal(t, result.Slideshow, expected)
}

func TestGetEmptyRootList(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetEmptyRootList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	if result.Bananas != nil {
		t.Fatal("expected nil slice")
	}
}

func TestGetEmptyWrappedLists(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetEmptyWrappedLists(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.AppleBarrel{}
	helpers.DeepEqualOrFatal(t, result.AppleBarrel, expected)
}

func TestGetHeaders(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetHeaders(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.CustomHeader, to.StringPtr("custom-value"))
}

func TestGetRootList(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetRootList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]xmlgroup.Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		{
			Name:       to.StringPtr("Plantain"),
			Flavor:     to.StringPtr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	helpers.DeepEqualOrFatal(t, result.Bananas, expected)
}

func TestGetRootListSingleItem(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetRootListSingleItem(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]xmlgroup.Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	helpers.DeepEqualOrFatal(t, result.Bananas, expected)
}

func TestGetServiceProperties(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetServiceProperties(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.StorageServiceProperties{
		HourMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(false),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		Logging: &xmlgroup.Logging{
			Version: to.StringPtr("1.0"),
			Delete:  to.BoolPtr(true),
			Read:    to.BoolPtr(false),
			Write:   to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		MinuteMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.StorageServiceProperties, expected)
}

func TestGetSimple(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetSimple(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: &[]xmlgroup.Slide{
			{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			{
				Items: &[]string{"Why WonderWidgets are great", "", "Who buys WonderWidgets"},
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.Slideshow, expected)
}

func TestGetWrappedLists(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.GetWrappedLists(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.AppleBarrel{
		BadApples:  &[]string{"Red Delicious"},
		GoodApples: &[]string{"Fuji", "Gala"},
	}
	helpers.DeepEqualOrFatal(t, result.AppleBarrel, expected)
}

func TestJSONInput(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.JSONInput(context.Background(), xmlgroup.JSONInput{
		ID: to.Int32Ptr(42),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestJSONOutput(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.JSONOutput(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected := xmlgroup.JSONOutput{
		ID: to.Int32Ptr(42),
	}
	helpers.DeepEqualOrFatal(t, result.JSONOutput, &expected)
}

func TestListBlobs(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.ListBlobs(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	blob1LM, err := time.Parse(time.RFC1123, "Wed, 09 Sep 2009 09:20:02 GMT")
	if err != nil {
		t.Fatal(err)
	}
	blob2LM, err := time.Parse(time.RFC1123, "Wed, 09 Sep 2009 09:20:03 GMT")
	if err != nil {
		t.Fatal(err)
	}
	expected := xmlgroup.ListBlobsResponse{
		Blobs: &xmlgroup.Blobs{
			Blob: &[]xmlgroup.Blob{
				{
					Metadata: &map[string]string{
						"color":            "blue",
						"blobnumber":       "01",
						"somemetadataname": "SomeMetadataValue",
					},
					Name: to.StringPtr("blob1.txt"),
					Properties: &xmlgroup.BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8A29A19"),
						ContentLength:   to.Int64Ptr(100),
						ContentType:     to.StringPtr("text/html"),
						ContentEncoding: to.StringPtr(""),
						ContentLanguage: to.StringPtr("en-US"),
						ContentMd5:      to.StringPtr(""),
						CacheControl:    to.StringPtr("no-cache"),
						BlobType:        xmlgroup.BlobTypeBlockBlob.ToPtr(),
						LeaseStatus:     xmlgroup.LeaseStatusTypeUnlocked.ToPtr(),
					},
				},
				{
					Metadata: &map[string]string{
						"color":             "green",
						"blobnumber":        "02",
						"somemetadataname":  "SomeMetadataValue",
						"x-ms-invalid-name": "nasdf$@#$$",
					},
					Name: to.StringPtr("blob2.txt"),
					Properties: &xmlgroup.BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Int64Ptr(5000),
						ContentType:     to.StringPtr("application/octet-stream"),
						ContentEncoding: to.StringPtr("gzip"),
						ContentLanguage: to.StringPtr(""),
						ContentMd5:      to.StringPtr(""),
						CacheControl:    to.StringPtr(""),
						BlobType:        xmlgroup.BlobTypeBlockBlob.ToPtr(),
					},
					Snapshot: to.StringPtr("2009-09-09T09:20:03.0427659Z"),
				},
				{
					Metadata: &map[string]string{
						"color":            "green",
						"blobnumber":       "02",
						"somemetadataname": "SomeMetadataValue",
					},
					Name: to.StringPtr("blob2.txt"),
					Properties: &xmlgroup.BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Int64Ptr(5000),
						ContentType:     to.StringPtr("application/octet-stream"),
						ContentEncoding: to.StringPtr("gzip"),
						ContentLanguage: to.StringPtr(""),
						ContentMd5:      to.StringPtr(""),
						CacheControl:    to.StringPtr(""),
						BlobType:        xmlgroup.BlobTypeBlockBlob.ToPtr(),
					},
					Snapshot: to.StringPtr("2009-09-09T09:20:03.1587543Z"),
				},
				{
					Metadata: &map[string]string{
						"color":            "green",
						"blobnumber":       "02",
						"somemetadataname": "SomeMetadataValue",
					},
					Name: to.StringPtr("blob2.txt"),
					Properties: &xmlgroup.BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Int64Ptr(5000),
						ContentType:     to.StringPtr("application/octet-stream"),
						ContentEncoding: to.StringPtr("gzip"),
						ContentLanguage: to.StringPtr(""),
						ContentMd5:      to.StringPtr(""),
						CacheControl:    to.StringPtr(""),
						BlobType:        xmlgroup.BlobTypeBlockBlob.ToPtr(),
						LeaseStatus:     xmlgroup.LeaseStatusTypeUnlocked.ToPtr(),
					},
				},
				{
					Metadata: &map[string]string{
						"color":            "yellow",
						"blobnumber":       "03",
						"somemetadataname": "SomeMetadataValue",
					},
					Name: to.StringPtr("blob3.txt"),
					Properties: &xmlgroup.BlobProperties{
						LastModified:       &blob2LM,
						Etag:               to.StringPtr("0x8CBFF45D911FADF"),
						ContentLength:      to.Int64Ptr(16384),
						ContentType:        to.StringPtr("image/jpeg"),
						ContentEncoding:    to.StringPtr(""),
						ContentLanguage:    to.StringPtr(""),
						ContentMd5:         to.StringPtr(""),
						CacheControl:       to.StringPtr(""),
						BlobSequenceNumber: to.Int32Ptr(3),
						BlobType:           xmlgroup.BlobTypePageBlob.ToPtr(),
						LeaseStatus:        xmlgroup.LeaseStatusTypeLocked.ToPtr(),
					},
				},
			},
		},
		ContainerName: to.StringPtr("https://myaccount.blob.core.windows.net/mycontainer"),
		NextMarker:    to.StringPtr(""),
	}
	helpers.DeepEqualOrFatal(t, result.EnumerationResults, &expected)
}

func TestListContainers(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.ListContainers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.ListContainersResponse{
		ServiceEndpoint: to.StringPtr("https://myaccount.blob.core.windows.net/"),
		MaxResults:      to.Int32Ptr(3),
		NextMarker:      to.StringPtr("video"),
		Containers: &[]xmlgroup.Container{
			{
				Name: to.StringPtr("audio"),
				Properties: &xmlgroup.ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7C6B1B2"),
					PublicAccess: xmlgroup.PublicAccessTypeContainer.ToPtr(),
				},
			},
			{
				Name: to.StringPtr("images"),
				Properties: &xmlgroup.ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7C1EEEC"),
				},
			},
			{
				Name: to.StringPtr("textfiles"),
				Properties: &xmlgroup.ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7BACAC3"),
				},
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.EnumerationResults, expected)
}

func TestPutACLs(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutACLs(context.Background(), []xmlgroup.SignedIDentifier{
		{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &xmlgroup.AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutComplexTypeRefNoMeta(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutComplexTypeRefNoMeta(context.Background(), xmlgroup.RootWithRefAndNoMeta{
		RefToModel: &xmlgroup.ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutComplexTypeRefWithMeta(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutComplexTypeRefWithMeta(context.Background(), xmlgroup.RootWithRefAndMeta{
		RefToModel: &xmlgroup.ComplexTypeWithMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyChildElement(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutEmptyChildElement(context.Background(), xmlgroup.Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyList(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutEmptyList(context.Background(), xmlgroup.Slideshow{
		Slides: &[]xmlgroup.Slide{},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyRootList(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutEmptyRootList(context.Background(), []xmlgroup.Banana{})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyWrappedLists(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutEmptyWrappedLists(context.Background(), xmlgroup.AppleBarrel{
		BadApples:  &[]string{},
		GoodApples: &[]string{},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutRootList(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutRootList(context.Background(), []xmlgroup.Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		{
			Name:       to.StringPtr("Plantain"),
			Flavor:     to.StringPtr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutRootListSingleItem(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutRootListSingleItem(context.Background(), []xmlgroup.Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutServiceProperties(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutServiceProperties(context.Background(), xmlgroup.StorageServiceProperties{
		HourMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(false),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		Logging: &xmlgroup.Logging{
			Version: to.StringPtr("1.0"),
			Delete:  to.BoolPtr(true),
			Read:    to.BoolPtr(false),
			Write:   to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		MinuteMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutSimple(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutSimple(context.Background(), xmlgroup.Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: &[]xmlgroup.Slide{
			{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			{
				Items: &[]string{"Why WonderWidgets are great", "", "Who buys WonderWidgets"},
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutWrappedLists(t *testing.T) {
	client := xmlgroup.NewDefaultClient(nil).XMLOperations()
	result, err := client.PutWrappedLists(context.Background(), xmlgroup.AppleBarrel{
		BadApples:  &[]string{"Red Delicious"},
		GoodApples: &[]string{"Fuji", "Gala"},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}