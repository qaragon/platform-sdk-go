# FansfeudGameView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentSource** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**PlayDefinitionRef**](PlayDefinitionRef.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**FansfeudGameResult**](FansfeudGameResult.md) |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Rush** | Pointer to [**FeudrushView**](FeudrushView.md) |  | [optional] 
**Showdown** | Pointer to [**FeudshowdownView**](FeudshowdownView.md) |  | [optional] 
**Stage** | Pointer to [**PlayStage**](PlayStage.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**ViewerRole** | Pointer to **string** |  | [optional] 

## Methods

### NewFansfeudGameView

`func NewFansfeudGameView(type_ string, ) *FansfeudGameView`

NewFansfeudGameView instantiates a new FansfeudGameView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFansfeudGameViewWithDefaults

`func NewFansfeudGameViewWithDefaults() *FansfeudGameView`

NewFansfeudGameViewWithDefaults instantiates a new FansfeudGameView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentSource

`func (o *FansfeudGameView) GetContentSource() string`

GetContentSource returns the ContentSource field if non-nil, zero value otherwise.

### GetContentSourceOk

`func (o *FansfeudGameView) GetContentSourceOk() (*string, bool)`

GetContentSourceOk returns a tuple with the ContentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSource

`func (o *FansfeudGameView) SetContentSource(v string)`

SetContentSource sets ContentSource field to given value.

### HasContentSource

`func (o *FansfeudGameView) HasContentSource() bool`

HasContentSource returns a boolean if a field has been set.

### GetDefinition

`func (o *FansfeudGameView) GetDefinition() PlayDefinitionRef`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *FansfeudGameView) GetDefinitionOk() (*PlayDefinitionRef, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *FansfeudGameView) SetDefinition(v PlayDefinitionRef)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *FansfeudGameView) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetId

`func (o *FansfeudGameView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FansfeudGameView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FansfeudGameView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FansfeudGameView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *FansfeudGameView) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *FansfeudGameView) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *FansfeudGameView) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *FansfeudGameView) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *FansfeudGameView) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *FansfeudGameView) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *FansfeudGameView) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *FansfeudGameView) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetResult

`func (o *FansfeudGameView) GetResult() FansfeudGameResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FansfeudGameView) GetResultOk() (*FansfeudGameResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FansfeudGameView) SetResult(v FansfeudGameResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *FansfeudGameView) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReward

`func (o *FansfeudGameView) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FansfeudGameView) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FansfeudGameView) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FansfeudGameView) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetRush

`func (o *FansfeudGameView) GetRush() FeudrushView`

GetRush returns the Rush field if non-nil, zero value otherwise.

### GetRushOk

`func (o *FansfeudGameView) GetRushOk() (*FeudrushView, bool)`

GetRushOk returns a tuple with the Rush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRush

`func (o *FansfeudGameView) SetRush(v FeudrushView)`

SetRush sets Rush field to given value.

### HasRush

`func (o *FansfeudGameView) HasRush() bool`

HasRush returns a boolean if a field has been set.

### GetShowdown

`func (o *FansfeudGameView) GetShowdown() FeudshowdownView`

GetShowdown returns the Showdown field if non-nil, zero value otherwise.

### GetShowdownOk

`func (o *FansfeudGameView) GetShowdownOk() (*FeudshowdownView, bool)`

GetShowdownOk returns a tuple with the Showdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowdown

`func (o *FansfeudGameView) SetShowdown(v FeudshowdownView)`

SetShowdown sets Showdown field to given value.

### HasShowdown

`func (o *FansfeudGameView) HasShowdown() bool`

HasShowdown returns a boolean if a field has been set.

### GetStage

`func (o *FansfeudGameView) GetStage() PlayStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FansfeudGameView) GetStageOk() (*PlayStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FansfeudGameView) SetStage(v PlayStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *FansfeudGameView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *FansfeudGameView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FansfeudGameView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FansfeudGameView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FansfeudGameView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *FansfeudGameView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FansfeudGameView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FansfeudGameView) SetType(v string)`

SetType sets Type field to given value.


### GetViewerRole

`func (o *FansfeudGameView) GetViewerRole() string`

GetViewerRole returns the ViewerRole field if non-nil, zero value otherwise.

### GetViewerRoleOk

`func (o *FansfeudGameView) GetViewerRoleOk() (*string, bool)`

GetViewerRoleOk returns a tuple with the ViewerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRole

`func (o *FansfeudGameView) SetViewerRole(v string)`

SetViewerRole sets ViewerRole field to given value.

### HasViewerRole

`func (o *FansfeudGameView) HasViewerRole() bool`

HasViewerRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


