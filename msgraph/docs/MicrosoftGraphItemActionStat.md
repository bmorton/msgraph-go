# MicrosoftGraphItemActionStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionCount** | Pointer to **NullableInt32** | The number of times the action took place. Read-only. | [optional] 
**ActorCount** | Pointer to **NullableInt32** | The number of distinct actors that performed the action. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphItemActionStat

`func NewMicrosoftGraphItemActionStat() *MicrosoftGraphItemActionStat`

NewMicrosoftGraphItemActionStat instantiates a new MicrosoftGraphItemActionStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphItemActionStatWithDefaults

`func NewMicrosoftGraphItemActionStatWithDefaults() *MicrosoftGraphItemActionStat`

NewMicrosoftGraphItemActionStatWithDefaults instantiates a new MicrosoftGraphItemActionStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionCount

`func (o *MicrosoftGraphItemActionStat) GetActionCount() int32`

GetActionCount returns the ActionCount field if non-nil, zero value otherwise.

### GetActionCountOk

`func (o *MicrosoftGraphItemActionStat) GetActionCountOk() (*int32, bool)`

GetActionCountOk returns a tuple with the ActionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCount

`func (o *MicrosoftGraphItemActionStat) SetActionCount(v int32)`

SetActionCount sets ActionCount field to given value.

### HasActionCount

`func (o *MicrosoftGraphItemActionStat) HasActionCount() bool`

HasActionCount returns a boolean if a field has been set.

### SetActionCountNil

`func (o *MicrosoftGraphItemActionStat) SetActionCountNil(b bool)`

 SetActionCountNil sets the value for ActionCount to be an explicit nil

### UnsetActionCount
`func (o *MicrosoftGraphItemActionStat) UnsetActionCount()`

UnsetActionCount ensures that no value is present for ActionCount, not even an explicit nil
### GetActorCount

`func (o *MicrosoftGraphItemActionStat) GetActorCount() int32`

GetActorCount returns the ActorCount field if non-nil, zero value otherwise.

### GetActorCountOk

`func (o *MicrosoftGraphItemActionStat) GetActorCountOk() (*int32, bool)`

GetActorCountOk returns a tuple with the ActorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorCount

`func (o *MicrosoftGraphItemActionStat) SetActorCount(v int32)`

SetActorCount sets ActorCount field to given value.

### HasActorCount

`func (o *MicrosoftGraphItemActionStat) HasActorCount() bool`

HasActorCount returns a boolean if a field has been set.

### SetActorCountNil

`func (o *MicrosoftGraphItemActionStat) SetActorCountNil(b bool)`

 SetActorCountNil sets the value for ActorCount to be an explicit nil

### UnsetActorCount
`func (o *MicrosoftGraphItemActionStat) UnsetActorCount()`

UnsetActorCount ensures that no value is present for ActorCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


