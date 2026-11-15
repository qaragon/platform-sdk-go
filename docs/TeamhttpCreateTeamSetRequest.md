# TeamhttpCreateTeamSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Teams** | [**[]TeamhttpCreateTeamRequest**](TeamhttpCreateTeamRequest.md) |  | 

## Methods

### NewTeamhttpCreateTeamSetRequest

`func NewTeamhttpCreateTeamSetRequest(name string, teams []TeamhttpCreateTeamRequest, ) *TeamhttpCreateTeamSetRequest`

NewTeamhttpCreateTeamSetRequest instantiates a new TeamhttpCreateTeamSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamhttpCreateTeamSetRequestWithDefaults

`func NewTeamhttpCreateTeamSetRequestWithDefaults() *TeamhttpCreateTeamSetRequest`

NewTeamhttpCreateTeamSetRequestWithDefaults instantiates a new TeamhttpCreateTeamSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TeamhttpCreateTeamSetRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TeamhttpCreateTeamSetRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TeamhttpCreateTeamSetRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TeamhttpCreateTeamSetRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *TeamhttpCreateTeamSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamhttpCreateTeamSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamhttpCreateTeamSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *TeamhttpCreateTeamSetRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TeamhttpCreateTeamSetRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TeamhttpCreateTeamSetRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TeamhttpCreateTeamSetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeams

`func (o *TeamhttpCreateTeamSetRequest) GetTeams() []TeamhttpCreateTeamRequest`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TeamhttpCreateTeamSetRequest) GetTeamsOk() (*[]TeamhttpCreateTeamRequest, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TeamhttpCreateTeamSetRequest) SetTeams(v []TeamhttpCreateTeamRequest)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


