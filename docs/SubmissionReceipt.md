# SubmissionReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | Pointer to **string** |  | [optional] 
**FailureCode** | Pointer to **string** |  | [optional] 
**PlatformIdentityId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**SubmissionResult**](SubmissionResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubmissionId** | Pointer to **string** |  | [optional] 
**SubmissionNumber** | Pointer to **int32** |  | [optional] 
**Target** | Pointer to [**SubmissionTarget**](SubmissionTarget.md) |  | [optional] 

## Methods

### NewSubmissionReceipt

`func NewSubmissionReceipt() *SubmissionReceipt`

NewSubmissionReceipt instantiates a new SubmissionReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionReceiptWithDefaults

`func NewSubmissionReceiptWithDefaults() *SubmissionReceipt`

NewSubmissionReceiptWithDefaults instantiates a new SubmissionReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *SubmissionReceipt) GetAcceptedAt() string`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *SubmissionReceipt) GetAcceptedAtOk() (*string, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *SubmissionReceipt) SetAcceptedAt(v string)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *SubmissionReceipt) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetFailureCode

`func (o *SubmissionReceipt) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *SubmissionReceipt) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *SubmissionReceipt) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *SubmissionReceipt) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetPlatformIdentityId

`func (o *SubmissionReceipt) GetPlatformIdentityId() string`

GetPlatformIdentityId returns the PlatformIdentityId field if non-nil, zero value otherwise.

### GetPlatformIdentityIdOk

`func (o *SubmissionReceipt) GetPlatformIdentityIdOk() (*string, bool)`

GetPlatformIdentityIdOk returns a tuple with the PlatformIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentityId

`func (o *SubmissionReceipt) SetPlatformIdentityId(v string)`

SetPlatformIdentityId sets PlatformIdentityId field to given value.

### HasPlatformIdentityId

`func (o *SubmissionReceipt) HasPlatformIdentityId() bool`

HasPlatformIdentityId returns a boolean if a field has been set.

### GetResult

`func (o *SubmissionReceipt) GetResult() SubmissionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubmissionReceipt) GetResultOk() (*SubmissionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubmissionReceipt) SetResult(v SubmissionResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SubmissionReceipt) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *SubmissionReceipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionReceipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionReceipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionReceipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmissionId

`func (o *SubmissionReceipt) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SubmissionReceipt) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SubmissionReceipt) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *SubmissionReceipt) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.

### GetSubmissionNumber

`func (o *SubmissionReceipt) GetSubmissionNumber() int32`

GetSubmissionNumber returns the SubmissionNumber field if non-nil, zero value otherwise.

### GetSubmissionNumberOk

`func (o *SubmissionReceipt) GetSubmissionNumberOk() (*int32, bool)`

GetSubmissionNumberOk returns a tuple with the SubmissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionNumber

`func (o *SubmissionReceipt) SetSubmissionNumber(v int32)`

SetSubmissionNumber sets SubmissionNumber field to given value.

### HasSubmissionNumber

`func (o *SubmissionReceipt) HasSubmissionNumber() bool`

HasSubmissionNumber returns a boolean if a field has been set.

### GetTarget

`func (o *SubmissionReceipt) GetTarget() SubmissionTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SubmissionReceipt) GetTargetOk() (*SubmissionTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SubmissionReceipt) SetTarget(v SubmissionTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SubmissionReceipt) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


