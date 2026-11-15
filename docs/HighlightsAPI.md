# \HighlightsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContestsByIdHighlight**](HighlightsAPI.md#DeleteContestsByIdHighlight) | **Delete** /v2/contests/{id}/highlight | Remove a Contest from the tenant&#39;s highlights
[**GetContestsHighlights**](HighlightsAPI.md#GetContestsHighlights) | **Get** /v2/contests/highlights | List the tenant&#39;s highlighted Contests
[**UpdateContestsByIdHighlight**](HighlightsAPI.md#UpdateContestsByIdHighlight) | **Put** /v2/contests/{id}/highlight | Highlight a Contest or update its display order



## DeleteContestsByIdHighlight

> DeleteContestsByIdHighlight(ctx, id).XActingAs(xActingAs).Execute()

Remove a Contest from the tenant's highlights

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
	r, err := apiClient.HighlightsAPI.DeleteContestsByIdHighlight(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.DeleteContestsByIdHighlight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContestsByIdHighlightRequest struct via the builder pattern


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


## GetContestsHighlights

> GetContestsHighlights200Response GetContestsHighlights(ctx).From(from).To(to).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List the tenant's highlighted Contests



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
	from := "from_example" // string | Schedule window start, inclusive (RFC3339) (optional)
	to := "to_example" // string | Schedule window end, exclusive (RFC3339) (optional)
	cursor := "cursor_example" // string | Next-page cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, maximum 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.GetContestsHighlights(context.Background()).From(from).To(to).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.GetContestsHighlights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsHighlights`: GetContestsHighlights200Response
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.GetContestsHighlights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Schedule window start, inclusive (RFC3339) | 
 **to** | **string** | Schedule window end, exclusive (RFC3339) | 
 **cursor** | **string** | Next-page cursor | 
 **limit** | **int32** | Page size (default 50, maximum 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsHighlights200Response**](GetContestsHighlights200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestsByIdHighlight

> UpdateContestsByIdHighlight200Response UpdateContestsByIdHighlight(ctx, id).HighlighthttpAddHighlightRequest(highlighthttpAddHighlightRequest).XActingAs(xActingAs).Execute()

Highlight a Contest or update its display order



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
	highlighthttpAddHighlightRequest := *openapiclient.NewHighlighthttpAddHighlightRequest() // HighlighthttpAddHighlightRequest | Display position; defaults to 1
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.UpdateContestsByIdHighlight(context.Background(), id).HighlighthttpAddHighlightRequest(highlighthttpAddHighlightRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.UpdateContestsByIdHighlight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestsByIdHighlight`: UpdateContestsByIdHighlight200Response
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.UpdateContestsByIdHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContestsByIdHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **highlighthttpAddHighlightRequest** | [**HighlighthttpAddHighlightRequest**](HighlighthttpAddHighlightRequest.md) | Display position; defaults to 1 | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**UpdateContestsByIdHighlight200Response**](UpdateContestsByIdHighlight200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

