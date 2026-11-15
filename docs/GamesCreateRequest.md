# GamesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceSpaceId** | Pointer to **string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**DraftId** | Pointer to **string** |  | [optional] 
**DraftRevision** | Pointer to **int32** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 
**FinalistIdentityIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGamesCreateRequest

`func NewGamesCreateRequest(type_ string, ) *GamesCreateRequest`

NewGamesCreateRequest instantiates a new GamesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamesCreateRequestWithDefaults

`func NewGamesCreateRequestWithDefaults() *GamesCreateRequest`

NewGamesCreateRequestWithDefaults instantiates a new GamesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceSpaceId

`func (o *GamesCreateRequest) GetAudienceSpaceId() string`

GetAudienceSpaceId returns the AudienceSpaceId field if non-nil, zero value otherwise.

### GetAudienceSpaceIdOk

`func (o *GamesCreateRequest) GetAudienceSpaceIdOk() (*string, bool)`

GetAudienceSpaceIdOk returns a tuple with the AudienceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceSpaceId

`func (o *GamesCreateRequest) SetAudienceSpaceId(v string)`

SetAudienceSpaceId sets AudienceSpaceId field to given value.

### HasAudienceSpaceId

`func (o *GamesCreateRequest) HasAudienceSpaceId() bool`

HasAudienceSpaceId returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *GamesCreateRequest) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *GamesCreateRequest) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *GamesCreateRequest) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *GamesCreateRequest) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetType

`func (o *GamesCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GamesCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GamesCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDraftId

`func (o *GamesCreateRequest) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *GamesCreateRequest) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *GamesCreateRequest) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *GamesCreateRequest) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetDraftRevision

`func (o *GamesCreateRequest) GetDraftRevision() int32`

GetDraftRevision returns the DraftRevision field if non-nil, zero value otherwise.

### GetDraftRevisionOk

`func (o *GamesCreateRequest) GetDraftRevisionOk() (*int32, bool)`

GetDraftRevisionOk returns a tuple with the DraftRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftRevision

`func (o *GamesCreateRequest) SetDraftRevision(v int32)`

SetDraftRevision sets DraftRevision field to given value.

### HasDraftRevision

`func (o *GamesCreateRequest) HasDraftRevision() bool`

HasDraftRevision returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *GamesCreateRequest) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *GamesCreateRequest) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *GamesCreateRequest) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *GamesCreateRequest) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetReward

`func (o *GamesCreateRequest) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *GamesCreateRequest) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *GamesCreateRequest) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *GamesCreateRequest) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetTeams

`func (o *GamesCreateRequest) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GamesCreateRequest) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GamesCreateRequest) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *GamesCreateRequest) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetFinalistIdentityIds

`func (o *GamesCreateRequest) GetFinalistIdentityIds() []string`

GetFinalistIdentityIds returns the FinalistIdentityIds field if non-nil, zero value otherwise.

### GetFinalistIdentityIdsOk

`func (o *GamesCreateRequest) GetFinalistIdentityIdsOk() (*[]string, bool)`

GetFinalistIdentityIdsOk returns a tuple with the FinalistIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistIdentityIds

`func (o *GamesCreateRequest) SetFinalistIdentityIds(v []string)`

SetFinalistIdentityIds sets FinalistIdentityIds field to given value.

### HasFinalistIdentityIds

`func (o *GamesCreateRequest) HasFinalistIdentityIds() bool`

HasFinalistIdentityIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


