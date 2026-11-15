# FansfeudActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ActionVersion** | Pointer to **int32** |  | [optional] 
**BoardId** | Pointer to **string** |  | [optional] 
**FinalistIdentityIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFansfeudActionRequest

`func NewFansfeudActionRequest(action string, ) *FansfeudActionRequest`

NewFansfeudActionRequest instantiates a new FansfeudActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFansfeudActionRequestWithDefaults

`func NewFansfeudActionRequestWithDefaults() *FansfeudActionRequest`

NewFansfeudActionRequestWithDefaults instantiates a new FansfeudActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FansfeudActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FansfeudActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FansfeudActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionVersion

`func (o *FansfeudActionRequest) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *FansfeudActionRequest) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *FansfeudActionRequest) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *FansfeudActionRequest) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetBoardId

`func (o *FansfeudActionRequest) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *FansfeudActionRequest) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *FansfeudActionRequest) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *FansfeudActionRequest) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetFinalistIdentityIds

`func (o *FansfeudActionRequest) GetFinalistIdentityIds() []string`

GetFinalistIdentityIds returns the FinalistIdentityIds field if non-nil, zero value otherwise.

### GetFinalistIdentityIdsOk

`func (o *FansfeudActionRequest) GetFinalistIdentityIdsOk() (*[]string, bool)`

GetFinalistIdentityIdsOk returns a tuple with the FinalistIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistIdentityIds

`func (o *FansfeudActionRequest) SetFinalistIdentityIds(v []string)`

SetFinalistIdentityIds sets FinalistIdentityIds field to given value.

### HasFinalistIdentityIds

`func (o *FansfeudActionRequest) HasFinalistIdentityIds() bool`

HasFinalistIdentityIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


