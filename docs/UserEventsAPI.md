# \UserEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2UserEventsControllerCreate**](UserEventsAPI.md#ContentWebV2UserEventsControllerCreate) | **Post** /v2/user-events | Create user event
[**ContentWebV2UserEventsControllerDelete**](UserEventsAPI.md#ContentWebV2UserEventsControllerDelete) | **Delete** /v2/user-events/{id} | Delete user event
[**ContentWebV2UserEventsControllerGetResultSuggestions**](UserEventsAPI.md#ContentWebV2UserEventsControllerGetResultSuggestions) | **Get** /v2/user-events/{user_events_id}/result-suggestions | Get result suggestions for a user event
[**ContentWebV2UserEventsControllerGetStatus**](UserEventsAPI.md#ContentWebV2UserEventsControllerGetStatus) | **Get** /v2/user-events/{user_events_id}/status | Get user event status
[**ContentWebV2UserEventsControllerIndex**](UserEventsAPI.md#ContentWebV2UserEventsControllerIndex) | **Get** /v2/user-events | List user events
[**ContentWebV2UserEventsControllerListTemplates**](UserEventsAPI.md#ContentWebV2UserEventsControllerListTemplates) | **Get** /v2/event-templates | List event templates
[**ContentWebV2UserEventsControllerShow**](UserEventsAPI.md#ContentWebV2UserEventsControllerShow) | **Get** /v2/user-events/{id} | Get user event
[**ContentWebV2UserEventsControllerUpdate**](UserEventsAPI.md#ContentWebV2UserEventsControllerUpdate) | **Patch** /v2/user-events/{id} | Update user event
[**ContentWebV2UserEventsControllerUseExistingResult**](UserEventsAPI.md#ContentWebV2UserEventsControllerUseExistingResult) | **Post** /v2/user-events/{user_events_id}/use-result | Copy result from another user event



## ContentWebV2UserEventsControllerCreate

> EventResponse ContentWebV2UserEventsControllerCreate(ctx).XActingAs(xActingAs).CreateUserEventRequest(createUserEventRequest).Execute()

Create user event

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
	createUserEventRequest := *openapiclient.NewCreateUserEventRequest("Name_example", "TemplateId_example") // CreateUserEventRequest | User event payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerCreate(context.Background()).XActingAs(xActingAs).CreateUserEventRequest(createUserEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerCreate`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createUserEventRequest** | [**CreateUserEventRequest**](CreateUserEventRequest.md) | User event payload | 

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


## ContentWebV2UserEventsControllerDelete

> ContentWebV2UserEventsControllerDelete(ctx, id).XActingAs(xActingAs).Execute()

Delete user event

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
	r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerDelete(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerDeleteRequest struct via the builder pattern


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


## ContentWebV2UserEventsControllerGetResultSuggestions

> ResultSuggestionsResponse ContentWebV2UserEventsControllerGetResultSuggestions(ctx, userEventsId).XActingAs(xActingAs).Execute()

Get result suggestions for a user event

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
	userEventsId := "userEventsId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerGetResultSuggestions(context.Background(), userEventsId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerGetResultSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerGetResultSuggestions`: ResultSuggestionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerGetResultSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEventsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerGetResultSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**ResultSuggestionsResponse**](ResultSuggestionsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2UserEventsControllerGetStatus

> UserEventStatusResponse ContentWebV2UserEventsControllerGetStatus(ctx, userEventsId).XActingAs(xActingAs).Execute()

Get user event status

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
	userEventsId := "userEventsId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerGetStatus(context.Background(), userEventsId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerGetStatus`: UserEventStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEventsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**UserEventStatusResponse**](UserEventStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2UserEventsControllerIndex

> ContentWebV2UserEventsControllerIndex200Response ContentWebV2UserEventsControllerIndex(ctx).All(all).Visibility(visibility).DateFrom(dateFrom).DateTo(dateTo).XActingAs(xActingAs).Execute()

List user events

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
	all := true // bool |  (optional)
	visibility := "visibility_example" // string |  (optional)
	dateFrom := time.Now() // time.Time |  (optional)
	dateTo := time.Now() // time.Time |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerIndex(context.Background()).All(all).Visibility(visibility).DateFrom(dateFrom).DateTo(dateTo).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerIndex`: ContentWebV2UserEventsControllerIndex200Response
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** |  | 
 **visibility** | **string** |  | 
 **dateFrom** | **time.Time** |  | 
 **dateTo** | **time.Time** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**ContentWebV2UserEventsControllerIndex200Response**](ContentWebV2UserEventsControllerIndex200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2UserEventsControllerListTemplates

> EventTemplatesListResponse ContentWebV2UserEventsControllerListTemplates(ctx).Category(category).EventType(eventType).XActingAs(xActingAs).Execute()

List event templates



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
	category := "category_example" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerListTemplates(context.Background()).Category(category).EventType(eventType).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerListTemplates`: EventTemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** |  | 
 **eventType** | **string** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventTemplatesListResponse**](EventTemplatesListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2UserEventsControllerShow

> EventResponse ContentWebV2UserEventsControllerShow(ctx, id).XActingAs(xActingAs).Execute()

Get user event

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
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerShow(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerShow`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerShowRequest struct via the builder pattern


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


## ContentWebV2UserEventsControllerUpdate

> EventResponse ContentWebV2UserEventsControllerUpdate(ctx, id).XActingAs(xActingAs).UpdateUserEventRequest(updateUserEventRequest).Execute()

Update user event

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
	updateUserEventRequest := *openapiclient.NewUpdateUserEventRequest() // UpdateUserEventRequest | User event update payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerUpdate(context.Background(), id).XActingAs(xActingAs).UpdateUserEventRequest(updateUserEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerUpdate`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **updateUserEventRequest** | [**UpdateUserEventRequest**](UpdateUserEventRequest.md) | User event update payload | 

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


## ContentWebV2UserEventsControllerUseExistingResult

> EventResponse ContentWebV2UserEventsControllerUseExistingResult(ctx, userEventsId).XActingAs(xActingAs).UseExistingResultRequest(useExistingResultRequest).Execute()

Copy result from another user event

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
	userEventsId := "userEventsId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	useExistingResultRequest := *openapiclient.NewUseExistingResultRequest("SourceEventId_example") // UseExistingResultRequest | Source event payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.ContentWebV2UserEventsControllerUseExistingResult(context.Background(), userEventsId).XActingAs(xActingAs).UseExistingResultRequest(useExistingResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.ContentWebV2UserEventsControllerUseExistingResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2UserEventsControllerUseExistingResult`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.ContentWebV2UserEventsControllerUseExistingResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEventsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2UserEventsControllerUseExistingResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **useExistingResultRequest** | [**UseExistingResultRequest**](UseExistingResultRequest.md) | Source event payload | 

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

