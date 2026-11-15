# SubmissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **map[string]interface{}** | Type-specific input: solutions for MCQ/Prediction, option_ids for Vote/Game, key for CheckIn, guess for Guess, or media_url/media_type/duration for Creative. | [optional] 
**PlatformIdentityId** | Pointer to **string** |  | [optional] 
**SpaceId** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**SubmissionTarget**](SubmissionTarget.md) |  | [optional] 

## Methods

### NewSubmissionRequest

`func NewSubmissionRequest() *SubmissionRequest`

NewSubmissionRequest instantiates a new SubmissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionRequestWithDefaults

`func NewSubmissionRequestWithDefaults() *SubmissionRequest`

NewSubmissionRequestWithDefaults instantiates a new SubmissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *SubmissionRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SubmissionRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SubmissionRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SubmissionRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPlatformIdentityId

`func (o *SubmissionRequest) GetPlatformIdentityId() string`

GetPlatformIdentityId returns the PlatformIdentityId field if non-nil, zero value otherwise.

### GetPlatformIdentityIdOk

`func (o *SubmissionRequest) GetPlatformIdentityIdOk() (*string, bool)`

GetPlatformIdentityIdOk returns a tuple with the PlatformIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentityId

`func (o *SubmissionRequest) SetPlatformIdentityId(v string)`

SetPlatformIdentityId sets PlatformIdentityId field to given value.

### HasPlatformIdentityId

`func (o *SubmissionRequest) HasPlatformIdentityId() bool`

HasPlatformIdentityId returns a boolean if a field has been set.

### GetSpaceId

`func (o *SubmissionRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SubmissionRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SubmissionRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *SubmissionRequest) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetTarget

`func (o *SubmissionRequest) GetTarget() SubmissionTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SubmissionRequest) GetTargetOk() (*SubmissionTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SubmissionRequest) SetTarget(v SubmissionTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SubmissionRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


