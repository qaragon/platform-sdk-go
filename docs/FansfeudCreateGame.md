# FansfeudCreateGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftId** | Pointer to **string** |  | [optional] 
**DraftRevision** | Pointer to **int32** |  | [optional] 
**FinalistIdentityIds** | Pointer to **[]string** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFansfeudCreateGame

`func NewFansfeudCreateGame(type_ string, ) *FansfeudCreateGame`

NewFansfeudCreateGame instantiates a new FansfeudCreateGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFansfeudCreateGameWithDefaults

`func NewFansfeudCreateGameWithDefaults() *FansfeudCreateGame`

NewFansfeudCreateGameWithDefaults instantiates a new FansfeudCreateGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftId

`func (o *FansfeudCreateGame) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *FansfeudCreateGame) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *FansfeudCreateGame) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *FansfeudCreateGame) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetDraftRevision

`func (o *FansfeudCreateGame) GetDraftRevision() int32`

GetDraftRevision returns the DraftRevision field if non-nil, zero value otherwise.

### GetDraftRevisionOk

`func (o *FansfeudCreateGame) GetDraftRevisionOk() (*int32, bool)`

GetDraftRevisionOk returns a tuple with the DraftRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftRevision

`func (o *FansfeudCreateGame) SetDraftRevision(v int32)`

SetDraftRevision sets DraftRevision field to given value.

### HasDraftRevision

`func (o *FansfeudCreateGame) HasDraftRevision() bool`

HasDraftRevision returns a boolean if a field has been set.

### GetFinalistIdentityIds

`func (o *FansfeudCreateGame) GetFinalistIdentityIds() []string`

GetFinalistIdentityIds returns the FinalistIdentityIds field if non-nil, zero value otherwise.

### GetFinalistIdentityIdsOk

`func (o *FansfeudCreateGame) GetFinalistIdentityIdsOk() (*[]string, bool)`

GetFinalistIdentityIdsOk returns a tuple with the FinalistIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistIdentityIds

`func (o *FansfeudCreateGame) SetFinalistIdentityIds(v []string)`

SetFinalistIdentityIds sets FinalistIdentityIds field to given value.

### HasFinalistIdentityIds

`func (o *FansfeudCreateGame) HasFinalistIdentityIds() bool`

HasFinalistIdentityIds returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *FansfeudCreateGame) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *FansfeudCreateGame) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *FansfeudCreateGame) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *FansfeudCreateGame) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *FansfeudCreateGame) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *FansfeudCreateGame) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *FansfeudCreateGame) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *FansfeudCreateGame) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetReward

`func (o *FansfeudCreateGame) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FansfeudCreateGame) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FansfeudCreateGame) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FansfeudCreateGame) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetTeams

`func (o *FansfeudCreateGame) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *FansfeudCreateGame) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *FansfeudCreateGame) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *FansfeudCreateGame) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetType

`func (o *FansfeudCreateGame) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FansfeudCreateGame) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FansfeudCreateGame) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


