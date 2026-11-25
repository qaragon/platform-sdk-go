# DomainBonusReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to **[]int32** |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 

## Methods

### NewDomainBonusReward

`func NewDomainBonusReward() *DomainBonusReward`

NewDomainBonusReward instantiates a new DomainBonusReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBonusRewardWithDefaults

`func NewDomainBonusRewardWithDefaults() *DomainBonusReward`

NewDomainBonusRewardWithDefaults instantiates a new DomainBonusReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainBonusReward) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainBonusReward) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainBonusReward) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainBonusReward) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DomainBonusReward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainBonusReward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainBonusReward) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainBonusReward) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPositions

`func (o *DomainBonusReward) GetPositions() []int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *DomainBonusReward) GetPositionsOk() (*[]int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *DomainBonusReward) SetPositions(v []int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *DomainBonusReward) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetReward

`func (o *DomainBonusReward) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *DomainBonusReward) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *DomainBonusReward) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *DomainBonusReward) HasReward() bool`

HasReward returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


