# WordchainCreateGame

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

### NewWordchainCreateGame

`func NewWordchainCreateGame(type_ string, ) *WordchainCreateGame`

NewWordchainCreateGame instantiates a new WordchainCreateGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordchainCreateGameWithDefaults

`func NewWordchainCreateGameWithDefaults() *WordchainCreateGame`

NewWordchainCreateGameWithDefaults instantiates a new WordchainCreateGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftId

`func (o *WordchainCreateGame) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *WordchainCreateGame) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *WordchainCreateGame) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *WordchainCreateGame) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetDraftRevision

`func (o *WordchainCreateGame) GetDraftRevision() int32`

GetDraftRevision returns the DraftRevision field if non-nil, zero value otherwise.

### GetDraftRevisionOk

`func (o *WordchainCreateGame) GetDraftRevisionOk() (*int32, bool)`

GetDraftRevisionOk returns a tuple with the DraftRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftRevision

`func (o *WordchainCreateGame) SetDraftRevision(v int32)`

SetDraftRevision sets DraftRevision field to given value.

### HasDraftRevision

`func (o *WordchainCreateGame) HasDraftRevision() bool`

HasDraftRevision returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *WordchainCreateGame) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *WordchainCreateGame) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *WordchainCreateGame) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *WordchainCreateGame) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *WordchainCreateGame) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *WordchainCreateGame) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *WordchainCreateGame) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *WordchainCreateGame) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetReward

`func (o *WordchainCreateGame) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *WordchainCreateGame) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *WordchainCreateGame) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *WordchainCreateGame) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetTeams

`func (o *WordchainCreateGame) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *WordchainCreateGame) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *WordchainCreateGame) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *WordchainCreateGame) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetType

`func (o *WordchainCreateGame) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WordchainCreateGame) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WordchainCreateGame) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


