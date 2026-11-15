# GamesDraftView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boards** | Pointer to [**[]WordchainBoardContent**](WordchainBoardContent.md) |  | [optional] 
**Brief** | Pointer to **string** |  | [optional] 
**Difficulty** | Pointer to **string** |  | [optional] 
**Finale** | Pointer to [**[]FanssayBoardContent**](FanssayBoardContent.md) |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Tiebreak** | Pointer to [**FansfeudBoardContent**](FansfeudBoardContent.md) |  | [optional] 
**Type** | **string** |  | 
**Rush** | Pointer to [**[]FansfeudBoardContent**](FansfeudBoardContent.md) |  | [optional] 

## Methods

### NewGamesDraftView

`func NewGamesDraftView(type_ string, ) *GamesDraftView`

NewGamesDraftView instantiates a new GamesDraftView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamesDraftViewWithDefaults

`func NewGamesDraftViewWithDefaults() *GamesDraftView`

NewGamesDraftViewWithDefaults instantiates a new GamesDraftView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoards

`func (o *GamesDraftView) GetBoards() []WordchainBoardContent`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *GamesDraftView) GetBoardsOk() (*[]WordchainBoardContent, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *GamesDraftView) SetBoards(v []WordchainBoardContent)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *GamesDraftView) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### GetBrief

`func (o *GamesDraftView) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *GamesDraftView) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *GamesDraftView) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *GamesDraftView) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetDifficulty

`func (o *GamesDraftView) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GamesDraftView) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GamesDraftView) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *GamesDraftView) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetFinale

`func (o *GamesDraftView) GetFinale() []FanssayBoardContent`

GetFinale returns the Finale field if non-nil, zero value otherwise.

### GetFinaleOk

`func (o *GamesDraftView) GetFinaleOk() (*[]FanssayBoardContent, bool)`

GetFinaleOk returns a tuple with the Finale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinale

`func (o *GamesDraftView) SetFinale(v []FanssayBoardContent)`

SetFinale sets Finale field to given value.

### HasFinale

`func (o *GamesDraftView) HasFinale() bool`

HasFinale returns a boolean if a field has been set.

### GetGameId

`func (o *GamesDraftView) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *GamesDraftView) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *GamesDraftView) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *GamesDraftView) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetId

`func (o *GamesDraftView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GamesDraftView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GamesDraftView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GamesDraftView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *GamesDraftView) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GamesDraftView) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GamesDraftView) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GamesDraftView) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRevision

`func (o *GamesDraftView) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GamesDraftView) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GamesDraftView) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GamesDraftView) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSource

`func (o *GamesDraftView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GamesDraftView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GamesDraftView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GamesDraftView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTiebreak

`func (o *GamesDraftView) GetTiebreak() FansfeudBoardContent`

GetTiebreak returns the Tiebreak field if non-nil, zero value otherwise.

### GetTiebreakOk

`func (o *GamesDraftView) GetTiebreakOk() (*FansfeudBoardContent, bool)`

GetTiebreakOk returns a tuple with the Tiebreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreak

`func (o *GamesDraftView) SetTiebreak(v FansfeudBoardContent)`

SetTiebreak sets Tiebreak field to given value.

### HasTiebreak

`func (o *GamesDraftView) HasTiebreak() bool`

HasTiebreak returns a boolean if a field has been set.

### GetType

`func (o *GamesDraftView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GamesDraftView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GamesDraftView) SetType(v string)`

SetType sets Type field to given value.


### GetRush

`func (o *GamesDraftView) GetRush() []FansfeudBoardContent`

GetRush returns the Rush field if non-nil, zero value otherwise.

### GetRushOk

`func (o *GamesDraftView) GetRushOk() (*[]FansfeudBoardContent, bool)`

GetRushOk returns a tuple with the Rush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRush

`func (o *GamesDraftView) SetRush(v []FansfeudBoardContent)`

SetRush sets Rush field to given value.

### HasRush

`func (o *GamesDraftView) HasRush() bool`

HasRush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


