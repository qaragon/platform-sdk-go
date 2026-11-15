# IdentitiesPlatformIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**IdentitiesCountry**](IdentitiesCountry.md) |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsOfficial** | Pointer to **bool** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata is an opaque, tenant-owned string→string map (ADR 0033). The platform stores and echoes it verbatim and never reads, indexes, or constrains it — it exists only so a tenant can hang its own reference (e.g. its member id) off a PI. Replaces the retired tenant_user_id scalar; omitted from JSON when empty. | [optional] 
**ProfilePic** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentitiesPlatformIdentity

`func NewIdentitiesPlatformIdentity() *IdentitiesPlatformIdentity`

NewIdentitiesPlatformIdentity instantiates a new IdentitiesPlatformIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitiesPlatformIdentityWithDefaults

`func NewIdentitiesPlatformIdentityWithDefaults() *IdentitiesPlatformIdentity`

NewIdentitiesPlatformIdentityWithDefaults instantiates a new IdentitiesPlatformIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *IdentitiesPlatformIdentity) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *IdentitiesPlatformIdentity) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *IdentitiesPlatformIdentity) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *IdentitiesPlatformIdentity) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *IdentitiesPlatformIdentity) GetCountry() IdentitiesCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IdentitiesPlatformIdentity) GetCountryOk() (*IdentitiesCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IdentitiesPlatformIdentity) SetCountry(v IdentitiesCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IdentitiesPlatformIdentity) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *IdentitiesPlatformIdentity) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *IdentitiesPlatformIdentity) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *IdentitiesPlatformIdentity) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *IdentitiesPlatformIdentity) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdentitiesPlatformIdentity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentitiesPlatformIdentity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentitiesPlatformIdentity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentitiesPlatformIdentity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentitiesPlatformIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentitiesPlatformIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentitiesPlatformIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentitiesPlatformIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFirstName

`func (o *IdentitiesPlatformIdentity) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IdentitiesPlatformIdentity) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IdentitiesPlatformIdentity) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IdentitiesPlatformIdentity) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *IdentitiesPlatformIdentity) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *IdentitiesPlatformIdentity) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *IdentitiesPlatformIdentity) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *IdentitiesPlatformIdentity) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetId

`func (o *IdentitiesPlatformIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentitiesPlatformIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentitiesPlatformIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentitiesPlatformIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOfficial

`func (o *IdentitiesPlatformIdentity) GetIsOfficial() bool`

GetIsOfficial returns the IsOfficial field if non-nil, zero value otherwise.

### GetIsOfficialOk

`func (o *IdentitiesPlatformIdentity) GetIsOfficialOk() (*bool, bool)`

GetIsOfficialOk returns a tuple with the IsOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficial

`func (o *IdentitiesPlatformIdentity) SetIsOfficial(v bool)`

SetIsOfficial sets IsOfficial field to given value.

### HasIsOfficial

`func (o *IdentitiesPlatformIdentity) HasIsOfficial() bool`

HasIsOfficial returns a boolean if a field has been set.

### GetLastName

`func (o *IdentitiesPlatformIdentity) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IdentitiesPlatformIdentity) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IdentitiesPlatformIdentity) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IdentitiesPlatformIdentity) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentitiesPlatformIdentity) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentitiesPlatformIdentity) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentitiesPlatformIdentity) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IdentitiesPlatformIdentity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProfilePic

`func (o *IdentitiesPlatformIdentity) GetProfilePic() string`

GetProfilePic returns the ProfilePic field if non-nil, zero value otherwise.

### GetProfilePicOk

`func (o *IdentitiesPlatformIdentity) GetProfilePicOk() (*string, bool)`

GetProfilePicOk returns a tuple with the ProfilePic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePic

`func (o *IdentitiesPlatformIdentity) SetProfilePic(v string)`

SetProfilePic sets ProfilePic field to given value.

### HasProfilePic

`func (o *IdentitiesPlatformIdentity) HasProfilePic() bool`

HasProfilePic returns a boolean if a field has been set.

### GetTenantId

`func (o *IdentitiesPlatformIdentity) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IdentitiesPlatformIdentity) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IdentitiesPlatformIdentity) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IdentitiesPlatformIdentity) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


