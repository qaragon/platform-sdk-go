# UpdatePlatformIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** | CountryCode is optional. Sent as ISO 3166-1 alpha-2; empty value preserves the current value on the row (mirrors the image-URL preservation pattern below). Added in slice 22.6 / uman#107. | [optional] 
**DisplayName** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**ProfilePic** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatePlatformIdentityRequest

`func NewUpdatePlatformIdentityRequest(displayName string, ) *UpdatePlatformIdentityRequest`

NewUpdatePlatformIdentityRequest instantiates a new UpdatePlatformIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlatformIdentityRequestWithDefaults

`func NewUpdatePlatformIdentityRequestWithDefaults() *UpdatePlatformIdentityRequest`

NewUpdatePlatformIdentityRequestWithDefaults instantiates a new UpdatePlatformIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *UpdatePlatformIdentityRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdatePlatformIdentityRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdatePlatformIdentityRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdatePlatformIdentityRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *UpdatePlatformIdentityRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UpdatePlatformIdentityRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UpdatePlatformIdentityRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UpdatePlatformIdentityRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdatePlatformIdentityRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdatePlatformIdentityRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdatePlatformIdentityRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFirstName

`func (o *UpdatePlatformIdentityRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdatePlatformIdentityRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdatePlatformIdentityRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdatePlatformIdentityRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *UpdatePlatformIdentityRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *UpdatePlatformIdentityRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *UpdatePlatformIdentityRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *UpdatePlatformIdentityRequest) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetLastName

`func (o *UpdatePlatformIdentityRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdatePlatformIdentityRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdatePlatformIdentityRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdatePlatformIdentityRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetProfilePic

`func (o *UpdatePlatformIdentityRequest) GetProfilePic() string`

GetProfilePic returns the ProfilePic field if non-nil, zero value otherwise.

### GetProfilePicOk

`func (o *UpdatePlatformIdentityRequest) GetProfilePicOk() (*string, bool)`

GetProfilePicOk returns a tuple with the ProfilePic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePic

`func (o *UpdatePlatformIdentityRequest) SetProfilePic(v string)`

SetProfilePic sets ProfilePic field to given value.

### HasProfilePic

`func (o *UpdatePlatformIdentityRequest) HasProfilePic() bool`

HasProfilePic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


