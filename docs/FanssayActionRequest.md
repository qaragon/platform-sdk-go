# FanssayActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ActionVersion** | Pointer to **int32** |  | [optional] 
**BoardId** | **string** |  | 
**SubmitterIdentityId** | Pointer to **string** |  | [optional] 
**Team** | Pointer to **int32** |  | [optional] 

## Methods

### NewFanssayActionRequest

`func NewFanssayActionRequest(action string, boardId string, ) *FanssayActionRequest`

NewFanssayActionRequest instantiates a new FanssayActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanssayActionRequestWithDefaults

`func NewFanssayActionRequestWithDefaults() *FanssayActionRequest`

NewFanssayActionRequestWithDefaults instantiates a new FanssayActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FanssayActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FanssayActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FanssayActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionVersion

`func (o *FanssayActionRequest) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *FanssayActionRequest) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *FanssayActionRequest) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *FanssayActionRequest) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetBoardId

`func (o *FanssayActionRequest) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *FanssayActionRequest) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *FanssayActionRequest) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.


### GetSubmitterIdentityId

`func (o *FanssayActionRequest) GetSubmitterIdentityId() string`

GetSubmitterIdentityId returns the SubmitterIdentityId field if non-nil, zero value otherwise.

### GetSubmitterIdentityIdOk

`func (o *FanssayActionRequest) GetSubmitterIdentityIdOk() (*string, bool)`

GetSubmitterIdentityIdOk returns a tuple with the SubmitterIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterIdentityId

`func (o *FanssayActionRequest) SetSubmitterIdentityId(v string)`

SetSubmitterIdentityId sets SubmitterIdentityId field to given value.

### HasSubmitterIdentityId

`func (o *FanssayActionRequest) HasSubmitterIdentityId() bool`

HasSubmitterIdentityId returns a boolean if a field has been set.

### GetTeam

`func (o *FanssayActionRequest) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *FanssayActionRequest) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *FanssayActionRequest) SetTeam(v int32)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *FanssayActionRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


