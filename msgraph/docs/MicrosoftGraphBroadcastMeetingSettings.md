# MicrosoftGraphBroadcastMeetingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAudience** | Pointer to [**AnyOfmicrosoftGraphBroadcastMeetingAudience**](anyOf&lt;microsoft.graph.broadcastMeetingAudience&gt;.md) | Defines who can join the Teams live event. Possible values are listed in the following table. | [optional] 
**IsAttendeeReportEnabled** | Pointer to **NullableBool** | Indicates whether attendee report is enabled for this Teams live event. Default value is false. | [optional] 
**IsQuestionAndAnswerEnabled** | Pointer to **NullableBool** | Indicates whether Q&amp;A is enabled for this Teams live event. Default value is false. | [optional] 
**IsRecordingEnabled** | Pointer to **NullableBool** | Indicates whether recording is enabled for this Teams live event. Default value is false. | [optional] 
**IsVideoOnDemandEnabled** | Pointer to **NullableBool** | Indicates whether video on demand is enabled for this Teams live event. Default value is false. | [optional] 

## Methods

### NewMicrosoftGraphBroadcastMeetingSettings

`func NewMicrosoftGraphBroadcastMeetingSettings() *MicrosoftGraphBroadcastMeetingSettings`

NewMicrosoftGraphBroadcastMeetingSettings instantiates a new MicrosoftGraphBroadcastMeetingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphBroadcastMeetingSettingsWithDefaults

`func NewMicrosoftGraphBroadcastMeetingSettingsWithDefaults() *MicrosoftGraphBroadcastMeetingSettings`

NewMicrosoftGraphBroadcastMeetingSettingsWithDefaults instantiates a new MicrosoftGraphBroadcastMeetingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAudience

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetAllowedAudience() AnyOfmicrosoftGraphBroadcastMeetingAudience`

GetAllowedAudience returns the AllowedAudience field if non-nil, zero value otherwise.

### GetAllowedAudienceOk

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetAllowedAudienceOk() (*AnyOfmicrosoftGraphBroadcastMeetingAudience, bool)`

GetAllowedAudienceOk returns a tuple with the AllowedAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAudience

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetAllowedAudience(v AnyOfmicrosoftGraphBroadcastMeetingAudience)`

SetAllowedAudience sets AllowedAudience field to given value.

### HasAllowedAudience

`func (o *MicrosoftGraphBroadcastMeetingSettings) HasAllowedAudience() bool`

HasAllowedAudience returns a boolean if a field has been set.

### SetAllowedAudienceNil

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetAllowedAudienceNil(b bool)`

 SetAllowedAudienceNil sets the value for AllowedAudience to be an explicit nil

### UnsetAllowedAudience
`func (o *MicrosoftGraphBroadcastMeetingSettings) UnsetAllowedAudience()`

UnsetAllowedAudience ensures that no value is present for AllowedAudience, not even an explicit nil
### GetIsAttendeeReportEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsAttendeeReportEnabled() bool`

GetIsAttendeeReportEnabled returns the IsAttendeeReportEnabled field if non-nil, zero value otherwise.

### GetIsAttendeeReportEnabledOk

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsAttendeeReportEnabledOk() (*bool, bool)`

GetIsAttendeeReportEnabledOk returns a tuple with the IsAttendeeReportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttendeeReportEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsAttendeeReportEnabled(v bool)`

SetIsAttendeeReportEnabled sets IsAttendeeReportEnabled field to given value.

### HasIsAttendeeReportEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) HasIsAttendeeReportEnabled() bool`

HasIsAttendeeReportEnabled returns a boolean if a field has been set.

### SetIsAttendeeReportEnabledNil

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsAttendeeReportEnabledNil(b bool)`

 SetIsAttendeeReportEnabledNil sets the value for IsAttendeeReportEnabled to be an explicit nil

### UnsetIsAttendeeReportEnabled
`func (o *MicrosoftGraphBroadcastMeetingSettings) UnsetIsAttendeeReportEnabled()`

UnsetIsAttendeeReportEnabled ensures that no value is present for IsAttendeeReportEnabled, not even an explicit nil
### GetIsQuestionAndAnswerEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsQuestionAndAnswerEnabled() bool`

GetIsQuestionAndAnswerEnabled returns the IsQuestionAndAnswerEnabled field if non-nil, zero value otherwise.

### GetIsQuestionAndAnswerEnabledOk

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsQuestionAndAnswerEnabledOk() (*bool, bool)`

GetIsQuestionAndAnswerEnabledOk returns a tuple with the IsQuestionAndAnswerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuestionAndAnswerEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsQuestionAndAnswerEnabled(v bool)`

SetIsQuestionAndAnswerEnabled sets IsQuestionAndAnswerEnabled field to given value.

### HasIsQuestionAndAnswerEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) HasIsQuestionAndAnswerEnabled() bool`

HasIsQuestionAndAnswerEnabled returns a boolean if a field has been set.

### SetIsQuestionAndAnswerEnabledNil

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsQuestionAndAnswerEnabledNil(b bool)`

 SetIsQuestionAndAnswerEnabledNil sets the value for IsQuestionAndAnswerEnabled to be an explicit nil

### UnsetIsQuestionAndAnswerEnabled
`func (o *MicrosoftGraphBroadcastMeetingSettings) UnsetIsQuestionAndAnswerEnabled()`

UnsetIsQuestionAndAnswerEnabled ensures that no value is present for IsQuestionAndAnswerEnabled, not even an explicit nil
### GetIsRecordingEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsRecordingEnabled() bool`

GetIsRecordingEnabled returns the IsRecordingEnabled field if non-nil, zero value otherwise.

### GetIsRecordingEnabledOk

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsRecordingEnabledOk() (*bool, bool)`

GetIsRecordingEnabledOk returns a tuple with the IsRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecordingEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsRecordingEnabled(v bool)`

SetIsRecordingEnabled sets IsRecordingEnabled field to given value.

### HasIsRecordingEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) HasIsRecordingEnabled() bool`

HasIsRecordingEnabled returns a boolean if a field has been set.

### SetIsRecordingEnabledNil

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsRecordingEnabledNil(b bool)`

 SetIsRecordingEnabledNil sets the value for IsRecordingEnabled to be an explicit nil

### UnsetIsRecordingEnabled
`func (o *MicrosoftGraphBroadcastMeetingSettings) UnsetIsRecordingEnabled()`

UnsetIsRecordingEnabled ensures that no value is present for IsRecordingEnabled, not even an explicit nil
### GetIsVideoOnDemandEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsVideoOnDemandEnabled() bool`

GetIsVideoOnDemandEnabled returns the IsVideoOnDemandEnabled field if non-nil, zero value otherwise.

### GetIsVideoOnDemandEnabledOk

`func (o *MicrosoftGraphBroadcastMeetingSettings) GetIsVideoOnDemandEnabledOk() (*bool, bool)`

GetIsVideoOnDemandEnabledOk returns a tuple with the IsVideoOnDemandEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideoOnDemandEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsVideoOnDemandEnabled(v bool)`

SetIsVideoOnDemandEnabled sets IsVideoOnDemandEnabled field to given value.

### HasIsVideoOnDemandEnabled

`func (o *MicrosoftGraphBroadcastMeetingSettings) HasIsVideoOnDemandEnabled() bool`

HasIsVideoOnDemandEnabled returns a boolean if a field has been set.

### SetIsVideoOnDemandEnabledNil

`func (o *MicrosoftGraphBroadcastMeetingSettings) SetIsVideoOnDemandEnabledNil(b bool)`

 SetIsVideoOnDemandEnabledNil sets the value for IsVideoOnDemandEnabled to be an explicit nil

### UnsetIsVideoOnDemandEnabled
`func (o *MicrosoftGraphBroadcastMeetingSettings) UnsetIsVideoOnDemandEnabled()`

UnsetIsVideoOnDemandEnabled ensures that no value is present for IsVideoOnDemandEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


