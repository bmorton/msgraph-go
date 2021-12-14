# ManagedEBookAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallIntent** | Pointer to [**AnyOfmicrosoftGraphInstallIntent**](anyOf&lt;microsoft.graph.installIntent&gt;.md) | The install intent for eBook. Possible values are: available, required, uninstall, availableWithoutEnrollment. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The assignment target for eBook. | [optional] 

## Methods

### NewManagedEBookAssignment

`func NewManagedEBookAssignment() *ManagedEBookAssignment`

NewManagedEBookAssignment instantiates a new ManagedEBookAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedEBookAssignmentWithDefaults

`func NewManagedEBookAssignmentWithDefaults() *ManagedEBookAssignment`

NewManagedEBookAssignmentWithDefaults instantiates a new ManagedEBookAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallIntent

`func (o *ManagedEBookAssignment) GetInstallIntent() AnyOfmicrosoftGraphInstallIntent`

GetInstallIntent returns the InstallIntent field if non-nil, zero value otherwise.

### GetInstallIntentOk

`func (o *ManagedEBookAssignment) GetInstallIntentOk() (*AnyOfmicrosoftGraphInstallIntent, bool)`

GetInstallIntentOk returns a tuple with the InstallIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallIntent

`func (o *ManagedEBookAssignment) SetInstallIntent(v AnyOfmicrosoftGraphInstallIntent)`

SetInstallIntent sets InstallIntent field to given value.

### HasInstallIntent

`func (o *ManagedEBookAssignment) HasInstallIntent() bool`

HasInstallIntent returns a boolean if a field has been set.

### SetInstallIntentNil

`func (o *ManagedEBookAssignment) SetInstallIntentNil(b bool)`

 SetInstallIntentNil sets the value for InstallIntent to be an explicit nil

### UnsetInstallIntent
`func (o *ManagedEBookAssignment) UnsetInstallIntent()`

UnsetInstallIntent ensures that no value is present for InstallIntent, not even an explicit nil
### GetTarget

`func (o *ManagedEBookAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ManagedEBookAssignment) GetTargetOk() (*AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ManagedEBookAssignment) SetTarget(v AnyOfobject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ManagedEBookAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ManagedEBookAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ManagedEBookAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


