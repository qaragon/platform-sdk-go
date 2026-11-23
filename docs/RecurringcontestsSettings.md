# RecurringcontestsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestTypes** | [**[]RecurringcontestsContestType**](RecurringcontestsContestType.md) |  | 
**CreationTime** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DurationMinutes** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**QuestionCount** | Pointer to **int32** |  | [optional] 
**SpaceId** | **string** |  | 
**StartTime** | **string** |  | 
**TimeZone** | **string** |  | 
**Title** | **string** |  | 
**Weekdays** | Pointer to **[]int32** |  | [optional] 
**WinnersPerGroup** | Pointer to **int32** |  | [optional] 

## Methods

### NewRecurringcontestsSettings

`func NewRecurringcontestsSettings(contestTypes []RecurringcontestsContestType, creationTime string, spaceId string, startTime string, timeZone string, title string, ) *RecurringcontestsSettings`

NewRecurringcontestsSettings instantiates a new RecurringcontestsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringcontestsSettingsWithDefaults

`func NewRecurringcontestsSettingsWithDefaults() *RecurringcontestsSettings`

NewRecurringcontestsSettingsWithDefaults instantiates a new RecurringcontestsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestTypes

`func (o *RecurringcontestsSettings) GetContestTypes() []RecurringcontestsContestType`

GetContestTypes returns the ContestTypes field if non-nil, zero value otherwise.

### GetContestTypesOk

`func (o *RecurringcontestsSettings) GetContestTypesOk() (*[]RecurringcontestsContestType, bool)`

GetContestTypesOk returns a tuple with the ContestTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestTypes

`func (o *RecurringcontestsSettings) SetContestTypes(v []RecurringcontestsContestType)`

SetContestTypes sets ContestTypes field to given value.


### GetCreationTime

`func (o *RecurringcontestsSettings) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecurringcontestsSettings) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecurringcontestsSettings) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.


### GetDescription

`func (o *RecurringcontestsSettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurringcontestsSettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurringcontestsSettings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurringcontestsSettings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *RecurringcontestsSettings) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *RecurringcontestsSettings) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *RecurringcontestsSettings) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *RecurringcontestsSettings) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetEnabled

`func (o *RecurringcontestsSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecurringcontestsSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecurringcontestsSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecurringcontestsSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLanguage

`func (o *RecurringcontestsSettings) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RecurringcontestsSettings) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RecurringcontestsSettings) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *RecurringcontestsSettings) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetQuestionCount

`func (o *RecurringcontestsSettings) GetQuestionCount() int32`

GetQuestionCount returns the QuestionCount field if non-nil, zero value otherwise.

### GetQuestionCountOk

`func (o *RecurringcontestsSettings) GetQuestionCountOk() (*int32, bool)`

GetQuestionCountOk returns a tuple with the QuestionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCount

`func (o *RecurringcontestsSettings) SetQuestionCount(v int32)`

SetQuestionCount sets QuestionCount field to given value.

### HasQuestionCount

`func (o *RecurringcontestsSettings) HasQuestionCount() bool`

HasQuestionCount returns a boolean if a field has been set.

### GetSpaceId

`func (o *RecurringcontestsSettings) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *RecurringcontestsSettings) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *RecurringcontestsSettings) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetStartTime

`func (o *RecurringcontestsSettings) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RecurringcontestsSettings) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RecurringcontestsSettings) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetTimeZone

`func (o *RecurringcontestsSettings) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *RecurringcontestsSettings) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *RecurringcontestsSettings) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetTitle

`func (o *RecurringcontestsSettings) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RecurringcontestsSettings) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RecurringcontestsSettings) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWeekdays

`func (o *RecurringcontestsSettings) GetWeekdays() []int32`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *RecurringcontestsSettings) GetWeekdaysOk() (*[]int32, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *RecurringcontestsSettings) SetWeekdays(v []int32)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *RecurringcontestsSettings) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetWinnersPerGroup

`func (o *RecurringcontestsSettings) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *RecurringcontestsSettings) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *RecurringcontestsSettings) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.

### HasWinnersPerGroup

`func (o *RecurringcontestsSettings) HasWinnersPerGroup() bool`

HasWinnersPerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


