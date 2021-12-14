# GroupLifecyclePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateNotificationEmails** | Pointer to **NullableString** | List of email address to send notifications for groups without owners. Multiple email address can be defined by separating email address with a semicolon. | [optional] 
**GroupLifetimeInDays** | Pointer to **NullableInt32** | Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the number of days defined. | [optional] 
**ManagedGroupTypes** | Pointer to **NullableString** | The group type for which the expiration policy applies. Possible values are All, Selected or None. | [optional] 

## Methods

### NewGroupLifecyclePolicy

`func NewGroupLifecyclePolicy() *GroupLifecyclePolicy`

NewGroupLifecyclePolicy instantiates a new GroupLifecyclePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupLifecyclePolicyWithDefaults

`func NewGroupLifecyclePolicyWithDefaults() *GroupLifecyclePolicy`

NewGroupLifecyclePolicyWithDefaults instantiates a new GroupLifecyclePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateNotificationEmails

`func (o *GroupLifecyclePolicy) GetAlternateNotificationEmails() string`

GetAlternateNotificationEmails returns the AlternateNotificationEmails field if non-nil, zero value otherwise.

### GetAlternateNotificationEmailsOk

`func (o *GroupLifecyclePolicy) GetAlternateNotificationEmailsOk() (*string, bool)`

GetAlternateNotificationEmailsOk returns a tuple with the AlternateNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateNotificationEmails

`func (o *GroupLifecyclePolicy) SetAlternateNotificationEmails(v string)`

SetAlternateNotificationEmails sets AlternateNotificationEmails field to given value.

### HasAlternateNotificationEmails

`func (o *GroupLifecyclePolicy) HasAlternateNotificationEmails() bool`

HasAlternateNotificationEmails returns a boolean if a field has been set.

### SetAlternateNotificationEmailsNil

`func (o *GroupLifecyclePolicy) SetAlternateNotificationEmailsNil(b bool)`

 SetAlternateNotificationEmailsNil sets the value for AlternateNotificationEmails to be an explicit nil

### UnsetAlternateNotificationEmails
`func (o *GroupLifecyclePolicy) UnsetAlternateNotificationEmails()`

UnsetAlternateNotificationEmails ensures that no value is present for AlternateNotificationEmails, not even an explicit nil
### GetGroupLifetimeInDays

`func (o *GroupLifecyclePolicy) GetGroupLifetimeInDays() int32`

GetGroupLifetimeInDays returns the GroupLifetimeInDays field if non-nil, zero value otherwise.

### GetGroupLifetimeInDaysOk

`func (o *GroupLifecyclePolicy) GetGroupLifetimeInDaysOk() (*int32, bool)`

GetGroupLifetimeInDaysOk returns a tuple with the GroupLifetimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLifetimeInDays

`func (o *GroupLifecyclePolicy) SetGroupLifetimeInDays(v int32)`

SetGroupLifetimeInDays sets GroupLifetimeInDays field to given value.

### HasGroupLifetimeInDays

`func (o *GroupLifecyclePolicy) HasGroupLifetimeInDays() bool`

HasGroupLifetimeInDays returns a boolean if a field has been set.

### SetGroupLifetimeInDaysNil

`func (o *GroupLifecyclePolicy) SetGroupLifetimeInDaysNil(b bool)`

 SetGroupLifetimeInDaysNil sets the value for GroupLifetimeInDays to be an explicit nil

### UnsetGroupLifetimeInDays
`func (o *GroupLifecyclePolicy) UnsetGroupLifetimeInDays()`

UnsetGroupLifetimeInDays ensures that no value is present for GroupLifetimeInDays, not even an explicit nil
### GetManagedGroupTypes

`func (o *GroupLifecyclePolicy) GetManagedGroupTypes() string`

GetManagedGroupTypes returns the ManagedGroupTypes field if non-nil, zero value otherwise.

### GetManagedGroupTypesOk

`func (o *GroupLifecyclePolicy) GetManagedGroupTypesOk() (*string, bool)`

GetManagedGroupTypesOk returns a tuple with the ManagedGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedGroupTypes

`func (o *GroupLifecyclePolicy) SetManagedGroupTypes(v string)`

SetManagedGroupTypes sets ManagedGroupTypes field to given value.

### HasManagedGroupTypes

`func (o *GroupLifecyclePolicy) HasManagedGroupTypes() bool`

HasManagedGroupTypes returns a boolean if a field has been set.

### SetManagedGroupTypesNil

`func (o *GroupLifecyclePolicy) SetManagedGroupTypesNil(b bool)`

 SetManagedGroupTypesNil sets the value for ManagedGroupTypes to be an explicit nil

### UnsetManagedGroupTypes
`func (o *GroupLifecyclePolicy) UnsetManagedGroupTypes()`

UnsetManagedGroupTypes ensures that no value is present for ManagedGroupTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


