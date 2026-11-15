# FeudrushView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ActiveIdentityId** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**FinalistIdentityIds** | Pointer to **[]string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**PromptId** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to [**[]FeudrushAttemptView**](FeudrushAttemptView.md) |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewFeudrushView

`func NewFeudrushView() *FeudrushView`

NewFeudrushView instantiates a new FeudrushView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeudrushViewWithDefaults

`func NewFeudrushViewWithDefaults() *FeudrushView`

NewFeudrushViewWithDefaults instantiates a new FeudrushView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionVersion

`func (o *FeudrushView) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *FeudrushView) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *FeudrushView) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *FeudrushView) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetActiveIdentityId

`func (o *FeudrushView) GetActiveIdentityId() string`

GetActiveIdentityId returns the ActiveIdentityId field if non-nil, zero value otherwise.

### GetActiveIdentityIdOk

`func (o *FeudrushView) GetActiveIdentityIdOk() (*string, bool)`

GetActiveIdentityIdOk returns a tuple with the ActiveIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentityId

`func (o *FeudrushView) SetActiveIdentityId(v string)`

SetActiveIdentityId sets ActiveIdentityId field to given value.

### HasActiveIdentityId

`func (o *FeudrushView) HasActiveIdentityId() bool`

HasActiveIdentityId returns a boolean if a field has been set.

### GetDeadline

`func (o *FeudrushView) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *FeudrushView) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *FeudrushView) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *FeudrushView) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetFinalistIdentityIds

`func (o *FeudrushView) GetFinalistIdentityIds() []string`

GetFinalistIdentityIds returns the FinalistIdentityIds field if non-nil, zero value otherwise.

### GetFinalistIdentityIdsOk

`func (o *FeudrushView) GetFinalistIdentityIdsOk() (*[]string, bool)`

GetFinalistIdentityIdsOk returns a tuple with the FinalistIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalistIdentityIds

`func (o *FeudrushView) SetFinalistIdentityIds(v []string)`

SetFinalistIdentityIds sets FinalistIdentityIds field to given value.

### HasFinalistIdentityIds

`func (o *FeudrushView) HasFinalistIdentityIds() bool`

HasFinalistIdentityIds returns a boolean if a field has been set.

### GetPrompt

`func (o *FeudrushView) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FeudrushView) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FeudrushView) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *FeudrushView) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetPromptId

`func (o *FeudrushView) GetPromptId() string`

GetPromptId returns the PromptId field if non-nil, zero value otherwise.

### GetPromptIdOk

`func (o *FeudrushView) GetPromptIdOk() (*string, bool)`

GetPromptIdOk returns a tuple with the PromptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptId

`func (o *FeudrushView) SetPromptId(v string)`

SetPromptId sets PromptId field to given value.

### HasPromptId

`func (o *FeudrushView) HasPromptId() bool`

HasPromptId returns a boolean if a field has been set.

### GetRuns

`func (o *FeudrushView) GetRuns() []FeudrushAttemptView`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *FeudrushView) GetRunsOk() (*[]FeudrushAttemptView, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *FeudrushView) SetRuns(v []FeudrushAttemptView)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *FeudrushView) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetStageId

`func (o *FeudrushView) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *FeudrushView) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *FeudrushView) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *FeudrushView) HasStageId() bool`

HasStageId returns a boolean if a field has been set.

### GetState

`func (o *FeudrushView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FeudrushView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FeudrushView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FeudrushView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTarget

`func (o *FeudrushView) GetTarget() int32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FeudrushView) GetTargetOk() (*int32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FeudrushView) SetTarget(v int32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *FeudrushView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTotal

`func (o *FeudrushView) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FeudrushView) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FeudrushView) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FeudrushView) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


