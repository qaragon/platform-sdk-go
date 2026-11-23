# \PredictionSlipsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2PredictionControllerClonePredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerClonePredictionSlip) | **Post** /v2/prediction-slips/{id}/clone | Clone prediction slip
[**ContentWebV2PredictionControllerCreatePredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerCreatePredictionSlip) | **Post** /v2/prediction-slips | Create prediction slip
[**ContentWebV2PredictionControllerDeletePredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerDeletePredictionSlip) | **Delete** /v2/prediction-slips/{id} | Delete prediction slip
[**ContentWebV2PredictionControllerListCommunityPredictionSlips**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerListCommunityPredictionSlips) | **Get** /v2/prediction-slips/community | List public/community prediction slips
[**ContentWebV2PredictionControllerListCuratedSlips**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerListCuratedSlips) | **Get** /v2/curated-slips | List curated prediction slips
[**ContentWebV2PredictionControllerListPredictionSlips**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerListPredictionSlips) | **Get** /v2/prediction-slips | List the caller&#39;s prediction slips
[**ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip) | **Post** /v2/prediction-slips/from-highlights | Prepare a Prediction slip from an upcoming highlighted event
[**ContentWebV2PredictionControllerShowOpenPredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerShowOpenPredictionSlip) | **Get** /v2/prediction-slips/open | Get the caller&#39;s open prediction slip
[**ContentWebV2PredictionControllerShowPredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerShowPredictionSlip) | **Get** /v2/prediction-slips/{id} | Get prediction slip
[**ContentWebV2PredictionControllerUpdatePredictionSlip**](PredictionSlipsAPI.md#ContentWebV2PredictionControllerUpdatePredictionSlip) | **Put** /v2/prediction-slips/{id} | Update prediction slip



## ContentWebV2PredictionControllerClonePredictionSlip

> PredictionSlipResponse ContentWebV2PredictionControllerClonePredictionSlip(ctx, id).XActingAs(xActingAs).ClonePredictionSlipRequest(clonePredictionSlipRequest).Execute()

Clone prediction slip

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	clonePredictionSlipRequest := *openapiclient.NewClonePredictionSlipRequest() // ClonePredictionSlipRequest | Clone payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerClonePredictionSlip(context.Background(), id).XActingAs(xActingAs).ClonePredictionSlipRequest(clonePredictionSlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerClonePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerClonePredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerClonePredictionSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerClonePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **clonePredictionSlipRequest** | [**ClonePredictionSlipRequest**](ClonePredictionSlipRequest.md) | Clone payload | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerCreatePredictionSlip

> PredictionSlipResponse ContentWebV2PredictionControllerCreatePredictionSlip(ctx).XActingAs(xActingAs).CreatePredictionSlipRequest(createPredictionSlipRequest).Execute()

Create prediction slip

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	createPredictionSlipRequest := *openapiclient.NewCreatePredictionSlipRequest([]openapiclient.CreatePredictionSlipRequestEventsInner{*openapiclient.NewCreatePredictionSlipRequestEventsInner("EsEventId_example", []openapiclient.CreatePredictionSlipRequestEventsInnerItemsInner{*openapiclient.NewCreatePredictionSlipRequestEventsInnerItemsInner()})}, "Name_example") // CreatePredictionSlipRequest | Prediction slip payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerCreatePredictionSlip(context.Background()).XActingAs(xActingAs).CreatePredictionSlipRequest(createPredictionSlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerCreatePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerCreatePredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerCreatePredictionSlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerCreatePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createPredictionSlipRequest** | [**CreatePredictionSlipRequest**](CreatePredictionSlipRequest.md) | Prediction slip payload | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerDeletePredictionSlip

> ContentWebV2PredictionControllerDeletePredictionSlip(ctx, id).XActingAs(xActingAs).Execute()

Delete prediction slip

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerDeletePredictionSlip(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerDeletePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerDeletePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

 (empty response body)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerListCommunityPredictionSlips

> PredictionSlipsListResponse ContentWebV2PredictionControllerListCommunityPredictionSlips(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List public/community prediction slips



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerListCommunityPredictionSlips(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerListCommunityPredictionSlips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerListCommunityPredictionSlips`: PredictionSlipsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerListCommunityPredictionSlips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerListCommunityPredictionSlipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipsListResponse**](PredictionSlipsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerListCuratedSlips

> PredictionSlipsListResponse ContentWebV2PredictionControllerListCuratedSlips(ctx).Family(family).Tags(tags).IsPublic(isPublic).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List curated prediction slips

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	family := "family_example" // string |  (optional)
	tags := "tags_example" // string |  (optional)
	isPublic := true // bool |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerListCuratedSlips(context.Background()).Family(family).Tags(tags).IsPublic(isPublic).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerListCuratedSlips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerListCuratedSlips`: PredictionSlipsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerListCuratedSlips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerListCuratedSlipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **family** | **string** |  | 
 **tags** | **string** |  | 
 **isPublic** | **bool** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipsListResponse**](PredictionSlipsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerListPredictionSlips

> PredictionSlipsListResponse ContentWebV2PredictionControllerListPredictionSlips(ctx).Locked(locked).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List the caller's prediction slips

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	locked := true // bool |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerListPredictionSlips(context.Background()).Locked(locked).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerListPredictionSlips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerListPredictionSlips`: PredictionSlipsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerListPredictionSlips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerListPredictionSlipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locked** | **bool** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipsListResponse**](PredictionSlipsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip

> HighlightedPredictionSlipResponse ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip(ctx).HighlightedPredictionSlipRequest(highlightedPredictionSlipRequest).XActingAs(xActingAs).Execute()

Prepare a Prediction slip from an upcoming highlighted event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	highlightedPredictionSlipRequest := *openapiclient.NewHighlightedPredictionSlipRequest(time.Now()) // HighlightedPredictionSlipRequest | Daily preparation window
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip(context.Background()).HighlightedPredictionSlipRequest(highlightedPredictionSlipRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip`: HighlightedPredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerPrepareHighlightedPredictionSlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerPrepareHighlightedPredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **highlightedPredictionSlipRequest** | [**HighlightedPredictionSlipRequest**](HighlightedPredictionSlipRequest.md) | Daily preparation window | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HighlightedPredictionSlipResponse**](HighlightedPredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerShowOpenPredictionSlip

> OpenPredictionSlipResponse ContentWebV2PredictionControllerShowOpenPredictionSlip(ctx).XActingAs(xActingAs).Execute()

Get the caller's open prediction slip



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerShowOpenPredictionSlip(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerShowOpenPredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerShowOpenPredictionSlip`: OpenPredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerShowOpenPredictionSlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerShowOpenPredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**OpenPredictionSlipResponse**](OpenPredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerShowPredictionSlip

> PredictionSlipResponse ContentWebV2PredictionControllerShowPredictionSlip(ctx, id).XActingAs(xActingAs).Execute()

Get prediction slip

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerShowPredictionSlip(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerShowPredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerShowPredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerShowPredictionSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerShowPredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PredictionControllerUpdatePredictionSlip

> PredictionSlipResponse ContentWebV2PredictionControllerUpdatePredictionSlip(ctx, id).XActingAs(xActingAs).UpdatePredictionSlipRequest(updatePredictionSlipRequest).Execute()

Update prediction slip

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qaragon/platform-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	updatePredictionSlipRequest := *openapiclient.NewUpdatePredictionSlipRequest() // UpdatePredictionSlipRequest | Prediction slip update payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.ContentWebV2PredictionControllerUpdatePredictionSlip(context.Background(), id).XActingAs(xActingAs).UpdatePredictionSlipRequest(updatePredictionSlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.ContentWebV2PredictionControllerUpdatePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PredictionControllerUpdatePredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.ContentWebV2PredictionControllerUpdatePredictionSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PredictionControllerUpdatePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **updatePredictionSlipRequest** | [**UpdatePredictionSlipRequest**](UpdatePredictionSlipRequest.md) | Prediction slip update payload | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

