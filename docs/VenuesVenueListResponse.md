# VenuesVenueListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VenuesVenue**](VenuesVenue.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewVenuesVenueListResponse

`func NewVenuesVenueListResponse(data []VenuesVenue, ) *VenuesVenueListResponse`

NewVenuesVenueListResponse instantiates a new VenuesVenueListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVenuesVenueListResponseWithDefaults

`func NewVenuesVenueListResponseWithDefaults() *VenuesVenueListResponse`

NewVenuesVenueListResponseWithDefaults instantiates a new VenuesVenueListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VenuesVenueListResponse) GetData() []VenuesVenue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VenuesVenueListResponse) GetDataOk() (*[]VenuesVenue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VenuesVenueListResponse) SetData(v []VenuesVenue)`

SetData sets Data field to given value.


### GetPagination

`func (o *VenuesVenueListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VenuesVenueListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VenuesVenueListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *VenuesVenueListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


