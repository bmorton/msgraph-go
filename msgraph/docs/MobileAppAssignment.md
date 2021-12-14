# MobileAppAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intent** | Pointer to [**AnyOfmicrosoftGraphInstallIntent**](anyOf&lt;microsoft.graph.installIntent&gt;.md) | The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment. | [optional] 
**Settings** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The settings for target assignment defined by the admin. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The target group assignment defined by the admin. | [optional] 

## Methods

### NewMobileAppAssignment

`func NewMobileAppAssignment() *MobileAppAssignment`

NewMobileAppAssignment instantiates a new MobileAppAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppAssignmentWithDefaults

`func NewMobileAppAssignmentWithDefaults() *MobileAppAssignment`

NewMobileAppAssignmentWithDefaults instantiates a new MobileAppAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntent

`func (o *MobileAppAssignment) GetIntent() AnyOfmicrosoftGraphInstallIntent`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *MobileAppAssignment) GetIntentOk() (*AnyOfmicrosoftGraphInstallIntent, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *MobileAppAssignment) SetIntent(v AnyOfmicrosoftGraphInstallIntent)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *MobileAppAssignment) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### SetIntentNil

`func (o *MobileAppAssignment) SetIntentNil(b bool)`

 SetIntentNil sets the value for Intent to be an explicit nil

### UnsetIntent
`func (o *MobileAppAssignment) UnsetIntent()`

UnsetIntent ensures that no value is present for Intent, not even an explicit nil
### GetSettings

`func (o *MobileAppAssignment) GetSettings() AnyOfobject`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MobileAppAssignment) GetSettingsOk() (*AnyOfobject, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MobileAppAssignment) SetSettings(v AnyOfobject)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MobileAppAssignment) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *MobileAppAssignment) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *MobileAppAssignment) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetTarget

`func (o *MobileAppAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MobileAppAssignment) GetTargetOk() (*AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MobileAppAssignment) SetTarget(v AnyOfobject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MobileAppAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MobileAppAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MobileAppAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


