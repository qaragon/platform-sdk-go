# ContesthttpCreatePlatformContestHostingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanParticipate** | Pointer to **bool** |  | [optional] 
**CanView** | Pointer to **bool** |  | [optional] 
**SpaceId** | **string** | venue_id is not a request field (core#177): the space determines the venue (spaces.venue_id is canonical), so the caller supplies only the space and the venue is derived from it. | 

## Methods

### NewContesthttpCreatePlatformContestHostingRequest

`func NewContesthttpCreatePlatformContestHostingRequest(spaceId string, ) *ContesthttpCreatePlatformContestHostingRequest`

NewContesthttpCreatePlatformContestHostingRequest instantiates a new ContesthttpCreatePlatformContestHostingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContesthttpCreatePlatformContestHostingRequestWithDefaults

`func NewContesthttpCreatePlatformContestHostingRequestWithDefaults() *ContesthttpCreatePlatformContestHostingRequest`

NewContesthttpCreatePlatformContestHostingRequestWithDefaults instantiates a new ContesthttpCreatePlatformContestHostingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanParticipate

`func (o *ContesthttpCreatePlatformContestHostingRequest) GetCanParticipate() bool`

GetCanParticipate returns the CanParticipate field if non-nil, zero value otherwise.

### GetCanParticipateOk

`func (o *ContesthttpCreatePlatformContestHostingRequest) GetCanParticipateOk() (*bool, bool)`

GetCanParticipateOk returns a tuple with the CanParticipate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanParticipate

`func (o *ContesthttpCreatePlatformContestHostingRequest) SetCanParticipate(v bool)`

SetCanParticipate sets CanParticipate field to given value.

### HasCanParticipate

`func (o *ContesthttpCreatePlatformContestHostingRequest) HasCanParticipate() bool`

HasCanParticipate returns a boolean if a field has been set.

### GetCanView

`func (o *ContesthttpCreatePlatformContestHostingRequest) GetCanView() bool`

GetCanView returns the CanView field if non-nil, zero value otherwise.

### GetCanViewOk

`func (o *ContesthttpCreatePlatformContestHostingRequest) GetCanViewOk() (*bool, bool)`

GetCanViewOk returns a tuple with the CanView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanView

`func (o *ContesthttpCreatePlatformContestHostingRequest) SetCanView(v bool)`

SetCanView sets CanView field to given value.

### HasCanView

`func (o *ContesthttpCreatePlatformContestHostingRequest) HasCanView() bool`

HasCanView returns a boolean if a field has been set.

### GetSpaceId

`func (o *ContesthttpCreatePlatformContestHostingRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ContesthttpCreatePlatformContestHostingRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ContesthttpCreatePlatformContestHostingRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


