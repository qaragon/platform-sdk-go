# SubmissionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedAnswerId** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 
**ParticipantScore** | Pointer to **int32** |  | [optional] 
**PointsAwarded** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Board** | Pointer to [**SevencluesBoardView**](SevencluesBoardView.md) |  | [optional] 
**Guess** | Pointer to [**GuesstypesResult**](GuesstypesResult.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Tiebreak** | Pointer to [**SevencluesTiebreakView**](SevencluesTiebreakView.md) |  | [optional] 
**Finale** | Pointer to [**FanssaytypesFinaleView**](FanssaytypesFinaleView.md) |  | [optional] 
**PromptId** | Pointer to **string** |  | [optional] 
**Game** | Pointer to [**WordchainGameView**](WordchainGameView.md) |  | [optional] 
**Showdown** | Pointer to [**FeudshowdownRoundView**](FeudshowdownRoundView.md) |  | [optional] 

## Methods

### NewSubmissionResult

`func NewSubmissionResult(type_ string, ) *SubmissionResult`

NewSubmissionResult instantiates a new SubmissionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionResultWithDefaults

`func NewSubmissionResultWithDefaults() *SubmissionResult`

NewSubmissionResultWithDefaults instantiates a new SubmissionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedAnswerId

`func (o *SubmissionResult) GetMatchedAnswerId() string`

GetMatchedAnswerId returns the MatchedAnswerId field if non-nil, zero value otherwise.

### GetMatchedAnswerIdOk

`func (o *SubmissionResult) GetMatchedAnswerIdOk() (*string, bool)`

GetMatchedAnswerIdOk returns a tuple with the MatchedAnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedAnswerId

`func (o *SubmissionResult) SetMatchedAnswerId(v string)`

SetMatchedAnswerId sets MatchedAnswerId field to given value.

### HasMatchedAnswerId

`func (o *SubmissionResult) HasMatchedAnswerId() bool`

HasMatchedAnswerId returns a boolean if a field has been set.

### GetOutcome

`func (o *SubmissionResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *SubmissionResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *SubmissionResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *SubmissionResult) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetParticipantScore

`func (o *SubmissionResult) GetParticipantScore() int32`

GetParticipantScore returns the ParticipantScore field if non-nil, zero value otherwise.

### GetParticipantScoreOk

`func (o *SubmissionResult) GetParticipantScoreOk() (*int32, bool)`

GetParticipantScoreOk returns a tuple with the ParticipantScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantScore

`func (o *SubmissionResult) SetParticipantScore(v int32)`

SetParticipantScore sets ParticipantScore field to given value.

### HasParticipantScore

`func (o *SubmissionResult) HasParticipantScore() bool`

HasParticipantScore returns a boolean if a field has been set.

### GetPointsAwarded

`func (o *SubmissionResult) GetPointsAwarded() int32`

GetPointsAwarded returns the PointsAwarded field if non-nil, zero value otherwise.

### GetPointsAwardedOk

`func (o *SubmissionResult) GetPointsAwardedOk() (*int32, bool)`

GetPointsAwardedOk returns a tuple with the PointsAwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsAwarded

`func (o *SubmissionResult) SetPointsAwarded(v int32)`

SetPointsAwarded sets PointsAwarded field to given value.

### HasPointsAwarded

`func (o *SubmissionResult) HasPointsAwarded() bool`

HasPointsAwarded returns a boolean if a field has been set.

### GetType

`func (o *SubmissionResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubmissionResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubmissionResult) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SubmissionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBoard

`func (o *SubmissionResult) GetBoard() SevencluesBoardView`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *SubmissionResult) GetBoardOk() (*SevencluesBoardView, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *SubmissionResult) SetBoard(v SevencluesBoardView)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *SubmissionResult) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetGuess

`func (o *SubmissionResult) GetGuess() GuesstypesResult`

GetGuess returns the Guess field if non-nil, zero value otherwise.

### GetGuessOk

`func (o *SubmissionResult) GetGuessOk() (*GuesstypesResult, bool)`

GetGuessOk returns a tuple with the Guess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuess

`func (o *SubmissionResult) SetGuess(v GuesstypesResult)`

SetGuess sets Guess field to given value.

### HasGuess

`func (o *SubmissionResult) HasGuess() bool`

HasGuess returns a boolean if a field has been set.

### GetAction

`func (o *SubmissionResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SubmissionResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SubmissionResult) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SubmissionResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTiebreak

`func (o *SubmissionResult) GetTiebreak() SevencluesTiebreakView`

GetTiebreak returns the Tiebreak field if non-nil, zero value otherwise.

### GetTiebreakOk

`func (o *SubmissionResult) GetTiebreakOk() (*SevencluesTiebreakView, bool)`

GetTiebreakOk returns a tuple with the Tiebreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreak

`func (o *SubmissionResult) SetTiebreak(v SevencluesTiebreakView)`

SetTiebreak sets Tiebreak field to given value.

### HasTiebreak

`func (o *SubmissionResult) HasTiebreak() bool`

HasTiebreak returns a boolean if a field has been set.

### GetFinale

`func (o *SubmissionResult) GetFinale() FanssaytypesFinaleView`

GetFinale returns the Finale field if non-nil, zero value otherwise.

### GetFinaleOk

`func (o *SubmissionResult) GetFinaleOk() (*FanssaytypesFinaleView, bool)`

GetFinaleOk returns a tuple with the Finale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinale

`func (o *SubmissionResult) SetFinale(v FanssaytypesFinaleView)`

SetFinale sets Finale field to given value.

### HasFinale

`func (o *SubmissionResult) HasFinale() bool`

HasFinale returns a boolean if a field has been set.

### GetPromptId

`func (o *SubmissionResult) GetPromptId() string`

GetPromptId returns the PromptId field if non-nil, zero value otherwise.

### GetPromptIdOk

`func (o *SubmissionResult) GetPromptIdOk() (*string, bool)`

GetPromptIdOk returns a tuple with the PromptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptId

`func (o *SubmissionResult) SetPromptId(v string)`

SetPromptId sets PromptId field to given value.

### HasPromptId

`func (o *SubmissionResult) HasPromptId() bool`

HasPromptId returns a boolean if a field has been set.

### GetGame

`func (o *SubmissionResult) GetGame() WordchainGameView`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *SubmissionResult) GetGameOk() (*WordchainGameView, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *SubmissionResult) SetGame(v WordchainGameView)`

SetGame sets Game field to given value.

### HasGame

`func (o *SubmissionResult) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetShowdown

`func (o *SubmissionResult) GetShowdown() FeudshowdownRoundView`

GetShowdown returns the Showdown field if non-nil, zero value otherwise.

### GetShowdownOk

`func (o *SubmissionResult) GetShowdownOk() (*FeudshowdownRoundView, bool)`

GetShowdownOk returns a tuple with the Showdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowdown

`func (o *SubmissionResult) SetShowdown(v FeudshowdownRoundView)`

SetShowdown sets Showdown field to given value.

### HasShowdown

`func (o *SubmissionResult) HasShowdown() bool`

HasShowdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


