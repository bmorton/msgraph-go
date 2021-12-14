# MicrosoftGraphDeviceActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **NullableString** | Action name | [optional] 
**ActionState** | Pointer to [**AnyOfmicrosoftGraphActionState**](anyOf&lt;microsoft.graph.actionState&gt;.md) | State of the action. Possible values are: none, pending, canceled, active, done, failed, notSupported. | [optional] 
**LastUpdatedDateTime** | Pointer to **time.Time** | Time the action state was last updated | [optional] 
**StartDateTime** | Pointer to **time.Time** | Time the action was initiated | [optional] 

## Methods

### NewMicrosoftGraphDeviceActionResult

`func NewMicrosoftGraphDeviceActionResult() *MicrosoftGraphDeviceActionResult`

NewMicrosoftGraphDeviceActionResult instantiates a new MicrosoftGraphDeviceActionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceActionResultWithDefaults

`func NewMicrosoftGraphDeviceActionResultWithDefaults() *MicrosoftGraphDeviceActionResult`

NewMicrosoftGraphDeviceActionResultWithDefaults instantiates a new MicrosoftGraphDeviceActionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionName

`func (o *MicrosoftGraphDeviceActionResult) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphDeviceActionResult) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *MicrosoftGraphDeviceActionResult) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *MicrosoftGraphDeviceActionResult) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionNameNil

`func (o *MicrosoftGraphDeviceActionResult) SetActionNameNil(b bool)`

 SetActionNameNil sets the value for ActionName to be an explicit nil

### UnsetActionName
`func (o *MicrosoftGraphDeviceActionResult) UnsetActionName()`

UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
### GetActionState

`func (o *MicrosoftGraphDeviceActionResult) GetActionState() AnyOfmicrosoftGraphActionState`

GetActionState returns the ActionState field if non-nil, zero value otherwise.

### GetActionStateOk

`func (o *MicrosoftGraphDeviceActionResult) GetActionStateOk() (*AnyOfmicrosoftGraphActionState, bool)`

GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionState

`func (o *MicrosoftGraphDeviceActionResult) SetActionState(v AnyOfmicrosoftGraphActionState)`

SetActionState sets ActionState field to given value.

### HasActionState

`func (o *MicrosoftGraphDeviceActionResult) HasActionState() bool`

HasActionState returns a boolean if a field has been set.

### SetActionStateNil

`func (o *MicrosoftGraphDeviceActionResult) SetActionStateNil(b bool)`

 SetActionStateNil sets the value for ActionState to be an explicit nil

### UnsetActionState
`func (o *MicrosoftGraphDeviceActionResult) UnsetActionState()`

UnsetActionState ensures that no value is present for ActionState, not even an explicit nil
### GetLastUpdatedDateTime

`func (o *MicrosoftGraphDeviceActionResult) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *MicrosoftGraphDeviceActionResult) GetLastUpdatedDateTimeOk() (*time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDateTime

`func (o *MicrosoftGraphDeviceActionResult) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime sets LastUpdatedDateTime field to given value.

### HasLastUpdatedDateTime

`func (o *MicrosoftGraphDeviceActionResult) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphDeviceActionResult) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphDeviceActionResult) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphDeviceActionResult) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphDeviceActionResult) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


