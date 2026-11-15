# \SubmissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubmissions**](SubmissionsAPI.md#CreateSubmissions) | **Post** /v2/submissions | Submit to a Contest or Game
[**GetSubmissions**](SubmissionsAPI.md#GetSubmissions) | **Get** /v2/submissions | List accepted submissions for a target
[**GetSubmissionsById**](SubmissionsAPI.md#GetSubmissionsById) | **Get** /v2/submissions/{id} | Read an accepted submission



## CreateSubmissions

> SubmissionReceiptResponse CreateSubmissions(ctx).SubmissionRequest(submissionRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Submit to a Contest or Game



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
	submissionRequest := *openapiclient.NewSubmissionRequest() // SubmissionRequest | Submission target and type-specific payload
	idempotencyKey := "idempotencyKey_example" // string | Retry key; reuse only for the same identity, target, and action (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubmissionsAPI.CreateSubmissions(context.Background()).SubmissionRequest(submissionRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubmissionsAPI.CreateSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubmissions`: SubmissionReceiptResponse
	fmt.Fprintf(os.Stdout, "Response from `SubmissionsAPI.CreateSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submissionRequest** | [**SubmissionRequest**](SubmissionRequest.md) | Submission target and type-specific payload | 
 **idempotencyKey** | **string** | Retry key; reuse only for the same identity, target, and action | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SubmissionReceiptResponse**](SubmissionReceiptResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubmissions

> SubmissionViewListResponse GetSubmissions(ctx).TargetType(targetType).TargetId(targetId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List accepted submissions for a target



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
	targetType := "targetType_example" // string | Target type
	targetId := "targetId_example" // string | External target ID
	cursor := "cursor_example" // string | Cursor (optional)
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubmissionsAPI.GetSubmissions(context.Background()).TargetType(targetType).TargetId(targetId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubmissionsAPI.GetSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubmissions`: SubmissionViewListResponse
	fmt.Fprintf(os.Stdout, "Response from `SubmissionsAPI.GetSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetType** | **string** | Target type | 
 **targetId** | **string** | External target ID | 
 **cursor** | **string** | Cursor | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SubmissionViewListResponse**](SubmissionViewListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubmissionsById

> SubmissionViewResponse GetSubmissionsById(ctx, id).XActingAs(xActingAs).Execute()

Read an accepted submission



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
	id := "id_example" // string | Submission ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubmissionsAPI.GetSubmissionsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubmissionsAPI.GetSubmissionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubmissionsById`: SubmissionViewResponse
	fmt.Fprintf(os.Stdout, "Response from `SubmissionsAPI.GetSubmissionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Submission ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmissionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SubmissionViewResponse**](SubmissionViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

