# EntitlementManagementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationUntilExternalUserDeletedAfterBlocked** | Pointer to **NullableString** | If externalUserLifecycleAction is blockSignInAndDelete, the duration, typically a number of days, after an external user is blocked from sign in before their account is deleted. | [optional] 
**ExternalUserLifecycleAction** | Pointer to [**AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction**](anyOf&lt;microsoft.graph.accessPackageExternalUserLifecycleAction&gt;.md) | Automatic action that the service should take when an external user&#39;s last access package assignment is removed. The possible values are: none, blockSignIn, blockSignInAndDelete, unknownFutureValue. | [optional] 

## Methods

### NewEntitlementManagementSettings

`func NewEntitlementManagementSettings() *EntitlementManagementSettings`

NewEntitlementManagementSettings instantiates a new EntitlementManagementSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementManagementSettingsWithDefaults

`func NewEntitlementManagementSettingsWithDefaults() *EntitlementManagementSettings`

NewEntitlementManagementSettingsWithDefaults instantiates a new EntitlementManagementSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationUntilExternalUserDeletedAfterBlocked

`func (o *EntitlementManagementSettings) GetDurationUntilExternalUserDeletedAfterBlocked() string`

GetDurationUntilExternalUserDeletedAfterBlocked returns the DurationUntilExternalUserDeletedAfterBlocked field if non-nil, zero value otherwise.

### GetDurationUntilExternalUserDeletedAfterBlockedOk

`func (o *EntitlementManagementSettings) GetDurationUntilExternalUserDeletedAfterBlockedOk() (*string, bool)`

GetDurationUntilExternalUserDeletedAfterBlockedOk returns a tuple with the DurationUntilExternalUserDeletedAfterBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUntilExternalUserDeletedAfterBlocked

`func (o *EntitlementManagementSettings) SetDurationUntilExternalUserDeletedAfterBlocked(v string)`

SetDurationUntilExternalUserDeletedAfterBlocked sets DurationUntilExternalUserDeletedAfterBlocked field to given value.

### HasDurationUntilExternalUserDeletedAfterBlocked

`func (o *EntitlementManagementSettings) HasDurationUntilExternalUserDeletedAfterBlocked() bool`

HasDurationUntilExternalUserDeletedAfterBlocked returns a boolean if a field has been set.

### SetDurationUntilExternalUserDeletedAfterBlockedNil

`func (o *EntitlementManagementSettings) SetDurationUntilExternalUserDeletedAfterBlockedNil(b bool)`

 SetDurationUntilExternalUserDeletedAfterBlockedNil sets the value for DurationUntilExternalUserDeletedAfterBlocked to be an explicit nil

### UnsetDurationUntilExternalUserDeletedAfterBlocked
`func (o *EntitlementManagementSettings) UnsetDurationUntilExternalUserDeletedAfterBlocked()`

UnsetDurationUntilExternalUserDeletedAfterBlocked ensures that no value is present for DurationUntilExternalUserDeletedAfterBlocked, not even an explicit nil
### GetExternalUserLifecycleAction

`func (o *EntitlementManagementSettings) GetExternalUserLifecycleAction() AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction`

GetExternalUserLifecycleAction returns the ExternalUserLifecycleAction field if non-nil, zero value otherwise.

### GetExternalUserLifecycleActionOk

`func (o *EntitlementManagementSettings) GetExternalUserLifecycleActionOk() (*AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction, bool)`

GetExternalUserLifecycleActionOk returns a tuple with the ExternalUserLifecycleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserLifecycleAction

`func (o *EntitlementManagementSettings) SetExternalUserLifecycleAction(v AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction)`

SetExternalUserLifecycleAction sets ExternalUserLifecycleAction field to given value.

### HasExternalUserLifecycleAction

`func (o *EntitlementManagementSettings) HasExternalUserLifecycleAction() bool`

HasExternalUserLifecycleAction returns a boolean if a field has been set.

### SetExternalUserLifecycleActionNil

`func (o *EntitlementManagementSettings) SetExternalUserLifecycleActionNil(b bool)`

 SetExternalUserLifecycleActionNil sets the value for ExternalUserLifecycleAction to be an explicit nil

### UnsetExternalUserLifecycleAction
`func (o *EntitlementManagementSettings) UnsetExternalUserLifecycleAction()`

UnsetExternalUserLifecycleAction ensures that no value is present for ExternalUserLifecycleAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


