# SubmissionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | Pointer to **string** |  | [optional] 
**FailureCode** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** | Complete normalized Guess or Game input. The target and action determine its shape; see submission.Input and wordchain.PlayerRequest. Preserve unknown action fields when decoding attempt history. | [optional] 
**PlatformIdentityId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**SubmissionResult**](SubmissionResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubmissionId** | Pointer to **string** |  | [optional] 
**SubmissionNumber** | Pointer to **int32** |  | [optional] 
**Target** | Pointer to [**SubmissionTarget**](SubmissionTarget.md) |  | [optional] 

## Methods

### NewSubmissionView

`func NewSubmissionView() *SubmissionView`

NewSubmissionView instantiates a new SubmissionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionViewWithDefaults

`func NewSubmissionViewWithDefaults() *SubmissionView`

NewSubmissionViewWithDefaults instantiates a new SubmissionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *SubmissionView) GetAcceptedAt() string`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *SubmissionView) GetAcceptedAtOk() (*string, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *SubmissionView) SetAcceptedAt(v string)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *SubmissionView) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetFailureCode

`func (o *SubmissionView) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *SubmissionView) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *SubmissionView) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *SubmissionView) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetPayload

`func (o *SubmissionView) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SubmissionView) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SubmissionView) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SubmissionView) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPlatformIdentityId

`func (o *SubmissionView) GetPlatformIdentityId() string`

GetPlatformIdentityId returns the PlatformIdentityId field if non-nil, zero value otherwise.

### GetPlatformIdentityIdOk

`func (o *SubmissionView) GetPlatformIdentityIdOk() (*string, bool)`

GetPlatformIdentityIdOk returns a tuple with the PlatformIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentityId

`func (o *SubmissionView) SetPlatformIdentityId(v string)`

SetPlatformIdentityId sets PlatformIdentityId field to given value.

### HasPlatformIdentityId

`func (o *SubmissionView) HasPlatformIdentityId() bool`

HasPlatformIdentityId returns a boolean if a field has been set.

### GetResult

`func (o *SubmissionView) GetResult() SubmissionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubmissionView) GetResultOk() (*SubmissionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubmissionView) SetResult(v SubmissionResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SubmissionView) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *SubmissionView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmissionId

`func (o *SubmissionView) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SubmissionView) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SubmissionView) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *SubmissionView) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.

### GetSubmissionNumber

`func (o *SubmissionView) GetSubmissionNumber() int32`

GetSubmissionNumber returns the SubmissionNumber field if non-nil, zero value otherwise.

### GetSubmissionNumberOk

`func (o *SubmissionView) GetSubmissionNumberOk() (*int32, bool)`

GetSubmissionNumberOk returns a tuple with the SubmissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionNumber

`func (o *SubmissionView) SetSubmissionNumber(v int32)`

SetSubmissionNumber sets SubmissionNumber field to given value.

### HasSubmissionNumber

`func (o *SubmissionView) HasSubmissionNumber() bool`

HasSubmissionNumber returns a boolean if a field has been set.

### GetTarget

`func (o *SubmissionView) GetTarget() SubmissionTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SubmissionView) GetTargetOk() (*SubmissionTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SubmissionView) SetTarget(v SubmissionTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SubmissionView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


