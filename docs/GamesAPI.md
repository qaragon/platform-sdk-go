# \GamesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGames**](GamesAPI.md#CreateGames) | **Post** /v2/games | Create a first-party Game
[**CreateGamesByIdActions**](GamesAPI.md#CreateGamesByIdActions) | **Post** /v2/games/{id}/actions | Apply a host action to a Game
[**CreateGamesByIdRounds**](GamesAPI.md#CreateGamesByIdRounds) | **Post** /v2/games/{id}/rounds | Start a Find the Minority round
[**CreateGamesDrafts**](GamesAPI.md#CreateGamesDrafts) | **Post** /v2/games/drafts | Generate a creator-only Game draft
[**CreateGamesDraftsByIdRotate**](GamesAPI.md#CreateGamesDraftsByIdRotate) | **Post** /v2/games/drafts/{id}/rotate | Replace one draft question or the entire draft
[**GetGames**](GamesAPI.md#GetGames) | **Get** /v2/games | List hosted and joined Games
[**GetGamesById**](GamesAPI.md#GetGamesById) | **Get** /v2/games/{id} | Read a Game&#39;s safe current state
[**GetGamesDraftsById**](GamesAPI.md#GetGamesDraftsById) | **Get** /v2/games/drafts/{id} | Read a creator-only Game draft



## CreateGames

> GamesGameViewResponse CreateGames(ctx).GamesCreateRequest(gamesCreateRequest).XActingAs(xActingAs).Execute()

Create a first-party Game



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
	gamesCreateRequest := openapiclient.games.CreateRequest{FansfeudCreateGame: openapiclient.NewFansfeudCreateGame("Type_example")} // GamesCreateRequest | Typed game configuration
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.CreateGames(context.Background()).GamesCreateRequest(gamesCreateRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.CreateGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGames`: GamesGameViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.CreateGames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gamesCreateRequest** | [**GamesCreateRequest**](GamesCreateRequest.md) | Typed game configuration | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesGameViewResponse**](GamesGameViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGamesByIdActions

> GamesGameViewResponse CreateGamesByIdActions(ctx, id).GamesActionRequest(gamesActionRequest).XActingAs(xActingAs).Execute()

Apply a host action to a Game



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
	id := "id_example" // string | Game ID
	gamesActionRequest := *openapiclient.NewGamesActionRequest("Action_example", "BoardId_example") // GamesActionRequest | Game-specific host action
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.CreateGamesByIdActions(context.Background(), id).GamesActionRequest(gamesActionRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.CreateGamesByIdActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGamesByIdActions`: GamesGameViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.CreateGamesByIdActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Game ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGamesByIdActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gamesActionRequest** | [**GamesActionRequest**](GamesActionRequest.md) | Game-specific host action | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesGameViewResponse**](GamesGameViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGamesByIdRounds

> FindtheminorityGameViewResponse CreateGamesByIdRounds(ctx, id).FindtheminorityStartRoundRequest(findtheminorityStartRoundRequest).XActingAs(xActingAs).Execute()

Start a Find the Minority round



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
	id := "id_example" // string | Game ID
	findtheminorityStartRoundRequest := *openapiclient.NewFindtheminorityStartRoundRequest() // FindtheminorityStartRoundRequest | Question and player roster
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.CreateGamesByIdRounds(context.Background(), id).FindtheminorityStartRoundRequest(findtheminorityStartRoundRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.CreateGamesByIdRounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGamesByIdRounds`: FindtheminorityGameViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.CreateGamesByIdRounds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Game ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGamesByIdRoundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findtheminorityStartRoundRequest** | [**FindtheminorityStartRoundRequest**](FindtheminorityStartRoundRequest.md) | Question and player roster | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**FindtheminorityGameViewResponse**](FindtheminorityGameViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGamesDrafts

> GamesDraftViewResponse CreateGamesDrafts(ctx).GamesDraftRequest(gamesDraftRequest).XActingAs(xActingAs).Execute()

Generate a creator-only Game draft



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
	gamesDraftRequest := openapiclient.games.DraftRequest{FansfeudDraftRequest: openapiclient.NewFansfeudDraftRequest("Type_example")} // GamesDraftRequest | Typed creator brief
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.CreateGamesDrafts(context.Background()).GamesDraftRequest(gamesDraftRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.CreateGamesDrafts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGamesDrafts`: GamesDraftViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.CreateGamesDrafts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGamesDraftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gamesDraftRequest** | [**GamesDraftRequest**](GamesDraftRequest.md) | Typed creator brief | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesDraftViewResponse**](GamesDraftViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGamesDraftsByIdRotate

> GamesDraftViewResponse CreateGamesDraftsByIdRotate(ctx, id).GamesRotateDraftRequest(gamesRotateDraftRequest).XActingAs(xActingAs).Execute()

Replace one draft question or the entire draft



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
	id := "id_example" // string | Draft ID
	gamesRotateDraftRequest := *openapiclient.NewGamesRotateDraftRequest() // GamesRotateDraftRequest | Draft revision and replacement selection
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.CreateGamesDraftsByIdRotate(context.Background(), id).GamesRotateDraftRequest(gamesRotateDraftRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.CreateGamesDraftsByIdRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGamesDraftsByIdRotate`: GamesDraftViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.CreateGamesDraftsByIdRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Draft ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGamesDraftsByIdRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gamesRotateDraftRequest** | [**GamesRotateDraftRequest**](GamesRotateDraftRequest.md) | Draft revision and replacement selection | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesDraftViewResponse**](GamesDraftViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGames

> GamesSummaryListResponse GetGames(ctx).Type_(type_).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List hosted and joined Games

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
	type_ := "type__example" // string | Optional game type (optional)
	cursor := "cursor_example" // string | Cursor (optional)
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetGames(context.Background()).Type_(type_).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGames`: GamesSummaryListResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetGames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Optional game type | 
 **cursor** | **string** | Cursor | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesSummaryListResponse**](GamesSummaryListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGamesById

> GamesGameViewResponse GetGamesById(ctx, id).XActingAs(xActingAs).Execute()

Read a Game's safe current state

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
	id := "id_example" // string | Game ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetGamesById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetGamesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGamesById`: GamesGameViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetGamesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Game ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGamesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesGameViewResponse**](GamesGameViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGamesDraftsById

> GamesDraftViewResponse GetGamesDraftsById(ctx, id).XActingAs(xActingAs).Execute()

Read a creator-only Game draft

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
	id := "id_example" // string | Draft ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetGamesDraftsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetGamesDraftsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGamesDraftsById`: GamesDraftViewResponse
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetGamesDraftsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Draft ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGamesDraftsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**GamesDraftViewResponse**](GamesDraftViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

