# HighlightedPredictionSlipResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of events in the prepared slip. | 
**Deadline** | **time.Time** | Submission deadline: the selected event&#39;s kickoff. | 
**Id** | **string** |  | 

## Methods

### NewHighlightedPredictionSlipResponseData

`func NewHighlightedPredictionSlipResponseData(count int32, deadline time.Time, id string, ) *HighlightedPredictionSlipResponseData`

NewHighlightedPredictionSlipResponseData instantiates a new HighlightedPredictionSlipResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightedPredictionSlipResponseDataWithDefaults

`func NewHighlightedPredictionSlipResponseDataWithDefaults() *HighlightedPredictionSlipResponseData`

NewHighlightedPredictionSlipResponseDataWithDefaults instantiates a new HighlightedPredictionSlipResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HighlightedPredictionSlipResponseData) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HighlightedPredictionSlipResponseData) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HighlightedPredictionSlipResponseData) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDeadline

`func (o *HighlightedPredictionSlipResponseData) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *HighlightedPredictionSlipResponseData) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *HighlightedPredictionSlipResponseData) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetId

`func (o *HighlightedPredictionSlipResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HighlightedPredictionSlipResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HighlightedPredictionSlipResponseData) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


