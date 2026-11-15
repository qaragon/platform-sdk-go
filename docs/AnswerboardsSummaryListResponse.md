# AnswerboardsSummaryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnswerboardsSummary**](AnswerboardsSummary.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewAnswerboardsSummaryListResponse

`func NewAnswerboardsSummaryListResponse(data []AnswerboardsSummary, ) *AnswerboardsSummaryListResponse`

NewAnswerboardsSummaryListResponse instantiates a new AnswerboardsSummaryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerboardsSummaryListResponseWithDefaults

`func NewAnswerboardsSummaryListResponseWithDefaults() *AnswerboardsSummaryListResponse`

NewAnswerboardsSummaryListResponseWithDefaults instantiates a new AnswerboardsSummaryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnswerboardsSummaryListResponse) GetData() []AnswerboardsSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnswerboardsSummaryListResponse) GetDataOk() (*[]AnswerboardsSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnswerboardsSummaryListResponse) SetData(v []AnswerboardsSummary)`

SetData sets Data field to given value.


### GetPagination

`func (o *AnswerboardsSummaryListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnswerboardsSummaryListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnswerboardsSummaryListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AnswerboardsSummaryListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


