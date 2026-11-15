# IdentitiesPlatformIdentityListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]IdentitiesPlatformIdentity**](IdentitiesPlatformIdentity.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewIdentitiesPlatformIdentityListResponse

`func NewIdentitiesPlatformIdentityListResponse(data []IdentitiesPlatformIdentity, ) *IdentitiesPlatformIdentityListResponse`

NewIdentitiesPlatformIdentityListResponse instantiates a new IdentitiesPlatformIdentityListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitiesPlatformIdentityListResponseWithDefaults

`func NewIdentitiesPlatformIdentityListResponseWithDefaults() *IdentitiesPlatformIdentityListResponse`

NewIdentitiesPlatformIdentityListResponseWithDefaults instantiates a new IdentitiesPlatformIdentityListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IdentitiesPlatformIdentityListResponse) GetData() []IdentitiesPlatformIdentity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IdentitiesPlatformIdentityListResponse) GetDataOk() (*[]IdentitiesPlatformIdentity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IdentitiesPlatformIdentityListResponse) SetData(v []IdentitiesPlatformIdentity)`

SetData sets Data field to given value.


### GetPagination

`func (o *IdentitiesPlatformIdentityListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IdentitiesPlatformIdentityListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IdentitiesPlatformIdentityListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IdentitiesPlatformIdentityListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


