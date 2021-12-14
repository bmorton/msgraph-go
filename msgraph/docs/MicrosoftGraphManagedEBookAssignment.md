# MicrosoftGraphManagedEBookAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**InstallIntent** | Pointer to [**AnyOfmicrosoftGraphInstallIntent**](anyOf&lt;microsoft.graph.installIntent&gt;.md) | The install intent for eBook. Possible values are: available, required, uninstall, availableWithoutEnrollment. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The assignment target for eBook. | [optional] 

## Methods

### NewMicrosoftGraphManagedEBookAssignment

`func NewMicrosoftGraphManagedEBookAssignment() *MicrosoftGraphManagedEBookAssignment`

NewMicrosoftGraphManagedEBookAssignment instantiates a new MicrosoftGraphManagedEBookAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedEBookAssignmentWithDefaults

`func NewMicrosoftGraphManagedEBookAssignmentWithDefaults() *MicrosoftGraphManagedEBookAssignment`

NewMicrosoftGraphManagedEBookAssignmentWithDefaults instantiates a new MicrosoftGraphManagedEBookAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphManagedEBookAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedEBookAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphManagedEBookAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphManagedEBookAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallIntent

`func (o *MicrosoftGraphManagedEBookAssignment) GetInstallIntent() AnyOfmicrosoftGraphInstallIntent`

GetInstallIntent returns the InstallIntent field if non-nil, zero value otherwise.

### GetInstallIntentOk

`func (o *MicrosoftGraphManagedEBookAssignment) GetInstallIntentOk() (*AnyOfmicrosoftGraphInstallIntent, bool)`

GetInstallIntentOk returns a tuple with the InstallIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallIntent

`func (o *MicrosoftGraphManagedEBookAssignment) SetInstallIntent(v AnyOfmicrosoftGraphInstallIntent)`

SetInstallIntent sets InstallIntent field to given value.

### HasInstallIntent

`func (o *MicrosoftGraphManagedEBookAssignment) HasInstallIntent() bool`

HasInstallIntent returns a boolean if a field has been set.

### SetInstallIntentNil

`func (o *MicrosoftGraphManagedEBookAssignment) SetInstallIntentNil(b bool)`

 SetInstallIntentNil sets the value for InstallIntent to be an explicit nil

### UnsetInstallIntent
`func (o *MicrosoftGraphManagedEBookAssignment) UnsetInstallIntent()`

UnsetInstallIntent ensures that no value is present for InstallIntent, not even an explicit nil
### GetTarget

`func (o *MicrosoftGraphManagedEBookAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphManagedEBookAssignment) GetTargetOk() (*AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphManagedEBookAssignment) SetTarget(v AnyOfobject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphManagedEBookAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MicrosoftGraphManagedEBookAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MicrosoftGraphManagedEBookAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


