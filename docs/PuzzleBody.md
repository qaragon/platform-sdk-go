# PuzzleBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**NullablePuzzleBodyMedia**](PuzzleBodyMedia.md) |  | [optional] 
**Number** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to **map[string]string** |  | [optional] 
**Points** | Pointer to **NullableInt32** |  | [optional] 
**PuzzleType** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **NullableString** | Required for Quiz; optional for Moment, whose media supplies the question. | [optional] 
**Solution** | Pointer to **string** |  | [optional] 

## Methods

### NewPuzzleBody

`func NewPuzzleBody() *PuzzleBody`

NewPuzzleBody instantiates a new PuzzleBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleBodyWithDefaults

`func NewPuzzleBodyWithDefaults() *PuzzleBody`

NewPuzzleBodyWithDefaults instantiates a new PuzzleBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *PuzzleBody) GetMedia() PuzzleBodyMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *PuzzleBody) GetMediaOk() (*PuzzleBodyMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *PuzzleBody) SetMedia(v PuzzleBodyMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *PuzzleBody) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *PuzzleBody) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *PuzzleBody) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetNumber

`func (o *PuzzleBody) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PuzzleBody) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PuzzleBody) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PuzzleBody) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *PuzzleBody) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *PuzzleBody) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetOptions

`func (o *PuzzleBody) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PuzzleBody) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PuzzleBody) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PuzzleBody) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPoints

`func (o *PuzzleBody) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *PuzzleBody) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *PuzzleBody) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *PuzzleBody) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPointsNil

`func (o *PuzzleBody) SetPointsNil(b bool)`

 SetPointsNil sets the value for Points to be an explicit nil

### UnsetPoints
`func (o *PuzzleBody) UnsetPoints()`

UnsetPoints ensures that no value is present for Points, not even an explicit nil
### GetPuzzleType

`func (o *PuzzleBody) GetPuzzleType() string`

GetPuzzleType returns the PuzzleType field if non-nil, zero value otherwise.

### GetPuzzleTypeOk

`func (o *PuzzleBody) GetPuzzleTypeOk() (*string, bool)`

GetPuzzleTypeOk returns a tuple with the PuzzleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzleType

`func (o *PuzzleBody) SetPuzzleType(v string)`

SetPuzzleType sets PuzzleType field to given value.

### HasPuzzleType

`func (o *PuzzleBody) HasPuzzleType() bool`

HasPuzzleType returns a boolean if a field has been set.

### GetQuestion

`func (o *PuzzleBody) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PuzzleBody) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PuzzleBody) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *PuzzleBody) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### SetQuestionNil

`func (o *PuzzleBody) SetQuestionNil(b bool)`

 SetQuestionNil sets the value for Question to be an explicit nil

### UnsetQuestion
`func (o *PuzzleBody) UnsetQuestion()`

UnsetQuestion ensures that no value is present for Question, not even an explicit nil
### GetSolution

`func (o *PuzzleBody) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *PuzzleBody) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *PuzzleBody) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *PuzzleBody) HasSolution() bool`

HasSolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


