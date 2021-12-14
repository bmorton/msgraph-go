# MicrosoftGraphAutomaticRepliesSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAudience** | Pointer to [**AnyOfmicrosoftGraphExternalAudienceScope**](anyOf&lt;microsoft.graph.externalAudienceScope&gt;.md) | The set of audience external to the signed-in user&#39;s organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all. | [optional] 
**ExternalReplyMessage** | Pointer to **NullableString** | The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled. | [optional] 
**InternalReplyMessage** | Pointer to **NullableString** | The automatic reply to send to the audience internal to the signed-in user&#39;s organization, if Status is AlwaysEnabled or Scheduled. | [optional] 
**ScheduledEndDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that automatic replies are set to end, if Status is set to Scheduled. | [optional] 
**ScheduledStartDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that automatic replies are set to begin, if Status is set to Scheduled. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphAutomaticRepliesStatus**](anyOf&lt;microsoft.graph.automaticRepliesStatus&gt;.md) | Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled. | [optional] 

## Methods

### NewMicrosoftGraphAutomaticRepliesSetting

`func NewMicrosoftGraphAutomaticRepliesSetting() *MicrosoftGraphAutomaticRepliesSetting`

NewMicrosoftGraphAutomaticRepliesSetting instantiates a new MicrosoftGraphAutomaticRepliesSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAutomaticRepliesSettingWithDefaults

`func NewMicrosoftGraphAutomaticRepliesSettingWithDefaults() *MicrosoftGraphAutomaticRepliesSetting`

NewMicrosoftGraphAutomaticRepliesSettingWithDefaults instantiates a new MicrosoftGraphAutomaticRepliesSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAudience

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalAudience() AnyOfmicrosoftGraphExternalAudienceScope`

GetExternalAudience returns the ExternalAudience field if non-nil, zero value otherwise.

### GetExternalAudienceOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalAudienceOk() (*AnyOfmicrosoftGraphExternalAudienceScope, bool)`

GetExternalAudienceOk returns a tuple with the ExternalAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAudience

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalAudience(v AnyOfmicrosoftGraphExternalAudienceScope)`

SetExternalAudience sets ExternalAudience field to given value.

### HasExternalAudience

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasExternalAudience() bool`

HasExternalAudience returns a boolean if a field has been set.

### SetExternalAudienceNil

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalAudienceNil(b bool)`

 SetExternalAudienceNil sets the value for ExternalAudience to be an explicit nil

### UnsetExternalAudience
`func (o *MicrosoftGraphAutomaticRepliesSetting) UnsetExternalAudience()`

UnsetExternalAudience ensures that no value is present for ExternalAudience, not even an explicit nil
### GetExternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalReplyMessage() string`

GetExternalReplyMessage returns the ExternalReplyMessage field if non-nil, zero value otherwise.

### GetExternalReplyMessageOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalReplyMessageOk() (*string, bool)`

GetExternalReplyMessageOk returns a tuple with the ExternalReplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalReplyMessage(v string)`

SetExternalReplyMessage sets ExternalReplyMessage field to given value.

### HasExternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasExternalReplyMessage() bool`

HasExternalReplyMessage returns a boolean if a field has been set.

### SetExternalReplyMessageNil

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalReplyMessageNil(b bool)`

 SetExternalReplyMessageNil sets the value for ExternalReplyMessage to be an explicit nil

### UnsetExternalReplyMessage
`func (o *MicrosoftGraphAutomaticRepliesSetting) UnsetExternalReplyMessage()`

UnsetExternalReplyMessage ensures that no value is present for ExternalReplyMessage, not even an explicit nil
### GetInternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetInternalReplyMessage() string`

GetInternalReplyMessage returns the InternalReplyMessage field if non-nil, zero value otherwise.

### GetInternalReplyMessageOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetInternalReplyMessageOk() (*string, bool)`

GetInternalReplyMessageOk returns a tuple with the InternalReplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetInternalReplyMessage(v string)`

SetInternalReplyMessage sets InternalReplyMessage field to given value.

### HasInternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasInternalReplyMessage() bool`

HasInternalReplyMessage returns a boolean if a field has been set.

### SetInternalReplyMessageNil

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetInternalReplyMessageNil(b bool)`

 SetInternalReplyMessageNil sets the value for InternalReplyMessage to be an explicit nil

### UnsetInternalReplyMessage
`func (o *MicrosoftGraphAutomaticRepliesSetting) UnsetInternalReplyMessage()`

UnsetInternalReplyMessage ensures that no value is present for InternalReplyMessage, not even an explicit nil
### GetScheduledEndDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledEndDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledEndDateTime returns the ScheduledEndDateTime field if non-nil, zero value otherwise.

### GetScheduledEndDateTimeOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledEndDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledEndDateTimeOk returns a tuple with the ScheduledEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledEndDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledEndDateTime sets ScheduledEndDateTime field to given value.

### HasScheduledEndDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasScheduledEndDateTime() bool`

HasScheduledEndDateTime returns a boolean if a field has been set.

### SetScheduledEndDateTimeNil

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledEndDateTimeNil(b bool)`

 SetScheduledEndDateTimeNil sets the value for ScheduledEndDateTime to be an explicit nil

### UnsetScheduledEndDateTime
`func (o *MicrosoftGraphAutomaticRepliesSetting) UnsetScheduledEndDateTime()`

UnsetScheduledEndDateTime ensures that no value is present for ScheduledEndDateTime, not even an explicit nil
### GetScheduledStartDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledStartDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledStartDateTime returns the ScheduledStartDateTime field if non-nil, zero value otherwise.

### GetScheduledStartDateTimeOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledStartDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledStartDateTimeOk returns a tuple with the ScheduledStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledStartDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledStartDateTime sets ScheduledStartDateTime field to given value.

### HasScheduledStartDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasScheduledStartDateTime() bool`

HasScheduledStartDateTime returns a boolean if a field has been set.

### SetScheduledStartDateTimeNil

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledStartDateTimeNil(b bool)`

 SetScheduledStartDateTimeNil sets the value for ScheduledStartDateTime to be an explicit nil

### UnsetScheduledStartDateTime
`func (o *MicrosoftGraphAutomaticRepliesSetting) UnsetScheduledStartDateTime()`

UnsetScheduledStartDateTime ensures that no value is present for ScheduledStartDateTime, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetStatus() AnyOfmicrosoftGraphAutomaticRepliesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetStatusOk() (*AnyOfmicrosoftGraphAutomaticRepliesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetStatus(v AnyOfmicrosoftGraphAutomaticRepliesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphAutomaticRepliesSetting) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


