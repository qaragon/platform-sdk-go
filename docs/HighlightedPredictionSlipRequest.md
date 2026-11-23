# HighlightedPredictionSlipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | **time.Time** | Exclusive kickoff deadline, in the future and at most 25 hours away to allow a local day across daylight saving changes. | 

## Methods

### NewHighlightedPredictionSlipRequest

`func NewHighlightedPredictionSlipRequest(before time.Time, ) *HighlightedPredictionSlipRequest`

NewHighlightedPredictionSlipRequest instantiates a new HighlightedPredictionSlipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightedPredictionSlipRequestWithDefaults

`func NewHighlightedPredictionSlipRequestWithDefaults() *HighlightedPredictionSlipRequest`

NewHighlightedPredictionSlipRequestWithDefaults instantiates a new HighlightedPredictionSlipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *HighlightedPredictionSlipRequest) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *HighlightedPredictionSlipRequest) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *HighlightedPredictionSlipRequest) SetBefore(v time.Time)`

SetBefore sets Before field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


