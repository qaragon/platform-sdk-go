# UpdateContestsByIdHighlight200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**HighlightsHighlight**](HighlightsHighlight.md) |  | [optional] 
**Meta** | Pointer to [**HttpxMetaInfo**](HttpxMetaInfo.md) |  | [optional] 
**Pagination** | Pointer to [**HttpxPaginationMeta**](HttpxPaginationMeta.md) | For v2 cursor-based pagination | [optional] 

## Methods

### NewUpdateContestsByIdHighlight200ResponseData

`func NewUpdateContestsByIdHighlight200ResponseData() *UpdateContestsByIdHighlight200ResponseData`

NewUpdateContestsByIdHighlight200ResponseData instantiates a new UpdateContestsByIdHighlight200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContestsByIdHighlight200ResponseDataWithDefaults

`func NewUpdateContestsByIdHighlight200ResponseDataWithDefaults() *UpdateContestsByIdHighlight200ResponseData`

NewUpdateContestsByIdHighlight200ResponseDataWithDefaults instantiates a new UpdateContestsByIdHighlight200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateContestsByIdHighlight200ResponseData) GetData() HighlightsHighlight`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateContestsByIdHighlight200ResponseData) GetDataOk() (*HighlightsHighlight, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateContestsByIdHighlight200ResponseData) SetData(v HighlightsHighlight)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateContestsByIdHighlight200ResponseData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateContestsByIdHighlight200ResponseData) GetMeta() HttpxMetaInfo`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateContestsByIdHighlight200ResponseData) GetMetaOk() (*HttpxMetaInfo, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateContestsByIdHighlight200ResponseData) SetMeta(v HttpxMetaInfo)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateContestsByIdHighlight200ResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPagination

`func (o *UpdateContestsByIdHighlight200ResponseData) GetPagination() HttpxPaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UpdateContestsByIdHighlight200ResponseData) GetPaginationOk() (*HttpxPaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UpdateContestsByIdHighlight200ResponseData) SetPagination(v HttpxPaginationMeta)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UpdateContestsByIdHighlight200ResponseData) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


