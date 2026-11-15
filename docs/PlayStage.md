# PlayStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Module** | Pointer to [**PlayModuleRef**](PlayModuleRef.md) |  | [optional] 

## Methods

### NewPlayStage

`func NewPlayStage() *PlayStage`

NewPlayStage instantiates a new PlayStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayStageWithDefaults

`func NewPlayStageWithDefaults() *PlayStage`

NewPlayStageWithDefaults instantiates a new PlayStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayStage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayStage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayStage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlayStage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModule

`func (o *PlayStage) GetModule() PlayModuleRef`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PlayStage) GetModuleOk() (*PlayModuleRef, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PlayStage) SetModule(v PlayModuleRef)`

SetModule sets Module field to given value.

### HasModule

`func (o *PlayStage) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


