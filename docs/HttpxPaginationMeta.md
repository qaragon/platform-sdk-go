# HttpxPaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** |  | [optional] 
**HasPrevious** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**NextCursor** | Pointer to **string** | Cursor-based pagination fields for v2 only Note: v1 does not include pagination metadata - it returns arrays directly All fields are always included (null for cursors when not available, false for booleans) | [optional] 
**PrevCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpxPaginationMeta

`func NewHttpxPaginationMeta() *HttpxPaginationMeta`

NewHttpxPaginationMeta instantiates a new HttpxPaginationMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpxPaginationMetaWithDefaults

`func NewHttpxPaginationMetaWithDefaults() *HttpxPaginationMeta`

NewHttpxPaginationMetaWithDefaults instantiates a new HttpxPaginationMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *HttpxPaginationMeta) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *HttpxPaginationMeta) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *HttpxPaginationMeta) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *HttpxPaginationMeta) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrevious

`func (o *HttpxPaginationMeta) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *HttpxPaginationMeta) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *HttpxPaginationMeta) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.

### HasHasPrevious

`func (o *HttpxPaginationMeta) HasHasPrevious() bool`

HasHasPrevious returns a boolean if a field has been set.

### GetLimit

`func (o *HttpxPaginationMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *HttpxPaginationMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *HttpxPaginationMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *HttpxPaginationMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *HttpxPaginationMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *HttpxPaginationMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *HttpxPaginationMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *HttpxPaginationMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetPrevCursor

`func (o *HttpxPaginationMeta) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *HttpxPaginationMeta) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *HttpxPaginationMeta) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.

### HasPrevCursor

`func (o *HttpxPaginationMeta) HasPrevCursor() bool`

HasPrevCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


