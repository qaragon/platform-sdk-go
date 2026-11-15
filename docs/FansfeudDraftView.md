# FansfeudDraftView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boards** | Pointer to [**[]FansfeudBoardContent**](FansfeudBoardContent.md) |  | [optional] 
**Brief** | Pointer to **string** |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Rush** | Pointer to [**[]FansfeudBoardContent**](FansfeudBoardContent.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Tiebreak** | Pointer to [**FansfeudBoardContent**](FansfeudBoardContent.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFansfeudDraftView

`func NewFansfeudDraftView(type_ string, ) *FansfeudDraftView`

NewFansfeudDraftView instantiates a new FansfeudDraftView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFansfeudDraftViewWithDefaults

`func NewFansfeudDraftViewWithDefaults() *FansfeudDraftView`

NewFansfeudDraftViewWithDefaults instantiates a new FansfeudDraftView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoards

`func (o *FansfeudDraftView) GetBoards() []FansfeudBoardContent`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *FansfeudDraftView) GetBoardsOk() (*[]FansfeudBoardContent, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *FansfeudDraftView) SetBoards(v []FansfeudBoardContent)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *FansfeudDraftView) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### GetBrief

`func (o *FansfeudDraftView) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *FansfeudDraftView) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *FansfeudDraftView) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *FansfeudDraftView) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetGameId

`func (o *FansfeudDraftView) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *FansfeudDraftView) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *FansfeudDraftView) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *FansfeudDraftView) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetId

`func (o *FansfeudDraftView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FansfeudDraftView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FansfeudDraftView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FansfeudDraftView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRevision

`func (o *FansfeudDraftView) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FansfeudDraftView) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FansfeudDraftView) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FansfeudDraftView) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRush

`func (o *FansfeudDraftView) GetRush() []FansfeudBoardContent`

GetRush returns the Rush field if non-nil, zero value otherwise.

### GetRushOk

`func (o *FansfeudDraftView) GetRushOk() (*[]FansfeudBoardContent, bool)`

GetRushOk returns a tuple with the Rush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRush

`func (o *FansfeudDraftView) SetRush(v []FansfeudBoardContent)`

SetRush sets Rush field to given value.

### HasRush

`func (o *FansfeudDraftView) HasRush() bool`

HasRush returns a boolean if a field has been set.

### GetSource

`func (o *FansfeudDraftView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FansfeudDraftView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FansfeudDraftView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FansfeudDraftView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTiebreak

`func (o *FansfeudDraftView) GetTiebreak() FansfeudBoardContent`

GetTiebreak returns the Tiebreak field if non-nil, zero value otherwise.

### GetTiebreakOk

`func (o *FansfeudDraftView) GetTiebreakOk() (*FansfeudBoardContent, bool)`

GetTiebreakOk returns a tuple with the Tiebreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreak

`func (o *FansfeudDraftView) SetTiebreak(v FansfeudBoardContent)`

SetTiebreak sets Tiebreak field to given value.

### HasTiebreak

`func (o *FansfeudDraftView) HasTiebreak() bool`

HasTiebreak returns a boolean if a field has been set.

### GetType

`func (o *FansfeudDraftView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FansfeudDraftView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FansfeudDraftView) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


