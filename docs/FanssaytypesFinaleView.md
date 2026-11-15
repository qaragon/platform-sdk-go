# FanssaytypesFinaleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ActiveIdentityId** | Pointer to **string** |  | [optional] 
**CurrentPromptId** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**PromptIndex** | Pointer to **int32** |  | [optional] 
**Prompts** | Pointer to [**[]TeninsixtyPromptView**](TeninsixtyPromptView.md) |  | [optional] 
**RemainingMs** | Pointer to **int32** |  | [optional] 
**SkipAvailable** | Pointer to **bool** |  | [optional] 
**SkipUsed** | Pointer to **bool** |  | [optional] 
**Stage** | Pointer to [**PlayStage**](PlayStage.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Team** | Pointer to **int32** |  | [optional] 

## Methods

### NewFanssaytypesFinaleView

`func NewFanssaytypesFinaleView() *FanssaytypesFinaleView`

NewFanssaytypesFinaleView instantiates a new FanssaytypesFinaleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanssaytypesFinaleViewWithDefaults

`func NewFanssaytypesFinaleViewWithDefaults() *FanssaytypesFinaleView`

NewFanssaytypesFinaleViewWithDefaults instantiates a new FanssaytypesFinaleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionVersion

`func (o *FanssaytypesFinaleView) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *FanssaytypesFinaleView) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *FanssaytypesFinaleView) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *FanssaytypesFinaleView) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetActiveIdentityId

`func (o *FanssaytypesFinaleView) GetActiveIdentityId() string`

GetActiveIdentityId returns the ActiveIdentityId field if non-nil, zero value otherwise.

### GetActiveIdentityIdOk

`func (o *FanssaytypesFinaleView) GetActiveIdentityIdOk() (*string, bool)`

GetActiveIdentityIdOk returns a tuple with the ActiveIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentityId

`func (o *FanssaytypesFinaleView) SetActiveIdentityId(v string)`

SetActiveIdentityId sets ActiveIdentityId field to given value.

### HasActiveIdentityId

`func (o *FanssaytypesFinaleView) HasActiveIdentityId() bool`

HasActiveIdentityId returns a boolean if a field has been set.

### GetCurrentPromptId

`func (o *FanssaytypesFinaleView) GetCurrentPromptId() string`

GetCurrentPromptId returns the CurrentPromptId field if non-nil, zero value otherwise.

### GetCurrentPromptIdOk

`func (o *FanssaytypesFinaleView) GetCurrentPromptIdOk() (*string, bool)`

GetCurrentPromptIdOk returns a tuple with the CurrentPromptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPromptId

`func (o *FanssaytypesFinaleView) SetCurrentPromptId(v string)`

SetCurrentPromptId sets CurrentPromptId field to given value.

### HasCurrentPromptId

`func (o *FanssaytypesFinaleView) HasCurrentPromptId() bool`

HasCurrentPromptId returns a boolean if a field has been set.

### GetDeadline

`func (o *FanssaytypesFinaleView) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *FanssaytypesFinaleView) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *FanssaytypesFinaleView) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *FanssaytypesFinaleView) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetPromptIndex

`func (o *FanssaytypesFinaleView) GetPromptIndex() int32`

GetPromptIndex returns the PromptIndex field if non-nil, zero value otherwise.

### GetPromptIndexOk

`func (o *FanssaytypesFinaleView) GetPromptIndexOk() (*int32, bool)`

GetPromptIndexOk returns a tuple with the PromptIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptIndex

`func (o *FanssaytypesFinaleView) SetPromptIndex(v int32)`

SetPromptIndex sets PromptIndex field to given value.

### HasPromptIndex

`func (o *FanssaytypesFinaleView) HasPromptIndex() bool`

HasPromptIndex returns a boolean if a field has been set.

### GetPrompts

`func (o *FanssaytypesFinaleView) GetPrompts() []TeninsixtyPromptView`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *FanssaytypesFinaleView) GetPromptsOk() (*[]TeninsixtyPromptView, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *FanssaytypesFinaleView) SetPrompts(v []TeninsixtyPromptView)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *FanssaytypesFinaleView) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetRemainingMs

`func (o *FanssaytypesFinaleView) GetRemainingMs() int32`

GetRemainingMs returns the RemainingMs field if non-nil, zero value otherwise.

### GetRemainingMsOk

`func (o *FanssaytypesFinaleView) GetRemainingMsOk() (*int32, bool)`

GetRemainingMsOk returns a tuple with the RemainingMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMs

`func (o *FanssaytypesFinaleView) SetRemainingMs(v int32)`

SetRemainingMs sets RemainingMs field to given value.

### HasRemainingMs

`func (o *FanssaytypesFinaleView) HasRemainingMs() bool`

HasRemainingMs returns a boolean if a field has been set.

### GetSkipAvailable

`func (o *FanssaytypesFinaleView) GetSkipAvailable() bool`

GetSkipAvailable returns the SkipAvailable field if non-nil, zero value otherwise.

### GetSkipAvailableOk

`func (o *FanssaytypesFinaleView) GetSkipAvailableOk() (*bool, bool)`

GetSkipAvailableOk returns a tuple with the SkipAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAvailable

`func (o *FanssaytypesFinaleView) SetSkipAvailable(v bool)`

SetSkipAvailable sets SkipAvailable field to given value.

### HasSkipAvailable

`func (o *FanssaytypesFinaleView) HasSkipAvailable() bool`

HasSkipAvailable returns a boolean if a field has been set.

### GetSkipUsed

`func (o *FanssaytypesFinaleView) GetSkipUsed() bool`

GetSkipUsed returns the SkipUsed field if non-nil, zero value otherwise.

### GetSkipUsedOk

`func (o *FanssaytypesFinaleView) GetSkipUsedOk() (*bool, bool)`

GetSkipUsedOk returns a tuple with the SkipUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUsed

`func (o *FanssaytypesFinaleView) SetSkipUsed(v bool)`

SetSkipUsed sets SkipUsed field to given value.

### HasSkipUsed

`func (o *FanssaytypesFinaleView) HasSkipUsed() bool`

HasSkipUsed returns a boolean if a field has been set.

### GetStage

`func (o *FanssaytypesFinaleView) GetStage() PlayStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FanssaytypesFinaleView) GetStageOk() (*PlayStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FanssaytypesFinaleView) SetStage(v PlayStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *FanssaytypesFinaleView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *FanssaytypesFinaleView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FanssaytypesFinaleView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FanssaytypesFinaleView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FanssaytypesFinaleView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTeam

`func (o *FanssaytypesFinaleView) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *FanssaytypesFinaleView) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *FanssaytypesFinaleView) SetTeam(v int32)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *FanssaytypesFinaleView) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


