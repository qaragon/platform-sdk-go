# WordchainActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 

## Methods

### NewWordchainActionRequest

`func NewWordchainActionRequest(action string, ) *WordchainActionRequest`

NewWordchainActionRequest instantiates a new WordchainActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordchainActionRequestWithDefaults

`func NewWordchainActionRequestWithDefaults() *WordchainActionRequest`

NewWordchainActionRequestWithDefaults instantiates a new WordchainActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WordchainActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WordchainActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WordchainActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionVersion

`func (o *WordchainActionRequest) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *WordchainActionRequest) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *WordchainActionRequest) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *WordchainActionRequest) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetChainId

`func (o *WordchainActionRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *WordchainActionRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *WordchainActionRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *WordchainActionRequest) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


