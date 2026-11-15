# FindtheminorityStartRoundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**[]VoteOption**](VoteOption.md) |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 

## Methods

### NewFindtheminorityStartRoundRequest

`func NewFindtheminorityStartRoundRequest() *FindtheminorityStartRoundRequest`

NewFindtheminorityStartRoundRequest instantiates a new FindtheminorityStartRoundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindtheminorityStartRoundRequestWithDefaults

`func NewFindtheminorityStartRoundRequestWithDefaults() *FindtheminorityStartRoundRequest`

NewFindtheminorityStartRoundRequestWithDefaults instantiates a new FindtheminorityStartRoundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *FindtheminorityStartRoundRequest) GetOptions() []VoteOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FindtheminorityStartRoundRequest) GetOptionsOk() (*[]VoteOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FindtheminorityStartRoundRequest) SetOptions(v []VoteOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FindtheminorityStartRoundRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *FindtheminorityStartRoundRequest) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *FindtheminorityStartRoundRequest) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *FindtheminorityStartRoundRequest) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *FindtheminorityStartRoundRequest) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetQuestion

`func (o *FindtheminorityStartRoundRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FindtheminorityStartRoundRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FindtheminorityStartRoundRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *FindtheminorityStartRoundRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


