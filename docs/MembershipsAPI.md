# \MembershipsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpacesByIdMemberships**](MembershipsAPI.md#CreateSpacesByIdMemberships) | **Post** /v2/spaces/{id}/memberships | Join a space
[**DeleteSpacesByIdMembershipsByMembershipId**](MembershipsAPI.md#DeleteSpacesByIdMembershipsByMembershipId) | **Delete** /v2/spaces/{id}/memberships/{membershipId} | Leave a space
[**GetIdentitiesByPiidMemberships**](MembershipsAPI.md#GetIdentitiesByPiidMemberships) | **Get** /v2/identities/{piid}/memberships | Get a platform identity&#39;s space memberships (v2)
[**GetSpacesByIdMemberships**](MembershipsAPI.md#GetSpacesByIdMemberships) | **Get** /v2/spaces/{id}/memberships | List space memberships (cursor-paginated)
[**GetSpacesByIdMembershipsByMembershipIdLogs**](MembershipsAPI.md#GetSpacesByIdMembershipsByMembershipIdLogs) | **Get** /v2/spaces/{id}/memberships/{membershipId}/logs | Get the participation log for a membership
[**GetSpacesMembershipsSearchKeywordByKeyword**](MembershipsAPI.md#GetSpacesMembershipsSearchKeywordByKeyword) | **Get** /v2/spaces/memberships/searchKeyword/{keyword} | Search memberships by keyword (cursor-paginated)
[**UpdateSpacesByIdMembershipsByMembershipIdStatus**](MembershipsAPI.md#UpdateSpacesByIdMembershipsByMembershipIdStatus) | **Put** /v2/spaces/{id}/memberships/{membershipId}/status | Approve or deny membership



## CreateSpacesByIdMemberships

> SpacesSpaceMembershipResponse CreateSpacesByIdMemberships(ctx, id).SpacehttpCreateMembershipRequest(spacehttpCreateMembershipRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Join a space

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
	spacehttpCreateMembershipRequest := *openapiclient.NewSpacehttpCreateMembershipRequest() // SpacehttpCreateMembershipRequest | Membership payload
	idempotencyKey := "idempotencyKey_example" // string | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.CreateSpacesByIdMemberships(context.Background(), id).SpacehttpCreateMembershipRequest(spacehttpCreateMembershipRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.CreateSpacesByIdMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpacesByIdMemberships`: SpacesSpaceMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.CreateSpacesByIdMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpacesByIdMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spacehttpCreateMembershipRequest** | [**SpacehttpCreateMembershipRequest**](SpacehttpCreateMembershipRequest.md) | Membership payload | 
 **idempotencyKey** | **string** | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceMembershipResponse**](SpacesSpaceMembershipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpacesByIdMembershipsByMembershipId

> DeleteSpacesByIdMembershipsByMembershipId(ctx, id, membershipId).XActingAs(xActingAs).Execute()

Leave a space



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
	membershipId := "membershipId_example" // string | Membership ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.DeleteSpacesByIdMembershipsByMembershipId(context.Background(), id, membershipId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.DeleteSpacesByIdMembershipsByMembershipId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 
**membershipId** | **string** | Membership ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpacesByIdMembershipsByMembershipIdRequest struct via the builder pattern


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


## GetIdentitiesByPiidMemberships

> SpacesSpaceMembershipListResponse GetIdentitiesByPiidMemberships(ctx, piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get a platform identity's space memberships (v2)



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
	piid := "piid_example" // string | Platform Identity ID
	cursor := "cursor_example" // string | Cursor for pagination (optional)
	limit := int32(56) // int32 | Items per page (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetIdentitiesByPiidMemberships(context.Background(), piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetIdentitiesByPiidMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidMemberships`: SpacesSpaceMembershipListResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetIdentitiesByPiidMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Cursor for pagination | 
 **limit** | **int32** | Items per page | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceMembershipListResponse**](SpacesSpaceMembershipListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacesByIdMemberships

> SpacesSpaceMembershipListResponse GetSpacesByIdMemberships(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List space memberships (cursor-paginated)



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetSpacesByIdMemberships(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetSpacesByIdMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesByIdMemberships`: SpacesSpaceMembershipListResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetSpacesByIdMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesByIdMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceMembershipListResponse**](SpacesSpaceMembershipListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacesByIdMembershipsByMembershipIdLogs

> DomainParticipationListResponse GetSpacesByIdMembershipsByMembershipIdLogs(ctx, id, membershipId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get the participation log for a membership

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
	membershipId := "membershipId_example" // string | Membership ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetSpacesByIdMembershipsByMembershipIdLogs(context.Background(), id, membershipId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetSpacesByIdMembershipsByMembershipIdLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesByIdMembershipsByMembershipIdLogs`: DomainParticipationListResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetSpacesByIdMembershipsByMembershipIdLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 
**membershipId** | **string** | Membership ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesByIdMembershipsByMembershipIdLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainParticipationListResponse**](DomainParticipationListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacesMembershipsSearchKeywordByKeyword

> SpacesSpaceMembershipListResponse GetSpacesMembershipsSearchKeywordByKeyword(ctx, keyword).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Search memberships by keyword (cursor-paginated)

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
	keyword := "keyword_example" // string | Search keyword
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetSpacesMembershipsSearchKeywordByKeyword(context.Background(), keyword).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetSpacesMembershipsSearchKeywordByKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesMembershipsSearchKeywordByKeyword`: SpacesSpaceMembershipListResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetSpacesMembershipsSearchKeywordByKeyword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string** | Search keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesMembershipsSearchKeywordByKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceMembershipListResponse**](SpacesSpaceMembershipListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpacesByIdMembershipsByMembershipIdStatus

> UpdateSpacesByIdMembershipsByMembershipIdStatus(ctx, id, membershipId).SpacehttpUpdateMembershipStatusRequest(spacehttpUpdateMembershipStatusRequest).XActingAs(xActingAs).Execute()

Approve or deny membership



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
	membershipId := "membershipId_example" // string | Membership ID
	spacehttpUpdateMembershipStatusRequest := *openapiclient.NewSpacehttpUpdateMembershipStatusRequest("Status_example") // SpacehttpUpdateMembershipStatusRequest | request
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.UpdateSpacesByIdMembershipsByMembershipIdStatus(context.Background(), id, membershipId).SpacehttpUpdateMembershipStatusRequest(spacehttpUpdateMembershipStatusRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.UpdateSpacesByIdMembershipsByMembershipIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 
**membershipId** | **string** | Membership ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpacesByIdMembershipsByMembershipIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spacehttpUpdateMembershipStatusRequest** | [**SpacehttpUpdateMembershipStatusRequest**](SpacehttpUpdateMembershipStatusRequest.md) | request | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

 (empty response body)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

