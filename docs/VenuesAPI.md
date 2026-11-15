# \VenuesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVenues**](VenuesAPI.md#CreateVenues) | **Post** /v2/venues | Create a venue under the calling tenant
[**CreateVenuesByVenueIdSpaces**](VenuesAPI.md#CreateVenuesByVenueIdSpaces) | **Post** /v2/venues/{venueId}/spaces | Create a space under one of the tenant&#39;s venues
[**GetVenues**](VenuesAPI.md#GetVenues) | **Get** /v2/venues | List venues owned by the calling tenant
[**GetVenuesByVenueIdSpaces**](VenuesAPI.md#GetVenuesByVenueIdSpaces) | **Get** /v2/venues/{venueId}/spaces | List spaces under one of the tenant&#39;s venues



## CreateVenues

> HttpxStatusResponseResponse CreateVenues(ctx).VenuehttpCreatePlatformVenueRequest(venuehttpCreatePlatformVenueRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Create a venue under the calling tenant

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
	venuehttpCreatePlatformVenueRequest := *openapiclient.NewVenuehttpCreatePlatformVenueRequest("Name_example", "Type_example") // VenuehttpCreatePlatformVenueRequest | Venue payload
	idempotencyKey := "idempotencyKey_example" // string | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.CreateVenues(context.Background()).VenuehttpCreatePlatformVenueRequest(venuehttpCreatePlatformVenueRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.CreateVenues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVenues`: HttpxStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.CreateVenues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVenuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **venuehttpCreatePlatformVenueRequest** | [**VenuehttpCreatePlatformVenueRequest**](VenuehttpCreatePlatformVenueRequest.md) | Venue payload | 
 **idempotencyKey** | **string** | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HttpxStatusResponseResponse**](HttpxStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVenuesByVenueIdSpaces

> HttpxStatusResponseResponse CreateVenuesByVenueIdSpaces(ctx, venueId).SpacehttpCreatePlatformSpaceRequest(spacehttpCreatePlatformSpaceRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Create a space under one of the tenant's venues

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
	venueId := "venueId_example" // string | Venue ID
	spacehttpCreatePlatformSpaceRequest := *openapiclient.NewSpacehttpCreatePlatformSpaceRequest("AccessType_example", "Name_example", int32(123)) // SpacehttpCreatePlatformSpaceRequest | Space payload
	idempotencyKey := "idempotencyKey_example" // string | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.CreateVenuesByVenueIdSpaces(context.Background(), venueId).SpacehttpCreatePlatformSpaceRequest(spacehttpCreatePlatformSpaceRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.CreateVenuesByVenueIdSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVenuesByVenueIdSpaces`: HttpxStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.CreateVenuesByVenueIdSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**venueId** | **string** | Venue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVenuesByVenueIdSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spacehttpCreatePlatformSpaceRequest** | [**SpacehttpCreatePlatformSpaceRequest**](SpacehttpCreatePlatformSpaceRequest.md) | Space payload | 
 **idempotencyKey** | **string** | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HttpxStatusResponseResponse**](HttpxStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVenues

> VenuesVenueListResponse GetVenues(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List venues owned by the calling tenant

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.GetVenues(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.GetVenues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVenues`: VenuesVenueListResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.GetVenues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVenuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**VenuesVenueListResponse**](VenuesVenueListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVenuesByVenueIdSpaces

> SpacesSpaceListResponse GetVenuesByVenueIdSpaces(ctx, venueId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List spaces under one of the tenant's venues

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
	venueId := "venueId_example" // string | Venue ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.GetVenuesByVenueIdSpaces(context.Background(), venueId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.GetVenuesByVenueIdSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVenuesByVenueIdSpaces`: SpacesSpaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.GetVenuesByVenueIdSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**venueId** | **string** | Venue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVenuesByVenueIdSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceListResponse**](SpacesSpaceListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

