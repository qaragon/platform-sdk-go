# DomainContestReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwardedTo** | **string** |  | 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to **[]int32** |  | [optional] 
**RedemptionCondition** | Pointer to **string** |  | [optional] 
**RewardType** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewDomainContestReward

`func NewDomainContestReward(awardedTo string, rewardType string, value string, ) *DomainContestReward`

NewDomainContestReward instantiates a new DomainContestReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContestRewardWithDefaults

`func NewDomainContestRewardWithDefaults() *DomainContestReward`

NewDomainContestRewardWithDefaults instantiates a new DomainContestReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwardedTo

`func (o *DomainContestReward) GetAwardedTo() string`

GetAwardedTo returns the AwardedTo field if non-nil, zero value otherwise.

### GetAwardedToOk

`func (o *DomainContestReward) GetAwardedToOk() (*string, bool)`

GetAwardedToOk returns a tuple with the AwardedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedTo

`func (o *DomainContestReward) SetAwardedTo(v string)`

SetAwardedTo sets AwardedTo field to given value.


### GetCurrency

`func (o *DomainContestReward) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DomainContestReward) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DomainContestReward) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DomainContestReward) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *DomainContestReward) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainContestReward) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainContestReward) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainContestReward) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DomainContestReward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainContestReward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainContestReward) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainContestReward) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPositions

`func (o *DomainContestReward) GetPositions() []int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *DomainContestReward) GetPositionsOk() (*[]int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *DomainContestReward) SetPositions(v []int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *DomainContestReward) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetRedemptionCondition

`func (o *DomainContestReward) GetRedemptionCondition() string`

GetRedemptionCondition returns the RedemptionCondition field if non-nil, zero value otherwise.

### GetRedemptionConditionOk

`func (o *DomainContestReward) GetRedemptionConditionOk() (*string, bool)`

GetRedemptionConditionOk returns a tuple with the RedemptionCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCondition

`func (o *DomainContestReward) SetRedemptionCondition(v string)`

SetRedemptionCondition sets RedemptionCondition field to given value.

### HasRedemptionCondition

`func (o *DomainContestReward) HasRedemptionCondition() bool`

HasRedemptionCondition returns a boolean if a field has been set.

### GetRewardType

`func (o *DomainContestReward) GetRewardType() string`

GetRewardType returns the RewardType field if non-nil, zero value otherwise.

### GetRewardTypeOk

`func (o *DomainContestReward) GetRewardTypeOk() (*string, bool)`

GetRewardTypeOk returns a tuple with the RewardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardType

`func (o *DomainContestReward) SetRewardType(v string)`

SetRewardType sets RewardType field to given value.


### GetValue

`func (o *DomainContestReward) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DomainContestReward) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DomainContestReward) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


