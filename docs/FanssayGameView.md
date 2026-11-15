# FanssayGameView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Board** | Pointer to [**SevencluesBoardView**](SevencluesBoardView.md) |  | [optional] 
**BoardIndex** | Pointer to **int32** |  | [optional] 
**CompletedBoards** | Pointer to [**[]SevencluesBoardView**](SevencluesBoardView.md) |  | [optional] 
**ContentSource** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**PlayDefinitionRef**](PlayDefinitionRef.md) |  | [optional] 
**Finale** | Pointer to [**FanssaytypesFinaleView**](FanssaytypesFinaleView.md) |  | [optional] 
**FinalistTeam** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastFinaleAction** | Pointer to [**FanssaytypesFinaleResult**](FanssaytypesFinaleResult.md) |  | [optional] 
**LastGuess** | Pointer to [**SevencluesSubmissionResult**](SevencluesSubmissionResult.md) |  | [optional] 
**LastTiebreakAction** | Pointer to [**SevencluesTiebreakResult**](SevencluesTiebreakResult.md) |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**FanssaytypesGameResult**](FanssaytypesGameResult.md) |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Round** | Pointer to **int32** |  | [optional] 
**Stage** | Pointer to [**PlayStage**](PlayStage.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TeamScores** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 
**Tiebreak** | Pointer to [**SevencluesTiebreakView**](SevencluesTiebreakView.md) |  | [optional] 
**Type** | **string** |  | 
**ViewerRole** | Pointer to **string** |  | [optional] 
**WinnerTeam** | Pointer to **int32** |  | [optional] 

## Methods

### NewFanssayGameView

`func NewFanssayGameView(type_ string, ) *FanssayGameView`

NewFanssayGameView instantiates a new FanssayGameView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanssayGameViewWithDefaults

`func NewFanssayGameViewWithDefaults() *FanssayGameView`

NewFanssayGameViewWithDefaults instantiates a new FanssayGameView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoard

`func (o *FanssayGameView) GetBoard() SevencluesBoardView`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *FanssayGameView) GetBoardOk() (*SevencluesBoardView, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *FanssayGameView) SetBoard(v SevencluesBoardView)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *FanssayGameView) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetBoardIndex

`func (o *FanssayGameView) GetBoardIndex() int32`

GetBoardIndex returns the BoardIndex field if non-nil, zero value otherwise.

### GetBoardIndexOk

`func (o *FanssayGameView) GetBoardIndexOk() (*int32, bool)`

GetBoardIndexOk returns a tuple with the BoardIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardIndex

`func (o *FanssayGameView) SetBoardIndex(v int32)`

SetBoardIndex sets BoardIndex field to given value.

### HasBoardIndex

`func (o *FanssayGameView) HasBoardIndex() bool`

HasBoardIndex returns a boolean if a field has been set.

### GetCompletedBoards

`func (o *FanssayGameView) GetCompletedBoards() []SevencluesBoardView`

GetCompletedBoards returns the CompletedBoards field if non-nil, zero value otherwise.

### GetCompletedBoardsOk

`func (o *FanssayGameView) GetCompletedBoardsOk() (*[]SevencluesBoardView, bool)`

GetCompletedBoardsOk returns a tuple with the CompletedBoards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBoards

`func (o *FanssayGameView) SetCompletedBoards(v []SevencluesBoardView)`

SetCompletedBoards sets CompletedBoards field to given value.

### HasCompletedBoards

`func (o *FanssayGameView) HasCompletedBoards() bool`

HasCompletedBoards returns a boolean if a field has been set.

### GetContentSource

`func (o *FanssayGameView) GetContentSource() string`

GetContentSource returns the ContentSource field if non-nil, zero value otherwise.

### GetContentSourceOk

`func (o *FanssayGameView) GetContentSourceOk() (*string, bool)`

GetContentSourceOk returns a tuple with the ContentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSource

`func (o *FanssayGameView) SetContentSource(v string)`

SetContentSource sets ContentSource field to given value.

### HasContentSource

`func (o *FanssayGameView) HasContentSource() bool`

HasContentSource returns a boolean if a field has been set.

### GetDefinition

`func (o *FanssayGameView) GetDefinition() PlayDefinitionRef`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *FanssayGameView) GetDefinitionOk() (*PlayDefinitionRef, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *FanssayGameView) SetDefinition(v PlayDefinitionRef)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *FanssayGameView) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetFinale

`func (o *FanssayGameView) GetFinale() FanssaytypesFinaleView`

GetFinale returns the Finale field if non-nil, zero value otherwise.

### GetFinaleOk

`func (o *FanssayGameView) GetFinaleOk() (*FanssaytypesFinaleView, bool)`

GetFinaleOk returns a tuple with the Finale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinale

`func (o *FanssayGameView) SetFinale(v FanssaytypesFinaleView)`

SetFinale sets Finale field to given value.

### HasFinale

`func (o *FanssayGameView) HasFinale() bool`

HasFinale returns a boolean if a field has been set.

### GetFinalistTeam

`func (o *FanssayGameView) GetFinalistTeam() int32`

GetFinalistTeam returns the FinalistTeam field if non-nil, zero value otherwise.

### GetFinalistTeamOk

`func (o *FanssayGameView) GetFinalistTeamOk() (*int32, bool)`

GetFinalistTeamOk returns a tuple with the FinalistTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistTeam

`func (o *FanssayGameView) SetFinalistTeam(v int32)`

SetFinalistTeam sets FinalistTeam field to given value.

### HasFinalistTeam

`func (o *FanssayGameView) HasFinalistTeam() bool`

HasFinalistTeam returns a boolean if a field has been set.

### GetId

`func (o *FanssayGameView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FanssayGameView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FanssayGameView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FanssayGameView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastFinaleAction

`func (o *FanssayGameView) GetLastFinaleAction() FanssaytypesFinaleResult`

GetLastFinaleAction returns the LastFinaleAction field if non-nil, zero value otherwise.

### GetLastFinaleActionOk

`func (o *FanssayGameView) GetLastFinaleActionOk() (*FanssaytypesFinaleResult, bool)`

GetLastFinaleActionOk returns a tuple with the LastFinaleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFinaleAction

`func (o *FanssayGameView) SetLastFinaleAction(v FanssaytypesFinaleResult)`

SetLastFinaleAction sets LastFinaleAction field to given value.

### HasLastFinaleAction

`func (o *FanssayGameView) HasLastFinaleAction() bool`

HasLastFinaleAction returns a boolean if a field has been set.

### GetLastGuess

`func (o *FanssayGameView) GetLastGuess() SevencluesSubmissionResult`

GetLastGuess returns the LastGuess field if non-nil, zero value otherwise.

### GetLastGuessOk

`func (o *FanssayGameView) GetLastGuessOk() (*SevencluesSubmissionResult, bool)`

GetLastGuessOk returns a tuple with the LastGuess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGuess

`func (o *FanssayGameView) SetLastGuess(v SevencluesSubmissionResult)`

SetLastGuess sets LastGuess field to given value.

### HasLastGuess

`func (o *FanssayGameView) HasLastGuess() bool`

HasLastGuess returns a boolean if a field has been set.

### GetLastTiebreakAction

`func (o *FanssayGameView) GetLastTiebreakAction() SevencluesTiebreakResult`

GetLastTiebreakAction returns the LastTiebreakAction field if non-nil, zero value otherwise.

### GetLastTiebreakActionOk

`func (o *FanssayGameView) GetLastTiebreakActionOk() (*SevencluesTiebreakResult, bool)`

GetLastTiebreakActionOk returns a tuple with the LastTiebreakAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTiebreakAction

`func (o *FanssayGameView) SetLastTiebreakAction(v SevencluesTiebreakResult)`

SetLastTiebreakAction sets LastTiebreakAction field to given value.

### HasLastTiebreakAction

`func (o *FanssayGameView) HasLastTiebreakAction() bool`

HasLastTiebreakAction returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *FanssayGameView) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *FanssayGameView) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *FanssayGameView) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *FanssayGameView) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *FanssayGameView) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *FanssayGameView) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *FanssayGameView) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *FanssayGameView) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetResult

`func (o *FanssayGameView) GetResult() FanssaytypesGameResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FanssayGameView) GetResultOk() (*FanssaytypesGameResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FanssayGameView) SetResult(v FanssaytypesGameResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *FanssayGameView) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReward

`func (o *FanssayGameView) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *FanssayGameView) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *FanssayGameView) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *FanssayGameView) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetRound

`func (o *FanssayGameView) GetRound() int32`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *FanssayGameView) GetRoundOk() (*int32, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *FanssayGameView) SetRound(v int32)`

SetRound sets Round field to given value.

### HasRound

`func (o *FanssayGameView) HasRound() bool`

HasRound returns a boolean if a field has been set.

### GetStage

`func (o *FanssayGameView) GetStage() PlayStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FanssayGameView) GetStageOk() (*PlayStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FanssayGameView) SetStage(v PlayStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *FanssayGameView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *FanssayGameView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FanssayGameView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FanssayGameView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FanssayGameView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTeamScores

`func (o *FanssayGameView) GetTeamScores() []int32`

GetTeamScores returns the TeamScores field if non-nil, zero value otherwise.

### GetTeamScoresOk

`func (o *FanssayGameView) GetTeamScoresOk() (*[]int32, bool)`

GetTeamScoresOk returns a tuple with the TeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScores

`func (o *FanssayGameView) SetTeamScores(v []int32)`

SetTeamScores sets TeamScores field to given value.

### HasTeamScores

`func (o *FanssayGameView) HasTeamScores() bool`

HasTeamScores returns a boolean if a field has been set.

### GetTeams

`func (o *FanssayGameView) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *FanssayGameView) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *FanssayGameView) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *FanssayGameView) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTiebreak

`func (o *FanssayGameView) GetTiebreak() SevencluesTiebreakView`

GetTiebreak returns the Tiebreak field if non-nil, zero value otherwise.

### GetTiebreakOk

`func (o *FanssayGameView) GetTiebreakOk() (*SevencluesTiebreakView, bool)`

GetTiebreakOk returns a tuple with the Tiebreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreak

`func (o *FanssayGameView) SetTiebreak(v SevencluesTiebreakView)`

SetTiebreak sets Tiebreak field to given value.

### HasTiebreak

`func (o *FanssayGameView) HasTiebreak() bool`

HasTiebreak returns a boolean if a field has been set.

### GetType

`func (o *FanssayGameView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FanssayGameView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FanssayGameView) SetType(v string)`

SetType sets Type field to given value.


### GetViewerRole

`func (o *FanssayGameView) GetViewerRole() string`

GetViewerRole returns the ViewerRole field if non-nil, zero value otherwise.

### GetViewerRoleOk

`func (o *FanssayGameView) GetViewerRoleOk() (*string, bool)`

GetViewerRoleOk returns a tuple with the ViewerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRole

`func (o *FanssayGameView) SetViewerRole(v string)`

SetViewerRole sets ViewerRole field to given value.

### HasViewerRole

`func (o *FanssayGameView) HasViewerRole() bool`

HasViewerRole returns a boolean if a field has been set.

### GetWinnerTeam

`func (o *FanssayGameView) GetWinnerTeam() int32`

GetWinnerTeam returns the WinnerTeam field if non-nil, zero value otherwise.

### GetWinnerTeamOk

`func (o *FanssayGameView) GetWinnerTeamOk() (*int32, bool)`

GetWinnerTeamOk returns a tuple with the WinnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerTeam

`func (o *FanssayGameView) SetWinnerTeam(v int32)`

SetWinnerTeam sets WinnerTeam field to given value.

### HasWinnerTeam

`func (o *FanssayGameView) HasWinnerTeam() bool`

HasWinnerTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


