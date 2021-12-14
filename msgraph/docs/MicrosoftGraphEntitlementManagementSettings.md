# MicrosoftGraphEntitlementManagementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DurationUntilExternalUserDeletedAfterBlocked** | Pointer to **NullableString** | If externalUserLifecycleAction is blockSignInAndDelete, the duration, typically a number of days, after an external user is blocked from sign in before their account is deleted. | [optional] 
**ExternalUserLifecycleAction** | Pointer to [**AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction**](anyOf&lt;microsoft.graph.accessPackageExternalUserLifecycleAction&gt;.md) | Automatic action that the service should take when an external user&#39;s last access package assignment is removed. The possible values are: none, blockSignIn, blockSignInAndDelete, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphEntitlementManagementSettings

`func NewMicrosoftGraphEntitlementManagementSettings() *MicrosoftGraphEntitlementManagementSettings`

NewMicrosoftGraphEntitlementManagementSettings instantiates a new MicrosoftGraphEntitlementManagementSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEntitlementManagementSettingsWithDefaults

`func NewMicrosoftGraphEntitlementManagementSettingsWithDefaults() *MicrosoftGraphEntitlementManagementSettings`

NewMicrosoftGraphEntitlementManagementSettingsWithDefaults instantiates a new MicrosoftGraphEntitlementManagementSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEntitlementManagementSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEntitlementManagementSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEntitlementManagementSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEntitlementManagementSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDurationUntilExternalUserDeletedAfterBlocked

`func (o *MicrosoftGraphEntitlementManagementSettings) GetDurationUntilExternalUserDeletedAfterBlocked() string`

GetDurationUntilExternalUserDeletedAfterBlocked returns the DurationUntilExternalUserDeletedAfterBlocked field if non-nil, zero value otherwise.

### GetDurationUntilExternalUserDeletedAfterBlockedOk

`func (o *MicrosoftGraphEntitlementManagementSettings) GetDurationUntilExternalUserDeletedAfterBlockedOk() (*string, bool)`

GetDurationUntilExternalUserDeletedAfterBlockedOk returns a tuple with the DurationUntilExternalUserDeletedAfterBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUntilExternalUserDeletedAfterBlocked

`func (o *MicrosoftGraphEntitlementManagementSettings) SetDurationUntilExternalUserDeletedAfterBlocked(v string)`

SetDurationUntilExternalUserDeletedAfterBlocked sets DurationUntilExternalUserDeletedAfterBlocked field to given value.

### HasDurationUntilExternalUserDeletedAfterBlocked

`func (o *MicrosoftGraphEntitlementManagementSettings) HasDurationUntilExternalUserDeletedAfterBlocked() bool`

HasDurationUntilExternalUserDeletedAfterBlocked returns a boolean if a field has been set.

### SetDurationUntilExternalUserDeletedAfterBlockedNil

`func (o *MicrosoftGraphEntitlementManagementSettings) SetDurationUntilExternalUserDeletedAfterBlockedNil(b bool)`

 SetDurationUntilExternalUserDeletedAfterBlockedNil sets the value for DurationUntilExternalUserDeletedAfterBlocked to be an explicit nil

### UnsetDurationUntilExternalUserDeletedAfterBlocked
`func (o *MicrosoftGraphEntitlementManagementSettings) UnsetDurationUntilExternalUserDeletedAfterBlocked()`

UnsetDurationUntilExternalUserDeletedAfterBlocked ensures that no value is present for DurationUntilExternalUserDeletedAfterBlocked, not even an explicit nil
### GetExternalUserLifecycleAction

`func (o *MicrosoftGraphEntitlementManagementSettings) GetExternalUserLifecycleAction() AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction`

GetExternalUserLifecycleAction returns the ExternalUserLifecycleAction field if non-nil, zero value otherwise.

### GetExternalUserLifecycleActionOk

`func (o *MicrosoftGraphEntitlementManagementSettings) GetExternalUserLifecycleActionOk() (*AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction, bool)`

GetExternalUserLifecycleActionOk returns a tuple with the ExternalUserLifecycleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserLifecycleAction

`func (o *MicrosoftGraphEntitlementManagementSettings) SetExternalUserLifecycleAction(v AnyOfmicrosoftGraphAccessPackageExternalUserLifecycleAction)`

SetExternalUserLifecycleAction sets ExternalUserLifecycleAction field to given value.

### HasExternalUserLifecycleAction

`func (o *MicrosoftGraphEntitlementManagementSettings) HasExternalUserLifecycleAction() bool`

HasExternalUserLifecycleAction returns a boolean if a field has been set.

### SetExternalUserLifecycleActionNil

`func (o *MicrosoftGraphEntitlementManagementSettings) SetExternalUserLifecycleActionNil(b bool)`

 SetExternalUserLifecycleActionNil sets the value for ExternalUserLifecycleAction to be an explicit nil

### UnsetExternalUserLifecycleAction
`func (o *MicrosoftGraphEntitlementManagementSettings) UnsetExternalUserLifecycleAction()`

UnsetExternalUserLifecycleAction ensures that no value is present for ExternalUserLifecycleAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


