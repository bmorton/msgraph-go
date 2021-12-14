# MicrosoftGraphDeviceConfigurationAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The assignment target for the device configuration. | [optional] 

## Methods

### NewMicrosoftGraphDeviceConfigurationAssignment

`func NewMicrosoftGraphDeviceConfigurationAssignment() *MicrosoftGraphDeviceConfigurationAssignment`

NewMicrosoftGraphDeviceConfigurationAssignment instantiates a new MicrosoftGraphDeviceConfigurationAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceConfigurationAssignmentWithDefaults

`func NewMicrosoftGraphDeviceConfigurationAssignmentWithDefaults() *MicrosoftGraphDeviceConfigurationAssignment`

NewMicrosoftGraphDeviceConfigurationAssignmentWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceConfigurationAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTarget

`func (o *MicrosoftGraphDeviceConfigurationAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphDeviceConfigurationAssignment) GetTargetOk() (*AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphDeviceConfigurationAssignment) SetTarget(v AnyOfobject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphDeviceConfigurationAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MicrosoftGraphDeviceConfigurationAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MicrosoftGraphDeviceConfigurationAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


