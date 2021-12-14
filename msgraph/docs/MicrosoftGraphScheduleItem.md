# MicrosoftGraphScheduleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date, time, and time zone that the corresponding event ends. | [optional] 
**IsPrivate** | Pointer to **NullableBool** | The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional. | [optional] 
**Location** | Pointer to **NullableString** | The location where the corresponding event is held or attended from. Optional. | [optional] 
**Start** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date, time, and time zone that the corresponding event starts. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) | The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown. | [optional] 
**Subject** | Pointer to **NullableString** | The corresponding event&#39;s subject line. Optional. | [optional] 

## Methods

### NewMicrosoftGraphScheduleItem

`func NewMicrosoftGraphScheduleItem() *MicrosoftGraphScheduleItem`

NewMicrosoftGraphScheduleItem instantiates a new MicrosoftGraphScheduleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScheduleItemWithDefaults

`func NewMicrosoftGraphScheduleItemWithDefaults() *MicrosoftGraphScheduleItem`

NewMicrosoftGraphScheduleItemWithDefaults instantiates a new MicrosoftGraphScheduleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MicrosoftGraphScheduleItem) GetEnd() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MicrosoftGraphScheduleItem) GetEndOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MicrosoftGraphScheduleItem) SetEnd(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MicrosoftGraphScheduleItem) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *MicrosoftGraphScheduleItem) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *MicrosoftGraphScheduleItem) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIsPrivate

`func (o *MicrosoftGraphScheduleItem) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *MicrosoftGraphScheduleItem) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *MicrosoftGraphScheduleItem) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *MicrosoftGraphScheduleItem) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### SetIsPrivateNil

`func (o *MicrosoftGraphScheduleItem) SetIsPrivateNil(b bool)`

 SetIsPrivateNil sets the value for IsPrivate to be an explicit nil

### UnsetIsPrivate
`func (o *MicrosoftGraphScheduleItem) UnsetIsPrivate()`

UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphScheduleItem) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphScheduleItem) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphScheduleItem) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphScheduleItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphScheduleItem) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphScheduleItem) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetStart

`func (o *MicrosoftGraphScheduleItem) GetStart() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MicrosoftGraphScheduleItem) GetStartOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MicrosoftGraphScheduleItem) SetStart(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStart sets Start field to given value.

### HasStart

`func (o *MicrosoftGraphScheduleItem) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *MicrosoftGraphScheduleItem) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *MicrosoftGraphScheduleItem) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphScheduleItem) GetStatus() AnyOfmicrosoftGraphFreeBusyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphScheduleItem) GetStatusOk() (*AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphScheduleItem) SetStatus(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphScheduleItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphScheduleItem) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphScheduleItem) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubject

`func (o *MicrosoftGraphScheduleItem) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphScheduleItem) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MicrosoftGraphScheduleItem) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MicrosoftGraphScheduleItem) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *MicrosoftGraphScheduleItem) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *MicrosoftGraphScheduleItem) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


