# VoteTallyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**[]VoteOptionTally**](VoteOptionTally.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubmissionCount** | Pointer to **int32** |  | [optional] 
**Winners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVoteTallyView

`func NewVoteTallyView() *VoteTallyView`

NewVoteTallyView instantiates a new VoteTallyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteTallyViewWithDefaults

`func NewVoteTallyViewWithDefaults() *VoteTallyView`

NewVoteTallyViewWithDefaults instantiates a new VoteTallyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *VoteTallyView) GetOptions() []VoteOptionTally`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *VoteTallyView) GetOptionsOk() (*[]VoteOptionTally, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *VoteTallyView) SetOptions(v []VoteOptionTally)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *VoteTallyView) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetStatus

`func (o *VoteTallyView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoteTallyView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoteTallyView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VoteTallyView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmissionCount

`func (o *VoteTallyView) GetSubmissionCount() int32`

GetSubmissionCount returns the SubmissionCount field if non-nil, zero value otherwise.

### GetSubmissionCountOk

`func (o *VoteTallyView) GetSubmissionCountOk() (*int32, bool)`

GetSubmissionCountOk returns a tuple with the SubmissionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionCount

`func (o *VoteTallyView) SetSubmissionCount(v int32)`

SetSubmissionCount sets SubmissionCount field to given value.

### HasSubmissionCount

`func (o *VoteTallyView) HasSubmissionCount() bool`

HasSubmissionCount returns a boolean if a field has been set.

### GetWinners

`func (o *VoteTallyView) GetWinners() []string`

GetWinners returns the Winners field if non-nil, zero value otherwise.

### GetWinnersOk

`func (o *VoteTallyView) GetWinnersOk() (*[]string, bool)`

GetWinnersOk returns a tuple with the Winners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinners

`func (o *VoteTallyView) SetWinners(v []string)`

SetWinners sets Winners field to given value.

### HasWinners

`func (o *VoteTallyView) HasWinners() bool`

HasWinners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


