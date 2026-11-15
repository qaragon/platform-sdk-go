# HttpxStandardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Meta** | Pointer to [**HttpxMetaInfo**](HttpxMetaInfo.md) |  | [optional] 
**Pagination** | Pointer to [**HttpxPaginationMeta**](HttpxPaginationMeta.md) | For v2 cursor-based pagination | [optional] 

## Methods

### NewHttpxStandardResponse

`func NewHttpxStandardResponse() *HttpxStandardResponse`

NewHttpxStandardResponse instantiates a new HttpxStandardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpxStandardResponseWithDefaults

`func NewHttpxStandardResponseWithDefaults() *HttpxStandardResponse`

NewHttpxStandardResponseWithDefaults instantiates a new HttpxStandardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HttpxStandardResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HttpxStandardResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HttpxStandardResponse) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *HttpxStandardResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *HttpxStandardResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *HttpxStandardResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *HttpxStandardResponse) GetMeta() HttpxMetaInfo`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpxStandardResponse) GetMetaOk() (*HttpxMetaInfo, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpxStandardResponse) SetMeta(v HttpxMetaInfo)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpxStandardResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPagination

`func (o *HttpxStandardResponse) GetPagination() HttpxPaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *HttpxStandardResponse) GetPaginationOk() (*HttpxPaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *HttpxStandardResponse) SetPagination(v HttpxPaginationMeta)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *HttpxStandardResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


