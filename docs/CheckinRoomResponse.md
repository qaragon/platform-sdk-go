# CheckinRoomResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **bool** |  | [optional] 
**ConfirmedAt** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]CheckinRoomMemberResponse**](CheckinRoomMemberResponse.md) |  | [optional] 
**RoomNumber** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewCheckinRoomResponse

`func NewCheckinRoomResponse() *CheckinRoomResponse`

NewCheckinRoomResponse instantiates a new CheckinRoomResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckinRoomResponseWithDefaults

`func NewCheckinRoomResponseWithDefaults() *CheckinRoomResponse`

NewCheckinRoomResponseWithDefaults instantiates a new CheckinRoomResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *CheckinRoomResponse) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *CheckinRoomResponse) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *CheckinRoomResponse) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *CheckinRoomResponse) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *CheckinRoomResponse) GetConfirmedAt() string`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *CheckinRoomResponse) GetConfirmedAtOk() (*string, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *CheckinRoomResponse) SetConfirmedAt(v string)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *CheckinRoomResponse) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetMembers

`func (o *CheckinRoomResponse) GetMembers() []CheckinRoomMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CheckinRoomResponse) GetMembersOk() (*[]CheckinRoomMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CheckinRoomResponse) SetMembers(v []CheckinRoomMemberResponse)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CheckinRoomResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetRoomNumber

`func (o *CheckinRoomResponse) GetRoomNumber() int32`

GetRoomNumber returns the RoomNumber field if non-nil, zero value otherwise.

### GetRoomNumberOk

`func (o *CheckinRoomResponse) GetRoomNumberOk() (*int32, bool)`

GetRoomNumberOk returns a tuple with the RoomNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomNumber

`func (o *CheckinRoomResponse) SetRoomNumber(v int32)`

SetRoomNumber sets RoomNumber field to given value.

### HasRoomNumber

`func (o *CheckinRoomResponse) HasRoomNumber() bool`

HasRoomNumber returns a boolean if a field has been set.

### GetSize

`func (o *CheckinRoomResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CheckinRoomResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CheckinRoomResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *CheckinRoomResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


