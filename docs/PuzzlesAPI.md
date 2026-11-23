# \PuzzlesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2PuzzleControllerCreate**](PuzzlesAPI.md#ContentWebV2PuzzleControllerCreate) | **Post** /v2/puzzles | Create puzzle
[**ContentWebV2PuzzleControllerDelete**](PuzzlesAPI.md#ContentWebV2PuzzleControllerDelete) | **Delete** /v2/puzzles/{id} | Delete puzzle
[**ContentWebV2PuzzleControllerGetTopicConfiguration**](PuzzlesAPI.md#ContentWebV2PuzzleControllerGetTopicConfiguration) | **Get** /v2/puzzles/get-topic-configuration | Get topic configuration
[**ContentWebV2PuzzleControllerGetUniqueCategories**](PuzzlesAPI.md#ContentWebV2PuzzleControllerGetUniqueCategories) | **Get** /v2/puzzles/categories | List unique categories
[**ContentWebV2PuzzleControllerIndex**](PuzzlesAPI.md#ContentWebV2PuzzleControllerIndex) | **Get** /v2/puzzles | List puzzles
[**ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch**](PuzzlesAPI.md#ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch) | **Get** /v2/puzzles/migrate | Migrate puzzles to OpenSearch
[**ContentWebV2PuzzleControllerSearchPuzzleBank**](PuzzlesAPI.md#ContentWebV2PuzzleControllerSearchPuzzleBank) | **Get** /v2/puzzles/search-puzzle-bank | Search puzzle bank
[**ContentWebV2PuzzleControllerShow**](PuzzlesAPI.md#ContentWebV2PuzzleControllerShow) | **Get** /v2/puzzles/{id} | Get puzzle
[**ContentWebV2PuzzleControllerUpdate**](PuzzlesAPI.md#ContentWebV2PuzzleControllerUpdate) | **Patch** /v2/puzzles/{id} | Update puzzle



## ContentWebV2PuzzleControllerCreate

> PuzzleResponse ContentWebV2PuzzleControllerCreate(ctx).XActingAs(xActingAs).CreatePuzzleRequest(createPuzzleRequest).Execute()

Create puzzle

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
	createPuzzleRequest := *openapiclient.NewCreatePuzzleRequest(map[string]interface{}{"key": interface{}(123)}) // CreatePuzzleRequest | Create puzzle request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerCreate(context.Background()).XActingAs(xActingAs).CreatePuzzleRequest(createPuzzleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerCreate`: PuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createPuzzleRequest** | [**CreatePuzzleRequest**](CreatePuzzleRequest.md) | Create puzzle request | 

### Return type

[**PuzzleResponse**](PuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerDelete

> ContentWebV2PuzzleControllerDelete(ctx, id).XActingAs(xActingAs).Execute()

Delete puzzle

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
	r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerDelete(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerDeleteRequest struct via the builder pattern


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


## ContentWebV2PuzzleControllerGetTopicConfiguration

> TopicConfigurationResponse ContentWebV2PuzzleControllerGetTopicConfiguration(ctx).Topic(topic).Format(format).XActingAs(xActingAs).Execute()

Get topic configuration



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
	topic := "topic_example" // string | 
	format := "format_example" // string |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerGetTopicConfiguration(context.Background()).Topic(topic).Format(format).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerGetTopicConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerGetTopicConfiguration`: TopicConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerGetTopicConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerGetTopicConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string** |  | 
 **format** | **string** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**TopicConfigurationResponse**](TopicConfigurationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerGetUniqueCategories

> CategoriesResponse ContentWebV2PuzzleControllerGetUniqueCategories(ctx).Format(format).XActingAs(xActingAs).Execute()

List unique categories



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
	format := "format_example" // string |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerGetUniqueCategories(context.Background()).Format(format).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerGetUniqueCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerGetUniqueCategories`: CategoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerGetUniqueCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerGetUniqueCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**CategoriesResponse**](CategoriesResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerIndex

> PuzzlesListResponse ContentWebV2PuzzleControllerIndex(ctx).Format(format).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List puzzles



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
	format := "format_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerIndex(context.Background()).Format(format).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerIndex`: PuzzlesListResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzlesListResponse**](PuzzlesListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch

> MessageResponse ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch(ctx).XActingAs(xActingAs).Execute()

Migrate puzzles to OpenSearch



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
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerMigratePuzzlesToOpensearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerMigratePuzzlesToOpensearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerSearchPuzzleBank

> PuzzlesListResponse ContentWebV2PuzzleControllerSearchPuzzleBank(ctx).Topic(topic).Format(format).Difficulty(difficulty).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Search puzzle bank



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
	topic := "topic_example" // string | 
	format := "format_example" // string |  (optional)
	difficulty := "difficulty_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerSearchPuzzleBank(context.Background()).Topic(topic).Format(format).Difficulty(difficulty).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerSearchPuzzleBank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerSearchPuzzleBank`: PuzzlesListResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerSearchPuzzleBank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerSearchPuzzleBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string** |  | 
 **format** | **string** |  | 
 **difficulty** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzlesListResponse**](PuzzlesListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerShow

> PuzzleResponse ContentWebV2PuzzleControllerShow(ctx, id).XActingAs(xActingAs).Execute()

Get puzzle

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
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerShow(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerShow`: PuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzleResponse**](PuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleControllerUpdate

> PuzzleResponse ContentWebV2PuzzleControllerUpdate(ctx, id).XActingAs(xActingAs).UpdatePuzzleRequest(updatePuzzleRequest).Execute()

Update puzzle

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
	updatePuzzleRequest := *openapiclient.NewUpdatePuzzleRequest() // UpdatePuzzleRequest | Update puzzle request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ContentWebV2PuzzleControllerUpdate(context.Background(), id).XActingAs(xActingAs).UpdatePuzzleRequest(updatePuzzleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ContentWebV2PuzzleControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleControllerUpdate`: PuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ContentWebV2PuzzleControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **updatePuzzleRequest** | [**UpdatePuzzleRequest**](UpdatePuzzleRequest.md) | Update puzzle request | 

### Return type

[**PuzzleResponse**](PuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

