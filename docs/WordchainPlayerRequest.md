# WordchainPlayerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ActionVersion** | Pointer to **int64** |  | [optional] 
**Answers** | Pointer to **[]string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Guess** | Pointer to **string** |  | [optional] 
**Play** | Pointer to **NullableBool** |  | [optional] 
**Wager** | Pointer to **int32** |  | [optional] 

## Methods

### NewWordchainPlayerRequest

`func NewWordchainPlayerRequest() *WordchainPlayerRequest`

NewWordchainPlayerRequest instantiates a new WordchainPlayerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordchainPlayerRequestWithDefaults

`func NewWordchainPlayerRequestWithDefaults() *WordchainPlayerRequest`

NewWordchainPlayerRequestWithDefaults instantiates a new WordchainPlayerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WordchainPlayerRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WordchainPlayerRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WordchainPlayerRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WordchainPlayerRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionVersion

`func (o *WordchainPlayerRequest) GetActionVersion() int64`

GetActionVersion returns the ActionVersion field if non-nil, zero value otherwise.

### GetActionVersionOk

`func (o *WordchainPlayerRequest) GetActionVersionOk() (*int64, bool)`

GetActionVersionOk returns a tuple with the ActionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionVersion

`func (o *WordchainPlayerRequest) SetActionVersion(v int64)`

SetActionVersion sets ActionVersion field to given value.

### HasActionVersion

`func (o *WordchainPlayerRequest) HasActionVersion() bool`

HasActionVersion returns a boolean if a field has been set.

### GetAnswers

`func (o *WordchainPlayerRequest) GetAnswers() []string`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *WordchainPlayerRequest) GetAnswersOk() (*[]string, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *WordchainPlayerRequest) SetAnswers(v []string)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *WordchainPlayerRequest) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetChainId

`func (o *WordchainPlayerRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *WordchainPlayerRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *WordchainPlayerRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *WordchainPlayerRequest) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetDirection

`func (o *WordchainPlayerRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *WordchainPlayerRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *WordchainPlayerRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *WordchainPlayerRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetGuess

`func (o *WordchainPlayerRequest) GetGuess() string`

GetGuess returns the Guess field if non-nil, zero value otherwise.

### GetGuessOk

`func (o *WordchainPlayerRequest) GetGuessOk() (*string, bool)`

GetGuessOk returns a tuple with the Guess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuess

`func (o *WordchainPlayerRequest) SetGuess(v string)`

SetGuess sets Guess field to given value.

### HasGuess

`func (o *WordchainPlayerRequest) HasGuess() bool`

HasGuess returns a boolean if a field has been set.

### GetPlay

`func (o *WordchainPlayerRequest) GetPlay() bool`

GetPlay returns the Play field if non-nil, zero value otherwise.

### GetPlayOk

`func (o *WordchainPlayerRequest) GetPlayOk() (*bool, bool)`

GetPlayOk returns a tuple with the Play field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlay

`func (o *WordchainPlayerRequest) SetPlay(v bool)`

SetPlay sets Play field to given value.

### HasPlay

`func (o *WordchainPlayerRequest) HasPlay() bool`

HasPlay returns a boolean if a field has been set.

### SetPlayNil

`func (o *WordchainPlayerRequest) SetPlayNil(b bool)`

 SetPlayNil sets the value for Play to be an explicit nil

### UnsetPlay
`func (o *WordchainPlayerRequest) UnsetPlay()`

UnsetPlay ensures that no value is present for Play, not even an explicit nil
### GetWager

`func (o *WordchainPlayerRequest) GetWager() int32`

GetWager returns the Wager field if non-nil, zero value otherwise.

### GetWagerOk

`func (o *WordchainPlayerRequest) GetWagerOk() (*int32, bool)`

GetWagerOk returns a tuple with the Wager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWager

`func (o *WordchainPlayerRequest) SetWager(v int32)`

SetWager sets Wager field to given value.

### HasWager

`func (o *WordchainPlayerRequest) HasWager() bool`

HasWager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


