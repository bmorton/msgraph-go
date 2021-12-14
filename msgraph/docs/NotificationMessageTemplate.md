# NotificationMessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandingOptions** | Pointer to [**AnyOfmicrosoftGraphNotificationTemplateBrandingOptions**](anyOf&lt;microsoft.graph.notificationTemplateBrandingOptions&gt;.md) | The Message Template Branding Options. Branding is defined in the Intune Admin Console. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation. | [optional] 
**DefaultLocale** | Pointer to **NullableString** | The default locale to fallback onto when the requested locale is not available. | [optional] 
**DisplayName** | Pointer to **string** | Display name for the Notification Message Template. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**LocalizedNotificationMessages** | Pointer to [**[]MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md) | The list of localized messages for this Notification Message Template. | [optional] 

## Methods

### NewNotificationMessageTemplate

`func NewNotificationMessageTemplate() *NotificationMessageTemplate`

NewNotificationMessageTemplate instantiates a new NotificationMessageTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMessageTemplateWithDefaults

`func NewNotificationMessageTemplateWithDefaults() *NotificationMessageTemplate`

NewNotificationMessageTemplateWithDefaults instantiates a new NotificationMessageTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandingOptions

`func (o *NotificationMessageTemplate) GetBrandingOptions() AnyOfmicrosoftGraphNotificationTemplateBrandingOptions`

GetBrandingOptions returns the BrandingOptions field if non-nil, zero value otherwise.

### GetBrandingOptionsOk

`func (o *NotificationMessageTemplate) GetBrandingOptionsOk() (*AnyOfmicrosoftGraphNotificationTemplateBrandingOptions, bool)`

GetBrandingOptionsOk returns a tuple with the BrandingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingOptions

`func (o *NotificationMessageTemplate) SetBrandingOptions(v AnyOfmicrosoftGraphNotificationTemplateBrandingOptions)`

SetBrandingOptions sets BrandingOptions field to given value.

### HasBrandingOptions

`func (o *NotificationMessageTemplate) HasBrandingOptions() bool`

HasBrandingOptions returns a boolean if a field has been set.

### SetBrandingOptionsNil

`func (o *NotificationMessageTemplate) SetBrandingOptionsNil(b bool)`

 SetBrandingOptionsNil sets the value for BrandingOptions to be an explicit nil

### UnsetBrandingOptions
`func (o *NotificationMessageTemplate) UnsetBrandingOptions()`

UnsetBrandingOptions ensures that no value is present for BrandingOptions, not even an explicit nil
### GetDefaultLocale

`func (o *NotificationMessageTemplate) GetDefaultLocale() string`

GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.

### GetDefaultLocaleOk

`func (o *NotificationMessageTemplate) GetDefaultLocaleOk() (*string, bool)`

GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocale

`func (o *NotificationMessageTemplate) SetDefaultLocale(v string)`

SetDefaultLocale sets DefaultLocale field to given value.

### HasDefaultLocale

`func (o *NotificationMessageTemplate) HasDefaultLocale() bool`

HasDefaultLocale returns a boolean if a field has been set.

### SetDefaultLocaleNil

`func (o *NotificationMessageTemplate) SetDefaultLocaleNil(b bool)`

 SetDefaultLocaleNil sets the value for DefaultLocale to be an explicit nil

### UnsetDefaultLocale
`func (o *NotificationMessageTemplate) UnsetDefaultLocale()`

UnsetDefaultLocale ensures that no value is present for DefaultLocale, not even an explicit nil
### GetDisplayName

`func (o *NotificationMessageTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NotificationMessageTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NotificationMessageTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NotificationMessageTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *NotificationMessageTemplate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *NotificationMessageTemplate) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *NotificationMessageTemplate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *NotificationMessageTemplate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetLocalizedNotificationMessages

`func (o *NotificationMessageTemplate) GetLocalizedNotificationMessages() []MicrosoftGraphLocalizedNotificationMessage`

GetLocalizedNotificationMessages returns the LocalizedNotificationMessages field if non-nil, zero value otherwise.

### GetLocalizedNotificationMessagesOk

`func (o *NotificationMessageTemplate) GetLocalizedNotificationMessagesOk() (*[]MicrosoftGraphLocalizedNotificationMessage, bool)`

GetLocalizedNotificationMessagesOk returns a tuple with the LocalizedNotificationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedNotificationMessages

`func (o *NotificationMessageTemplate) SetLocalizedNotificationMessages(v []MicrosoftGraphLocalizedNotificationMessage)`

SetLocalizedNotificationMessages sets LocalizedNotificationMessages field to given value.

### HasLocalizedNotificationMessages

`func (o *NotificationMessageTemplate) HasLocalizedNotificationMessages() bool`

HasLocalizedNotificationMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


