# FeudshowdownRoundView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ActiveIdentityId** | Pointer to **string** |  | [optional] 
**Answers** | Pointer to [**[]FeudshowdownAnswerView**](FeudshowdownAnswerView.md) |  | [optional] 
**Award** | Pointer to **int32** |  | [optional] 
**ChoiceIdentityId** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**FaceoffIdentityIds** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Multiplier** | Pointer to **int32** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Pot** | Pointer to **int32** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Strikes** | Pointer to **int32** |  | [optional] 
**WinnerTeam** | Pointer to **int32** |  | [optional] 

## Methods

### NewFeudshowdownRoundView

`func NewFeudshowdownRoundView() *FeudshowdownRoundView`

NewFeudshowdownRoundView instantiates a new FeudshowdownRoundView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeudshowdownRoundViewWithDefaults

`func NewFeudshowdownRoundViewWithDefaults() *FeudshowdownRoundView`

NewFeudshowdownRoundViewWithDefaults instantiates a new FeudshowdownRoundView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionVersion

`func (o *FeudshowdownRoundView) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *FeudshowdownRoundView) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *FeudshowdownRoundView) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *FeudshowdownRoundView) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetActiveIdentityId

`func (o *FeudshowdownRoundView) GetActiveIdentityId() string`

GetActiveIdentityId returns the ActiveIdentityId field if non-nil, zero value otherwise.

### GetActiveIdentityIdOk

`func (o *FeudshowdownRoundView) GetActiveIdentityIdOk() (*string, bool)`

GetActiveIdentityIdOk returns a tuple with the ActiveIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentityId

`func (o *FeudshowdownRoundView) SetActiveIdentityId(v string)`

SetActiveIdentityId sets ActiveIdentityId field to given value.

### HasActiveIdentityId

`func (o *FeudshowdownRoundView) HasActiveIdentityId() bool`

HasActiveIdentityId returns a boolean if a field has been set.

### GetAnswers

`func (o *FeudshowdownRoundView) GetAnswers() []FeudshowdownAnswerView`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *FeudshowdownRoundView) GetAnswersOk() (*[]FeudshowdownAnswerView, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *FeudshowdownRoundView) SetAnswers(v []FeudshowdownAnswerView)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *FeudshowdownRoundView) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAward

`func (o *FeudshowdownRoundView) GetAward() int32`

GetAward returns the Award field if non-nil, zero value otherwise.

### GetAwardOk

`func (o *FeudshowdownRoundView) GetAwardOk() (*int32, bool)`

GetAwardOk returns a tuple with the Award field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAward

`func (o *FeudshowdownRoundView) SetAward(v int32)`

SetAward sets Award field to given value.

### HasAward

`func (o *FeudshowdownRoundView) HasAward() bool`

HasAward returns a boolean if a field has been set.

### GetChoiceIdentityId

`func (o *FeudshowdownRoundView) GetChoiceIdentityId() string`

GetChoiceIdentityId returns the ChoiceIdentityId field if non-nil, zero value otherwise.

### GetChoiceIdentityIdOk

`func (o *FeudshowdownRoundView) GetChoiceIdentityIdOk() (*string, bool)`

GetChoiceIdentityIdOk returns a tuple with the ChoiceIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceIdentityId

`func (o *FeudshowdownRoundView) SetChoiceIdentityId(v string)`

SetChoiceIdentityId sets ChoiceIdentityId field to given value.

### HasChoiceIdentityId

`func (o *FeudshowdownRoundView) HasChoiceIdentityId() bool`

HasChoiceIdentityId returns a boolean if a field has been set.

### GetDeadline

`func (o *FeudshowdownRoundView) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *FeudshowdownRoundView) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *FeudshowdownRoundView) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *FeudshowdownRoundView) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetFaceoffIdentityIds

`func (o *FeudshowdownRoundView) GetFaceoffIdentityIds() []string`

GetFaceoffIdentityIds returns the FaceoffIdentityIds field if non-nil, zero value otherwise.

### GetFaceoffIdentityIdsOk

`func (o *FeudshowdownRoundView) GetFaceoffIdentityIdsOk() (*[]string, bool)`

GetFaceoffIdentityIdsOk returns a tuple with the FaceoffIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceoffIdentityIds

`func (o *FeudshowdownRoundView) SetFaceoffIdentityIds(v []string)`

SetFaceoffIdentityIds sets FaceoffIdentityIds field to given value.

### HasFaceoffIdentityIds

`func (o *FeudshowdownRoundView) HasFaceoffIdentityIds() bool`

HasFaceoffIdentityIds returns a boolean if a field has been set.

### GetId

`func (o *FeudshowdownRoundView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeudshowdownRoundView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeudshowdownRoundView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeudshowdownRoundView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *FeudshowdownRoundView) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FeudshowdownRoundView) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FeudshowdownRoundView) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FeudshowdownRoundView) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMultiplier

`func (o *FeudshowdownRoundView) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *FeudshowdownRoundView) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *FeudshowdownRoundView) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *FeudshowdownRoundView) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetPhase

`func (o *FeudshowdownRoundView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *FeudshowdownRoundView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *FeudshowdownRoundView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *FeudshowdownRoundView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPot

`func (o *FeudshowdownRoundView) GetPot() int32`

GetPot returns the Pot field if non-nil, zero value otherwise.

### GetPotOk

`func (o *FeudshowdownRoundView) GetPotOk() (*int32, bool)`

GetPotOk returns a tuple with the Pot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPot

`func (o *FeudshowdownRoundView) SetPot(v int32)`

SetPot sets Pot field to given value.

### HasPot

`func (o *FeudshowdownRoundView) HasPot() bool`

HasPot returns a boolean if a field has been set.

### GetPrompt

`func (o *FeudshowdownRoundView) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FeudshowdownRoundView) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FeudshowdownRoundView) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *FeudshowdownRoundView) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetStrikes

`func (o *FeudshowdownRoundView) GetStrikes() int32`

GetStrikes returns the Strikes field if non-nil, zero value otherwise.

### GetStrikesOk

`func (o *FeudshowdownRoundView) GetStrikesOk() (*int32, bool)`

GetStrikesOk returns a tuple with the Strikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikes

`func (o *FeudshowdownRoundView) SetStrikes(v int32)`

SetStrikes sets Strikes field to given value.

### HasStrikes

`func (o *FeudshowdownRoundView) HasStrikes() bool`

HasStrikes returns a boolean if a field has been set.

### GetWinnerTeam

`func (o *FeudshowdownRoundView) GetWinnerTeam() int32`

GetWinnerTeam returns the WinnerTeam field if non-nil, zero value otherwise.

### GetWinnerTeamOk

`func (o *FeudshowdownRoundView) GetWinnerTeamOk() (*int32, bool)`

GetWinnerTeamOk returns a tuple with the WinnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerTeam

`func (o *FeudshowdownRoundView) SetWinnerTeam(v int32)`

SetWinnerTeam sets WinnerTeam field to given value.

### HasWinnerTeam

`func (o *FeudshowdownRoundView) HasWinnerTeam() bool`

HasWinnerTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


