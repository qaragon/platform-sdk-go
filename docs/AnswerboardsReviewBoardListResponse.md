# AnswerboardsReviewBoardListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnswerboardsReviewBoard**](AnswerboardsReviewBoard.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewAnswerboardsReviewBoardListResponse

`func NewAnswerboardsReviewBoardListResponse(data []AnswerboardsReviewBoard, ) *AnswerboardsReviewBoardListResponse`

NewAnswerboardsReviewBoardListResponse instantiates a new AnswerboardsReviewBoardListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerboardsReviewBoardListResponseWithDefaults

`func NewAnswerboardsReviewBoardListResponseWithDefaults() *AnswerboardsReviewBoardListResponse`

NewAnswerboardsReviewBoardListResponseWithDefaults instantiates a new AnswerboardsReviewBoardListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnswerboardsReviewBoardListResponse) GetData() []AnswerboardsReviewBoard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnswerboardsReviewBoardListResponse) GetDataOk() (*[]AnswerboardsReviewBoard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnswerboardsReviewBoardListResponse) SetData(v []AnswerboardsReviewBoard)`

SetData sets Data field to given value.


### GetPagination

`func (o *AnswerboardsReviewBoardListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnswerboardsReviewBoardListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnswerboardsReviewBoardListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AnswerboardsReviewBoardListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


