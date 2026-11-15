# GamesGameView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceSpaceId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]FindtheminorityPlayer**](FindtheminorityPlayer.md) |  | [optional] 
**Round** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**ViewerRole** | Pointer to **string** |  | [optional] 
**Board** | Pointer to [**SevencluesBoardView**](SevencluesBoardView.md) |  | [optional] 
**BoardIndex** | Pointer to **int32** |  | [optional] 
**CompletedBoards** | Pointer to [**[]SevencluesBoardView**](SevencluesBoardView.md) |  | [optional] 
**ContentSource** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**PlayDefinitionRef**](PlayDefinitionRef.md) |  | [optional] 
**Finale** | Pointer to [**FanssaytypesFinaleView**](FanssaytypesFinaleView.md) |  | [optional] 
**FinalistTeam** | Pointer to **int32** |  | [optional] 
**LastFinaleAction** | Pointer to [**FanssaytypesFinaleResult**](FanssaytypesFinaleResult.md) |  | [optional] 
**LastGuess** | Pointer to [**SevencluesSubmissionResult**](SevencluesSubmissionResult.md) |  | [optional] 
**LastTiebreakAction** | Pointer to [**SevencluesTiebreakResult**](SevencluesTiebreakResult.md) |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**Result** | Pointer to [**WordchainGameResult**](WordchainGameResult.md) |  | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Stage** | Pointer to [**PlayStage**](PlayStage.md) |  | [optional] 
**TeamScores** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 
**Tiebreak** | Pointer to [**SevencluesTiebreakView**](SevencluesTiebreakView.md) |  | [optional] 
**WinnerTeam** | Pointer to **int32** |  | [optional] 
**Rush** | Pointer to [**LinkrushView**](LinkrushView.md) |  | [optional] 
**Showdown** | Pointer to [**FeudshowdownView**](FeudshowdownView.md) |  | [optional] 
**Clash** | Pointer to [**LinkclashView**](LinkclashView.md) |  | [optional] 

## Methods

### NewGamesGameView

`func NewGamesGameView(type_ string, ) *GamesGameView`

NewGamesGameView instantiates a new GamesGameView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamesGameViewWithDefaults

`func NewGamesGameViewWithDefaults() *GamesGameView`

NewGamesGameViewWithDefaults instantiates a new GamesGameView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceSpaceId

`func (o *GamesGameView) GetAudienceSpaceId() string`

GetAudienceSpaceId returns the AudienceSpaceId field if non-nil, zero value otherwise.

### GetAudienceSpaceIdOk

`func (o *GamesGameView) GetAudienceSpaceIdOk() (*string, bool)`

GetAudienceSpaceIdOk returns a tuple with the AudienceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceSpaceId

`func (o *GamesGameView) SetAudienceSpaceId(v string)`

SetAudienceSpaceId sets AudienceSpaceId field to given value.

### HasAudienceSpaceId

`func (o *GamesGameView) HasAudienceSpaceId() bool`

HasAudienceSpaceId returns a boolean if a field has been set.

### GetId

`func (o *GamesGameView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GamesGameView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GamesGameView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GamesGameView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *GamesGameView) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *GamesGameView) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *GamesGameView) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *GamesGameView) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetPlayers

`func (o *GamesGameView) GetPlayers() []FindtheminorityPlayer`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GamesGameView) GetPlayersOk() (*[]FindtheminorityPlayer, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GamesGameView) SetPlayers(v []FindtheminorityPlayer)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *GamesGameView) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetRound

`func (o *GamesGameView) GetRound() int32`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *GamesGameView) GetRoundOk() (*int32, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *GamesGameView) SetRound(v int32)`

SetRound sets Round field to given value.

### HasRound

`func (o *GamesGameView) HasRound() bool`

HasRound returns a boolean if a field has been set.

### GetState

`func (o *GamesGameView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GamesGameView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GamesGameView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GamesGameView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *GamesGameView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GamesGameView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GamesGameView) SetType(v string)`

SetType sets Type field to given value.


### GetViewerRole

`func (o *GamesGameView) GetViewerRole() string`

GetViewerRole returns the ViewerRole field if non-nil, zero value otherwise.

### GetViewerRoleOk

`func (o *GamesGameView) GetViewerRoleOk() (*string, bool)`

GetViewerRoleOk returns a tuple with the ViewerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRole

`func (o *GamesGameView) SetViewerRole(v string)`

SetViewerRole sets ViewerRole field to given value.

### HasViewerRole

`func (o *GamesGameView) HasViewerRole() bool`

HasViewerRole returns a boolean if a field has been set.

### GetBoard

`func (o *GamesGameView) GetBoard() SevencluesBoardView`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *GamesGameView) GetBoardOk() (*SevencluesBoardView, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *GamesGameView) SetBoard(v SevencluesBoardView)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *GamesGameView) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetBoardIndex

`func (o *GamesGameView) GetBoardIndex() int32`

GetBoardIndex returns the BoardIndex field if non-nil, zero value otherwise.

### GetBoardIndexOk

`func (o *GamesGameView) GetBoardIndexOk() (*int32, bool)`

GetBoardIndexOk returns a tuple with the BoardIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardIndex

`func (o *GamesGameView) SetBoardIndex(v int32)`

SetBoardIndex sets BoardIndex field to given value.

### HasBoardIndex

`func (o *GamesGameView) HasBoardIndex() bool`

HasBoardIndex returns a boolean if a field has been set.

### GetCompletedBoards

`func (o *GamesGameView) GetCompletedBoards() []SevencluesBoardView`

GetCompletedBoards returns the CompletedBoards field if non-nil, zero value otherwise.

### GetCompletedBoardsOk

`func (o *GamesGameView) GetCompletedBoardsOk() (*[]SevencluesBoardView, bool)`

GetCompletedBoardsOk returns a tuple with the CompletedBoards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBoards

`func (o *GamesGameView) SetCompletedBoards(v []SevencluesBoardView)`

SetCompletedBoards sets CompletedBoards field to given value.

### HasCompletedBoards

`func (o *GamesGameView) HasCompletedBoards() bool`

HasCompletedBoards returns a boolean if a field has been set.

### GetContentSource

`func (o *GamesGameView) GetContentSource() string`

GetContentSource returns the ContentSource field if non-nil, zero value otherwise.

### GetContentSourceOk

`func (o *GamesGameView) GetContentSourceOk() (*string, bool)`

GetContentSourceOk returns a tuple with the ContentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSource

`func (o *GamesGameView) SetContentSource(v string)`

SetContentSource sets ContentSource field to given value.

### HasContentSource

`func (o *GamesGameView) HasContentSource() bool`

HasContentSource returns a boolean if a field has been set.

### GetDefinition

`func (o *GamesGameView) GetDefinition() PlayDefinitionRef`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *GamesGameView) GetDefinitionOk() (*PlayDefinitionRef, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *GamesGameView) SetDefinition(v PlayDefinitionRef)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *GamesGameView) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetFinale

`func (o *GamesGameView) GetFinale() FanssaytypesFinaleView`

GetFinale returns the Finale field if non-nil, zero value otherwise.

### GetFinaleOk

`func (o *GamesGameView) GetFinaleOk() (*FanssaytypesFinaleView, bool)`

GetFinaleOk returns a tuple with the Finale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinale

`func (o *GamesGameView) SetFinale(v FanssaytypesFinaleView)`

SetFinale sets Finale field to given value.

### HasFinale

`func (o *GamesGameView) HasFinale() bool`

HasFinale returns a boolean if a field has been set.

### GetFinalistTeam

`func (o *GamesGameView) GetFinalistTeam() int32`

GetFinalistTeam returns the FinalistTeam field if non-nil, zero value otherwise.

### GetFinalistTeamOk

`func (o *GamesGameView) GetFinalistTeamOk() (*int32, bool)`

GetFinalistTeamOk returns a tuple with the FinalistTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistTeam

`func (o *GamesGameView) SetFinalistTeam(v int32)`

SetFinalistTeam sets FinalistTeam field to given value.

### HasFinalistTeam

`func (o *GamesGameView) HasFinalistTeam() bool`

HasFinalistTeam returns a boolean if a field has been set.

### GetLastFinaleAction

`func (o *GamesGameView) GetLastFinaleAction() FanssaytypesFinaleResult`

GetLastFinaleAction returns the LastFinaleAction field if non-nil, zero value otherwise.

### GetLastFinaleActionOk

`func (o *GamesGameView) GetLastFinaleActionOk() (*FanssaytypesFinaleResult, bool)`

GetLastFinaleActionOk returns a tuple with the LastFinaleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFinaleAction

`func (o *GamesGameView) SetLastFinaleAction(v FanssaytypesFinaleResult)`

SetLastFinaleAction sets LastFinaleAction field to given value.

### HasLastFinaleAction

`func (o *GamesGameView) HasLastFinaleAction() bool`

HasLastFinaleAction returns a boolean if a field has been set.

### GetLastGuess

`func (o *GamesGameView) GetLastGuess() SevencluesSubmissionResult`

GetLastGuess returns the LastGuess field if non-nil, zero value otherwise.

### GetLastGuessOk

`func (o *GamesGameView) GetLastGuessOk() (*SevencluesSubmissionResult, bool)`

GetLastGuessOk returns a tuple with the LastGuess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGuess

`func (o *GamesGameView) SetLastGuess(v SevencluesSubmissionResult)`

SetLastGuess sets LastGuess field to given value.

### HasLastGuess

`func (o *GamesGameView) HasLastGuess() bool`

HasLastGuess returns a boolean if a field has been set.

### GetLastTiebreakAction

`func (o *GamesGameView) GetLastTiebreakAction() SevencluesTiebreakResult`

GetLastTiebreakAction returns the LastTiebreakAction field if non-nil, zero value otherwise.

### GetLastTiebreakActionOk

`func (o *GamesGameView) GetLastTiebreakActionOk() (*SevencluesTiebreakResult, bool)`

GetLastTiebreakActionOk returns a tuple with the LastTiebreakAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTiebreakAction

`func (o *GamesGameView) SetLastTiebreakAction(v SevencluesTiebreakResult)`

SetLastTiebreakAction sets LastTiebreakAction field to given value.

### HasLastTiebreakAction

`func (o *GamesGameView) HasLastTiebreakAction() bool`

HasLastTiebreakAction returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *GamesGameView) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *GamesGameView) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *GamesGameView) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *GamesGameView) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetResult

`func (o *GamesGameView) GetResult() WordchainGameResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GamesGameView) GetResultOk() (*WordchainGameResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GamesGameView) SetResult(v WordchainGameResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GamesGameView) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReward

`func (o *GamesGameView) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *GamesGameView) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *GamesGameView) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *GamesGameView) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetStage

`func (o *GamesGameView) GetStage() PlayStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *GamesGameView) GetStageOk() (*PlayStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *GamesGameView) SetStage(v PlayStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *GamesGameView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetTeamScores

`func (o *GamesGameView) GetTeamScores() []int32`

GetTeamScores returns the TeamScores field if non-nil, zero value otherwise.

### GetTeamScoresOk

`func (o *GamesGameView) GetTeamScoresOk() (*[]int32, bool)`

GetTeamScoresOk returns a tuple with the TeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScores

`func (o *GamesGameView) SetTeamScores(v []int32)`

SetTeamScores sets TeamScores field to given value.

### HasTeamScores

`func (o *GamesGameView) HasTeamScores() bool`

HasTeamScores returns a boolean if a field has been set.

### GetTeams

`func (o *GamesGameView) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GamesGameView) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GamesGameView) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *GamesGameView) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTiebreak

`func (o *GamesGameView) GetTiebreak() SevencluesTiebreakView`

GetTiebreak returns the Tiebreak field if non-nil, zero value otherwise.

### GetTiebreakOk

`func (o *GamesGameView) GetTiebreakOk() (*SevencluesTiebreakView, bool)`

GetTiebreakOk returns a tuple with the Tiebreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreak

`func (o *GamesGameView) SetTiebreak(v SevencluesTiebreakView)`

SetTiebreak sets Tiebreak field to given value.

### HasTiebreak

`func (o *GamesGameView) HasTiebreak() bool`

HasTiebreak returns a boolean if a field has been set.

### GetWinnerTeam

`func (o *GamesGameView) GetWinnerTeam() int32`

GetWinnerTeam returns the WinnerTeam field if non-nil, zero value otherwise.

### GetWinnerTeamOk

`func (o *GamesGameView) GetWinnerTeamOk() (*int32, bool)`

GetWinnerTeamOk returns a tuple with the WinnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerTeam

`func (o *GamesGameView) SetWinnerTeam(v int32)`

SetWinnerTeam sets WinnerTeam field to given value.

### HasWinnerTeam

`func (o *GamesGameView) HasWinnerTeam() bool`

HasWinnerTeam returns a boolean if a field has been set.

### GetRush

`func (o *GamesGameView) GetRush() LinkrushView`

GetRush returns the Rush field if non-nil, zero value otherwise.

### GetRushOk

`func (o *GamesGameView) GetRushOk() (*LinkrushView, bool)`

GetRushOk returns a tuple with the Rush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRush

`func (o *GamesGameView) SetRush(v LinkrushView)`

SetRush sets Rush field to given value.

### HasRush

`func (o *GamesGameView) HasRush() bool`

HasRush returns a boolean if a field has been set.

### GetShowdown

`func (o *GamesGameView) GetShowdown() FeudshowdownView`

GetShowdown returns the Showdown field if non-nil, zero value otherwise.

### GetShowdownOk

`func (o *GamesGameView) GetShowdownOk() (*FeudshowdownView, bool)`

GetShowdownOk returns a tuple with the Showdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowdown

`func (o *GamesGameView) SetShowdown(v FeudshowdownView)`

SetShowdown sets Showdown field to given value.

### HasShowdown

`func (o *GamesGameView) HasShowdown() bool`

HasShowdown returns a boolean if a field has been set.

### GetClash

`func (o *GamesGameView) GetClash() LinkclashView`

GetClash returns the Clash field if non-nil, zero value otherwise.

### GetClashOk

`func (o *GamesGameView) GetClashOk() (*LinkclashView, bool)`

GetClashOk returns a tuple with the Clash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClash

`func (o *GamesGameView) SetClash(v LinkclashView)`

SetClash sets Clash field to given value.

### HasClash

`func (o *GamesGameView) HasClash() bool`

HasClash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


