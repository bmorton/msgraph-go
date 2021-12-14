# MicrosoftGraphLocalizedNotificationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**IsDefault** | Pointer to **bool** | Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**Locale** | Pointer to **string** | The Locale for which this message is destined. | [optional] 
**MessageTemplate** | Pointer to **string** | The Message Template content. | [optional] 
**Subject** | Pointer to **string** | The Message Template Subject. | [optional] 

## Methods

### NewMicrosoftGraphLocalizedNotificationMessage

`func NewMicrosoftGraphLocalizedNotificationMessage() *MicrosoftGraphLocalizedNotificationMessage`

NewMicrosoftGraphLocalizedNotificationMessage instantiates a new MicrosoftGraphLocalizedNotificationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLocalizedNotificationMessageWithDefaults

`func NewMicrosoftGraphLocalizedNotificationMessageWithDefaults() *MicrosoftGraphLocalizedNotificationMessage`

NewMicrosoftGraphLocalizedNotificationMessageWithDefaults instantiates a new MicrosoftGraphLocalizedNotificationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphLocalizedNotificationMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphLocalizedNotificationMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphLocalizedNotificationMessage) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphLocalizedNotificationMessage) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphLocalizedNotificationMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphLocalizedNotificationMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetLocale

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MicrosoftGraphLocalizedNotificationMessage) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MicrosoftGraphLocalizedNotificationMessage) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMessageTemplate

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetMessageTemplate() string`

GetMessageTemplate returns the MessageTemplate field if non-nil, zero value otherwise.

### GetMessageTemplateOk

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetMessageTemplateOk() (*string, bool)`

GetMessageTemplateOk returns a tuple with the MessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTemplate

`func (o *MicrosoftGraphLocalizedNotificationMessage) SetMessageTemplate(v string)`

SetMessageTemplate sets MessageTemplate field to given value.

### HasMessageTemplate

`func (o *MicrosoftGraphLocalizedNotificationMessage) HasMessageTemplate() bool`

HasMessageTemplate returns a boolean if a field has been set.

### GetSubject

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphLocalizedNotificationMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MicrosoftGraphLocalizedNotificationMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MicrosoftGraphLocalizedNotificationMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


