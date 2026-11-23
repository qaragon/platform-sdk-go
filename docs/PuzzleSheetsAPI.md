# \PuzzleSheetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2PuzzleSheetControllerClonePuzzleSheet**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerClonePuzzleSheet) | **Post** /v2/puzzle-sheets/{id}/clone | Clone puzzle sheet
[**ContentWebV2PuzzleSheetControllerCreate**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerCreate) | **Post** /v2/puzzle-sheets | Create puzzle sheet
[**ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet) | **Post** /v2/puzzle-sheets/adhoc-community-sheet | Create adhoc community puzzle sheet
[**ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet) | **Post** /v2/puzzle-sheets/adhoc-user-sheet | Create adhoc user puzzle sheet
[**ContentWebV2PuzzleSheetControllerDelete**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerDelete) | **Delete** /v2/puzzle-sheets/{id} | Delete puzzle sheet
[**ContentWebV2PuzzleSheetControllerIndex**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerIndex) | **Get** /v2/puzzle-sheets | List puzzle sheets
[**ContentWebV2PuzzleSheetControllerShow**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerShow) | **Get** /v2/puzzle-sheets/{id} | Get puzzle sheet
[**ContentWebV2PuzzleSheetControllerUpdate**](PuzzleSheetsAPI.md#ContentWebV2PuzzleSheetControllerUpdate) | **Patch** /v2/puzzle-sheets/{id} | Update puzzle sheet



## ContentWebV2PuzzleSheetControllerClonePuzzleSheet

> PuzzleSheetResponse ContentWebV2PuzzleSheetControllerClonePuzzleSheet(ctx, id).XActingAs(xActingAs).ClonePuzzleSheetRequest(clonePuzzleSheetRequest).Execute()

Clone puzzle sheet

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
	clonePuzzleSheetRequest := *openapiclient.NewClonePuzzleSheetRequest() // ClonePuzzleSheetRequest | Clone payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerClonePuzzleSheet(context.Background(), id).XActingAs(xActingAs).ClonePuzzleSheetRequest(clonePuzzleSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerClonePuzzleSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerClonePuzzleSheet`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerClonePuzzleSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerClonePuzzleSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **clonePuzzleSheetRequest** | [**ClonePuzzleSheetRequest**](ClonePuzzleSheetRequest.md) | Clone payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleSheetControllerCreate

> PuzzleSheetResponse ContentWebV2PuzzleSheetControllerCreate(ctx).XActingAs(xActingAs).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()

Create puzzle sheet

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
	createPuzzleSheetRequest := *openapiclient.NewCreatePuzzleSheetRequest("Name_example", "private") // CreatePuzzleSheetRequest | Puzzle sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreate(context.Background()).XActingAs(xActingAs).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerCreate`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createPuzzleSheetRequest** | [**CreatePuzzleSheetRequest**](CreatePuzzleSheetRequest.md) | Puzzle sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet

> PuzzleSheetResponse ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet(ctx).XActingAs(xActingAs).CreateAdhocCommunitySheetRequest(createAdhocCommunitySheetRequest).Execute()

Create adhoc community puzzle sheet



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
	createAdhocCommunitySheetRequest := *openapiclient.NewCreateAdhocCommunitySheetRequest("Difficulty_example", int32(123), "Topic_example") // CreateAdhocCommunitySheetRequest | Adhoc community sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet(context.Background()).XActingAs(xActingAs).CreateAdhocCommunitySheetRequest(createAdhocCommunitySheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerCreateAdhocCommunitySheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createAdhocCommunitySheetRequest** | [**CreateAdhocCommunitySheetRequest**](CreateAdhocCommunitySheetRequest.md) | Adhoc community sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet

> PuzzleSheetResponse ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet(ctx).XActingAs(xActingAs).CreateAdhocUserSheetRequest(createAdhocUserSheetRequest).Execute()

Create adhoc user puzzle sheet



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
	createAdhocUserSheetRequest := *openapiclient.NewCreateAdhocUserSheetRequest(int32(123), "PuzzleSheetId_example") // CreateAdhocUserSheetRequest | Adhoc user sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet(context.Background()).XActingAs(xActingAs).CreateAdhocUserSheetRequest(createAdhocUserSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerCreateAdhocUserSheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerCreateAdhocUserSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createAdhocUserSheetRequest** | [**CreateAdhocUserSheetRequest**](CreateAdhocUserSheetRequest.md) | Adhoc user sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleSheetControllerDelete

> ContentWebV2PuzzleSheetControllerDelete(ctx, id).XActingAs(xActingAs).Execute()

Delete puzzle sheet

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
	r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerDelete(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerDeleteRequest struct via the builder pattern


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


## ContentWebV2PuzzleSheetControllerIndex

> PuzzleSheetsListResponse ContentWebV2PuzzleSheetControllerIndex(ctx).Format(format).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List puzzle sheets



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
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerIndex(context.Background()).Format(format).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerIndex`: PuzzleSheetsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzleSheetsListResponse**](PuzzleSheetsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleSheetControllerShow

> PuzzleSheetResponse ContentWebV2PuzzleSheetControllerShow(ctx, id).XActingAs(xActingAs).Execute()

Get puzzle sheet

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
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerShow(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerShow`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2PuzzleSheetControllerUpdate

> PuzzleSheetResponse ContentWebV2PuzzleSheetControllerUpdate(ctx, id).XActingAs(xActingAs).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()

Update puzzle sheet

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
	createPuzzleSheetRequest := *openapiclient.NewCreatePuzzleSheetRequest("Name_example", "private") // CreatePuzzleSheetRequest | Puzzle sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerUpdate(context.Background(), id).XActingAs(xActingAs).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2PuzzleSheetControllerUpdate`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.ContentWebV2PuzzleSheetControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2PuzzleSheetControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createPuzzleSheetRequest** | [**CreatePuzzleSheetRequest**](CreatePuzzleSheetRequest.md) | Puzzle sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

