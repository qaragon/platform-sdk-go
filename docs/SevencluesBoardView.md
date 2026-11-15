# SevencluesBoardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ActiveIdentityId** | Pointer to **string** |  | [optional] 
**ActiveTeam** | Pointer to **int32** |  | [optional] 
**Answers** | Pointer to [**[]SevencluesAnswerView**](SevencluesAnswerView.md) |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TeamScores** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSevencluesBoardView

`func NewSevencluesBoardView() *SevencluesBoardView`

NewSevencluesBoardView instantiates a new SevencluesBoardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSevencluesBoardViewWithDefaults

`func NewSevencluesBoardViewWithDefaults() *SevencluesBoardView`

NewSevencluesBoardViewWithDefaults instantiates a new SevencluesBoardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionVersion

`func (o *SevencluesBoardView) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *SevencluesBoardView) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *SevencluesBoardView) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *SevencluesBoardView) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetActiveIdentityId

`func (o *SevencluesBoardView) GetActiveIdentityId() string`

GetActiveIdentityId returns the ActiveIdentityId field if non-nil, zero value otherwise.

### GetActiveIdentityIdOk

`func (o *SevencluesBoardView) GetActiveIdentityIdOk() (*string, bool)`

GetActiveIdentityIdOk returns a tuple with the ActiveIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentityId

`func (o *SevencluesBoardView) SetActiveIdentityId(v string)`

SetActiveIdentityId sets ActiveIdentityId field to given value.

### HasActiveIdentityId

`func (o *SevencluesBoardView) HasActiveIdentityId() bool`

HasActiveIdentityId returns a boolean if a field has been set.

### GetActiveTeam

`func (o *SevencluesBoardView) GetActiveTeam() int32`

GetActiveTeam returns the ActiveTeam field if non-nil, zero value otherwise.

### GetActiveTeamOk

`func (o *SevencluesBoardView) GetActiveTeamOk() (*int32, bool)`

GetActiveTeamOk returns a tuple with the ActiveTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTeam

`func (o *SevencluesBoardView) SetActiveTeam(v int32)`

SetActiveTeam sets ActiveTeam field to given value.

### HasActiveTeam

`func (o *SevencluesBoardView) HasActiveTeam() bool`

HasActiveTeam returns a boolean if a field has been set.

### GetAnswers

`func (o *SevencluesBoardView) GetAnswers() []SevencluesAnswerView`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *SevencluesBoardView) GetAnswersOk() (*[]SevencluesAnswerView, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *SevencluesBoardView) SetAnswers(v []SevencluesAnswerView)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *SevencluesBoardView) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetDeadline

`func (o *SevencluesBoardView) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *SevencluesBoardView) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *SevencluesBoardView) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *SevencluesBoardView) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetId

`func (o *SevencluesBoardView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SevencluesBoardView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SevencluesBoardView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SevencluesBoardView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrompt

`func (o *SevencluesBoardView) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *SevencluesBoardView) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *SevencluesBoardView) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *SevencluesBoardView) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetScore

`func (o *SevencluesBoardView) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SevencluesBoardView) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SevencluesBoardView) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SevencluesBoardView) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetState

`func (o *SevencluesBoardView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SevencluesBoardView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SevencluesBoardView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SevencluesBoardView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTeamScores

`func (o *SevencluesBoardView) GetTeamScores() []int32`

GetTeamScores returns the TeamScores field if non-nil, zero value otherwise.

### GetTeamScoresOk

`func (o *SevencluesBoardView) GetTeamScoresOk() (*[]int32, bool)`

GetTeamScoresOk returns a tuple with the TeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScores

`func (o *SevencluesBoardView) SetTeamScores(v []int32)`

SetTeamScores sets TeamScores field to given value.

### HasTeamScores

`func (o *SevencluesBoardView) HasTeamScores() bool`

HasTeamScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


