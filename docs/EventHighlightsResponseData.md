# EventHighlightsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]EventHighlightsResponseDataItemsInner**](EventHighlightsResponseDataItemsInner.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewEventHighlightsResponseData

`func NewEventHighlightsResponseData() *EventHighlightsResponseData`

NewEventHighlightsResponseData instantiates a new EventHighlightsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHighlightsResponseDataWithDefaults

`func NewEventHighlightsResponseDataWithDefaults() *EventHighlightsResponseData`

NewEventHighlightsResponseDataWithDefaults instantiates a new EventHighlightsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EventHighlightsResponseData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EventHighlightsResponseData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EventHighlightsResponseData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EventHighlightsResponseData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetItems

`func (o *EventHighlightsResponseData) GetItems() []EventHighlightsResponseDataItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventHighlightsResponseData) GetItemsOk() (*[]EventHighlightsResponseDataItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventHighlightsResponseData) SetItems(v []EventHighlightsResponseDataItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *EventHighlightsResponseData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKey

`func (o *EventHighlightsResponseData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventHighlightsResponseData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventHighlightsResponseData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventHighlightsResponseData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTitle

`func (o *EventHighlightsResponseData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventHighlightsResponseData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventHighlightsResponseData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventHighlightsResponseData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


