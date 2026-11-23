# \LeaderboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpacesByIdStandings**](LeaderboardAPI.md#GetSpacesByIdStandings) | **Get** /v2/spaces/{id}/standings | Get a space&#39;s season standings (ranked by cumulative participation points)



## GetSpacesByIdStandings

> DomainLeaderboardResponse GetSpacesByIdStandings(ctx, id).Window(window).Actor(actor).XActingAs(xActingAs).Execute()

Get a space's season standings (ranked by cumulative participation points)



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
	id := "id_example" // string | Space ID
	window := "window_example" // string | Standings window: week | month | all_time (default all_time) (optional)
	actor := "actor_example" // string | member (truncated view) | owner (full board, default) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetSpacesByIdStandings(context.Background(), id).Window(window).Actor(actor).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetSpacesByIdStandings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesByIdStandings`: DomainLeaderboardResponse
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetSpacesByIdStandings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesByIdStandingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **window** | **string** | Standings window: week | month | all_time (default all_time) | 
 **actor** | **string** | member (truncated view) | owner (full board, default) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainLeaderboardResponse**](DomainLeaderboardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

