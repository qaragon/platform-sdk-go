# SevencluesSubmissionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Board** | Pointer to [**SevencluesBoardView**](SevencluesBoardView.md) |  | [optional] 
**Guess** | Pointer to [**GuesstypesResult**](GuesstypesResult.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewSevencluesSubmissionResult

`func NewSevencluesSubmissionResult(type_ string, ) *SevencluesSubmissionResult`

NewSevencluesSubmissionResult instantiates a new SevencluesSubmissionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSevencluesSubmissionResultWithDefaults

`func NewSevencluesSubmissionResultWithDefaults() *SevencluesSubmissionResult`

NewSevencluesSubmissionResultWithDefaults instantiates a new SevencluesSubmissionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoard

`func (o *SevencluesSubmissionResult) GetBoard() SevencluesBoardView`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *SevencluesSubmissionResult) GetBoardOk() (*SevencluesBoardView, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *SevencluesSubmissionResult) SetBoard(v SevencluesBoardView)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *SevencluesSubmissionResult) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetGuess

`func (o *SevencluesSubmissionResult) GetGuess() GuesstypesResult`

GetGuess returns the Guess field if non-nil, zero value otherwise.

### GetGuessOk

`func (o *SevencluesSubmissionResult) GetGuessOk() (*GuesstypesResult, bool)`

GetGuessOk returns a tuple with the Guess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuess

`func (o *SevencluesSubmissionResult) SetGuess(v GuesstypesResult)`

SetGuess sets Guess field to given value.

### HasGuess

`func (o *SevencluesSubmissionResult) HasGuess() bool`

HasGuess returns a boolean if a field has been set.

### GetType

`func (o *SevencluesSubmissionResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SevencluesSubmissionResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SevencluesSubmissionResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


