# \EventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2EventControllerApproveResult**](EventsAPI.md#ContentWebV2EventControllerApproveResult) | **Post** /v2/events/{id}/approve | Approve event result
[**ContentWebV2EventControllerEnterResult**](EventsAPI.md#ContentWebV2EventControllerEnterResult) | **Patch** /v2/events/{id}/result | Enter event result
[**ContentWebV2EventControllerPendingApproval**](EventsAPI.md#ContentWebV2EventControllerPendingApproval) | **Get** /v2/events/pending-approval | List events with results pending approval
[**ContentWebV2EventControllerSearch**](EventsAPI.md#ContentWebV2EventControllerSearch) | **Get** /v2/events/search | Search events
[**ContentWebV2EventControllerShow**](EventsAPI.md#ContentWebV2EventControllerShow) | **Get** /v2/events/{id} | Get event by es_event_id
[**ContentWebV2EventControllerShowDetails**](EventsAPI.md#ContentWebV2EventControllerShowDetails) | **Get** /v2/events/{id}/details | Get event details (from Postgres)



## ContentWebV2EventControllerApproveResult

> EventResponse ContentWebV2EventControllerApproveResult(ctx, id).XActingAs(xActingAs).Execute()

Approve event result



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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ContentWebV2EventControllerApproveResult(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ContentWebV2EventControllerApproveResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventControllerApproveResult`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ContentWebV2EventControllerApproveResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventControllerApproveResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2EventControllerEnterResult

> EventResponse ContentWebV2EventControllerEnterResult(ctx, id).XActingAs(xActingAs).EnterEventResultRequest(enterEventResultRequest).Execute()

Enter event result

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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	enterEventResultRequest := *openapiclient.NewEnterEventResultRequest(map[string]interface{}{"key": interface{}(123)}) // EnterEventResultRequest | Result payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ContentWebV2EventControllerEnterResult(context.Background(), id).XActingAs(xActingAs).EnterEventResultRequest(enterEventResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ContentWebV2EventControllerEnterResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventControllerEnterResult`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ContentWebV2EventControllerEnterResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventControllerEnterResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **enterEventResultRequest** | [**EnterEventResultRequest**](EnterEventResultRequest.md) | Result payload | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2EventControllerPendingApproval

> EventsListResponse ContentWebV2EventControllerPendingApproval(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List events with results pending approval



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
	resp, r, err := apiClient.EventsAPI.ContentWebV2EventControllerPendingApproval(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ContentWebV2EventControllerPendingApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventControllerPendingApproval`: EventsListResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ContentWebV2EventControllerPendingApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventControllerPendingApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventsListResponse**](EventsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2EventControllerSearch

> EventsListResponse ContentWebV2EventControllerSearch(ctx).Q(q).DateFrom(dateFrom).DateTo(dateTo).Category(category).Competition(competition).IsUserEvent(isUserEvent).Size(size).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Search events



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
	q := "q_example" // string |  (optional)
	dateFrom := time.Now() // string |  (optional)
	dateTo := time.Now() // string |  (optional)
	category := "category_example" // string |  (optional)
	competition := "competition_example" // string |  (optional)
	isUserEvent := true // bool |  (optional)
	size := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ContentWebV2EventControllerSearch(context.Background()).Q(q).DateFrom(dateFrom).DateTo(dateTo).Category(category).Competition(competition).IsUserEvent(isUserEvent).Size(size).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ContentWebV2EventControllerSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventControllerSearch`: EventsListResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ContentWebV2EventControllerSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventControllerSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **category** | **string** |  | 
 **competition** | **string** |  | 
 **isUserEvent** | **bool** |  | 
 **size** | **int32** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventsListResponse**](EventsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2EventControllerShow

> EventResponse ContentWebV2EventControllerShow(ctx, id).XActingAs(xActingAs).Execute()

Get event by es_event_id

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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ContentWebV2EventControllerShow(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ContentWebV2EventControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventControllerShow`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ContentWebV2EventControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventControllerShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2EventControllerShowDetails

> EventResponse ContentWebV2EventControllerShowDetails(ctx, id).XActingAs(xActingAs).Execute()

Get event details (from Postgres)



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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ContentWebV2EventControllerShowDetails(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ContentWebV2EventControllerShowDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventControllerShowDetails`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ContentWebV2EventControllerShowDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventControllerShowDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

