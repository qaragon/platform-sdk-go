# ContesthttpUpdateContestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**GroupingType** | Pointer to **string** |  | [optional] 
**NoOfGroups** | **int32** |  | 
**Rewards** | Pointer to [**[]DomainContestReward**](DomainContestReward.md) |  | [optional] 
**Timing** | [**DomainTiming**](DomainTiming.md) |  | 
**Title** | **string** |  | 
**WinnersPerGroup** | **int32** |  | 

## Methods

### NewContesthttpUpdateContestRequest

`func NewContesthttpUpdateContestRequest(description string, noOfGroups int32, timing DomainTiming, title string, winnersPerGroup int32, ) *ContesthttpUpdateContestRequest`

NewContesthttpUpdateContestRequest instantiates a new ContesthttpUpdateContestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContesthttpUpdateContestRequestWithDefaults

`func NewContesthttpUpdateContestRequestWithDefaults() *ContesthttpUpdateContestRequest`

NewContesthttpUpdateContestRequestWithDefaults instantiates a new ContesthttpUpdateContestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ContesthttpUpdateContestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContesthttpUpdateContestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContesthttpUpdateContestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroupingType

`func (o *ContesthttpUpdateContestRequest) GetGroupingType() string`

GetGroupingType returns the GroupingType field if non-nil, zero value otherwise.

### GetGroupingTypeOk

`func (o *ContesthttpUpdateContestRequest) GetGroupingTypeOk() (*string, bool)`

GetGroupingTypeOk returns a tuple with the GroupingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingType

`func (o *ContesthttpUpdateContestRequest) SetGroupingType(v string)`

SetGroupingType sets GroupingType field to given value.

### HasGroupingType

`func (o *ContesthttpUpdateContestRequest) HasGroupingType() bool`

HasGroupingType returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *ContesthttpUpdateContestRequest) GetNoOfGroups() int32`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *ContesthttpUpdateContestRequest) GetNoOfGroupsOk() (*int32, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *ContesthttpUpdateContestRequest) SetNoOfGroups(v int32)`

SetNoOfGroups sets NoOfGroups field to given value.


### GetRewards

`func (o *ContesthttpUpdateContestRequest) GetRewards() []DomainContestReward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *ContesthttpUpdateContestRequest) GetRewardsOk() (*[]DomainContestReward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *ContesthttpUpdateContestRequest) SetRewards(v []DomainContestReward)`

SetRewards sets Rewards field to given value.

### HasRewards

`func (o *ContesthttpUpdateContestRequest) HasRewards() bool`

HasRewards returns a boolean if a field has been set.

### GetTiming

`func (o *ContesthttpUpdateContestRequest) GetTiming() DomainTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *ContesthttpUpdateContestRequest) GetTimingOk() (*DomainTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *ContesthttpUpdateContestRequest) SetTiming(v DomainTiming)`

SetTiming sets Timing field to given value.


### GetTitle

`func (o *ContesthttpUpdateContestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContesthttpUpdateContestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContesthttpUpdateContestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWinnersPerGroup

`func (o *ContesthttpUpdateContestRequest) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *ContesthttpUpdateContestRequest) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *ContesthttpUpdateContestRequest) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


