# SpacesSpaceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SpacesSpace**](SpacesSpace.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewSpacesSpaceListResponse

`func NewSpacesSpaceListResponse(data []SpacesSpace, ) *SpacesSpaceListResponse`

NewSpacesSpaceListResponse instantiates a new SpacesSpaceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacesSpaceListResponseWithDefaults

`func NewSpacesSpaceListResponseWithDefaults() *SpacesSpaceListResponse`

NewSpacesSpaceListResponseWithDefaults instantiates a new SpacesSpaceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SpacesSpaceListResponse) GetData() []SpacesSpace`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SpacesSpaceListResponse) GetDataOk() (*[]SpacesSpace, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SpacesSpaceListResponse) SetData(v []SpacesSpace)`

SetData sets Data field to given value.


### GetPagination

`func (o *SpacesSpaceListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SpacesSpaceListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SpacesSpaceListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SpacesSpaceListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


