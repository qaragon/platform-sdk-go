# FanssayCreateGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftId** | Pointer to **string** |  | [optional] 
**DraftRevision** | Pointer to **int32** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFanssayCreateGame

`func NewFanssayCreateGame(type_ string, ) *FanssayCreateGame`

NewFanssayCreateGame instantiates a new FanssayCreateGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanssayCreateGameWithDefaults

`func NewFanssayCreateGameWithDefaults() *FanssayCreateGame`

NewFanssayCreateGameWithDefaults instantiates a new FanssayCreateGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftId

`func (o *FanssayCreateGame) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *FanssayCreateGame) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *FanssayCreateGame) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *FanssayCreateGame) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetDraftRevision

`func (o *FanssayCreateGame) GetDraftRevision() int32`

GetDraftRevision returns the DraftRevision field if non-nil, zero value otherwise.

### GetDraftRevisionOk

`func (o *FanssayCreateGame) GetDraftRevisionOk() (*int32, bool)`

GetDraftRevisionOk returns a tuple with the DraftRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftRevision

`func (o *FanssayCreateGame) SetDraftRevision(v int32)`

SetDraftRevision sets DraftRevision field to given value.

### HasDraftRevision

`func (o *FanssayCreateGame) HasDraftRevision() bool`

HasDraftRevision returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *FanssayCreateGame) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *FanssayCreateGame) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *FanssayCreateGame) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *FanssayCreateGame) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *FanssayCreateGame) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *FanssayCreateGame) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *FanssayCreateGame) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *FanssayCreateGame) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetReward

`func (o *FanssayCreateGame) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FanssayCreateGame) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FanssayCreateGame) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FanssayCreateGame) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetTeams

`func (o *FanssayCreateGame) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *FanssayCreateGame) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *FanssayCreateGame) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *FanssayCreateGame) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetType

`func (o *FanssayCreateGame) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FanssayCreateGame) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FanssayCreateGame) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


