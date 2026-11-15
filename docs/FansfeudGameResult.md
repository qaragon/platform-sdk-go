# FansfeudGameResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**WinnerPlatformIdentityIds** | Pointer to **[]string** |  | [optional] 
**WinnerTeam** | Pointer to **int32** |  | [optional] 

## Methods

### NewFansfeudGameResult

`func NewFansfeudGameResult() *FansfeudGameResult`

NewFansfeudGameResult instantiates a new FansfeudGameResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFansfeudGameResultWithDefaults

`func NewFansfeudGameResultWithDefaults() *FansfeudGameResult`

NewFansfeudGameResultWithDefaults instantiates a new FansfeudGameResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *FansfeudGameResult) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *FansfeudGameResult) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *FansfeudGameResult) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *FansfeudGameResult) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetOutcome

`func (o *FansfeudGameResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *FansfeudGameResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *FansfeudGameResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *FansfeudGameResult) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetReward

`func (o *FansfeudGameResult) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FansfeudGameResult) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FansfeudGameResult) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FansfeudGameResult) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetWinnerPlatformIdentityIds

`func (o *FansfeudGameResult) GetWinnerPlatformIdentityIds() []string`

GetWinnerPlatformIdentityIds returns the WinnerPlatformIdentityIds field if non-nil, zero value otherwise.

### GetWinnerPlatformIdentityIdsOk

`func (o *FansfeudGameResult) GetWinnerPlatformIdentityIdsOk() (*[]string, bool)`

GetWinnerPlatformIdentityIdsOk returns a tuple with the WinnerPlatformIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerPlatformIdentityIds

`func (o *FansfeudGameResult) SetWinnerPlatformIdentityIds(v []string)`

SetWinnerPlatformIdentityIds sets WinnerPlatformIdentityIds field to given value.

### HasWinnerPlatformIdentityIds

`func (o *FansfeudGameResult) HasWinnerPlatformIdentityIds() bool`

HasWinnerPlatformIdentityIds returns a boolean if a field has been set.

### GetWinnerTeam

`func (o *FansfeudGameResult) GetWinnerTeam() int32`

GetWinnerTeam returns the WinnerTeam field if non-nil, zero value otherwise.

### GetWinnerTeamOk

`func (o *FansfeudGameResult) GetWinnerTeamOk() (*int32, bool)`

GetWinnerTeamOk returns a tuple with the WinnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerTeam

`func (o *FansfeudGameResult) SetWinnerTeam(v int32)`

SetWinnerTeam sets WinnerTeam field to given value.

### HasWinnerTeam

`func (o *FansfeudGameResult) HasWinnerTeam() bool`

HasWinnerTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


