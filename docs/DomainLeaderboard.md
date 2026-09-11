# DomainLeaderboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**DomainBoardContext**](DomainBoardContext.md) |  | [optional] 
**LeaderboardVisibility** | Pointer to **string** | LeaderboardVisibility echoes the channel&#39;s setting (\&quot;visible\&quot;/\&quot;hidden\&quot;). When \&quot;hidden\&quot; the board is withheld from non-owners (empty Rankings/MemberRank); the owner still gets the full board so their client can show a banner. | [optional] 
**MemberRank** | Pointer to [**DomainLeaderboardRank**](DomainLeaderboardRank.md) |  | [optional] 
**Owner** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Period** | Pointer to [**DomainPeriod**](DomainPeriod.md) |  | [optional] 
**Rankings** | Pointer to [**[]DomainLeaderboardRank**](DomainLeaderboardRank.md) |  | [optional] 

## Methods

### NewDomainLeaderboard

`func NewDomainLeaderboard() *DomainLeaderboard`

NewDomainLeaderboard instantiates a new DomainLeaderboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainLeaderboardWithDefaults

`func NewDomainLeaderboardWithDefaults() *DomainLeaderboard`

NewDomainLeaderboardWithDefaults instantiates a new DomainLeaderboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *DomainLeaderboard) GetContext() DomainBoardContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DomainLeaderboard) GetContextOk() (*DomainBoardContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DomainLeaderboard) SetContext(v DomainBoardContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *DomainLeaderboard) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLeaderboardVisibility

`func (o *DomainLeaderboard) GetLeaderboardVisibility() string`

GetLeaderboardVisibility returns the LeaderboardVisibility field if non-nil, zero value otherwise.

### GetLeaderboardVisibilityOk

`func (o *DomainLeaderboard) GetLeaderboardVisibilityOk() (*string, bool)`

GetLeaderboardVisibilityOk returns a tuple with the LeaderboardVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardVisibility

`func (o *DomainLeaderboard) SetLeaderboardVisibility(v string)`

SetLeaderboardVisibility sets LeaderboardVisibility field to given value.

### HasLeaderboardVisibility

`func (o *DomainLeaderboard) HasLeaderboardVisibility() bool`

HasLeaderboardVisibility returns a boolean if a field has been set.

### GetMemberRank

`func (o *DomainLeaderboard) GetMemberRank() DomainLeaderboardRank`

GetMemberRank returns the MemberRank field if non-nil, zero value otherwise.

### GetMemberRankOk

`func (o *DomainLeaderboard) GetMemberRankOk() (*DomainLeaderboardRank, bool)`

GetMemberRankOk returns a tuple with the MemberRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRank

`func (o *DomainLeaderboard) SetMemberRank(v DomainLeaderboardRank)`

SetMemberRank sets MemberRank field to given value.

### HasMemberRank

`func (o *DomainLeaderboard) HasMemberRank() bool`

HasMemberRank returns a boolean if a field has been set.

### GetOwner

`func (o *DomainLeaderboard) GetOwner() DomainPlatformIdentity`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DomainLeaderboard) GetOwnerOk() (*DomainPlatformIdentity, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DomainLeaderboard) SetOwner(v DomainPlatformIdentity)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DomainLeaderboard) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPeriod

`func (o *DomainLeaderboard) GetPeriod() DomainPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainLeaderboard) GetPeriodOk() (*DomainPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainLeaderboard) SetPeriod(v DomainPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainLeaderboard) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRankings

`func (o *DomainLeaderboard) GetRankings() []DomainLeaderboardRank`

GetRankings returns the Rankings field if non-nil, zero value otherwise.

### GetRankingsOk

`func (o *DomainLeaderboard) GetRankingsOk() (*[]DomainLeaderboardRank, bool)`

GetRankingsOk returns a tuple with the Rankings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankings

`func (o *DomainLeaderboard) SetRankings(v []DomainLeaderboardRank)`

SetRankings sets Rankings field to given value.

### HasRankings

`func (o *DomainLeaderboard) HasRankings() bool`

HasRankings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


