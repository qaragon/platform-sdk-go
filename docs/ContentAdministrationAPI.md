# \ContentAdministrationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdminAnswerBoards**](ContentAdministrationAPI.md#GetAdminAnswerBoards) | **Get** /v2/admin/answer-boards | Browse answer boards for publication review
[**GetAdminAnswerBoardsById**](ContentAdministrationAPI.md#GetAdminAnswerBoardsById) | **Get** /v2/admin/answer-boards/{id} | Read a candidate answer board with hidden answers
[**GetAdminAnswerBoardsGeneration**](ContentAdministrationAPI.md#GetAdminAnswerBoardsGeneration) | **Get** /v2/admin/answer-boards/generation | Read platform answer-board library generation settings
[**UpdateAdminAnswerBoardsByIdReview**](ContentAdministrationAPI.md#UpdateAdminAnswerBoardsByIdReview) | **Put** /v2/admin/answer-boards/{id}/review | Approve or reject an answer board for library publication
[**UpdateAdminAnswerBoardsGenerationByKind**](ContentAdministrationAPI.md#UpdateAdminAnswerBoardsGenerationByKind) | **Put** /v2/admin/answer-boards/generation/{kind} | Configure scheduled library generation



## GetAdminAnswerBoards

> AnswerboardsReviewBoardListResponse GetAdminAnswerBoards(ctx).Kind(kind).Status(status).After(after).Limit(limit).XActingAs(xActingAs).Execute()

Browse answer boards for publication review



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
	kind := "kind_example" // string | Board kind
	status := "status_example" // string | Review status (optional) (default to "pending_review")
	after := "after_example" // string | Last board ID from the previous page (optional)
	limit := int32(56) // int32 | Page size (1–100) (optional) (default to 20)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAdministrationAPI.GetAdminAnswerBoards(context.Background()).Kind(kind).Status(status).After(after).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAdministrationAPI.GetAdminAnswerBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAnswerBoards`: AnswerboardsReviewBoardListResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAdministrationAPI.GetAdminAnswerBoards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAnswerBoardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Board kind | 
 **status** | **string** | Review status | [default to &quot;pending_review&quot;]
 **after** | **string** | Last board ID from the previous page | 
 **limit** | **int32** | Page size (1–100) | [default to 20]
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**AnswerboardsReviewBoardListResponse**](AnswerboardsReviewBoardListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminAnswerBoardsById

> AnswerboardsReviewBoardResponse GetAdminAnswerBoardsById(ctx, id).XActingAs(xActingAs).Execute()

Read a candidate answer board with hidden answers



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
	id := "id_example" // string | Namespaced board ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAdministrationAPI.GetAdminAnswerBoardsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAdministrationAPI.GetAdminAnswerBoardsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAnswerBoardsById`: AnswerboardsReviewBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAdministrationAPI.GetAdminAnswerBoardsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespaced board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAnswerBoardsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**AnswerboardsReviewBoardResponse**](AnswerboardsReviewBoardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminAnswerBoardsGeneration

> AnswerboardsGenerationSettingsListResponse GetAdminAnswerBoardsGeneration(ctx).XActingAs(xActingAs).Execute()

Read platform answer-board library generation settings

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
	resp, r, err := apiClient.ContentAdministrationAPI.GetAdminAnswerBoardsGeneration(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAdministrationAPI.GetAdminAnswerBoardsGeneration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAnswerBoardsGeneration`: AnswerboardsGenerationSettingsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAdministrationAPI.GetAdminAnswerBoardsGeneration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAnswerBoardsGenerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**AnswerboardsGenerationSettingsListResponse**](AnswerboardsGenerationSettingsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminAnswerBoardsByIdReview

> AnswerboardsReviewBoardResponse UpdateAdminAnswerBoardsByIdReview(ctx, id).AnswerboardsReviewRequest(answerboardsReviewRequest).XActingAs(xActingAs).Execute()

Approve or reject an answer board for library publication



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
	id := "id_example" // string | Namespaced board ID
	answerboardsReviewRequest := *openapiclient.NewAnswerboardsReviewRequest() // AnswerboardsReviewRequest | Publication decision
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAdministrationAPI.UpdateAdminAnswerBoardsByIdReview(context.Background(), id).AnswerboardsReviewRequest(answerboardsReviewRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAdministrationAPI.UpdateAdminAnswerBoardsByIdReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdminAnswerBoardsByIdReview`: AnswerboardsReviewBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAdministrationAPI.UpdateAdminAnswerBoardsByIdReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespaced board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminAnswerBoardsByIdReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **answerboardsReviewRequest** | [**AnswerboardsReviewRequest**](AnswerboardsReviewRequest.md) | Publication decision | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**AnswerboardsReviewBoardResponse**](AnswerboardsReviewBoardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminAnswerBoardsGenerationByKind

> AnswerboardsGenerationSettingsResponse UpdateAdminAnswerBoardsGenerationByKind(ctx, kind).AnswerboardsGenerationSettings(answerboardsGenerationSettings).XActingAs(xActingAs).Execute()

Configure scheduled library generation



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
	kind := "kind_example" // string | Board kind
	answerboardsGenerationSettings := *openapiclient.NewAnswerboardsGenerationSettings("Brief_example", "Language_example") // AnswerboardsGenerationSettings | Operating settings
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAdministrationAPI.UpdateAdminAnswerBoardsGenerationByKind(context.Background(), kind).AnswerboardsGenerationSettings(answerboardsGenerationSettings).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAdministrationAPI.UpdateAdminAnswerBoardsGenerationByKind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdminAnswerBoardsGenerationByKind`: AnswerboardsGenerationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAdministrationAPI.UpdateAdminAnswerBoardsGenerationByKind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | Board kind | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminAnswerBoardsGenerationByKindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **answerboardsGenerationSettings** | [**AnswerboardsGenerationSettings**](AnswerboardsGenerationSettings.md) | Operating settings | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**AnswerboardsGenerationSettingsResponse**](AnswerboardsGenerationSettingsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

