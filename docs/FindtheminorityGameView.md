# FindtheminorityGameView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceSpaceId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlayerSpaceId** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]FindtheminorityPlayer**](FindtheminorityPlayer.md) |  | [optional] 
**Round** | Pointer to [**FindtheminorityRoundView**](FindtheminorityRoundView.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**ViewerRole** | Pointer to **string** |  | [optional] 

## Methods

### NewFindtheminorityGameView

`func NewFindtheminorityGameView(type_ string, ) *FindtheminorityGameView`

NewFindtheminorityGameView instantiates a new FindtheminorityGameView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindtheminorityGameViewWithDefaults

`func NewFindtheminorityGameViewWithDefaults() *FindtheminorityGameView`

NewFindtheminorityGameViewWithDefaults instantiates a new FindtheminorityGameView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceSpaceId

`func (o *FindtheminorityGameView) GetAudienceSpaceId() string`

GetAudienceSpaceId returns the AudienceSpaceId field if non-nil, zero value otherwise.

### GetAudienceSpaceIdOk

`func (o *FindtheminorityGameView) GetAudienceSpaceIdOk() (*string, bool)`

GetAudienceSpaceIdOk returns a tuple with the AudienceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceSpaceId

`func (o *FindtheminorityGameView) SetAudienceSpaceId(v string)`

SetAudienceSpaceId sets AudienceSpaceId field to given value.

### HasAudienceSpaceId

`func (o *FindtheminorityGameView) HasAudienceSpaceId() bool`

HasAudienceSpaceId returns a boolean if a field has been set.

### GetId

`func (o *FindtheminorityGameView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindtheminorityGameView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindtheminorityGameView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindtheminorityGameView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlayerSpaceId

`func (o *FindtheminorityGameView) GetPlayerSpaceId() string`

GetPlayerSpaceId returns the PlayerSpaceId field if non-nil, zero value otherwise.

### GetPlayerSpaceIdOk

`func (o *FindtheminorityGameView) GetPlayerSpaceIdOk() (*string, bool)`

GetPlayerSpaceIdOk returns a tuple with the PlayerSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSpaceId

`func (o *FindtheminorityGameView) SetPlayerSpaceId(v string)`

SetPlayerSpaceId sets PlayerSpaceId field to given value.

### HasPlayerSpaceId

`func (o *FindtheminorityGameView) HasPlayerSpaceId() bool`

HasPlayerSpaceId returns a boolean if a field has been set.

### GetPlayers

`func (o *FindtheminorityGameView) GetPlayers() []FindtheminorityPlayer`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *FindtheminorityGameView) GetPlayersOk() (*[]FindtheminorityPlayer, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *FindtheminorityGameView) SetPlayers(v []FindtheminorityPlayer)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *FindtheminorityGameView) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetRound

`func (o *FindtheminorityGameView) GetRound() FindtheminorityRoundView`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *FindtheminorityGameView) GetRoundOk() (*FindtheminorityRoundView, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *FindtheminorityGameView) SetRound(v FindtheminorityRoundView)`

SetRound sets Round field to given value.

### HasRound

`func (o *FindtheminorityGameView) HasRound() bool`

HasRound returns a boolean if a field has been set.

### GetState

`func (o *FindtheminorityGameView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindtheminorityGameView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindtheminorityGameView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindtheminorityGameView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *FindtheminorityGameView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindtheminorityGameView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindtheminorityGameView) SetType(v string)`

SetType sets Type field to given value.


### GetViewerRole

`func (o *FindtheminorityGameView) GetViewerRole() string`

GetViewerRole returns the ViewerRole field if non-nil, zero value otherwise.

### GetViewerRoleOk

`func (o *FindtheminorityGameView) GetViewerRoleOk() (*string, bool)`

GetViewerRoleOk returns a tuple with the ViewerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRole

`func (o *FindtheminorityGameView) SetViewerRole(v string)`

SetViewerRole sets ViewerRole field to given value.

### HasViewerRole

`func (o *FindtheminorityGameView) HasViewerRole() bool`

HasViewerRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


