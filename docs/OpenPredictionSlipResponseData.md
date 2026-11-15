# OpenPredictionSlipResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CuratedTags** | Pointer to **[]string** |  | [optional] 
**Events** | Pointer to [**[]PredictionEvent**](PredictionEvent.md) |  | [optional] 
**Id** | **string** |  | 
**IsCurated** | Pointer to **NullableBool** |  | [optional] 
**IsPublic** | Pointer to **NullableBool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOpenPredictionSlipResponseData

`func NewOpenPredictionSlipResponseData(id string, name string, ) *OpenPredictionSlipResponseData`

NewOpenPredictionSlipResponseData instantiates a new OpenPredictionSlipResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPredictionSlipResponseDataWithDefaults

`func NewOpenPredictionSlipResponseDataWithDefaults() *OpenPredictionSlipResponseData`

NewOpenPredictionSlipResponseDataWithDefaults instantiates a new OpenPredictionSlipResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OpenPredictionSlipResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenPredictionSlipResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenPredictionSlipResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpenPredictionSlipResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCuratedTags

`func (o *OpenPredictionSlipResponseData) GetCuratedTags() []string`

GetCuratedTags returns the CuratedTags field if non-nil, zero value otherwise.

### GetCuratedTagsOk

`func (o *OpenPredictionSlipResponseData) GetCuratedTagsOk() (*[]string, bool)`

GetCuratedTagsOk returns a tuple with the CuratedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuratedTags

`func (o *OpenPredictionSlipResponseData) SetCuratedTags(v []string)`

SetCuratedTags sets CuratedTags field to given value.

### HasCuratedTags

`func (o *OpenPredictionSlipResponseData) HasCuratedTags() bool`

HasCuratedTags returns a boolean if a field has been set.

### GetEvents

`func (o *OpenPredictionSlipResponseData) GetEvents() []PredictionEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OpenPredictionSlipResponseData) GetEventsOk() (*[]PredictionEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OpenPredictionSlipResponseData) SetEvents(v []PredictionEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OpenPredictionSlipResponseData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *OpenPredictionSlipResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenPredictionSlipResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenPredictionSlipResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetIsCurated

`func (o *OpenPredictionSlipResponseData) GetIsCurated() bool`

GetIsCurated returns the IsCurated field if non-nil, zero value otherwise.

### GetIsCuratedOk

`func (o *OpenPredictionSlipResponseData) GetIsCuratedOk() (*bool, bool)`

GetIsCuratedOk returns a tuple with the IsCurated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurated

`func (o *OpenPredictionSlipResponseData) SetIsCurated(v bool)`

SetIsCurated sets IsCurated field to given value.

### HasIsCurated

`func (o *OpenPredictionSlipResponseData) HasIsCurated() bool`

HasIsCurated returns a boolean if a field has been set.

### SetIsCuratedNil

`func (o *OpenPredictionSlipResponseData) SetIsCuratedNil(b bool)`

 SetIsCuratedNil sets the value for IsCurated to be an explicit nil

### UnsetIsCurated
`func (o *OpenPredictionSlipResponseData) UnsetIsCurated()`

UnsetIsCurated ensures that no value is present for IsCurated, not even an explicit nil
### GetIsPublic

`func (o *OpenPredictionSlipResponseData) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *OpenPredictionSlipResponseData) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *OpenPredictionSlipResponseData) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *OpenPredictionSlipResponseData) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### SetIsPublicNil

`func (o *OpenPredictionSlipResponseData) SetIsPublicNil(b bool)`

 SetIsPublicNil sets the value for IsPublic to be an explicit nil

### UnsetIsPublic
`func (o *OpenPredictionSlipResponseData) UnsetIsPublic()`

UnsetIsPublic ensures that no value is present for IsPublic, not even an explicit nil
### GetLocked

`func (o *OpenPredictionSlipResponseData) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *OpenPredictionSlipResponseData) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *OpenPredictionSlipResponseData) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *OpenPredictionSlipResponseData) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *OpenPredictionSlipResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenPredictionSlipResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenPredictionSlipResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *OpenPredictionSlipResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpenPredictionSlipResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpenPredictionSlipResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpenPredictionSlipResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


