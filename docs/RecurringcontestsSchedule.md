# RecurringcontestsSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestTypes** | [**[]RecurringcontestsContestType**](RecurringcontestsContestType.md) |  | 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreationTime** | **string** |  | 
**CreatorIdentityId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DurationMinutes** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**QuestionCount** | Pointer to **int32** |  | [optional] 
**SpaceId** | **string** |  | 
**StartTime** | **string** |  | 
**TimeZone** | **string** |  | 
**Title** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Weekdays** | Pointer to **[]int32** |  | [optional] 
**WinnersPerGroup** | Pointer to **int32** |  | [optional] 

## Methods

### NewRecurringcontestsSchedule

`func NewRecurringcontestsSchedule(contestTypes []RecurringcontestsContestType, creationTime string, spaceId string, startTime string, timeZone string, title string, ) *RecurringcontestsSchedule`

NewRecurringcontestsSchedule instantiates a new RecurringcontestsSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringcontestsScheduleWithDefaults

`func NewRecurringcontestsScheduleWithDefaults() *RecurringcontestsSchedule`

NewRecurringcontestsScheduleWithDefaults instantiates a new RecurringcontestsSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestTypes

`func (o *RecurringcontestsSchedule) GetContestTypes() []RecurringcontestsContestType`

GetContestTypes returns the ContestTypes field if non-nil, zero value otherwise.

### GetContestTypesOk

`func (o *RecurringcontestsSchedule) GetContestTypesOk() (*[]RecurringcontestsContestType, bool)`

GetContestTypesOk returns a tuple with the ContestTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestTypes

`func (o *RecurringcontestsSchedule) SetContestTypes(v []RecurringcontestsContestType)`

SetContestTypes sets ContestTypes field to given value.


### GetCreatedAt

`func (o *RecurringcontestsSchedule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RecurringcontestsSchedule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RecurringcontestsSchedule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RecurringcontestsSchedule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecurringcontestsSchedule) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecurringcontestsSchedule) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecurringcontestsSchedule) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.


### GetCreatorIdentityId

`func (o *RecurringcontestsSchedule) GetCreatorIdentityId() string`

GetCreatorIdentityId returns the CreatorIdentityId field if non-nil, zero value otherwise.

### GetCreatorIdentityIdOk

`func (o *RecurringcontestsSchedule) GetCreatorIdentityIdOk() (*string, bool)`

GetCreatorIdentityIdOk returns a tuple with the CreatorIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorIdentityId

`func (o *RecurringcontestsSchedule) SetCreatorIdentityId(v string)`

SetCreatorIdentityId sets CreatorIdentityId field to given value.

### HasCreatorIdentityId

`func (o *RecurringcontestsSchedule) HasCreatorIdentityId() bool`

HasCreatorIdentityId returns a boolean if a field has been set.

### GetDescription

`func (o *RecurringcontestsSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurringcontestsSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurringcontestsSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurringcontestsSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *RecurringcontestsSchedule) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *RecurringcontestsSchedule) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *RecurringcontestsSchedule) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *RecurringcontestsSchedule) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetEnabled

`func (o *RecurringcontestsSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecurringcontestsSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecurringcontestsSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecurringcontestsSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *RecurringcontestsSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurringcontestsSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurringcontestsSchedule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecurringcontestsSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *RecurringcontestsSchedule) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RecurringcontestsSchedule) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RecurringcontestsSchedule) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *RecurringcontestsSchedule) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetQuestionCount

`func (o *RecurringcontestsSchedule) GetQuestionCount() int32`

GetQuestionCount returns the QuestionCount field if non-nil, zero value otherwise.

### GetQuestionCountOk

`func (o *RecurringcontestsSchedule) GetQuestionCountOk() (*int32, bool)`

GetQuestionCountOk returns a tuple with the QuestionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCount

`func (o *RecurringcontestsSchedule) SetQuestionCount(v int32)`

SetQuestionCount sets QuestionCount field to given value.

### HasQuestionCount

`func (o *RecurringcontestsSchedule) HasQuestionCount() bool`

HasQuestionCount returns a boolean if a field has been set.

### GetSpaceId

`func (o *RecurringcontestsSchedule) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *RecurringcontestsSchedule) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *RecurringcontestsSchedule) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetStartTime

`func (o *RecurringcontestsSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RecurringcontestsSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RecurringcontestsSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetTimeZone

`func (o *RecurringcontestsSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *RecurringcontestsSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *RecurringcontestsSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetTitle

`func (o *RecurringcontestsSchedule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RecurringcontestsSchedule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RecurringcontestsSchedule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *RecurringcontestsSchedule) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RecurringcontestsSchedule) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RecurringcontestsSchedule) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RecurringcontestsSchedule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWeekdays

`func (o *RecurringcontestsSchedule) GetWeekdays() []int32`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *RecurringcontestsSchedule) GetWeekdaysOk() (*[]int32, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *RecurringcontestsSchedule) SetWeekdays(v []int32)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *RecurringcontestsSchedule) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetWinnersPerGroup

`func (o *RecurringcontestsSchedule) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *RecurringcontestsSchedule) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *RecurringcontestsSchedule) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.

### HasWinnersPerGroup

`func (o *RecurringcontestsSchedule) HasWinnersPerGroup() bool`

HasWinnersPerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


