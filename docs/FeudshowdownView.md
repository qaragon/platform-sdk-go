# FeudshowdownView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedRounds** | Pointer to [**[]FeudshowdownRoundView**](FeudshowdownRoundView.md) |  | [optional] 
**Round** | Pointer to [**FeudshowdownRoundView**](FeudshowdownRoundView.md) |  | [optional] 
**RoundIndex** | Pointer to **int32** |  | [optional] 
**TeamScores** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewFeudshowdownView

`func NewFeudshowdownView() *FeudshowdownView`

NewFeudshowdownView instantiates a new FeudshowdownView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeudshowdownViewWithDefaults

`func NewFeudshowdownViewWithDefaults() *FeudshowdownView`

NewFeudshowdownViewWithDefaults instantiates a new FeudshowdownView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedRounds

`func (o *FeudshowdownView) GetCompletedRounds() []FeudshowdownRoundView`

GetCompletedRounds returns the CompletedRounds field if non-nil, zero value otherwise.

### GetCompletedRoundsOk

`func (o *FeudshowdownView) GetCompletedRoundsOk() (*[]FeudshowdownRoundView, bool)`

GetCompletedRoundsOk returns a tuple with the CompletedRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedRounds

`func (o *FeudshowdownView) SetCompletedRounds(v []FeudshowdownRoundView)`

SetCompletedRounds sets CompletedRounds field to given value.

### HasCompletedRounds

`func (o *FeudshowdownView) HasCompletedRounds() bool`

HasCompletedRounds returns a boolean if a field has been set.

### GetRound

`func (o *FeudshowdownView) GetRound() FeudshowdownRoundView`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *FeudshowdownView) GetRoundOk() (*FeudshowdownRoundView, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *FeudshowdownView) SetRound(v FeudshowdownRoundView)`

SetRound sets Round field to given value.

### HasRound

`func (o *FeudshowdownView) HasRound() bool`

HasRound returns a boolean if a field has been set.

### GetRoundIndex

`func (o *FeudshowdownView) GetRoundIndex() int32`

GetRoundIndex returns the RoundIndex field if non-nil, zero value otherwise.

### GetRoundIndexOk

`func (o *FeudshowdownView) GetRoundIndexOk() (*int32, bool)`

GetRoundIndexOk returns a tuple with the RoundIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundIndex

`func (o *FeudshowdownView) SetRoundIndex(v int32)`

SetRoundIndex sets RoundIndex field to given value.

### HasRoundIndex

`func (o *FeudshowdownView) HasRoundIndex() bool`

HasRoundIndex returns a boolean if a field has been set.

### GetTeamScores

`func (o *FeudshowdownView) GetTeamScores() []int32`

GetTeamScores returns the TeamScores field if non-nil, zero value otherwise.

### GetTeamScoresOk

`func (o *FeudshowdownView) GetTeamScoresOk() (*[]int32, bool)`

GetTeamScoresOk returns a tuple with the TeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScores

`func (o *FeudshowdownView) SetTeamScores(v []int32)`

SetTeamScores sets TeamScores field to given value.

### HasTeamScores

`func (o *FeudshowdownView) HasTeamScores() bool`

HasTeamScores returns a boolean if a field has been set.

### GetTeams

`func (o *FeudshowdownView) GetTeams() [][]string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *FeudshowdownView) GetTeamsOk() (*[][]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *FeudshowdownView) SetTeams(v [][]string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *FeudshowdownView) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


