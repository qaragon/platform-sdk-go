# EventPresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwayTeam** | Pointer to [**EventTeam**](EventTeam.md) |  | [optional] 
**Competition** | Pointer to [**EventTeam**](EventTeam.md) |  | [optional] 
**HomeTeam** | Pointer to [**EventTeam**](EventTeam.md) |  | [optional] 

## Methods

### NewEventPresentation

`func NewEventPresentation() *EventPresentation`

NewEventPresentation instantiates a new EventPresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPresentationWithDefaults

`func NewEventPresentationWithDefaults() *EventPresentation`

NewEventPresentationWithDefaults instantiates a new EventPresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwayTeam

`func (o *EventPresentation) GetAwayTeam() EventTeam`

GetAwayTeam returns the AwayTeam field if non-nil, zero value otherwise.

### GetAwayTeamOk

`func (o *EventPresentation) GetAwayTeamOk() (*EventTeam, bool)`

GetAwayTeamOk returns a tuple with the AwayTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwayTeam

`func (o *EventPresentation) SetAwayTeam(v EventTeam)`

SetAwayTeam sets AwayTeam field to given value.

### HasAwayTeam

`func (o *EventPresentation) HasAwayTeam() bool`

HasAwayTeam returns a boolean if a field has been set.

### GetCompetition

`func (o *EventPresentation) GetCompetition() EventTeam`

GetCompetition returns the Competition field if non-nil, zero value otherwise.

### GetCompetitionOk

`func (o *EventPresentation) GetCompetitionOk() (*EventTeam, bool)`

GetCompetitionOk returns a tuple with the Competition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetition

`func (o *EventPresentation) SetCompetition(v EventTeam)`

SetCompetition sets Competition field to given value.

### HasCompetition

`func (o *EventPresentation) HasCompetition() bool`

HasCompetition returns a boolean if a field has been set.

### GetHomeTeam

`func (o *EventPresentation) GetHomeTeam() EventTeam`

GetHomeTeam returns the HomeTeam field if non-nil, zero value otherwise.

### GetHomeTeamOk

`func (o *EventPresentation) GetHomeTeamOk() (*EventTeam, bool)`

GetHomeTeamOk returns a tuple with the HomeTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTeam

`func (o *EventPresentation) SetHomeTeam(v EventTeam)`

SetHomeTeam sets HomeTeam field to given value.

### HasHomeTeam

`func (o *EventPresentation) HasHomeTeam() bool`

HasHomeTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


