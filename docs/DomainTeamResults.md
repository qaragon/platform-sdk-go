# DomainTeamResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | Pointer to [**[]DomainTeamResult**](DomainTeamResult.md) |  | [optional] 
**TotalWinners** | Pointer to **int32** |  | [optional] 

## Methods

### NewDomainTeamResults

`func NewDomainTeamResults() *DomainTeamResults`

NewDomainTeamResults instantiates a new DomainTeamResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTeamResultsWithDefaults

`func NewDomainTeamResultsWithDefaults() *DomainTeamResults`

NewDomainTeamResultsWithDefaults instantiates a new DomainTeamResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *DomainTeamResults) GetTeams() []DomainTeamResult`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *DomainTeamResults) GetTeamsOk() (*[]DomainTeamResult, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *DomainTeamResults) SetTeams(v []DomainTeamResult)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *DomainTeamResults) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTotalWinners

`func (o *DomainTeamResults) GetTotalWinners() int32`

GetTotalWinners returns the TotalWinners field if non-nil, zero value otherwise.

### GetTotalWinnersOk

`func (o *DomainTeamResults) GetTotalWinnersOk() (*int32, bool)`

GetTotalWinnersOk returns a tuple with the TotalWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWinners

`func (o *DomainTeamResults) SetTotalWinners(v int32)`

SetTotalWinners sets TotalWinners field to given value.

### HasTotalWinners

`func (o *DomainTeamResults) HasTotalWinners() bool`

HasTotalWinners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


