# GuessGuessView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptLimit** | Pointer to **int32** |  | [optional] 
**BoardId** | Pointer to **string** |  | [optional] 
**ClaimScope** | Pointer to **string** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**ClueStrategy** | Pointer to **string** |  | [optional] 
**MatchingRule** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Slots** | Pointer to [**[]GuessSlotView**](GuessSlotView.md) |  | [optional] 
**SubmissionPolicy** | Pointer to **string** |  | [optional] 
**Submitted** | Pointer to **bool** |  | [optional] 

## Methods

### NewGuessGuessView

`func NewGuessGuessView() *GuessGuessView`

NewGuessGuessView instantiates a new GuessGuessView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuessGuessViewWithDefaults

`func NewGuessGuessViewWithDefaults() *GuessGuessView`

NewGuessGuessViewWithDefaults instantiates a new GuessGuessView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptLimit

`func (o *GuessGuessView) GetAttemptLimit() int32`

GetAttemptLimit returns the AttemptLimit field if non-nil, zero value otherwise.

### GetAttemptLimitOk

`func (o *GuessGuessView) GetAttemptLimitOk() (*int32, bool)`

GetAttemptLimitOk returns a tuple with the AttemptLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptLimit

`func (o *GuessGuessView) SetAttemptLimit(v int32)`

SetAttemptLimit sets AttemptLimit field to given value.

### HasAttemptLimit

`func (o *GuessGuessView) HasAttemptLimit() bool`

HasAttemptLimit returns a boolean if a field has been set.

### GetBoardId

`func (o *GuessGuessView) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *GuessGuessView) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *GuessGuessView) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *GuessGuessView) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetClaimScope

`func (o *GuessGuessView) GetClaimScope() string`

GetClaimScope returns the ClaimScope field if non-nil, zero value otherwise.

### GetClaimScopeOk

`func (o *GuessGuessView) GetClaimScopeOk() (*string, bool)`

GetClaimScopeOk returns a tuple with the ClaimScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimScope

`func (o *GuessGuessView) SetClaimScope(v string)`

SetClaimScope sets ClaimScope field to given value.

### HasClaimScope

`func (o *GuessGuessView) HasClaimScope() bool`

HasClaimScope returns a boolean if a field has been set.

### GetClosed

`func (o *GuessGuessView) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *GuessGuessView) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *GuessGuessView) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *GuessGuessView) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClueStrategy

`func (o *GuessGuessView) GetClueStrategy() string`

GetClueStrategy returns the ClueStrategy field if non-nil, zero value otherwise.

### GetClueStrategyOk

`func (o *GuessGuessView) GetClueStrategyOk() (*string, bool)`

GetClueStrategyOk returns a tuple with the ClueStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClueStrategy

`func (o *GuessGuessView) SetClueStrategy(v string)`

SetClueStrategy sets ClueStrategy field to given value.

### HasClueStrategy

`func (o *GuessGuessView) HasClueStrategy() bool`

HasClueStrategy returns a boolean if a field has been set.

### GetMatchingRule

`func (o *GuessGuessView) GetMatchingRule() string`

GetMatchingRule returns the MatchingRule field if non-nil, zero value otherwise.

### GetMatchingRuleOk

`func (o *GuessGuessView) GetMatchingRuleOk() (*string, bool)`

GetMatchingRuleOk returns a tuple with the MatchingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRule

`func (o *GuessGuessView) SetMatchingRule(v string)`

SetMatchingRule sets MatchingRule field to given value.

### HasMatchingRule

`func (o *GuessGuessView) HasMatchingRule() bool`

HasMatchingRule returns a boolean if a field has been set.

### GetPrompt

`func (o *GuessGuessView) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GuessGuessView) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GuessGuessView) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *GuessGuessView) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetSlots

`func (o *GuessGuessView) GetSlots() []GuessSlotView`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *GuessGuessView) GetSlotsOk() (*[]GuessSlotView, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *GuessGuessView) SetSlots(v []GuessSlotView)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *GuessGuessView) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetSubmissionPolicy

`func (o *GuessGuessView) GetSubmissionPolicy() string`

GetSubmissionPolicy returns the SubmissionPolicy field if non-nil, zero value otherwise.

### GetSubmissionPolicyOk

`func (o *GuessGuessView) GetSubmissionPolicyOk() (*string, bool)`

GetSubmissionPolicyOk returns a tuple with the SubmissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionPolicy

`func (o *GuessGuessView) SetSubmissionPolicy(v string)`

SetSubmissionPolicy sets SubmissionPolicy field to given value.

### HasSubmissionPolicy

`func (o *GuessGuessView) HasSubmissionPolicy() bool`

HasSubmissionPolicy returns a boolean if a field has been set.

### GetSubmitted

`func (o *GuessGuessView) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *GuessGuessView) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *GuessGuessView) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *GuessGuessView) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


