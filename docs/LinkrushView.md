# LinkrushView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionVersion** | Pointer to **int32** |  | [optional] 
**ActiveIdentityId** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to [**[]WordchainWordView**](WordchainWordView.md) |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**ChainNumber** | Pointer to **int32** |  | [optional] 
**Deadline** | Pointer to **NullableString** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**RemainingMilliseconds** | Pointer to **int32** |  | [optional] 
**SelectedPosition** | Pointer to **NullableInt32** |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewLinkrushView

`func NewLinkrushView() *LinkrushView`

NewLinkrushView instantiates a new LinkrushView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkrushViewWithDefaults

`func NewLinkrushViewWithDefaults() *LinkrushView`

NewLinkrushViewWithDefaults instantiates a new LinkrushView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionVersion

`func (o *LinkrushView) GetActionVersion() int32`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *LinkrushView) GetActionVersionOk() (*int32, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *LinkrushView) SetActionVersion(v int32)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *LinkrushView) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetActiveIdentityId

`func (o *LinkrushView) GetActiveIdentityId() string`

GetActiveIdentityId returns the ActiveIdentityId field if non-nil, zero value otherwise.

### GetActiveIdentityIdOk

`func (o *LinkrushView) GetActiveIdentityIdOk() (*string, bool)`

GetActiveIdentityIdOk returns a tuple with the ActiveIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentityId

`func (o *LinkrushView) SetActiveIdentityId(v string)`

SetActiveIdentityId sets ActiveIdentityId field to given value.

### HasActiveIdentityId

`func (o *LinkrushView) HasActiveIdentityId() bool`

HasActiveIdentityId returns a boolean if a field has been set.

### GetChain

`func (o *LinkrushView) GetChain() []WordchainWordView`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *LinkrushView) GetChainOk() (*[]WordchainWordView, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *LinkrushView) SetChain(v []WordchainWordView)`

SetChain sets Chain field to given value.

### HasChain

`func (o *LinkrushView) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *LinkrushView) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *LinkrushView) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *LinkrushView) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *LinkrushView) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainNumber

`func (o *LinkrushView) GetChainNumber() int32`

GetChainNumber returns the ChainNumber field if non-nil, zero value otherwise.

### GetChainNumberOk

`func (o *LinkrushView) GetChainNumberOk() (*int32, bool)`

GetChainNumberOk returns a tuple with the ChainNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainNumber

`func (o *LinkrushView) SetChainNumber(v int32)`

SetChainNumber sets ChainNumber field to given value.

### HasChainNumber

`func (o *LinkrushView) HasChainNumber() bool`

HasChainNumber returns a boolean if a field has been set.

### GetDeadline

`func (o *LinkrushView) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *LinkrushView) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *LinkrushView) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *LinkrushView) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *LinkrushView) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *LinkrushView) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetPlayerIdentityIds

`func (o *LinkrushView) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *LinkrushView) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *LinkrushView) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *LinkrushView) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetRemainingMilliseconds

`func (o *LinkrushView) GetRemainingMilliseconds() int32`

GetRemainingMilliseconds returns the RemainingMilliseconds field if non-nil, zero value otherwise.

### GetRemainingMillisecondsOk

`func (o *LinkrushView) GetRemainingMillisecondsOk() (*int32, bool)`

GetRemainingMillisecondsOk returns a tuple with the RemainingMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMilliseconds

`func (o *LinkrushView) SetRemainingMilliseconds(v int32)`

SetRemainingMilliseconds sets RemainingMilliseconds field to given value.

### HasRemainingMilliseconds

`func (o *LinkrushView) HasRemainingMilliseconds() bool`

HasRemainingMilliseconds returns a boolean if a field has been set.

### GetSelectedPosition

`func (o *LinkrushView) GetSelectedPosition() int32`

GetSelectedPosition returns the SelectedPosition field if non-nil, zero value otherwise.

### GetSelectedPositionOk

`func (o *LinkrushView) GetSelectedPositionOk() (*int32, bool)`

GetSelectedPositionOk returns a tuple with the SelectedPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPosition

`func (o *LinkrushView) SetSelectedPosition(v int32)`

SetSelectedPosition sets SelectedPosition field to given value.

### HasSelectedPosition

`func (o *LinkrushView) HasSelectedPosition() bool`

HasSelectedPosition returns a boolean if a field has been set.

### SetSelectedPositionNil

`func (o *LinkrushView) SetSelectedPositionNil(b bool)`

 SetSelectedPositionNil sets the value for SelectedPosition to be an explicit nil

### UnsetSelectedPosition
`func (o *LinkrushView) UnsetSelectedPosition()`

UnsetSelectedPosition ensures that no value is present for SelectedPosition, not even an explicit nil
### GetStageId

`func (o *LinkrushView) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *LinkrushView) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *LinkrushView) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *LinkrushView) HasStageId() bool`

HasStageId returns a boolean if a field has been set.

### GetState

`func (o *LinkrushView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkrushView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LinkrushView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LinkrushView) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


