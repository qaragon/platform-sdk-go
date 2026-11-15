# TeamsTeamSetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TeamsTeamSet**](TeamsTeamSet.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewTeamsTeamSetListResponse

`func NewTeamsTeamSetListResponse(data []TeamsTeamSet, ) *TeamsTeamSetListResponse`

NewTeamsTeamSetListResponse instantiates a new TeamsTeamSetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsTeamSetListResponseWithDefaults

`func NewTeamsTeamSetListResponseWithDefaults() *TeamsTeamSetListResponse`

NewTeamsTeamSetListResponseWithDefaults instantiates a new TeamsTeamSetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TeamsTeamSetListResponse) GetData() []TeamsTeamSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TeamsTeamSetListResponse) GetDataOk() (*[]TeamsTeamSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TeamsTeamSetListResponse) SetData(v []TeamsTeamSet)`

SetData sets Data field to given value.


### GetPagination

`func (o *TeamsTeamSetListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TeamsTeamSetListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TeamsTeamSetListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TeamsTeamSetListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


