# MicrosoftGraphAutomaticRepliesMailTips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | The automatic reply message. | [optional] 
**MessageLanguage** | Pointer to [**AnyOfmicrosoftGraphLocaleInfo**](anyOf&lt;microsoft.graph.localeInfo&gt;.md) | The language that the automatic reply message is in. | [optional] 
**ScheduledEndTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that automatic replies are set to end. | [optional] 
**ScheduledStartTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that automatic replies are set to begin. | [optional] 

## Methods

### NewMicrosoftGraphAutomaticRepliesMailTips

`func NewMicrosoftGraphAutomaticRepliesMailTips() *MicrosoftGraphAutomaticRepliesMailTips`

NewMicrosoftGraphAutomaticRepliesMailTips instantiates a new MicrosoftGraphAutomaticRepliesMailTips object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAutomaticRepliesMailTipsWithDefaults

`func NewMicrosoftGraphAutomaticRepliesMailTipsWithDefaults() *MicrosoftGraphAutomaticRepliesMailTips`

NewMicrosoftGraphAutomaticRepliesMailTipsWithDefaults instantiates a new MicrosoftGraphAutomaticRepliesMailTips object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphAutomaticRepliesMailTips) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMessageLanguage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessageLanguage() AnyOfmicrosoftGraphLocaleInfo`

GetMessageLanguage returns the MessageLanguage field if non-nil, zero value otherwise.

### GetMessageLanguageOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessageLanguageOk() (*AnyOfmicrosoftGraphLocaleInfo, bool)`

GetMessageLanguageOk returns a tuple with the MessageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageLanguage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessageLanguage(v AnyOfmicrosoftGraphLocaleInfo)`

SetMessageLanguage sets MessageLanguage field to given value.

### HasMessageLanguage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasMessageLanguage() bool`

HasMessageLanguage returns a boolean if a field has been set.

### SetMessageLanguageNil

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessageLanguageNil(b bool)`

 SetMessageLanguageNil sets the value for MessageLanguage to be an explicit nil

### UnsetMessageLanguage
`func (o *MicrosoftGraphAutomaticRepliesMailTips) UnsetMessageLanguage()`

UnsetMessageLanguage ensures that no value is present for MessageLanguage, not even an explicit nil
### GetScheduledEndTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledEndTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledEndTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *MicrosoftGraphAutomaticRepliesMailTips) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetScheduledStartTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledStartTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledStartTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *MicrosoftGraphAutomaticRepliesMailTips) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


