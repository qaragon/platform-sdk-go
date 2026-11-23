# \SpacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpaces**](SpacesAPI.md#CreateSpaces) | **Post** /v2/spaces | Create a space
[**CreateSpacesByIdReport**](SpacesAPI.md#CreateSpacesByIdReport) | **Post** /v2/spaces/{id}/report | Report a space
[**CreateSpacesByIdTeamSets**](SpacesAPI.md#CreateSpacesByIdTeamSets) | **Post** /v2/spaces/{id}/team-sets | Create a space-owned team set
[**CreateSpacesByIdValidateAccessCode**](SpacesAPI.md#CreateSpacesByIdValidateAccessCode) | **Post** /v2/spaces/{id}/validate-access-code | Validate a space access code
[**DeleteSpacesById**](SpacesAPI.md#DeleteSpacesById) | **Delete** /v2/spaces/{id} | Delete Space
[**GetIdentitiesByPiidSpaces**](SpacesAPI.md#GetIdentitiesByPiidSpaces) | **Get** /v2/identities/{piid}/spaces | List spaces for a platform identity (cursor-paginated)
[**GetSpaces**](SpacesAPI.md#GetSpaces) | **Get** /v2/spaces | List the caller&#39;s spaces (cursor-paginated)
[**GetSpacesById**](SpacesAPI.md#GetSpacesById) | **Get** /v2/spaces/{id} | Get a space by ID
[**GetSpacesByIdTeamSets**](SpacesAPI.md#GetSpacesByIdTeamSets) | **Get** /v2/spaces/{id}/team-sets | List team sets owned by a space
[**GetSpacesByRegion**](SpacesAPI.md#GetSpacesByRegion) | **Get** /v2/spaces/by-region | Get Spaces by Region (v2)
[**GetSpacesSearch**](SpacesAPI.md#GetSpacesSearch) | **Get** /v2/spaces/search | List public spaces the caller is not subscribed to (cursor-paginated)
[**GetSpacesSearchKeywordByKeyword**](SpacesAPI.md#GetSpacesSearchKeywordByKeyword) | **Get** /v2/spaces/searchKeyword/{keyword} | Search spaces by keyword (cursor-paginated)
[**GetTeamSetsSystem**](SpacesAPI.md#GetTeamSetsSystem) | **Get** /v2/team-sets/system | List the calling tenant&#39;s system team sets
[**PatchSpacesByIdAccessCode**](SpacesAPI.md#PatchSpacesByIdAccessCode) | **Patch** /v2/spaces/{id}/access-code | Update Space Access Code
[**UpdateSpacesById**](SpacesAPI.md#UpdateSpacesById) | **Put** /v2/spaces/{id} | Update Space
[**UpdateSpacesByIdTeamSetsByTeamSetId**](SpacesAPI.md#UpdateSpacesByIdTeamSetsByTeamSetId) | **Put** /v2/spaces/{id}/team-sets/{teamSetId} | Replace a space team set&#39;s lineup



## CreateSpaces

> SpacesSpaceResponse CreateSpaces(ctx).SpacehttpCreateSpaceRequest(spacehttpCreateSpaceRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Create a space

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
	spacehttpCreateSpaceRequest := *openapiclient.NewSpacehttpCreateSpaceRequest("Name_example", int32(123)) // SpacehttpCreateSpaceRequest | Space payload
	idempotencyKey := "idempotencyKey_example" // string | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateSpaces(context.Background()).SpacehttpCreateSpaceRequest(spacehttpCreateSpaceRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpaces`: SpacesSpaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateSpaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spacehttpCreateSpaceRequest** | [**SpacehttpCreateSpaceRequest**](SpacehttpCreateSpaceRequest.md) | Space payload | 
 **idempotencyKey** | **string** | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceResponse**](SpacesSpaceResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpacesByIdReport

> DeleteContestsById200Response CreateSpacesByIdReport(ctx, id).XActingAs(xActingAs).Execute()

Report a space



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateSpacesByIdReport(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateSpacesByIdReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpacesByIdReport`: DeleteContestsById200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateSpacesByIdReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpacesByIdReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DeleteContestsById200Response**](DeleteContestsById200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpacesByIdTeamSets

> TeamsTeamSetResponse CreateSpacesByIdTeamSets(ctx, id).TeamhttpCreateTeamSetRequest(teamhttpCreateTeamSetRequest).XActingAs(xActingAs).Execute()

Create a space-owned team set

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
	teamhttpCreateTeamSetRequest := *openapiclient.NewTeamhttpCreateTeamSetRequest("Name_example", []openapiclient.TeamhttpCreateTeamRequest{*openapiclient.NewTeamhttpCreateTeamRequest("Name_example")}) // TeamhttpCreateTeamSetRequest | Team set payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateSpacesByIdTeamSets(context.Background(), id).TeamhttpCreateTeamSetRequest(teamhttpCreateTeamSetRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateSpacesByIdTeamSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpacesByIdTeamSets`: TeamsTeamSetResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateSpacesByIdTeamSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpacesByIdTeamSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamhttpCreateTeamSetRequest** | [**TeamhttpCreateTeamSetRequest**](TeamhttpCreateTeamSetRequest.md) | Team set payload | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**TeamsTeamSetResponse**](TeamsTeamSetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpacesByIdValidateAccessCode

> HttpxStatusResponseResponse CreateSpacesByIdValidateAccessCode(ctx, id).SpacehttpValidateAccessCodeRequest(spacehttpValidateAccessCodeRequest).XActingAs(xActingAs).Execute()

Validate a space access code

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
	spacehttpValidateAccessCodeRequest := *openapiclient.NewSpacehttpValidateAccessCodeRequest("AccessCode_example") // SpacehttpValidateAccessCodeRequest | Access code body
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.CreateSpacesByIdValidateAccessCode(context.Background(), id).SpacehttpValidateAccessCodeRequest(spacehttpValidateAccessCodeRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.CreateSpacesByIdValidateAccessCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpacesByIdValidateAccessCode`: HttpxStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.CreateSpacesByIdValidateAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpacesByIdValidateAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spacehttpValidateAccessCodeRequest** | [**SpacehttpValidateAccessCodeRequest**](SpacehttpValidateAccessCodeRequest.md) | Access code body | 
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


## DeleteSpacesById

> SpacesSpaceResponse DeleteSpacesById(ctx, id).XActingAs(xActingAs).Execute()

Delete Space



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.DeleteSpacesById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.DeleteSpacesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSpacesById`: SpacesSpaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.DeleteSpacesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpacesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceResponse**](SpacesSpaceResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidSpaces

> SpacesSpaceListResponse GetIdentitiesByPiidSpaces(ctx, piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List spaces for a platform identity (cursor-paginated)

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetIdentitiesByPiidSpaces(context.Background(), piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetIdentitiesByPiidSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidSpaces`: SpacesSpaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetIdentitiesByPiidSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
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


## GetSpaces

> SpacesSpaceListResponse GetSpaces(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List the caller's spaces (cursor-paginated)

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 20, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpaces(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaces`: SpacesSpaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 20, max 200) | 
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


## GetSpacesById

> SpacesSpaceResponse GetSpacesById(ctx, id).XActingAs(xActingAs).Execute()

Get a space by ID

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpacesById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpacesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesById`: SpacesSpaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpacesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceResponse**](SpacesSpaceResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacesByIdTeamSets

> TeamsTeamSetListResponse GetSpacesByIdTeamSets(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List team sets owned by a space



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page, name-sorted) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page, name-sorted) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpacesByIdTeamSets(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpacesByIdTeamSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesByIdTeamSets`: TeamsTeamSetListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpacesByIdTeamSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesByIdTeamSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page, name-sorted) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page, name-sorted) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**TeamsTeamSetListResponse**](TeamsTeamSetListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacesByRegion

> SpacesSpaceListResponse GetSpacesByRegion(ctx).Region(region).Cursor(cursor).Limit(limit).Sort(sort).XActingAs(xActingAs).Execute()

Get Spaces by Region (v2)



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
	region := "region_example" // string | Region Code
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (optional)
	sort := "sort_example" // string | Space sort order (optional) (default to "newest")
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpacesByRegion(context.Background()).Region(region).Cursor(cursor).Limit(limit).Sort(sort).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpacesByRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesByRegion`: SpacesSpaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpacesByRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Region Code | 
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **sort** | **string** | Space sort order | [default to &quot;newest&quot;]
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


## GetSpacesSearch

> SpacesSpaceListResponse GetSpacesSearch(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List public spaces the caller is not subscribed to (cursor-paginated)

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpacesSearch(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpacesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesSearch`: SpacesSpaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpacesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
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


## GetSpacesSearchKeywordByKeyword

> SpacesSpaceListResponse GetSpacesSearchKeywordByKeyword(ctx, keyword).Cursor(cursor).Limit(limit).Sort(sort).XActingAs(xActingAs).Execute()

Search spaces by keyword (cursor-paginated)

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
	sort := "sort_example" // string | Space sort order (optional) (default to "newest")
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetSpacesSearchKeywordByKeyword(context.Background(), keyword).Cursor(cursor).Limit(limit).Sort(sort).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpacesSearchKeywordByKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacesSearchKeywordByKeyword`: SpacesSpaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpacesSearchKeywordByKeyword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string** | Search keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesSearchKeywordByKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **sort** | **string** | Space sort order | [default to &quot;newest&quot;]
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


## GetTeamSetsSystem

> TeamsTeamSetListResponse GetTeamSetsSystem(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List the calling tenant's system team sets



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page, name-sorted) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page, name-sorted) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetTeamSetsSystem(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetTeamSetsSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamSetsSystem`: TeamsTeamSetListResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetTeamSetsSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamSetsSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page, name-sorted) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page, name-sorted) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**TeamsTeamSetListResponse**](TeamsTeamSetListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSpacesByIdAccessCode

> SpacesSpaceResponse PatchSpacesByIdAccessCode(ctx, id).SpacehttpUpdateSpaceAccessCodeRequest(spacehttpUpdateSpaceAccessCodeRequest).XActingAs(xActingAs).Execute()

Update Space Access Code



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
	spacehttpUpdateSpaceAccessCodeRequest := *openapiclient.NewSpacehttpUpdateSpaceAccessCodeRequest("AccessCode_example") // SpacehttpUpdateSpaceAccessCodeRequest | request
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.PatchSpacesByIdAccessCode(context.Background(), id).SpacehttpUpdateSpaceAccessCodeRequest(spacehttpUpdateSpaceAccessCodeRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.PatchSpacesByIdAccessCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSpacesByIdAccessCode`: SpacesSpaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.PatchSpacesByIdAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSpacesByIdAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spacehttpUpdateSpaceAccessCodeRequest** | [**SpacehttpUpdateSpaceAccessCodeRequest**](SpacehttpUpdateSpaceAccessCodeRequest.md) | request | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceResponse**](SpacesSpaceResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpacesById

> SpacesSpaceResponse UpdateSpacesById(ctx, id).SpacehttpUpdateSpaceRequest(spacehttpUpdateSpaceRequest).XActingAs(xActingAs).Execute()

Update Space



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
	spacehttpUpdateSpaceRequest := *openapiclient.NewSpacehttpUpdateSpaceRequest("Name_example", int32(123)) // SpacehttpUpdateSpaceRequest | request
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateSpacesById(context.Background(), id).SpacehttpUpdateSpaceRequest(spacehttpUpdateSpaceRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateSpacesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpacesById`: SpacesSpaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateSpacesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpacesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spacehttpUpdateSpaceRequest** | [**SpacehttpUpdateSpaceRequest**](SpacehttpUpdateSpaceRequest.md) | request | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**SpacesSpaceResponse**](SpacesSpaceResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpacesByIdTeamSetsByTeamSetId

> TeamsTeamSetResponse UpdateSpacesByIdTeamSetsByTeamSetId(ctx, id, teamSetId).TeamhttpCreateTeamSetRequest(teamhttpCreateTeamSetRequest).XActingAs(xActingAs).Execute()

Replace a space team set's lineup



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
	teamSetId := "teamSetId_example" // string | Team Set ID
	teamhttpCreateTeamSetRequest := *openapiclient.NewTeamhttpCreateTeamSetRequest("Name_example", []openapiclient.TeamhttpCreateTeamRequest{*openapiclient.NewTeamhttpCreateTeamRequest("Name_example")}) // TeamhttpCreateTeamSetRequest | Replacement team set payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.UpdateSpacesByIdTeamSetsByTeamSetId(context.Background(), id, teamSetId).TeamhttpCreateTeamSetRequest(teamhttpCreateTeamSetRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.UpdateSpacesByIdTeamSetsByTeamSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpacesByIdTeamSetsByTeamSetId`: TeamsTeamSetResponse
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.UpdateSpacesByIdTeamSetsByTeamSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Space ID | 
**teamSetId** | **string** | Team Set ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpacesByIdTeamSetsByTeamSetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamhttpCreateTeamSetRequest** | [**TeamhttpCreateTeamSetRequest**](TeamhttpCreateTeamSetRequest.md) | Replacement team set payload | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**TeamsTeamSetResponse**](TeamsTeamSetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

