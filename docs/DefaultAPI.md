# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentWebV2EventHighlightsControllerGenerate**](DefaultAPI.md#ContentWebV2EventHighlightsControllerGenerate) | **Post** /v2/event-highlights/generate | Queue shared daily event highlight preparation
[**ContentWebV2EventHighlightsControllerIndex**](DefaultAPI.md#ContentWebV2EventHighlightsControllerIndex) | **Get** /v2/event-highlights | Read Qaragon&#39;s shared daily event Highlights



## ContentWebV2EventHighlightsControllerGenerate

> EventHighlightsGenerationResponse ContentWebV2EventHighlightsControllerGenerate(ctx).XActingAs(xActingAs).EventHighlightsGenerationRequest(eventHighlightsGenerationRequest).Execute()

Queue shared daily event highlight preparation



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
	eventHighlightsGenerationRequest := *openapiclient.NewEventHighlightsGenerationRequest() // EventHighlightsGenerationRequest | UTC day and optional refresh (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ContentWebV2EventHighlightsControllerGenerate(context.Background()).XActingAs(xActingAs).EventHighlightsGenerationRequest(eventHighlightsGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContentWebV2EventHighlightsControllerGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventHighlightsControllerGenerate`: EventHighlightsGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContentWebV2EventHighlightsControllerGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventHighlightsControllerGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **eventHighlightsGenerationRequest** | [**EventHighlightsGenerationRequest**](EventHighlightsGenerationRequest.md) | UTC day and optional refresh | 

### Return type

[**EventHighlightsGenerationResponse**](EventHighlightsGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentWebV2EventHighlightsControllerIndex

> EventHighlightsResponse ContentWebV2EventHighlightsControllerIndex(ctx).Date(date).XActingAs(xActingAs).Execute()

Read Qaragon's shared daily event Highlights



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
	date := time.Now() // string |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ContentWebV2EventHighlightsControllerIndex(context.Background()).Date(date).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContentWebV2EventHighlightsControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentWebV2EventHighlightsControllerIndex`: EventHighlightsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContentWebV2EventHighlightsControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentWebV2EventHighlightsControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventHighlightsResponse**](EventHighlightsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

