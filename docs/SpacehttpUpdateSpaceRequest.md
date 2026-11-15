# SpacehttpUpdateSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**IsFanSpace** | Pointer to **bool** |  | [optional] 
**LeaderboardVisibility** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Regions** | Pointer to [**[]SpacehttpRegion**](SpacehttpRegion.md) | Omit to retain the current regions; [] makes the Space permanently worldwide. | [optional] 
**Size** | **int32** |  | 

## Methods

### NewSpacehttpUpdateSpaceRequest

`func NewSpacehttpUpdateSpaceRequest(name string, size int32, ) *SpacehttpUpdateSpaceRequest`

NewSpacehttpUpdateSpaceRequest instantiates a new SpacehttpUpdateSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacehttpUpdateSpaceRequestWithDefaults

`func NewSpacehttpUpdateSpaceRequestWithDefaults() *SpacehttpUpdateSpaceRequest`

NewSpacehttpUpdateSpaceRequestWithDefaults instantiates a new SpacehttpUpdateSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *SpacehttpUpdateSpaceRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SpacehttpUpdateSpaceRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SpacehttpUpdateSpaceRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SpacehttpUpdateSpaceRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SpacehttpUpdateSpaceRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SpacehttpUpdateSpaceRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SpacehttpUpdateSpaceRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SpacehttpUpdateSpaceRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *SpacehttpUpdateSpaceRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *SpacehttpUpdateSpaceRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *SpacehttpUpdateSpaceRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *SpacehttpUpdateSpaceRequest) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetIsFanSpace

`func (o *SpacehttpUpdateSpaceRequest) GetIsFanSpace() bool`

GetIsFanSpace returns the IsFanSpace field if non-nil, zero value otherwise.

### GetIsFanSpaceOk

`func (o *SpacehttpUpdateSpaceRequest) GetIsFanSpaceOk() (*bool, bool)`

GetIsFanSpaceOk returns a tuple with the IsFanSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFanSpace

`func (o *SpacehttpUpdateSpaceRequest) SetIsFanSpace(v bool)`

SetIsFanSpace sets IsFanSpace field to given value.

### HasIsFanSpace

`func (o *SpacehttpUpdateSpaceRequest) HasIsFanSpace() bool`

HasIsFanSpace returns a boolean if a field has been set.

### GetLeaderboardVisibility

`func (o *SpacehttpUpdateSpaceRequest) GetLeaderboardVisibility() string`

GetLeaderboardVisibility returns the LeaderboardVisibility field if non-nil, zero value otherwise.

### GetLeaderboardVisibilityOk

`func (o *SpacehttpUpdateSpaceRequest) GetLeaderboardVisibilityOk() (*string, bool)`

GetLeaderboardVisibilityOk returns a tuple with the LeaderboardVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardVisibility

`func (o *SpacehttpUpdateSpaceRequest) SetLeaderboardVisibility(v string)`

SetLeaderboardVisibility sets LeaderboardVisibility field to given value.

### HasLeaderboardVisibility

`func (o *SpacehttpUpdateSpaceRequest) HasLeaderboardVisibility() bool`

HasLeaderboardVisibility returns a boolean if a field has been set.

### GetName

`func (o *SpacehttpUpdateSpaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpacehttpUpdateSpaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpacehttpUpdateSpaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegions

`func (o *SpacehttpUpdateSpaceRequest) GetRegions() []SpacehttpRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SpacehttpUpdateSpaceRequest) GetRegionsOk() (*[]SpacehttpRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SpacehttpUpdateSpaceRequest) SetRegions(v []SpacehttpRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SpacehttpUpdateSpaceRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *SpacehttpUpdateSpaceRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SpacehttpUpdateSpaceRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SpacehttpUpdateSpaceRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


