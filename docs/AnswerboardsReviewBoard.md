# AnswerboardsReviewBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]AnswerboardsAnswer**](AnswerboardsAnswer.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Culture** | Pointer to **string** |  | [optional] 
**Difficulty** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**MinimumAge** | Pointer to **int32** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**ReviewStatus** | Pointer to **string** |  | [optional] 
**ReviewedAt** | Pointer to **string** |  | [optional] 
**ReviewedBy** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAnswerboardsReviewBoard

`func NewAnswerboardsReviewBoard() *AnswerboardsReviewBoard`

NewAnswerboardsReviewBoard instantiates a new AnswerboardsReviewBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerboardsReviewBoardWithDefaults

`func NewAnswerboardsReviewBoardWithDefaults() *AnswerboardsReviewBoard`

NewAnswerboardsReviewBoardWithDefaults instantiates a new AnswerboardsReviewBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *AnswerboardsReviewBoard) GetAnswers() []AnswerboardsAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *AnswerboardsReviewBoard) GetAnswersOk() (*[]AnswerboardsAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *AnswerboardsReviewBoard) SetAnswers(v []AnswerboardsAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *AnswerboardsReviewBoard) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnswerboardsReviewBoard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnswerboardsReviewBoard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnswerboardsReviewBoard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnswerboardsReviewBoard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCulture

`func (o *AnswerboardsReviewBoard) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AnswerboardsReviewBoard) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AnswerboardsReviewBoard) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AnswerboardsReviewBoard) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### GetDifficulty

`func (o *AnswerboardsReviewBoard) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *AnswerboardsReviewBoard) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *AnswerboardsReviewBoard) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *AnswerboardsReviewBoard) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetId

`func (o *AnswerboardsReviewBoard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnswerboardsReviewBoard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnswerboardsReviewBoard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnswerboardsReviewBoard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AnswerboardsReviewBoard) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AnswerboardsReviewBoard) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AnswerboardsReviewBoard) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AnswerboardsReviewBoard) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguage

`func (o *AnswerboardsReviewBoard) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AnswerboardsReviewBoard) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AnswerboardsReviewBoard) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AnswerboardsReviewBoard) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMinimumAge

`func (o *AnswerboardsReviewBoard) GetMinimumAge() int32`

GetMinimumAge returns the MinimumAge field if non-nil, zero value otherwise.

### GetMinimumAgeOk

`func (o *AnswerboardsReviewBoard) GetMinimumAgeOk() (*int32, bool)`

GetMinimumAgeOk returns a tuple with the MinimumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAge

`func (o *AnswerboardsReviewBoard) SetMinimumAge(v int32)`

SetMinimumAge sets MinimumAge field to given value.

### HasMinimumAge

`func (o *AnswerboardsReviewBoard) HasMinimumAge() bool`

HasMinimumAge returns a boolean if a field has been set.

### GetPrompt

`func (o *AnswerboardsReviewBoard) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AnswerboardsReviewBoard) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AnswerboardsReviewBoard) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AnswerboardsReviewBoard) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetReviewStatus

`func (o *AnswerboardsReviewBoard) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *AnswerboardsReviewBoard) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *AnswerboardsReviewBoard) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.

### HasReviewStatus

`func (o *AnswerboardsReviewBoard) HasReviewStatus() bool`

HasReviewStatus returns a boolean if a field has been set.

### GetReviewedAt

`func (o *AnswerboardsReviewBoard) GetReviewedAt() string`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *AnswerboardsReviewBoard) GetReviewedAtOk() (*string, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *AnswerboardsReviewBoard) SetReviewedAt(v string)`

SetReviewedAt sets ReviewedAt field to given value.

### HasReviewedAt

`func (o *AnswerboardsReviewBoard) HasReviewedAt() bool`

HasReviewedAt returns a boolean if a field has been set.

### GetReviewedBy

`func (o *AnswerboardsReviewBoard) GetReviewedBy() string`

GetReviewedBy returns the ReviewedBy field if non-nil, zero value otherwise.

### GetReviewedByOk

`func (o *AnswerboardsReviewBoard) GetReviewedByOk() (*string, bool)`

GetReviewedByOk returns a tuple with the ReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBy

`func (o *AnswerboardsReviewBoard) SetReviewedBy(v string)`

SetReviewedBy sets ReviewedBy field to given value.

### HasReviewedBy

`func (o *AnswerboardsReviewBoard) HasReviewedBy() bool`

HasReviewedBy returns a boolean if a field has been set.

### GetSource

`func (o *AnswerboardsReviewBoard) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnswerboardsReviewBoard) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnswerboardsReviewBoard) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AnswerboardsReviewBoard) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTopics

`func (o *AnswerboardsReviewBoard) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *AnswerboardsReviewBoard) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *AnswerboardsReviewBoard) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *AnswerboardsReviewBoard) HasTopics() bool`

HasTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


