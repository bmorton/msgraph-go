# LocalizedNotificationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**Locale** | Pointer to **string** | The Locale for which this message is destined. | [optional] 
**MessageTemplate** | Pointer to **string** | The Message Template content. | [optional] 
**Subject** | Pointer to **string** | The Message Template Subject. | [optional] 

## Methods

### NewLocalizedNotificationMessage

`func NewLocalizedNotificationMessage() *LocalizedNotificationMessage`

NewLocalizedNotificationMessage instantiates a new LocalizedNotificationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizedNotificationMessageWithDefaults

`func NewLocalizedNotificationMessageWithDefaults() *LocalizedNotificationMessage`

NewLocalizedNotificationMessageWithDefaults instantiates a new LocalizedNotificationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *LocalizedNotificationMessage) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *LocalizedNotificationMessage) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *LocalizedNotificationMessage) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *LocalizedNotificationMessage) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *LocalizedNotificationMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *LocalizedNotificationMessage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *LocalizedNotificationMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *LocalizedNotificationMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetLocale

`func (o *LocalizedNotificationMessage) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LocalizedNotificationMessage) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *LocalizedNotificationMessage) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *LocalizedNotificationMessage) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMessageTemplate

`func (o *LocalizedNotificationMessage) GetMessageTemplate() string`

GetMessageTemplate returns the MessageTemplate field if non-nil, zero value otherwise.

### GetMessageTemplateOk

`func (o *LocalizedNotificationMessage) GetMessageTemplateOk() (*string, bool)`

GetMessageTemplateOk returns a tuple with the MessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTemplate

`func (o *LocalizedNotificationMessage) SetMessageTemplate(v string)`

SetMessageTemplate sets MessageTemplate field to given value.

### HasMessageTemplate

`func (o *LocalizedNotificationMessage) HasMessageTemplate() bool`

HasMessageTemplate returns a boolean if a field has been set.

### GetSubject

`func (o *LocalizedNotificationMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LocalizedNotificationMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LocalizedNotificationMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LocalizedNotificationMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


