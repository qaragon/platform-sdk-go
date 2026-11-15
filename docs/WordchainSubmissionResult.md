# WordchainSubmissionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Game** | Pointer to [**WordchainGameView**](WordchainGameView.md) |  | [optional] 
**Guess** | Pointer to [**GuesstypesResult**](GuesstypesResult.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewWordchainSubmissionResult

`func NewWordchainSubmissionResult(type_ string, ) *WordchainSubmissionResult`

NewWordchainSubmissionResult instantiates a new WordchainSubmissionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordchainSubmissionResultWithDefaults

`func NewWordchainSubmissionResultWithDefaults() *WordchainSubmissionResult`

NewWordchainSubmissionResultWithDefaults instantiates a new WordchainSubmissionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WordchainSubmissionResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WordchainSubmissionResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WordchainSubmissionResult) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WordchainSubmissionResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetGame

`func (o *WordchainSubmissionResult) GetGame() WordchainGameView`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *WordchainSubmissionResult) GetGameOk() (*WordchainGameView, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *WordchainSubmissionResult) SetGame(v WordchainGameView)`

SetGame sets Game field to given value.

### HasGame

`func (o *WordchainSubmissionResult) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetGuess

`func (o *WordchainSubmissionResult) GetGuess() GuesstypesResult`

GetGuess returns the Guess field if non-nil, zero value otherwise.

### GetGuessOk

`func (o *WordchainSubmissionResult) GetGuessOk() (*GuesstypesResult, bool)`

GetGuessOk returns a tuple with the Guess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuess

`func (o *WordchainSubmissionResult) SetGuess(v GuesstypesResult)`

SetGuess sets Guess field to given value.

### HasGuess

`func (o *WordchainSubmissionResult) HasGuess() bool`

HasGuess returns a boolean if a field has been set.

### GetType

`func (o *WordchainSubmissionResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WordchainSubmissionResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WordchainSubmissionResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


