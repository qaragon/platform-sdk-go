# WordchainGameView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clash** | Pointer to [**LinkclashView**](LinkclashView.md) |  | [optional] 
**ContentSource** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**PlayDefinitionRef**](PlayDefinitionRef.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**WordchainGameResult**](WordchainGameResult.md) |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Rush** | Pointer to [**LinkrushView**](LinkrushView.md) |  | [optional] 
**Stage** | Pointer to [**PlayStage**](PlayStage.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**ViewerRole** | Pointer to **string** |  | [optional] 

## Methods

### NewWordchainGameView

`func NewWordchainGameView(type_ string, ) *WordchainGameView`

NewWordchainGameView instantiates a new WordchainGameView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordchainGameViewWithDefaults

`func NewWordchainGameViewWithDefaults() *WordchainGameView`

NewWordchainGameViewWithDefaults instantiates a new WordchainGameView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClash

`func (o *WordchainGameView) GetClash() LinkclashView`

GetClash returns the Clash field if non-nil, zero value otherwise.

### GetClashOk

`func (o *WordchainGameView) GetClashOk() (*LinkclashView, bool)`

GetClashOk returns a tuple with the Clash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClash

`func (o *WordchainGameView) SetClash(v LinkclashView)`

SetClash sets Clash field to given value.

### HasClash

`func (o *WordchainGameView) HasClash() bool`

HasClash returns a boolean if a field has been set.

### GetContentSource

`func (o *WordchainGameView) GetContentSource() string`

GetContentSource returns the ContentSource field if non-nil, zero value otherwise.

### GetContentSourceOk

`func (o *WordchainGameView) GetContentSourceOk() (*string, bool)`

GetContentSourceOk returns a tuple with the ContentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSource

`func (o *WordchainGameView) SetContentSource(v string)`

SetContentSource sets ContentSource field to given value.

### HasContentSource

`func (o *WordchainGameView) HasContentSource() bool`

HasContentSource returns a boolean if a field has been set.

### GetDefinition

`func (o *WordchainGameView) GetDefinition() PlayDefinitionRef`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *WordchainGameView) GetDefinitionOk() (*PlayDefinitionRef, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *WordchainGameView) SetDefinition(v PlayDefinitionRef)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *WordchainGameView) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetId

`func (o *WordchainGameView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WordchainGameView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WordchainGameView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WordchainGameView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *WordchainGameView) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *WordchainGameView) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *WordchainGameView) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *WordchainGameView) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *WordchainGameView) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *WordchainGameView) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *WordchainGameView) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *WordchainGameView) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetResult

`func (o *WordchainGameView) GetResult() WordchainGameResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WordchainGameView) GetResultOk() (*WordchainGameResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WordchainGameView) SetResult(v WordchainGameResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *WordchainGameView) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReward

`func (o *WordchainGameView) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *WordchainGameView) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *WordchainGameView) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *WordchainGameView) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetRush

`func (o *WordchainGameView) GetRush() LinkrushView`

GetRush returns the Rush field if non-nil, zero value otherwise.

### GetRushOk

`func (o *WordchainGameView) GetRushOk() (*LinkrushView, bool)`

GetRushOk returns a tuple with the Rush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRush

`func (o *WordchainGameView) SetRush(v LinkrushView)`

SetRush sets Rush field to given value.

### HasRush

`func (o *WordchainGameView) HasRush() bool`

HasRush returns a boolean if a field has been set.

### GetStage

`func (o *WordchainGameView) GetStage() PlayStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *WordchainGameView) GetStageOk() (*PlayStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *WordchainGameView) SetStage(v PlayStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *WordchainGameView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *WordchainGameView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WordchainGameView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WordchainGameView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WordchainGameView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *WordchainGameView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WordchainGameView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WordchainGameView) SetType(v string)`

SetType sets Type field to given value.


### GetViewerRole

`func (o *WordchainGameView) GetViewerRole() string`

GetViewerRole returns the ViewerRole field if non-nil, zero value otherwise.

### GetViewerRoleOk

`func (o *WordchainGameView) GetViewerRoleOk() (*string, bool)`

GetViewerRoleOk returns a tuple with the ViewerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRole

`func (o *WordchainGameView) SetViewerRole(v string)`

SetViewerRole sets ViewerRole field to given value.

### HasViewerRole

`func (o *WordchainGameView) HasViewerRole() bool`

HasViewerRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


