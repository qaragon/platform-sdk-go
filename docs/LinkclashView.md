# LinkclashView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ActiveIdentityId** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to [**[]WordchainWordView**](WordchainWordView.md) |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**ControlTeam** | Pointer to **int32** |  | [optional] 
**Deadline** | Pointer to **NullableString** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**RoundNumber** | Pointer to **int32** |  | [optional] 
**SelectedPosition** | Pointer to **NullableInt32** |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 
**TeamScores** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 
**Wager** | Pointer to **int32** |  | [optional] 
**WinnerTeam** | Pointer to **int32** |  | [optional] 

## Methods

### NewLinkclashView

`func NewLinkclashView() *LinkclashView`

NewLinkclashView instantiates a new LinkclashView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkclashViewWithDefaults

`func NewLinkclashViewWithDefaults() *LinkclashView`

NewLinkclashViewWithDefaults instantiates a new LinkclashView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionVersion

`func (o *LinkclashView) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *LinkclashView) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *LinkclashView) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *LinkclashView) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetActiveIdentityId

`func (o *LinkclashView) GetActiveIdentityId() string`

GetActiveIdentityId returns the ActiveIdentityId field if non-nil, zero value otherwise.

### GetActiveIdentityIdOk

`func (o *LinkclashView) GetActiveIdentityIdOk() (*string, bool)`

GetActiveIdentityIdOk returns a tuple with the ActiveIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentityId

`func (o *LinkclashView) SetActiveIdentityId(v string)`

SetActiveIdentityId sets ActiveIdentityId field to given value.

### HasActiveIdentityId

`func (o *LinkclashView) HasActiveIdentityId() bool`

HasActiveIdentityId returns a boolean if a field has been set.

### GetChain

`func (o *LinkclashView) GetChain() []WordchainWordView`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *LinkclashView) GetChainOk() (*[]WordchainWordView, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *LinkclashView) SetChain(v []WordchainWordView)`

SetChain sets Chain field to given value.

### HasChain

`func (o *LinkclashView) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *LinkclashView) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *LinkclashView) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *LinkclashView) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *LinkclashView) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetControlTeam

`func (o *LinkclashView) GetControlTeam() int32`

GetControlTeam returns the ControlTeam field if non-nil, zero value otherwise.

### GetControlTeamOk

`func (o *LinkclashView) GetControlTeamOk() (*int32, bool)`

GetControlTeamOk returns a tuple with the ControlTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTeam

`func (o *LinkclashView) SetControlTeam(v int32)`

SetControlTeam sets ControlTeam field to given value.

### HasControlTeam

`func (o *LinkclashView) HasControlTeam() bool`

HasControlTeam returns a boolean if a field has been set.

### GetDeadline

`func (o *LinkclashView) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *LinkclashView) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *LinkclashView) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *LinkclashView) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *LinkclashView) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *LinkclashView) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetPhase

`func (o *LinkclashView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *LinkclashView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *LinkclashView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *LinkclashView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRoundNumber

`func (o *LinkclashView) GetRoundNumber() int32`

GetRoundNumber returns the RoundNumber field if non-nil, zero value otherwise.

### GetRoundNumberOk

`func (o *LinkclashView) GetRoundNumberOk() (*int32, bool)`

GetRoundNumberOk returns a tuple with the RoundNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundNumber

`func (o *LinkclashView) SetRoundNumber(v int32)`

SetRoundNumber sets RoundNumber field to given value.

### HasRoundNumber

`func (o *LinkclashView) HasRoundNumber() bool`

HasRoundNumber returns a boolean if a field has been set.

### GetSelectedPosition

`func (o *LinkclashView) GetSelectedPosition() int32`

GetSelectedPosition returns the SelectedPosition field if non-nil, zero value otherwise.

### GetSelectedPositionOk

`func (o *LinkclashView) GetSelectedPositionOk() (*int32, bool)`

GetSelectedPositionOk returns a tuple with the SelectedPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPosition

`func (o *LinkclashView) SetSelectedPosition(v int32)`

SetSelectedPosition sets SelectedPosition field to given value.

### HasSelectedPosition

`func (o *LinkclashView) HasSelectedPosition() bool`

HasSelectedPosition returns a boolean if a field has been set.

### SetSelectedPositionNil

`func (o *LinkclashView) SetSelectedPositionNil(b bool)`

 SetSelectedPositionNil sets the value for SelectedPosition to be an explicit nil

### UnsetSelectedPosition
`func (o *LinkclashView) UnsetSelectedPosition()`

UnsetSelectedPosition ensures that no value is present for SelectedPosition, not even an explicit nil
### GetStageId

`func (o *LinkclashView) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *LinkclashView) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *LinkclashView) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *LinkclashView) HasStageId() bool`

HasStageId returns a boolean if a field has been set.

### GetTeamScores

`func (o *LinkclashView) GetTeamScores() []int32`

GetTeamScores returns the TeamScores field if non-nil, zero value otherwise.

### GetTeamScoresOk

`func (o *LinkclashView) GetTeamScoresOk() (*[]int32, bool)`

GetTeamScoresOk returns a tuple with the TeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScores

`func (o *LinkclashView) SetTeamScores(v []int32)`

SetTeamScores sets TeamScores field to given value.

### HasTeamScores

`func (o *LinkclashView) HasTeamScores() bool`

HasTeamScores returns a boolean if a field has been set.

### GetTeams

`func (o *LinkclashView) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *LinkclashView) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *LinkclashView) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *LinkclashView) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetWager

`func (o *LinkclashView) GetWager() int32`

GetWager returns the Wager field if non-nil, zero value otherwise.

### GetWagerOk

`func (o *LinkclashView) GetWagerOk() (*int32, bool)`

GetWagerOk returns a tuple with the Wager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWager

`func (o *LinkclashView) SetWager(v int32)`

SetWager sets Wager field to given value.

### HasWager

`func (o *LinkclashView) HasWager() bool`

HasWager returns a boolean if a field has been set.

### GetWinnerTeam

`func (o *LinkclashView) GetWinnerTeam() int32`

GetWinnerTeam returns the WinnerTeam field if non-nil, zero value otherwise.

### GetWinnerTeamOk

`func (o *LinkclashView) GetWinnerTeamOk() (*int32, bool)`

GetWinnerTeamOk returns a tuple with the WinnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerTeam

`func (o *LinkclashView) SetWinnerTeam(v int32)`

SetWinnerTeam sets WinnerTeam field to given value.

### HasWinnerTeam

`func (o *LinkclashView) HasWinnerTeam() bool`

HasWinnerTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


