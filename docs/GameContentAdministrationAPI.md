# \GameContentAdministrationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdminGamesFansFeudLibraryGenerate**](GameContentAdministrationAPI.md#CreateAdminGamesFansFeudLibraryGenerate) | **Post** /v2/admin/games/fans-feud/library/generate | Generate fresh Fans Feud library content
[**CreateAdminGamesFansSayLibraryGenerate**](GameContentAdministrationAPI.md#CreateAdminGamesFansSayLibraryGenerate) | **Post** /v2/admin/games/fans-say/library/generate | Generate fresh Fans Say library content
[**CreateAdminGamesWordChainLibraryGenerate**](GameContentAdministrationAPI.md#CreateAdminGamesWordChainLibraryGenerate) | **Post** /v2/admin/games/word-chain/library/generate | Generate fresh Word Chain library content
[**GetAdminGamesFansFeudLibraryGenerationsById**](GameContentAdministrationAPI.md#GetAdminGamesFansFeudLibraryGenerationsById) | **Get** /v2/admin/games/fans-feud/library/generations/{id} | Read a Fans Feud library generation status
[**GetAdminGamesFansSayLibraryGenerationsById**](GameContentAdministrationAPI.md#GetAdminGamesFansSayLibraryGenerationsById) | **Get** /v2/admin/games/fans-say/library/generations/{id} | Read a Fans Say library generation status
[**GetAdminGamesWordChainLibraryGenerationsById**](GameContentAdministrationAPI.md#GetAdminGamesWordChainLibraryGenerationsById) | **Get** /v2/admin/games/word-chain/library/generations/{id} | Read a Word Chain library generation status



## CreateAdminGamesFansFeudLibraryGenerate

> FansfeudLibraryGenerationResponse CreateAdminGamesFansFeudLibraryGenerate(ctx).FansfeudLibraryGenerationRequest(fansfeudLibraryGenerationRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Generate fresh Fans Feud library content



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
	fansfeudLibraryGenerationRequest := *openapiclient.NewFansfeudLibraryGenerationRequest() // FansfeudLibraryGenerationRequest | Content brief and requested set count
	idempotencyKey := "idempotencyKey_example" // string | Replay an identical request instead of enqueuing another paid batch (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameContentAdministrationAPI.CreateAdminGamesFansFeudLibraryGenerate(context.Background()).FansfeudLibraryGenerationRequest(fansfeudLibraryGenerationRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameContentAdministrationAPI.CreateAdminGamesFansFeudLibraryGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdminGamesFansFeudLibraryGenerate`: FansfeudLibraryGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `GameContentAdministrationAPI.CreateAdminGamesFansFeudLibraryGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminGamesFansFeudLibraryGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fansfeudLibraryGenerationRequest** | [**FansfeudLibraryGenerationRequest**](FansfeudLibraryGenerationRequest.md) | Content brief and requested set count | 
 **idempotencyKey** | **string** | Replay an identical request instead of enqueuing another paid batch | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**FansfeudLibraryGenerationResponse**](FansfeudLibraryGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdminGamesFansSayLibraryGenerate

> FanssayLibraryGenerationResponse CreateAdminGamesFansSayLibraryGenerate(ctx).FanssayLibraryGenerationRequest(fanssayLibraryGenerationRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Generate fresh Fans Say library content



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
	fanssayLibraryGenerationRequest := *openapiclient.NewFanssayLibraryGenerationRequest() // FanssayLibraryGenerationRequest | Content brief and requested set count
	idempotencyKey := "idempotencyKey_example" // string | Replay an identical request instead of enqueuing another paid batch (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameContentAdministrationAPI.CreateAdminGamesFansSayLibraryGenerate(context.Background()).FanssayLibraryGenerationRequest(fanssayLibraryGenerationRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameContentAdministrationAPI.CreateAdminGamesFansSayLibraryGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdminGamesFansSayLibraryGenerate`: FanssayLibraryGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `GameContentAdministrationAPI.CreateAdminGamesFansSayLibraryGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminGamesFansSayLibraryGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fanssayLibraryGenerationRequest** | [**FanssayLibraryGenerationRequest**](FanssayLibraryGenerationRequest.md) | Content brief and requested set count | 
 **idempotencyKey** | **string** | Replay an identical request instead of enqueuing another paid batch | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**FanssayLibraryGenerationResponse**](FanssayLibraryGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdminGamesWordChainLibraryGenerate

> WordchainLibraryGenerationResponse CreateAdminGamesWordChainLibraryGenerate(ctx).WordchainLibraryGenerationRequest(wordchainLibraryGenerationRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Generate fresh Word Chain library content



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
	wordchainLibraryGenerationRequest := *openapiclient.NewWordchainLibraryGenerationRequest() // WordchainLibraryGenerationRequest | Content brief and requested set count
	idempotencyKey := "idempotencyKey_example" // string | Replay an identical request instead of enqueuing another paid batch (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameContentAdministrationAPI.CreateAdminGamesWordChainLibraryGenerate(context.Background()).WordchainLibraryGenerationRequest(wordchainLibraryGenerationRequest).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameContentAdministrationAPI.CreateAdminGamesWordChainLibraryGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdminGamesWordChainLibraryGenerate`: WordchainLibraryGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `GameContentAdministrationAPI.CreateAdminGamesWordChainLibraryGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminGamesWordChainLibraryGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wordchainLibraryGenerationRequest** | [**WordchainLibraryGenerationRequest**](WordchainLibraryGenerationRequest.md) | Content brief and requested set count | 
 **idempotencyKey** | **string** | Replay an identical request instead of enqueuing another paid batch | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**WordchainLibraryGenerationResponse**](WordchainLibraryGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminGamesFansFeudLibraryGenerationsById

> FansfeudLibraryGenerationResponse GetAdminGamesFansFeudLibraryGenerationsById(ctx, id).XActingAs(xActingAs).Execute()

Read a Fans Feud library generation status



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
	id := "id_example" // string | Generation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameContentAdministrationAPI.GetAdminGamesFansFeudLibraryGenerationsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameContentAdministrationAPI.GetAdminGamesFansFeudLibraryGenerationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminGamesFansFeudLibraryGenerationsById`: FansfeudLibraryGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `GameContentAdministrationAPI.GetAdminGamesFansFeudLibraryGenerationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Generation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminGamesFansFeudLibraryGenerationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**FansfeudLibraryGenerationResponse**](FansfeudLibraryGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminGamesFansSayLibraryGenerationsById

> FanssayLibraryGenerationResponse GetAdminGamesFansSayLibraryGenerationsById(ctx, id).XActingAs(xActingAs).Execute()

Read a Fans Say library generation status



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
	id := "id_example" // string | Generation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameContentAdministrationAPI.GetAdminGamesFansSayLibraryGenerationsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameContentAdministrationAPI.GetAdminGamesFansSayLibraryGenerationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminGamesFansSayLibraryGenerationsById`: FanssayLibraryGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `GameContentAdministrationAPI.GetAdminGamesFansSayLibraryGenerationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Generation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminGamesFansSayLibraryGenerationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**FanssayLibraryGenerationResponse**](FanssayLibraryGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminGamesWordChainLibraryGenerationsById

> WordchainLibraryGenerationResponse GetAdminGamesWordChainLibraryGenerationsById(ctx, id).XActingAs(xActingAs).Execute()

Read a Word Chain library generation status



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
	id := "id_example" // string | Generation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameContentAdministrationAPI.GetAdminGamesWordChainLibraryGenerationsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameContentAdministrationAPI.GetAdminGamesWordChainLibraryGenerationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminGamesWordChainLibraryGenerationsById`: WordchainLibraryGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `GameContentAdministrationAPI.GetAdminGamesWordChainLibraryGenerationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Generation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminGamesWordChainLibraryGenerationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**WordchainLibraryGenerationResponse**](WordchainLibraryGenerationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

