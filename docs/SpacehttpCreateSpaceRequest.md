# SpacehttpCreateSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCode** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**IsFanSpace** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Regions** | Pointer to [**[]SpacehttpRegion**](SpacehttpRegion.md) |  | [optional] 
**Size** | **int32** |  | 

## Methods

### NewSpacehttpCreateSpaceRequest

`func NewSpacehttpCreateSpaceRequest(name string, size int32, ) *SpacehttpCreateSpaceRequest`

NewSpacehttpCreateSpaceRequest instantiates a new SpacehttpCreateSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacehttpCreateSpaceRequestWithDefaults

`func NewSpacehttpCreateSpaceRequestWithDefaults() *SpacehttpCreateSpaceRequest`

NewSpacehttpCreateSpaceRequestWithDefaults instantiates a new SpacehttpCreateSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCode

`func (o *SpacehttpCreateSpaceRequest) GetAccessCode() string`

GetAccessCode returns the AccessCode field if non-nil, zero value otherwise.

### GetAccessCodeOk

`func (o *SpacehttpCreateSpaceRequest) GetAccessCodeOk() (*string, bool)`

GetAccessCodeOk returns a tuple with the AccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCode

`func (o *SpacehttpCreateSpaceRequest) SetAccessCode(v string)`

SetAccessCode sets AccessCode field to given value.

### HasAccessCode

`func (o *SpacehttpCreateSpaceRequest) HasAccessCode() bool`

HasAccessCode returns a boolean if a field has been set.

### GetAccessType

`func (o *SpacehttpCreateSpaceRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SpacehttpCreateSpaceRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SpacehttpCreateSpaceRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SpacehttpCreateSpaceRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetCode

`func (o *SpacehttpCreateSpaceRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpacehttpCreateSpaceRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpacehttpCreateSpaceRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SpacehttpCreateSpaceRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SpacehttpCreateSpaceRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SpacehttpCreateSpaceRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SpacehttpCreateSpaceRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SpacehttpCreateSpaceRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *SpacehttpCreateSpaceRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *SpacehttpCreateSpaceRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *SpacehttpCreateSpaceRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *SpacehttpCreateSpaceRequest) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetIsFanSpace

`func (o *SpacehttpCreateSpaceRequest) GetIsFanSpace() bool`

GetIsFanSpace returns the IsFanSpace field if non-nil, zero value otherwise.

### GetIsFanSpaceOk

`func (o *SpacehttpCreateSpaceRequest) GetIsFanSpaceOk() (*bool, bool)`

GetIsFanSpaceOk returns a tuple with the IsFanSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFanSpace

`func (o *SpacehttpCreateSpaceRequest) SetIsFanSpace(v bool)`

SetIsFanSpace sets IsFanSpace field to given value.

### HasIsFanSpace

`func (o *SpacehttpCreateSpaceRequest) HasIsFanSpace() bool`

HasIsFanSpace returns a boolean if a field has been set.

### GetName

`func (o *SpacehttpCreateSpaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpacehttpCreateSpaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpacehttpCreateSpaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegions

`func (o *SpacehttpCreateSpaceRequest) GetRegions() []SpacehttpRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SpacehttpCreateSpaceRequest) GetRegionsOk() (*[]SpacehttpRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SpacehttpCreateSpaceRequest) SetRegions(v []SpacehttpRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SpacehttpCreateSpaceRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *SpacehttpCreateSpaceRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SpacehttpCreateSpaceRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SpacehttpCreateSpaceRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


