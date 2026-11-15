# VoteTallyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | Pointer to **map[string]int32** |  | [optional] 
**SubmissionCount** | Pointer to **int32** |  | [optional] 
**Winners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVoteTallyResult

`func NewVoteTallyResult() *VoteTallyResult`

NewVoteTallyResult instantiates a new VoteTallyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteTallyResultWithDefaults

`func NewVoteTallyResultWithDefaults() *VoteTallyResult`

NewVoteTallyResultWithDefaults instantiates a new VoteTallyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *VoteTallyResult) GetCounts() map[string]int32`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *VoteTallyResult) GetCountsOk() (*map[string]int32, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *VoteTallyResult) SetCounts(v map[string]int32)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *VoteTallyResult) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetSubmissionCount

`func (o *VoteTallyResult) GetSubmissionCount() int32`

GetSubmissionCount returns the SubmissionCount field if non-nil, zero value otherwise.

### GetSubmissionCountOk

`func (o *VoteTallyResult) GetSubmissionCountOk() (*int32, bool)`

GetSubmissionCountOk returns a tuple with the SubmissionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionCount

`func (o *VoteTallyResult) SetSubmissionCount(v int32)`

SetSubmissionCount sets SubmissionCount field to given value.

### HasSubmissionCount

`func (o *VoteTallyResult) HasSubmissionCount() bool`

HasSubmissionCount returns a boolean if a field has been set.

### GetWinners

`func (o *VoteTallyResult) GetWinners() []string`

GetWinners returns the Winners field if non-nil, zero value otherwise.

### GetWinnersOk

`func (o *VoteTallyResult) GetWinnersOk() (*[]string, bool)`

GetWinnersOk returns a tuple with the Winners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinners

`func (o *VoteTallyResult) SetWinners(v []string)`

SetWinners sets Winners field to given value.

### HasWinners

`func (o *VoteTallyResult) HasWinners() bool`

HasWinners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


