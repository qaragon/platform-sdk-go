# GetContestsHighlights200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]HighlightsHighlight**](HighlightsHighlight.md) |  | [optional] 
**Meta** | Pointer to [**HttpxMetaInfo**](HttpxMetaInfo.md) |  | [optional] 
**Pagination** | Pointer to [**HttpxPaginationMeta**](HttpxPaginationMeta.md) | For v2 cursor-based pagination | [optional] 

## Methods

### NewGetContestsHighlights200ResponseData

`func NewGetContestsHighlights200ResponseData() *GetContestsHighlights200ResponseData`

NewGetContestsHighlights200ResponseData instantiates a new GetContestsHighlights200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContestsHighlights200ResponseDataWithDefaults

`func NewGetContestsHighlights200ResponseDataWithDefaults() *GetContestsHighlights200ResponseData`

NewGetContestsHighlights200ResponseDataWithDefaults instantiates a new GetContestsHighlights200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContestsHighlights200ResponseData) GetData() []HighlightsHighlight`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContestsHighlights200ResponseData) GetDataOk() (*[]HighlightsHighlight, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContestsHighlights200ResponseData) SetData(v []HighlightsHighlight)`

SetData sets Data field to given value.

### HasData

`func (o *GetContestsHighlights200ResponseData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetContestsHighlights200ResponseData) GetMeta() HttpxMetaInfo`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetContestsHighlights200ResponseData) GetMetaOk() (*HttpxMetaInfo, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetContestsHighlights200ResponseData) SetMeta(v HttpxMetaInfo)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetContestsHighlights200ResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPagination

`func (o *GetContestsHighlights200ResponseData) GetPagination() HttpxPaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetContestsHighlights200ResponseData) GetPaginationOk() (*HttpxPaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetContestsHighlights200ResponseData) SetPagination(v HttpxPaginationMeta)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetContestsHighlights200ResponseData) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


