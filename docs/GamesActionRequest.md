# GamesActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**BoardToAudience** | Pointer to **bool** |  | [optional] 
**BoardToPlayers** | Pointer to **bool** |  | [optional] 
**ActionVersion** | Pointer to **int32** |  | [optional] 
**BoardId** | **string** |  | 
**SubmitterIdentityId** | Pointer to **string** |  | [optional] 
**Team** | Pointer to **int32** |  | [optional] 
**FinalistIdentityIds** | Pointer to **[]string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 

## Methods

### NewGamesActionRequest

`func NewGamesActionRequest(action string, boardId string, ) *GamesActionRequest`

NewGamesActionRequest instantiates a new GamesActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamesActionRequestWithDefaults

`func NewGamesActionRequestWithDefaults() *GamesActionRequest`

NewGamesActionRequestWithDefaults instantiates a new GamesActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GamesActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GamesActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GamesActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetBoardToAudience

`func (o *GamesActionRequest) GetBoardToAudience() bool`

GetBoardToAudience returns the BoardToAudience field if non-nil, zero value otherwise.

### GetBoardToAudienceOk

`func (o *GamesActionRequest) GetBoardToAudienceOk() (*bool, bool)`

GetBoardToAudienceOk returns a tuple with the BoardToAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardToAudience

`func (o *GamesActionRequest) SetBoardToAudience(v bool)`

SetBoardToAudience sets BoardToAudience field to given value.

### HasBoardToAudience

`func (o *GamesActionRequest) HasBoardToAudience() bool`

HasBoardToAudience returns a boolean if a field has been set.

### GetBoardToPlayers

`func (o *GamesActionRequest) GetBoardToPlayers() bool`

GetBoardToPlayers returns the BoardToPlayers field if non-nil, zero value otherwise.

### GetBoardToPlayersOk

`func (o *GamesActionRequest) GetBoardToPlayersOk() (*bool, bool)`

GetBoardToPlayersOk returns a tuple with the BoardToPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardToPlayers

`func (o *GamesActionRequest) SetBoardToPlayers(v bool)`

SetBoardToPlayers sets BoardToPlayers field to given value.

### HasBoardToPlayers

`func (o *GamesActionRequest) HasBoardToPlayers() bool`

HasBoardToPlayers returns a boolean if a field has been set.

### GetActionVersion

`func (o *GamesActionRequest) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *GamesActionRequest) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *GamesActionRequest) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *GamesActionRequest) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetBoardId

`func (o *GamesActionRequest) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *GamesActionRequest) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *GamesActionRequest) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.


### GetSubmitterIdentityId

`func (o *GamesActionRequest) GetSubmitterIdentityId() string`

GetSubmitterIdentityId returns the SubmitterIdentityId field if non-nil, zero value otherwise.

### GetSubmitterIdentityIdOk

`func (o *GamesActionRequest) GetSubmitterIdentityIdOk() (*string, bool)`

GetSubmitterIdentityIdOk returns a tuple with the SubmitterIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterIdentityId

`func (o *GamesActionRequest) SetSubmitterIdentityId(v string)`

SetSubmitterIdentityId sets SubmitterIdentityId field to given value.

### HasSubmitterIdentityId

`func (o *GamesActionRequest) HasSubmitterIdentityId() bool`

HasSubmitterIdentityId returns a boolean if a field has been set.

### GetTeam

`func (o *GamesActionRequest) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GamesActionRequest) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GamesActionRequest) SetTeam(v int32)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GamesActionRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFinalistIdentityIds

`func (o *GamesActionRequest) GetFinalistIdentityIds() []string`

GetFinalistIdentityIds returns the FinalistIdentityIds field if non-nil, zero value otherwise.

### GetFinalistIdentityIdsOk

`func (o *GamesActionRequest) GetFinalistIdentityIdsOk() (*[]string, bool)`

GetFinalistIdentityIdsOk returns a tuple with the FinalistIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistIdentityIds

`func (o *GamesActionRequest) SetFinalistIdentityIds(v []string)`

SetFinalistIdentityIds sets FinalistIdentityIds field to given value.

### HasFinalistIdentityIds

`func (o *GamesActionRequest) HasFinalistIdentityIds() bool`

HasFinalistIdentityIds returns a boolean if a field has been set.

### GetChainId

`func (o *GamesActionRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GamesActionRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GamesActionRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GamesActionRequest) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


