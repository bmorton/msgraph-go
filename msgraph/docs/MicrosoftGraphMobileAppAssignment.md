# MicrosoftGraphMobileAppAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Intent** | Pointer to [**AnyOfmicrosoftGraphInstallIntent**](anyOf&lt;microsoft.graph.installIntent&gt;.md) | The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment. | [optional] 
**Settings** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The settings for target assignment defined by the admin. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The target group assignment defined by the admin. | [optional] 

## Methods

### NewMicrosoftGraphMobileAppAssignment

`func NewMicrosoftGraphMobileAppAssignment() *MicrosoftGraphMobileAppAssignment`

NewMicrosoftGraphMobileAppAssignment instantiates a new MicrosoftGraphMobileAppAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMobileAppAssignmentWithDefaults

`func NewMicrosoftGraphMobileAppAssignmentWithDefaults() *MicrosoftGraphMobileAppAssignment`

NewMicrosoftGraphMobileAppAssignmentWithDefaults instantiates a new MicrosoftGraphMobileAppAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMobileAppAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMobileAppAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMobileAppAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMobileAppAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntent

`func (o *MicrosoftGraphMobileAppAssignment) GetIntent() AnyOfmicrosoftGraphInstallIntent`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *MicrosoftGraphMobileAppAssignment) GetIntentOk() (*AnyOfmicrosoftGraphInstallIntent, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *MicrosoftGraphMobileAppAssignment) SetIntent(v AnyOfmicrosoftGraphInstallIntent)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *MicrosoftGraphMobileAppAssignment) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### SetIntentNil

`func (o *MicrosoftGraphMobileAppAssignment) SetIntentNil(b bool)`

 SetIntentNil sets the value for Intent to be an explicit nil

### UnsetIntent
`func (o *MicrosoftGraphMobileAppAssignment) UnsetIntent()`

UnsetIntent ensures that no value is present for Intent, not even an explicit nil
### GetSettings

`func (o *MicrosoftGraphMobileAppAssignment) GetSettings() AnyOfobject`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphMobileAppAssignment) GetSettingsOk() (*AnyOfobject, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MicrosoftGraphMobileAppAssignment) SetSettings(v AnyOfobject)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MicrosoftGraphMobileAppAssignment) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *MicrosoftGraphMobileAppAssignment) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *MicrosoftGraphMobileAppAssignment) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetTarget

`func (o *MicrosoftGraphMobileAppAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphMobileAppAssignment) GetTargetOk() (*AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphMobileAppAssignment) SetTarget(v AnyOfobject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphMobileAppAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MicrosoftGraphMobileAppAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MicrosoftGraphMobileAppAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


