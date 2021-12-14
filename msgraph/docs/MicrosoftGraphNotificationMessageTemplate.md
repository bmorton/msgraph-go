# MicrosoftGraphNotificationMessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**BrandingOptions** | Pointer to [**AnyOfmicrosoftGraphNotificationTemplateBrandingOptions**](anyOf&lt;microsoft.graph.notificationTemplateBrandingOptions&gt;.md) | The Message Template Branding Options. Branding is defined in the Intune Admin Console. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation. | [optional] 
**DefaultLocale** | Pointer to **NullableString** | The default locale to fallback onto when the requested locale is not available. | [optional] 
**DisplayName** | Pointer to **string** | Display name for the Notification Message Template. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**LocalizedNotificationMessages** | Pointer to [**[]MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md) | The list of localized messages for this Notification Message Template. | [optional] 

## Methods

### NewMicrosoftGraphNotificationMessageTemplate

`func NewMicrosoftGraphNotificationMessageTemplate() *MicrosoftGraphNotificationMessageTemplate`

NewMicrosoftGraphNotificationMessageTemplate instantiates a new MicrosoftGraphNotificationMessageTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphNotificationMessageTemplateWithDefaults

`func NewMicrosoftGraphNotificationMessageTemplateWithDefaults() *MicrosoftGraphNotificationMessageTemplate`

NewMicrosoftGraphNotificationMessageTemplateWithDefaults instantiates a new MicrosoftGraphNotificationMessageTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphNotificationMessageTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphNotificationMessageTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphNotificationMessageTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphNotificationMessageTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBrandingOptions

`func (o *MicrosoftGraphNotificationMessageTemplate) GetBrandingOptions() AnyOfmicrosoftGraphNotificationTemplateBrandingOptions`

GetBrandingOptions returns the BrandingOptions field if non-nil, zero value otherwise.

### GetBrandingOptionsOk

`func (o *MicrosoftGraphNotificationMessageTemplate) GetBrandingOptionsOk() (*AnyOfmicrosoftGraphNotificationTemplateBrandingOptions, bool)`

GetBrandingOptionsOk returns a tuple with the BrandingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingOptions

`func (o *MicrosoftGraphNotificationMessageTemplate) SetBrandingOptions(v AnyOfmicrosoftGraphNotificationTemplateBrandingOptions)`

SetBrandingOptions sets BrandingOptions field to given value.

### HasBrandingOptions

`func (o *MicrosoftGraphNotificationMessageTemplate) HasBrandingOptions() bool`

HasBrandingOptions returns a boolean if a field has been set.

### SetBrandingOptionsNil

`func (o *MicrosoftGraphNotificationMessageTemplate) SetBrandingOptionsNil(b bool)`

 SetBrandingOptionsNil sets the value for BrandingOptions to be an explicit nil

### UnsetBrandingOptions
`func (o *MicrosoftGraphNotificationMessageTemplate) UnsetBrandingOptions()`

UnsetBrandingOptions ensures that no value is present for BrandingOptions, not even an explicit nil
### GetDefaultLocale

`func (o *MicrosoftGraphNotificationMessageTemplate) GetDefaultLocale() string`

GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.

### GetDefaultLocaleOk

`func (o *MicrosoftGraphNotificationMessageTemplate) GetDefaultLocaleOk() (*string, bool)`

GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocale

`func (o *MicrosoftGraphNotificationMessageTemplate) SetDefaultLocale(v string)`

SetDefaultLocale sets DefaultLocale field to given value.

### HasDefaultLocale

`func (o *MicrosoftGraphNotificationMessageTemplate) HasDefaultLocale() bool`

HasDefaultLocale returns a boolean if a field has been set.

### SetDefaultLocaleNil

`func (o *MicrosoftGraphNotificationMessageTemplate) SetDefaultLocaleNil(b bool)`

 SetDefaultLocaleNil sets the value for DefaultLocale to be an explicit nil

### UnsetDefaultLocale
`func (o *MicrosoftGraphNotificationMessageTemplate) UnsetDefaultLocale()`

UnsetDefaultLocale ensures that no value is present for DefaultLocale, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphNotificationMessageTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphNotificationMessageTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphNotificationMessageTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphNotificationMessageTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphNotificationMessageTemplate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphNotificationMessageTemplate) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphNotificationMessageTemplate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphNotificationMessageTemplate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetLocalizedNotificationMessages

`func (o *MicrosoftGraphNotificationMessageTemplate) GetLocalizedNotificationMessages() []MicrosoftGraphLocalizedNotificationMessage`

GetLocalizedNotificationMessages returns the LocalizedNotificationMessages field if non-nil, zero value otherwise.

### GetLocalizedNotificationMessagesOk

`func (o *MicrosoftGraphNotificationMessageTemplate) GetLocalizedNotificationMessagesOk() (*[]MicrosoftGraphLocalizedNotificationMessage, bool)`

GetLocalizedNotificationMessagesOk returns a tuple with the LocalizedNotificationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedNotificationMessages

`func (o *MicrosoftGraphNotificationMessageTemplate) SetLocalizedNotificationMessages(v []MicrosoftGraphLocalizedNotificationMessage)`

SetLocalizedNotificationMessages sets LocalizedNotificationMessages field to given value.

### HasLocalizedNotificationMessages

`func (o *MicrosoftGraphNotificationMessageTemplate) HasLocalizedNotificationMessages() bool`

HasLocalizedNotificationMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


