# EventHighlightsGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Refresh** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEventHighlightsGenerationRequest

`func NewEventHighlightsGenerationRequest() *EventHighlightsGenerationRequest`

NewEventHighlightsGenerationRequest instantiates a new EventHighlightsGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHighlightsGenerationRequestWithDefaults

`func NewEventHighlightsGenerationRequestWithDefaults() *EventHighlightsGenerationRequest`

NewEventHighlightsGenerationRequestWithDefaults instantiates a new EventHighlightsGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EventHighlightsGenerationRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EventHighlightsGenerationRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EventHighlightsGenerationRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EventHighlightsGenerationRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetRefresh

`func (o *EventHighlightsGenerationRequest) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *EventHighlightsGenerationRequest) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *EventHighlightsGenerationRequest) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *EventHighlightsGenerationRequest) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


