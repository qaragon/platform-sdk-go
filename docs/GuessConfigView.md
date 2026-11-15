# GuessConfigView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]GuessGuessAnswer**](GuessGuessAnswer.md) |  | [optional] 
**AttemptLimit** | Pointer to **int32** |  | [optional] 
**BoardId** | Pointer to **string** |  | [optional] 
**ClaimScope** | Pointer to **string** |  | [optional] 
**ClueStrategy** | Pointer to **string** |  | [optional] 
**MatchingRule** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**SubmissionPolicy** | Pointer to **string** |  | [optional] 

## Methods

### NewGuessConfigView

`func NewGuessConfigView() *GuessConfigView`

NewGuessConfigView instantiates a new GuessConfigView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuessConfigViewWithDefaults

`func NewGuessConfigViewWithDefaults() *GuessConfigView`

NewGuessConfigViewWithDefaults instantiates a new GuessConfigView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *GuessConfigView) GetAnswers() []GuessGuessAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *GuessConfigView) GetAnswersOk() (*[]GuessGuessAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *GuessConfigView) SetAnswers(v []GuessGuessAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *GuessConfigView) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAttemptLimit

`func (o *GuessConfigView) GetAttemptLimit() int32`

GetAttemptLimit returns the AttemptLimit field if non-nil, zero value otherwise.

### GetAttemptLimitOk

`func (o *GuessConfigView) GetAttemptLimitOk() (*int32, bool)`

GetAttemptLimitOk returns a tuple with the AttemptLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptLimit

`func (o *GuessConfigView) SetAttemptLimit(v int32)`

SetAttemptLimit sets AttemptLimit field to given value.

### HasAttemptLimit

`func (o *GuessConfigView) HasAttemptLimit() bool`

HasAttemptLimit returns a boolean if a field has been set.

### GetBoardId

`func (o *GuessConfigView) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *GuessConfigView) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *GuessConfigView) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *GuessConfigView) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetClaimScope

`func (o *GuessConfigView) GetClaimScope() string`

GetClaimScope returns the ClaimScope field if non-nil, zero value otherwise.

### GetClaimScopeOk

`func (o *GuessConfigView) GetClaimScopeOk() (*string, bool)`

GetClaimScopeOk returns a tuple with the ClaimScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimScope

`func (o *GuessConfigView) SetClaimScope(v string)`

SetClaimScope sets ClaimScope field to given value.

### HasClaimScope

`func (o *GuessConfigView) HasClaimScope() bool`

HasClaimScope returns a boolean if a field has been set.

### GetClueStrategy

`func (o *GuessConfigView) GetClueStrategy() string`

GetClueStrategy returns the ClueStrategy field if non-nil, zero value otherwise.

### GetClueStrategyOk

`func (o *GuessConfigView) GetClueStrategyOk() (*string, bool)`

GetClueStrategyOk returns a tuple with the ClueStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClueStrategy

`func (o *GuessConfigView) SetClueStrategy(v string)`

SetClueStrategy sets ClueStrategy field to given value.

### HasClueStrategy

`func (o *GuessConfigView) HasClueStrategy() bool`

HasClueStrategy returns a boolean if a field has been set.

### GetMatchingRule

`func (o *GuessConfigView) GetMatchingRule() string`

GetMatchingRule returns the MatchingRule field if non-nil, zero value otherwise.

### GetMatchingRuleOk

`func (o *GuessConfigView) GetMatchingRuleOk() (*string, bool)`

GetMatchingRuleOk returns a tuple with the MatchingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRule

`func (o *GuessConfigView) SetMatchingRule(v string)`

SetMatchingRule sets MatchingRule field to given value.

### HasMatchingRule

`func (o *GuessConfigView) HasMatchingRule() bool`

HasMatchingRule returns a boolean if a field has been set.

### GetPrompt

`func (o *GuessConfigView) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GuessConfigView) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GuessConfigView) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *GuessConfigView) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetSubmissionPolicy

`func (o *GuessConfigView) GetSubmissionPolicy() string`

GetSubmissionPolicy returns the SubmissionPolicy field if non-nil, zero value otherwise.

### GetSubmissionPolicyOk

`func (o *GuessConfigView) GetSubmissionPolicyOk() (*string, bool)`

GetSubmissionPolicyOk returns a tuple with the SubmissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionPolicy

`func (o *GuessConfigView) SetSubmissionPolicy(v string)`

SetSubmissionPolicy sets SubmissionPolicy field to given value.

### HasSubmissionPolicy

`func (o *GuessConfigView) HasSubmissionPolicy() bool`

HasSubmissionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


