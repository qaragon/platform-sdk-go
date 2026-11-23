# DomainTeamResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leading** | Pointer to **bool** |  | [optional] 
**Team** | Pointer to [**TeamsTeam**](TeamsTeam.md) |  | [optional] 
**WinnerSharePct** | Pointer to **float32** |  | [optional] 
**Winners** | Pointer to **int32** |  | [optional] 

## Methods

### NewDomainTeamResult

`func NewDomainTeamResult() *DomainTeamResult`

NewDomainTeamResult instantiates a new DomainTeamResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTeamResultWithDefaults

`func NewDomainTeamResultWithDefaults() *DomainTeamResult`

NewDomainTeamResultWithDefaults instantiates a new DomainTeamResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeading

`func (o *DomainTeamResult) GetLeading() bool`

GetLeading returns the Leading field if non-nil, zero value otherwise.

### GetLeadingOk

`func (o *DomainTeamResult) GetLeadingOk() (*bool, bool)`

GetLeadingOk returns a tuple with the Leading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeading

`func (o *DomainTeamResult) SetLeading(v bool)`

SetLeading sets Leading field to given value.

### HasLeading

`func (o *DomainTeamResult) HasLeading() bool`

HasLeading returns a boolean if a field has been set.

### GetTeam

`func (o *DomainTeamResult) GetTeam() TeamsTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DomainTeamResult) GetTeamOk() (*TeamsTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DomainTeamResult) SetTeam(v TeamsTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DomainTeamResult) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetWinnerSharePct

`func (o *DomainTeamResult) GetWinnerSharePct() float32`

GetWinnerSharePct returns the WinnerSharePct field if non-nil, zero value otherwise.

### GetWinnerSharePctOk

`func (o *DomainTeamResult) GetWinnerSharePctOk() (*float32, bool)`

GetWinnerSharePctOk returns a tuple with the WinnerSharePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerSharePct

`func (o *DomainTeamResult) SetWinnerSharePct(v float32)`

SetWinnerSharePct sets WinnerSharePct field to given value.

### HasWinnerSharePct

`func (o *DomainTeamResult) HasWinnerSharePct() bool`

HasWinnerSharePct returns a boolean if a field has been set.

### GetWinners

`func (o *DomainTeamResult) GetWinners() int32`

GetWinners returns the Winners field if non-nil, zero value otherwise.

### GetWinnersOk

`func (o *DomainTeamResult) GetWinnersOk() (*int32, bool)`

GetWinnersOk returns a tuple with the Winners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinners

`func (o *DomainTeamResult) SetWinners(v int32)`

SetWinners sets Winners field to given value.

### HasWinners

`func (o *DomainTeamResult) HasWinners() bool`

HasWinners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


