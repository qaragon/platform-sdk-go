# \CheckInAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContestsByIdAssignments**](CheckInAPI.md#GetContestsByIdAssignments) | **Get** /v2/contests/{id}/assignments | Get the caller&#39;s check-in assignment status
[**GetContestsByIdRooms**](CheckInAPI.md#GetContestsByIdRooms) | **Get** /v2/contests/{id}/rooms | List every room in a check-in contest
[**GetContestsByIdRoomsByRoomNumber**](CheckInAPI.md#GetContestsByIdRoomsByRoomNumber) | **Get** /v2/contests/{id}/rooms/{roomNumber} | Get one room in a check-in contest



## GetContestsByIdAssignments

> CheckinAssignmentResponseResponse GetContestsByIdAssignments(ctx, id).XActingAs(xActingAs).Execute()

Get the caller's check-in assignment status



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
	resp, r, err := apiClient.CheckInAPI.GetContestsByIdAssignments(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckInAPI.GetContestsByIdAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdAssignments`: CheckinAssignmentResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckInAPI.GetContestsByIdAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**CheckinAssignmentResponseResponse**](CheckinAssignmentResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdRooms

> CheckinRoomListResponseResponse GetContestsByIdRooms(ctx, id).After(after).Limit(limit).XActingAs(xActingAs).Execute()

List every room in a check-in contest



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
	after := int32(56) // int32 | Last room_number seen (0 for the first page) (optional)
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckInAPI.GetContestsByIdRooms(context.Background(), id).After(after).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckInAPI.GetContestsByIdRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdRooms`: CheckinRoomListResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckInAPI.GetContestsByIdRooms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdRoomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int32** | Last room_number seen (0 for the first page) | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**CheckinRoomListResponseResponse**](CheckinRoomListResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdRoomsByRoomNumber

> CheckinRoomResponseResponse GetContestsByIdRoomsByRoomNumber(ctx, id, roomNumber).XActingAs(xActingAs).Execute()

Get one room in a check-in contest



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
	roomNumber := int32(56) // int32 | Room number
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckInAPI.GetContestsByIdRoomsByRoomNumber(context.Background(), id, roomNumber).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckInAPI.GetContestsByIdRoomsByRoomNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdRoomsByRoomNumber`: CheckinRoomResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckInAPI.GetContestsByIdRoomsByRoomNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**roomNumber** | **int32** | Room number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdRoomsByRoomNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**CheckinRoomResponseResponse**](CheckinRoomResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

