# \AutomatedContestSchedulerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestSchedulers**](AutomatedContestSchedulerAPI.md#CreateContestSchedulers) | **Post** /v2/contest-schedulers | Configure the Automated Contest Scheduler for a Space
[**GetContestSchedulers**](AutomatedContestSchedulerAPI.md#GetContestSchedulers) | **Get** /v2/contest-schedulers | List the caller&#39;s scheduler settings
[**GetContestSchedulersContests**](AutomatedContestSchedulerAPI.md#GetContestSchedulersContests) | **Get** /v2/contest-schedulers/contests | Read today&#39;s Contests in tenant-configured Daily Spaces
[**UpdateContestSchedulersById**](AutomatedContestSchedulerAPI.md#UpdateContestSchedulersById) | **Put** /v2/contest-schedulers/{id} | Update or disable the caller&#39;s scheduler settings



## CreateContestSchedulers

> RecurringcontestsScheduleResponse CreateContestSchedulers(ctx).RecurringcontestsSettings(recurringcontestsSettings).XActingAs(xActingAs).Execute()

Configure the Automated Contest Scheduler for a Space



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
	recurringcontestsSettings := *openapiclient.NewRecurringcontestsSettings([]openapiclient.RecurringcontestsContestType{*openapiclient.NewRecurringcontestsContestType("SourceType_example")}, "CreationTime_example", "SpaceId_example", "StartTime_example", "TimeZone_example", "Title_example") // RecurringcontestsSettings | Tenant scheduler settings; weekdays are Sunday=0 through Saturday=6
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedContestSchedulerAPI.CreateContestSchedulers(context.Background()).RecurringcontestsSettings(recurringcontestsSettings).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedContestSchedulerAPI.CreateContestSchedulers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestSchedulers`: RecurringcontestsScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomatedContestSchedulerAPI.CreateContestSchedulers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestSchedulersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recurringcontestsSettings** | [**RecurringcontestsSettings**](RecurringcontestsSettings.md) | Tenant scheduler settings; weekdays are Sunday&#x3D;0 through Saturday&#x3D;6 | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**RecurringcontestsScheduleResponse**](RecurringcontestsScheduleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestSchedulers

> RecurringcontestsScheduleListResponse GetContestSchedulers(ctx).XActingAs(xActingAs).Execute()

List the caller's scheduler settings

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedContestSchedulerAPI.GetContestSchedulers(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedContestSchedulerAPI.GetContestSchedulers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestSchedulers`: RecurringcontestsScheduleListResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomatedContestSchedulerAPI.GetContestSchedulers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContestSchedulersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**RecurringcontestsScheduleListResponse**](RecurringcontestsScheduleListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestSchedulersContests

> DomainContestListResponse GetContestSchedulersContests(ctx).Region(region).XActingAs(xActingAs).Execute()

Read today's Contests in tenant-configured Daily Spaces



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
	region := "region_example" // string | ISO country code; omit for worldwide only (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedContestSchedulerAPI.GetContestSchedulersContests(context.Background()).Region(region).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedContestSchedulerAPI.GetContestSchedulersContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestSchedulersContests`: DomainContestListResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomatedContestSchedulerAPI.GetContestSchedulersContests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContestSchedulersContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | ISO country code; omit for worldwide only | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestListResponse**](DomainContestListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestSchedulersById

> RecurringcontestsScheduleResponse UpdateContestSchedulersById(ctx, id).RecurringcontestsSettings(recurringcontestsSettings).XActingAs(xActingAs).Execute()

Update or disable the caller's scheduler settings

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
	id := "id_example" // string | Scheduler settings ID
	recurringcontestsSettings := *openapiclient.NewRecurringcontestsSettings([]openapiclient.RecurringcontestsContestType{*openapiclient.NewRecurringcontestsContestType("SourceType_example")}, "CreationTime_example", "SpaceId_example", "StartTime_example", "TimeZone_example", "Title_example") // RecurringcontestsSettings | Replacement settings; existing Contests remain unchanged
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedContestSchedulerAPI.UpdateContestSchedulersById(context.Background(), id).RecurringcontestsSettings(recurringcontestsSettings).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedContestSchedulerAPI.UpdateContestSchedulersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestSchedulersById`: RecurringcontestsScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomatedContestSchedulerAPI.UpdateContestSchedulersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Scheduler settings ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContestSchedulersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recurringcontestsSettings** | [**RecurringcontestsSettings**](RecurringcontestsSettings.md) | Replacement settings; existing Contests remain unchanged | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**RecurringcontestsScheduleResponse**](RecurringcontestsScheduleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

