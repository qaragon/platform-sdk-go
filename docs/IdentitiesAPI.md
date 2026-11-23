# \IdentitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentities**](IdentitiesAPI.md#CreateIdentities) | **Post** /v2/identities | Register a platform identity (platform mints the id)
[**DeleteIdentitiesById**](IdentitiesAPI.md#DeleteIdentitiesById) | **Delete** /v2/identities/{id} | Soft-delete a platform identity, scoped to the calling tenant
[**GetIdentitiesById**](IdentitiesAPI.md#GetIdentitiesById) | **Get** /v2/identities/{id} | Get a platform identity by ID, scoped to the calling tenant
[**UpdateIdentitiesById**](IdentitiesAPI.md#UpdateIdentitiesById) | **Put** /v2/identities/{id} | Replace a platform identity&#39;s profile, scoped to the calling tenant



## CreateIdentities

> IdentitiesPlatformIdentityResponse CreateIdentities(ctx).CreatePlatformIdentityRequest(createPlatformIdentityRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Register a platform identity (platform mints the id)



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
	createPlatformIdentityRequest := *openapiclient.NewCreatePlatformIdentityRequest() // CreatePlatformIdentityRequest | Platform identity payload
	idempotencyKey := "idempotencyKey_example" // string | Retry key (~24h). Same key + same body replays the original response; a different body returns 422. (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.CreateIdentities(context.Background()).CreatePlatformIdentityRequest(createPlatformIdentityRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.CreateIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentities`: IdentitiesPlatformIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.CreateIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlatformIdentityRequest** | [**CreatePlatformIdentityRequest**](CreatePlatformIdentityRequest.md) | Platform identity payload | 
 **idempotencyKey** | **string** | Retry key (~24h). Same key + same body replays the original response; a different body returns 422. | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**IdentitiesPlatformIdentityResponse**](IdentitiesPlatformIdentityResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentitiesById

> DeleteIdentitiesById(ctx, id).XActingAs(xActingAs).Execute()

Soft-delete a platform identity, scoped to the calling tenant



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
	id := "id_example" // string | Platform Identity ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitiesAPI.DeleteIdentitiesById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.DeleteIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitiesByIdRequest struct via the builder pattern


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


## GetIdentitiesById

> IdentitiesPlatformIdentityResponse GetIdentitiesById(ctx, id).XActingAs(xActingAs).Execute()

Get a platform identity by ID, scoped to the calling tenant



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
	id := "id_example" // string | Platform Identity ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetIdentitiesById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesById`: IdentitiesPlatformIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetIdentitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**IdentitiesPlatformIdentityResponse**](IdentitiesPlatformIdentityResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentitiesById

> IdentitiesPlatformIdentityResponse UpdateIdentitiesById(ctx, id).UpdatePlatformIdentityRequest(updatePlatformIdentityRequest).XActingAs(xActingAs).Execute()

Replace a platform identity's profile, scoped to the calling tenant



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
	id := "id_example" // string | Platform Identity ID
	updatePlatformIdentityRequest := *openapiclient.NewUpdatePlatformIdentityRequest("DisplayName_example") // UpdatePlatformIdentityRequest | Profile fields
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.UpdateIdentitiesById(context.Background(), id).UpdatePlatformIdentityRequest(updatePlatformIdentityRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.UpdateIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentitiesById`: IdentitiesPlatformIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.UpdateIdentitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePlatformIdentityRequest** | [**UpdatePlatformIdentityRequest**](UpdatePlatformIdentityRequest.md) | Profile fields | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**IdentitiesPlatformIdentityResponse**](IdentitiesPlatformIdentityResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

