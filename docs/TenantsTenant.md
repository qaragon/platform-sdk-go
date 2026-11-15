# TenantsTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**EnforcesParticipationEligibility** | Pointer to **bool** | EnforcesParticipationEligibility enables configured Contest entry rules. Access, enrollment and experience qualification remain mandatory (ADR 0046). | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantsTenant

`func NewTenantsTenant() *TenantsTenant`

NewTenantsTenant instantiates a new TenantsTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantsTenantWithDefaults

`func NewTenantsTenantWithDefaults() *TenantsTenant`

NewTenantsTenantWithDefaults instantiates a new TenantsTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TenantsTenant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantsTenant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantsTenant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TenantsTenant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnforcesParticipationEligibility

`func (o *TenantsTenant) GetEnforcesParticipationEligibility() bool`

GetEnforcesParticipationEligibility returns the EnforcesParticipationEligibility field if non-nil, zero value otherwise.

### GetEnforcesParticipationEligibilityOk

`func (o *TenantsTenant) GetEnforcesParticipationEligibilityOk() (*bool, bool)`

GetEnforcesParticipationEligibilityOk returns a tuple with the EnforcesParticipationEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcesParticipationEligibility

`func (o *TenantsTenant) SetEnforcesParticipationEligibility(v bool)`

SetEnforcesParticipationEligibility sets EnforcesParticipationEligibility field to given value.

### HasEnforcesParticipationEligibility

`func (o *TenantsTenant) HasEnforcesParticipationEligibility() bool`

HasEnforcesParticipationEligibility returns a boolean if a field has been set.

### GetId

`func (o *TenantsTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantsTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantsTenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantsTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TenantsTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantsTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantsTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantsTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *TenantsTenant) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *TenantsTenant) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *TenantsTenant) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *TenantsTenant) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetStatus

`func (o *TenantsTenant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantsTenant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantsTenant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TenantsTenant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


