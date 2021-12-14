# MicrosoftGraphAccessReviewNotificationRecipientItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationRecipientScope** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Determines the recipient of the notification email. | [optional] 
**NotificationTemplateType** | Pointer to **NullableString** | Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients. | [optional] 

## Methods

### NewMicrosoftGraphAccessReviewNotificationRecipientItem

`func NewMicrosoftGraphAccessReviewNotificationRecipientItem() *MicrosoftGraphAccessReviewNotificationRecipientItem`

NewMicrosoftGraphAccessReviewNotificationRecipientItem instantiates a new MicrosoftGraphAccessReviewNotificationRecipientItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessReviewNotificationRecipientItemWithDefaults

`func NewMicrosoftGraphAccessReviewNotificationRecipientItemWithDefaults() *MicrosoftGraphAccessReviewNotificationRecipientItem`

NewMicrosoftGraphAccessReviewNotificationRecipientItemWithDefaults instantiates a new MicrosoftGraphAccessReviewNotificationRecipientItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationRecipientScope

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) GetNotificationRecipientScope() AnyOfobject`

GetNotificationRecipientScope returns the NotificationRecipientScope field if non-nil, zero value otherwise.

### GetNotificationRecipientScopeOk

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) GetNotificationRecipientScopeOk() (*AnyOfobject, bool)`

GetNotificationRecipientScopeOk returns a tuple with the NotificationRecipientScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRecipientScope

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) SetNotificationRecipientScope(v AnyOfobject)`

SetNotificationRecipientScope sets NotificationRecipientScope field to given value.

### HasNotificationRecipientScope

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) HasNotificationRecipientScope() bool`

HasNotificationRecipientScope returns a boolean if a field has been set.

### SetNotificationRecipientScopeNil

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) SetNotificationRecipientScopeNil(b bool)`

 SetNotificationRecipientScopeNil sets the value for NotificationRecipientScope to be an explicit nil

### UnsetNotificationRecipientScope
`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) UnsetNotificationRecipientScope()`

UnsetNotificationRecipientScope ensures that no value is present for NotificationRecipientScope, not even an explicit nil
### GetNotificationTemplateType

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) GetNotificationTemplateType() string`

GetNotificationTemplateType returns the NotificationTemplateType field if non-nil, zero value otherwise.

### GetNotificationTemplateTypeOk

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) GetNotificationTemplateTypeOk() (*string, bool)`

GetNotificationTemplateTypeOk returns a tuple with the NotificationTemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTemplateType

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) SetNotificationTemplateType(v string)`

SetNotificationTemplateType sets NotificationTemplateType field to given value.

### HasNotificationTemplateType

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) HasNotificationTemplateType() bool`

HasNotificationTemplateType returns a boolean if a field has been set.

### SetNotificationTemplateTypeNil

`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) SetNotificationTemplateTypeNil(b bool)`

 SetNotificationTemplateTypeNil sets the value for NotificationTemplateType to be an explicit nil

### UnsetNotificationTemplateType
`func (o *MicrosoftGraphAccessReviewNotificationRecipientItem) UnsetNotificationTemplateType()`

UnsetNotificationTemplateType ensures that no value is present for NotificationTemplateType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


