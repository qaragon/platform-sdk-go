# FansfeudSubmissionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Game** | Pointer to [**FansfeudGameView**](FansfeudGameView.md) |  | [optional] 
**Guess** | Pointer to [**GuesstypesResult**](GuesstypesResult.md) |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 
**Showdown** | Pointer to [**FeudshowdownRoundView**](FeudshowdownRoundView.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFansfeudSubmissionResult

`func NewFansfeudSubmissionResult(type_ string, ) *FansfeudSubmissionResult`

NewFansfeudSubmissionResult instantiates a new FansfeudSubmissionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFansfeudSubmissionResultWithDefaults

`func NewFansfeudSubmissionResultWithDefaults() *FansfeudSubmissionResult`

NewFansfeudSubmissionResultWithDefaults instantiates a new FansfeudSubmissionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FansfeudSubmissionResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FansfeudSubmissionResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FansfeudSubmissionResult) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FansfeudSubmissionResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetGame

`func (o *FansfeudSubmissionResult) GetGame() FansfeudGameView`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *FansfeudSubmissionResult) GetGameOk() (*FansfeudGameView, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *FansfeudSubmissionResult) SetGame(v FansfeudGameView)`

SetGame sets Game field to given value.

### HasGame

`func (o *FansfeudSubmissionResult) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetGuess

`func (o *FansfeudSubmissionResult) GetGuess() GuesstypesResult`

GetGuess returns the Guess field if non-nil, zero value otherwise.

### GetGuessOk

`func (o *FansfeudSubmissionResult) GetGuessOk() (*GuesstypesResult, bool)`

GetGuessOk returns a tuple with the Guess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuess

`func (o *FansfeudSubmissionResult) SetGuess(v GuesstypesResult)`

SetGuess sets Guess field to given value.

### HasGuess

`func (o *FansfeudSubmissionResult) HasGuess() bool`

HasGuess returns a boolean if a field has been set.

### GetOutcome

`func (o *FansfeudSubmissionResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *FansfeudSubmissionResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *FansfeudSubmissionResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *FansfeudSubmissionResult) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetShowdown

`func (o *FansfeudSubmissionResult) GetShowdown() FeudshowdownRoundView`

GetShowdown returns the Showdown field if non-nil, zero value otherwise.

### GetShowdownOk

`func (o *FansfeudSubmissionResult) GetShowdownOk() (*FeudshowdownRoundView, bool)`

GetShowdownOk returns a tuple with the Showdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowdown

`func (o *FansfeudSubmissionResult) SetShowdown(v FeudshowdownRoundView)`

SetShowdown sets Showdown field to given value.

### HasShowdown

`func (o *FansfeudSubmissionResult) HasShowdown() bool`

HasShowdown returns a boolean if a field has been set.

### GetType

`func (o *FansfeudSubmissionResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FansfeudSubmissionResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FansfeudSubmissionResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


