# \GuessAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByIdGuessesClose**](GuessAPI.md#CreateContestsByIdGuessesClose) | **Post** /v2/contests/{id}/guesses/close | Close a Guess contest
[**GetAnswerBoards**](GuessAPI.md#GetAnswerBoards) | **Get** /v2/answer-boards | Browse prepared community answer boards
[**GetContestsByIdGuessConfig**](GuessAPI.md#GetContestsByIdGuessConfig) | **Get** /v2/contests/{id}/guess-config | Read Guess answer configuration
[**GetContestsByIdGuesses**](GuessAPI.md#GetContestsByIdGuesses) | **Get** /v2/contests/{id}/guesses | Read Guess rules



## CreateContestsByIdGuessesClose

> HttpxStatusResponseResponse CreateContestsByIdGuessesClose(ctx, id).XActingAs(xActingAs).Execute()

Close a Guess contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuessAPI.CreateContestsByIdGuessesClose(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuessAPI.CreateContestsByIdGuessesClose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdGuessesClose`: HttpxStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `GuessAPI.CreateContestsByIdGuessesClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdGuessesCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HttpxStatusResponseResponse**](HttpxStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnswerBoards

> AnswerboardsSummaryListResponse GetAnswerBoards(ctx).Kind(kind).After(after).Limit(limit).XActingAs(xActingAs).Execute()

Browse prepared community answer boards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	kind := "kind_example" // string | Board kind
	after := "after_example" // string | Last board ID from the preceding page (optional)
	limit := int32(56) // int32 | Page size (1–100) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuessAPI.GetAnswerBoards(context.Background()).Kind(kind).After(after).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuessAPI.GetAnswerBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnswerBoards`: AnswerboardsSummaryListResponse
	fmt.Fprintf(os.Stdout, "Response from `GuessAPI.GetAnswerBoards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnswerBoardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Board kind | 
 **after** | **string** | Last board ID from the preceding page | 
 **limit** | **int32** | Page size (1–100) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**AnswerboardsSummaryListResponse**](AnswerboardsSummaryListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdGuessConfig

> GuessConfigViewResponse GetContestsByIdGuessConfig(ctx, id).XActingAs(xActingAs).Execute()

Read Guess answer configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuessAPI.GetContestsByIdGuessConfig(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuessAPI.GetContestsByIdGuessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdGuessConfig`: GuessConfigViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GuessAPI.GetContestsByIdGuessConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdGuessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GuessConfigViewResponse**](GuessConfigViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdGuesses

> GuessGuessViewResponse GetContestsByIdGuesses(ctx, id).XActingAs(xActingAs).Execute()

Read Guess rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuessAPI.GetContestsByIdGuesses(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuessAPI.GetContestsByIdGuesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdGuesses`: GuessGuessViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GuessAPI.GetContestsByIdGuesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdGuessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GuessGuessViewResponse**](GuessGuessViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

