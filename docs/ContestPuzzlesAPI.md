# \ContestPuzzlesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2ContestPuzzleControllerGetPuzzle**](ContestPuzzlesAPI.md#ContentWebV2ContestPuzzleControllerGetPuzzle) | **Get** /v2/contest-puzzles/{contest_id} | Get contest source
[**ContentWebV2ContestPuzzleControllerRankings**](ContestPuzzlesAPI.md#ContentWebV2ContestPuzzleControllerRankings) | **Get** /v2/contest-puzzles/{contest_id}/rankings | Get contest rankings
[**ContentWebV2ContestPuzzleControllerStatus**](ContestPuzzlesAPI.md#ContentWebV2ContestPuzzleControllerStatus) | **Get** /v2/contest-puzzles/{contest_id}/status | Get contest status



## ContentWebV2ContestPuzzleControllerGetPuzzle

> ContestPuzzleResponse ContentWebV2ContestPuzzleControllerGetPuzzle(ctx, contestId).XActingAs(xActingAs).Execute()

Get contest source



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
	contestId := "contestId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerGetPuzzle(context.Background(), contestId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerGetPuzzle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2ContestPuzzleControllerGetPuzzle`: ContestPuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerGetPuzzle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2ContestPuzzleControllerGetPuzzleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**ContestPuzzleResponse**](ContestPuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2ContestPuzzleControllerRankings

> RankingsResponse ContentWebV2ContestPuzzleControllerRankings(ctx, contestId).XActingAs(xActingAs).Execute()

Get contest rankings



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
	contestId := "contestId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerRankings(context.Background(), contestId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerRankings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2ContestPuzzleControllerRankings`: RankingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerRankings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2ContestPuzzleControllerRankingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**RankingsResponse**](RankingsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2ContestPuzzleControllerStatus

> ContestStatusResponse ContentWebV2ContestPuzzleControllerStatus(ctx, contestId).XActingAs(xActingAs).Execute()

Get contest status



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
	contestId := "contestId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerStatus(context.Background(), contestId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2ContestPuzzleControllerStatus`: ContestStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestPuzzlesAPI.ContentWebV2ContestPuzzleControllerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2ContestPuzzleControllerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**ContestStatusResponse**](ContestStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

