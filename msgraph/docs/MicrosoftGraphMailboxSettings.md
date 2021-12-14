# MicrosoftGraphMailboxSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFolder** | Pointer to **NullableString** | Folder ID of an archive folder for the user. | [optional] 
**AutomaticRepliesSetting** | Pointer to [**AnyOfmicrosoftGraphAutomaticRepliesSetting**](anyOf&lt;microsoft.graph.automaticRepliesSetting&gt;.md) | Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user. | [optional] 
**DateFormat** | Pointer to **NullableString** | The date format for the user&#39;s mailbox. | [optional] 
**DelegateMeetingMessageDeliveryOptions** | Pointer to [**AnyOfmicrosoftGraphDelegateMeetingMessageDeliveryOptions**](anyOf&lt;microsoft.graph.delegateMeetingMessageDeliveryOptions&gt;.md) | If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly. | [optional] 
**Language** | Pointer to [**AnyOfmicrosoftGraphLocaleInfo**](anyOf&lt;microsoft.graph.localeInfo&gt;.md) | The locale information for the user, including the preferred language and country/region. | [optional] 
**TimeFormat** | Pointer to **NullableString** | The time format for the user&#39;s mailbox. | [optional] 
**TimeZone** | Pointer to **NullableString** | The default time zone for the user&#39;s mailbox. | [optional] 
**WorkingHours** | Pointer to [**AnyOfmicrosoftGraphWorkingHours**](anyOf&lt;microsoft.graph.workingHours&gt;.md) | The days of the week and hours in a specific time zone that the user works. | [optional] 

## Methods

### NewMicrosoftGraphMailboxSettings

`func NewMicrosoftGraphMailboxSettings() *MicrosoftGraphMailboxSettings`

NewMicrosoftGraphMailboxSettings instantiates a new MicrosoftGraphMailboxSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMailboxSettingsWithDefaults

`func NewMicrosoftGraphMailboxSettingsWithDefaults() *MicrosoftGraphMailboxSettings`

NewMicrosoftGraphMailboxSettingsWithDefaults instantiates a new MicrosoftGraphMailboxSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFolder

`func (o *MicrosoftGraphMailboxSettings) GetArchiveFolder() string`

GetArchiveFolder returns the ArchiveFolder field if non-nil, zero value otherwise.

### GetArchiveFolderOk

`func (o *MicrosoftGraphMailboxSettings) GetArchiveFolderOk() (*string, bool)`

GetArchiveFolderOk returns a tuple with the ArchiveFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFolder

`func (o *MicrosoftGraphMailboxSettings) SetArchiveFolder(v string)`

SetArchiveFolder sets ArchiveFolder field to given value.

### HasArchiveFolder

`func (o *MicrosoftGraphMailboxSettings) HasArchiveFolder() bool`

HasArchiveFolder returns a boolean if a field has been set.

### SetArchiveFolderNil

`func (o *MicrosoftGraphMailboxSettings) SetArchiveFolderNil(b bool)`

 SetArchiveFolderNil sets the value for ArchiveFolder to be an explicit nil

### UnsetArchiveFolder
`func (o *MicrosoftGraphMailboxSettings) UnsetArchiveFolder()`

UnsetArchiveFolder ensures that no value is present for ArchiveFolder, not even an explicit nil
### GetAutomaticRepliesSetting

`func (o *MicrosoftGraphMailboxSettings) GetAutomaticRepliesSetting() AnyOfmicrosoftGraphAutomaticRepliesSetting`

GetAutomaticRepliesSetting returns the AutomaticRepliesSetting field if non-nil, zero value otherwise.

### GetAutomaticRepliesSettingOk

`func (o *MicrosoftGraphMailboxSettings) GetAutomaticRepliesSettingOk() (*AnyOfmicrosoftGraphAutomaticRepliesSetting, bool)`

GetAutomaticRepliesSettingOk returns a tuple with the AutomaticRepliesSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRepliesSetting

`func (o *MicrosoftGraphMailboxSettings) SetAutomaticRepliesSetting(v AnyOfmicrosoftGraphAutomaticRepliesSetting)`

SetAutomaticRepliesSetting sets AutomaticRepliesSetting field to given value.

### HasAutomaticRepliesSetting

`func (o *MicrosoftGraphMailboxSettings) HasAutomaticRepliesSetting() bool`

HasAutomaticRepliesSetting returns a boolean if a field has been set.

### SetAutomaticRepliesSettingNil

`func (o *MicrosoftGraphMailboxSettings) SetAutomaticRepliesSettingNil(b bool)`

 SetAutomaticRepliesSettingNil sets the value for AutomaticRepliesSetting to be an explicit nil

### UnsetAutomaticRepliesSetting
`func (o *MicrosoftGraphMailboxSettings) UnsetAutomaticRepliesSetting()`

UnsetAutomaticRepliesSetting ensures that no value is present for AutomaticRepliesSetting, not even an explicit nil
### GetDateFormat

`func (o *MicrosoftGraphMailboxSettings) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *MicrosoftGraphMailboxSettings) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *MicrosoftGraphMailboxSettings) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *MicrosoftGraphMailboxSettings) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormatNil

`func (o *MicrosoftGraphMailboxSettings) SetDateFormatNil(b bool)`

 SetDateFormatNil sets the value for DateFormat to be an explicit nil

### UnsetDateFormat
`func (o *MicrosoftGraphMailboxSettings) UnsetDateFormat()`

UnsetDateFormat ensures that no value is present for DateFormat, not even an explicit nil
### GetDelegateMeetingMessageDeliveryOptions

`func (o *MicrosoftGraphMailboxSettings) GetDelegateMeetingMessageDeliveryOptions() AnyOfmicrosoftGraphDelegateMeetingMessageDeliveryOptions`

GetDelegateMeetingMessageDeliveryOptions returns the DelegateMeetingMessageDeliveryOptions field if non-nil, zero value otherwise.

### GetDelegateMeetingMessageDeliveryOptionsOk

`func (o *MicrosoftGraphMailboxSettings) GetDelegateMeetingMessageDeliveryOptionsOk() (*AnyOfmicrosoftGraphDelegateMeetingMessageDeliveryOptions, bool)`

GetDelegateMeetingMessageDeliveryOptionsOk returns a tuple with the DelegateMeetingMessageDeliveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateMeetingMessageDeliveryOptions

`func (o *MicrosoftGraphMailboxSettings) SetDelegateMeetingMessageDeliveryOptions(v AnyOfmicrosoftGraphDelegateMeetingMessageDeliveryOptions)`

SetDelegateMeetingMessageDeliveryOptions sets DelegateMeetingMessageDeliveryOptions field to given value.

### HasDelegateMeetingMessageDeliveryOptions

`func (o *MicrosoftGraphMailboxSettings) HasDelegateMeetingMessageDeliveryOptions() bool`

HasDelegateMeetingMessageDeliveryOptions returns a boolean if a field has been set.

### SetDelegateMeetingMessageDeliveryOptionsNil

`func (o *MicrosoftGraphMailboxSettings) SetDelegateMeetingMessageDeliveryOptionsNil(b bool)`

 SetDelegateMeetingMessageDeliveryOptionsNil sets the value for DelegateMeetingMessageDeliveryOptions to be an explicit nil

### UnsetDelegateMeetingMessageDeliveryOptions
`func (o *MicrosoftGraphMailboxSettings) UnsetDelegateMeetingMessageDeliveryOptions()`

UnsetDelegateMeetingMessageDeliveryOptions ensures that no value is present for DelegateMeetingMessageDeliveryOptions, not even an explicit nil
### GetLanguage

`func (o *MicrosoftGraphMailboxSettings) GetLanguage() AnyOfmicrosoftGraphLocaleInfo`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphMailboxSettings) GetLanguageOk() (*AnyOfmicrosoftGraphLocaleInfo, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MicrosoftGraphMailboxSettings) SetLanguage(v AnyOfmicrosoftGraphLocaleInfo)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MicrosoftGraphMailboxSettings) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *MicrosoftGraphMailboxSettings) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *MicrosoftGraphMailboxSettings) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetTimeFormat

`func (o *MicrosoftGraphMailboxSettings) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *MicrosoftGraphMailboxSettings) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *MicrosoftGraphMailboxSettings) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *MicrosoftGraphMailboxSettings) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### SetTimeFormatNil

`func (o *MicrosoftGraphMailboxSettings) SetTimeFormatNil(b bool)`

 SetTimeFormatNil sets the value for TimeFormat to be an explicit nil

### UnsetTimeFormat
`func (o *MicrosoftGraphMailboxSettings) UnsetTimeFormat()`

UnsetTimeFormat ensures that no value is present for TimeFormat, not even an explicit nil
### GetTimeZone

`func (o *MicrosoftGraphMailboxSettings) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphMailboxSettings) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MicrosoftGraphMailboxSettings) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MicrosoftGraphMailboxSettings) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *MicrosoftGraphMailboxSettings) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *MicrosoftGraphMailboxSettings) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetWorkingHours

`func (o *MicrosoftGraphMailboxSettings) GetWorkingHours() AnyOfmicrosoftGraphWorkingHours`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *MicrosoftGraphMailboxSettings) GetWorkingHoursOk() (*AnyOfmicrosoftGraphWorkingHours, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *MicrosoftGraphMailboxSettings) SetWorkingHours(v AnyOfmicrosoftGraphWorkingHours)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *MicrosoftGraphMailboxSettings) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.

### SetWorkingHoursNil

`func (o *MicrosoftGraphMailboxSettings) SetWorkingHoursNil(b bool)`

 SetWorkingHoursNil sets the value for WorkingHours to be an explicit nil

### UnsetWorkingHours
`func (o *MicrosoftGraphMailboxSettings) UnsetWorkingHours()`

UnsetWorkingHours ensures that no value is present for WorkingHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


