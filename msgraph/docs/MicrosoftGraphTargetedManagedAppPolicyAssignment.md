# MicrosoftGraphTargetedManagedAppPolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Identifier for deployment to a group or app | [optional] 

## Methods

### NewMicrosoftGraphTargetedManagedAppPolicyAssignment

`func NewMicrosoftGraphTargetedManagedAppPolicyAssignment() *MicrosoftGraphTargetedManagedAppPolicyAssignment`

NewMicrosoftGraphTargetedManagedAppPolicyAssignment instantiates a new MicrosoftGraphTargetedManagedAppPolicyAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTargetedManagedAppPolicyAssignmentWithDefaults

`func NewMicrosoftGraphTargetedManagedAppPolicyAssignmentWithDefaults() *MicrosoftGraphTargetedManagedAppPolicyAssignment`

NewMicrosoftGraphTargetedManagedAppPolicyAssignmentWithDefaults instantiates a new MicrosoftGraphTargetedManagedAppPolicyAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTarget

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetTargetOk() (*AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) SetTarget(v AnyOfobject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


