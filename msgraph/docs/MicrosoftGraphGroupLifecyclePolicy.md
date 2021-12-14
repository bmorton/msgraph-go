# MicrosoftGraphGroupLifecyclePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AlternateNotificationEmails** | Pointer to **NullableString** | List of email address to send notifications for groups without owners. Multiple email address can be defined by separating email address with a semicolon. | [optional] 
**GroupLifetimeInDays** | Pointer to **NullableInt32** | Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the number of days defined. | [optional] 
**ManagedGroupTypes** | Pointer to **NullableString** | The group type for which the expiration policy applies. Possible values are All, Selected or None. | [optional] 

## Methods

### NewMicrosoftGraphGroupLifecyclePolicy

`func NewMicrosoftGraphGroupLifecyclePolicy() *MicrosoftGraphGroupLifecyclePolicy`

NewMicrosoftGraphGroupLifecyclePolicy instantiates a new MicrosoftGraphGroupLifecyclePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphGroupLifecyclePolicyWithDefaults

`func NewMicrosoftGraphGroupLifecyclePolicyWithDefaults() *MicrosoftGraphGroupLifecyclePolicy`

NewMicrosoftGraphGroupLifecyclePolicyWithDefaults instantiates a new MicrosoftGraphGroupLifecyclePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphGroupLifecyclePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlternateNotificationEmails

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetAlternateNotificationEmails() string`

GetAlternateNotificationEmails returns the AlternateNotificationEmails field if non-nil, zero value otherwise.

### GetAlternateNotificationEmailsOk

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetAlternateNotificationEmailsOk() (*string, bool)`

GetAlternateNotificationEmailsOk returns a tuple with the AlternateNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateNotificationEmails

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetAlternateNotificationEmails(v string)`

SetAlternateNotificationEmails sets AlternateNotificationEmails field to given value.

### HasAlternateNotificationEmails

`func (o *MicrosoftGraphGroupLifecyclePolicy) HasAlternateNotificationEmails() bool`

HasAlternateNotificationEmails returns a boolean if a field has been set.

### SetAlternateNotificationEmailsNil

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetAlternateNotificationEmailsNil(b bool)`

 SetAlternateNotificationEmailsNil sets the value for AlternateNotificationEmails to be an explicit nil

### UnsetAlternateNotificationEmails
`func (o *MicrosoftGraphGroupLifecyclePolicy) UnsetAlternateNotificationEmails()`

UnsetAlternateNotificationEmails ensures that no value is present for AlternateNotificationEmails, not even an explicit nil
### GetGroupLifetimeInDays

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetGroupLifetimeInDays() int32`

GetGroupLifetimeInDays returns the GroupLifetimeInDays field if non-nil, zero value otherwise.

### GetGroupLifetimeInDaysOk

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetGroupLifetimeInDaysOk() (*int32, bool)`

GetGroupLifetimeInDaysOk returns a tuple with the GroupLifetimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLifetimeInDays

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetGroupLifetimeInDays(v int32)`

SetGroupLifetimeInDays sets GroupLifetimeInDays field to given value.

### HasGroupLifetimeInDays

`func (o *MicrosoftGraphGroupLifecyclePolicy) HasGroupLifetimeInDays() bool`

HasGroupLifetimeInDays returns a boolean if a field has been set.

### SetGroupLifetimeInDaysNil

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetGroupLifetimeInDaysNil(b bool)`

 SetGroupLifetimeInDaysNil sets the value for GroupLifetimeInDays to be an explicit nil

### UnsetGroupLifetimeInDays
`func (o *MicrosoftGraphGroupLifecyclePolicy) UnsetGroupLifetimeInDays()`

UnsetGroupLifetimeInDays ensures that no value is present for GroupLifetimeInDays, not even an explicit nil
### GetManagedGroupTypes

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetManagedGroupTypes() string`

GetManagedGroupTypes returns the ManagedGroupTypes field if non-nil, zero value otherwise.

### GetManagedGroupTypesOk

`func (o *MicrosoftGraphGroupLifecyclePolicy) GetManagedGroupTypesOk() (*string, bool)`

GetManagedGroupTypesOk returns a tuple with the ManagedGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedGroupTypes

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetManagedGroupTypes(v string)`

SetManagedGroupTypes sets ManagedGroupTypes field to given value.

### HasManagedGroupTypes

`func (o *MicrosoftGraphGroupLifecyclePolicy) HasManagedGroupTypes() bool`

HasManagedGroupTypes returns a boolean if a field has been set.

### SetManagedGroupTypesNil

`func (o *MicrosoftGraphGroupLifecyclePolicy) SetManagedGroupTypesNil(b bool)`

 SetManagedGroupTypesNil sets the value for ManagedGroupTypes to be an explicit nil

### UnsetManagedGroupTypes
`func (o *MicrosoftGraphGroupLifecyclePolicy) UnsetManagedGroupTypes()`

UnsetManagedGroupTypes ensures that no value is present for ManagedGroupTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


