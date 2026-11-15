# DomainBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to [**IdentitiesPlatformIdentity**](IdentitiesPlatformIdentity.md) |  | [optional] 
**Blocker** | Pointer to [**IdentitiesPlatformIdentity**](IdentitiesPlatformIdentity.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Space** | Pointer to [**SpacesSpace**](SpacesSpace.md) |  | [optional] 
**SpaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainBlock

`func NewDomainBlock() *DomainBlock`

NewDomainBlock instantiates a new DomainBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBlockWithDefaults

`func NewDomainBlockWithDefaults() *DomainBlock`

NewDomainBlockWithDefaults instantiates a new DomainBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *DomainBlock) GetBlocked() IdentitiesPlatformIdentity`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *DomainBlock) GetBlockedOk() (*IdentitiesPlatformIdentity, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *DomainBlock) SetBlocked(v IdentitiesPlatformIdentity)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *DomainBlock) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlocker

`func (o *DomainBlock) GetBlocker() IdentitiesPlatformIdentity`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DomainBlock) GetBlockerOk() (*IdentitiesPlatformIdentity, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DomainBlock) SetBlocker(v IdentitiesPlatformIdentity)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DomainBlock) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainBlock) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainBlock) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainBlock) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *DomainBlock) GetSpace() SpacesSpace`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *DomainBlock) GetSpaceOk() (*SpacesSpace, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *DomainBlock) SetSpace(v SpacesSpace)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *DomainBlock) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetSpaceId

`func (o *DomainBlock) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DomainBlock) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DomainBlock) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DomainBlock) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


