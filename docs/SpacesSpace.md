# SpacesSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCode** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | Human-facing lookup code; persisted references use Id. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsFanSpace** | Pointer to **bool** |  | [optional] 
**LeaderboardVisibility** | Pointer to **string** | Owner-set: \&quot;visible\&quot; (default) or \&quot;hidden\&quot; (standings withheld from non-owners) | [optional] 
**Membership** | Pointer to [**SpacesMembershipSummary**](SpacesMembershipSummary.md) | Set when authenticated; omitted when not subscribed | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**IdentitiesPlatformIdentity**](IdentitiesPlatformIdentity.md) |  | [optional] 
**Regions** | Pointer to [**[]SpacesRegion**](SpacesRegion.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Stats** | Pointer to [**SpacesStats**](SpacesStats.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**VenueId** | Pointer to **string** |  | [optional] 

## Methods

### NewSpacesSpace

`func NewSpacesSpace() *SpacesSpace`

NewSpacesSpace instantiates a new SpacesSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacesSpaceWithDefaults

`func NewSpacesSpaceWithDefaults() *SpacesSpace`

NewSpacesSpaceWithDefaults instantiates a new SpacesSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCode

`func (o *SpacesSpace) GetAccessCode() string`

GetAccessCode returns the AccessCode field if non-nil, zero value otherwise.

### GetAccessCodeOk

`func (o *SpacesSpace) GetAccessCodeOk() (*string, bool)`

GetAccessCodeOk returns a tuple with the AccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCode

`func (o *SpacesSpace) SetAccessCode(v string)`

SetAccessCode sets AccessCode field to given value.

### HasAccessCode

`func (o *SpacesSpace) HasAccessCode() bool`

HasAccessCode returns a boolean if a field has been set.

### GetAccessType

`func (o *SpacesSpace) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SpacesSpace) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SpacesSpace) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SpacesSpace) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetCode

`func (o *SpacesSpace) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpacesSpace) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpacesSpace) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SpacesSpace) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SpacesSpace) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpacesSpace) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpacesSpace) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpacesSpace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SpacesSpace) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SpacesSpace) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SpacesSpace) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SpacesSpace) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *SpacesSpace) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *SpacesSpace) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *SpacesSpace) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *SpacesSpace) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetId

`func (o *SpacesSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpacesSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpacesSpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpacesSpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFanSpace

`func (o *SpacesSpace) GetIsFanSpace() bool`

GetIsFanSpace returns the IsFanSpace field if non-nil, zero value otherwise.

### GetIsFanSpaceOk

`func (o *SpacesSpace) GetIsFanSpaceOk() (*bool, bool)`

GetIsFanSpaceOk returns a tuple with the IsFanSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFanSpace

`func (o *SpacesSpace) SetIsFanSpace(v bool)`

SetIsFanSpace sets IsFanSpace field to given value.

### HasIsFanSpace

`func (o *SpacesSpace) HasIsFanSpace() bool`

HasIsFanSpace returns a boolean if a field has been set.

### GetLeaderboardVisibility

`func (o *SpacesSpace) GetLeaderboardVisibility() string`

GetLeaderboardVisibility returns the LeaderboardVisibility field if non-nil, zero value otherwise.

### GetLeaderboardVisibilityOk

`func (o *SpacesSpace) GetLeaderboardVisibilityOk() (*string, bool)`

GetLeaderboardVisibilityOk returns a tuple with the LeaderboardVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardVisibility

`func (o *SpacesSpace) SetLeaderboardVisibility(v string)`

SetLeaderboardVisibility sets LeaderboardVisibility field to given value.

### HasLeaderboardVisibility

`func (o *SpacesSpace) HasLeaderboardVisibility() bool`

HasLeaderboardVisibility returns a boolean if a field has been set.

### GetMembership

`func (o *SpacesSpace) GetMembership() SpacesMembershipSummary`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *SpacesSpace) GetMembershipOk() (*SpacesMembershipSummary, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *SpacesSpace) SetMembership(v SpacesMembershipSummary)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *SpacesSpace) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetName

`func (o *SpacesSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpacesSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpacesSpace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpacesSpace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *SpacesSpace) GetOwner() IdentitiesPlatformIdentity`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SpacesSpace) GetOwnerOk() (*IdentitiesPlatformIdentity, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SpacesSpace) SetOwner(v IdentitiesPlatformIdentity)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SpacesSpace) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRegions

`func (o *SpacesSpace) GetRegions() []SpacesRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SpacesSpace) GetRegionsOk() (*[]SpacesRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SpacesSpace) SetRegions(v []SpacesRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SpacesSpace) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *SpacesSpace) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SpacesSpace) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SpacesSpace) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SpacesSpace) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStats

`func (o *SpacesSpace) GetStats() SpacesStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SpacesSpace) GetStatsOk() (*SpacesStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SpacesSpace) SetStats(v SpacesStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SpacesSpace) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTenantId

`func (o *SpacesSpace) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SpacesSpace) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SpacesSpace) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SpacesSpace) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetVenueId

`func (o *SpacesSpace) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *SpacesSpace) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *SpacesSpace) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.

### HasVenueId

`func (o *SpacesSpace) HasVenueId() bool`

HasVenueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


