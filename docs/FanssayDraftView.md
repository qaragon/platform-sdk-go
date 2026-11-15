# FanssayDraftView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boards** | Pointer to [**[]FanssayBoardContent**](FanssayBoardContent.md) |  | [optional] 
**Brief** | Pointer to **string** |  | [optional] 
**Difficulty** | Pointer to **string** |  | [optional] 
**Finale** | Pointer to [**[]FanssayBoardContent**](FanssayBoardContent.md) |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Tiebreak** | Pointer to [**FanssayBoardContent**](FanssayBoardContent.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFanssayDraftView

`func NewFanssayDraftView(type_ string, ) *FanssayDraftView`

NewFanssayDraftView instantiates a new FanssayDraftView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanssayDraftViewWithDefaults

`func NewFanssayDraftViewWithDefaults() *FanssayDraftView`

NewFanssayDraftViewWithDefaults instantiates a new FanssayDraftView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoards

`func (o *FanssayDraftView) GetBoards() []FanssayBoardContent`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *FanssayDraftView) GetBoardsOk() (*[]FanssayBoardContent, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *FanssayDraftView) SetBoards(v []FanssayBoardContent)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *FanssayDraftView) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### GetBrief

`func (o *FanssayDraftView) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *FanssayDraftView) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *FanssayDraftView) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *FanssayDraftView) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetDifficulty

`func (o *FanssayDraftView) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *FanssayDraftView) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *FanssayDraftView) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *FanssayDraftView) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetFinale

`func (o *FanssayDraftView) GetFinale() []FanssayBoardContent`

GetFinale returns the Finale field if non-nil, zero value otherwise.

### GetFinaleOk

`func (o *FanssayDraftView) GetFinaleOk() (*[]FanssayBoardContent, bool)`

GetFinaleOk returns a tuple with the Finale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinale

`func (o *FanssayDraftView) SetFinale(v []FanssayBoardContent)`

SetFinale sets Finale field to given value.

### HasFinale

`func (o *FanssayDraftView) HasFinale() bool`

HasFinale returns a boolean if a field has been set.

### GetGameId

`func (o *FanssayDraftView) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *FanssayDraftView) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *FanssayDraftView) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *FanssayDraftView) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetId

`func (o *FanssayDraftView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FanssayDraftView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FanssayDraftView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FanssayDraftView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *FanssayDraftView) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FanssayDraftView) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FanssayDraftView) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FanssayDraftView) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRevision

`func (o *FanssayDraftView) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FanssayDraftView) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FanssayDraftView) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FanssayDraftView) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSource

`func (o *FanssayDraftView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FanssayDraftView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FanssayDraftView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FanssayDraftView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTiebreak

`func (o *FanssayDraftView) GetTiebreak() FanssayBoardContent`

GetTiebreak returns the Tiebreak field if non-nil, zero value otherwise.

### GetTiebreakOk

`func (o *FanssayDraftView) GetTiebreakOk() (*FanssayBoardContent, bool)`

GetTiebreakOk returns a tuple with the Tiebreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreak

`func (o *FanssayDraftView) SetTiebreak(v FanssayBoardContent)`

SetTiebreak sets Tiebreak field to given value.

### HasTiebreak

`func (o *FanssayDraftView) HasTiebreak() bool`

HasTiebreak returns a boolean if a field has been set.

### GetType

`func (o *FanssayDraftView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FanssayDraftView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FanssayDraftView) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


