# MicrosoftGraphAttendee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | The recipient&#39;s email address. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphAttendeeType**](anyOf&lt;microsoft.graph.attendeeType&gt;.md) | The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type. | [optional] 
**ProposedNewTime** | Pointer to [**AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) | An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn&#39;t proposed another time, then this property is not included in a response of a GET event. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphResponseStatus**](anyOf&lt;microsoft.graph.responseStatus&gt;.md) | The attendee&#39;s response (none, accepted, declined, etc.) for the event and date-time that the response was sent. | [optional] 

## Methods

### NewMicrosoftGraphAttendee

`func NewMicrosoftGraphAttendee() *MicrosoftGraphAttendee`

NewMicrosoftGraphAttendee instantiates a new MicrosoftGraphAttendee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAttendeeWithDefaults

`func NewMicrosoftGraphAttendeeWithDefaults() *MicrosoftGraphAttendee`

NewMicrosoftGraphAttendeeWithDefaults instantiates a new MicrosoftGraphAttendee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *MicrosoftGraphAttendee) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphAttendee) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MicrosoftGraphAttendee) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MicrosoftGraphAttendee) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *MicrosoftGraphAttendee) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *MicrosoftGraphAttendee) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetType

`func (o *MicrosoftGraphAttendee) GetType() AnyOfmicrosoftGraphAttendeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphAttendee) GetTypeOk() (*AnyOfmicrosoftGraphAttendeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphAttendee) SetType(v AnyOfmicrosoftGraphAttendeeType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphAttendee) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphAttendee) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphAttendee) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProposedNewTime

`func (o *MicrosoftGraphAttendee) GetProposedNewTime() AnyOfmicrosoftGraphTimeSlot`

GetProposedNewTime returns the ProposedNewTime field if non-nil, zero value otherwise.

### GetProposedNewTimeOk

`func (o *MicrosoftGraphAttendee) GetProposedNewTimeOk() (*AnyOfmicrosoftGraphTimeSlot, bool)`

GetProposedNewTimeOk returns a tuple with the ProposedNewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedNewTime

`func (o *MicrosoftGraphAttendee) SetProposedNewTime(v AnyOfmicrosoftGraphTimeSlot)`

SetProposedNewTime sets ProposedNewTime field to given value.

### HasProposedNewTime

`func (o *MicrosoftGraphAttendee) HasProposedNewTime() bool`

HasProposedNewTime returns a boolean if a field has been set.

### SetProposedNewTimeNil

`func (o *MicrosoftGraphAttendee) SetProposedNewTimeNil(b bool)`

 SetProposedNewTimeNil sets the value for ProposedNewTime to be an explicit nil

### UnsetProposedNewTime
`func (o *MicrosoftGraphAttendee) UnsetProposedNewTime()`

UnsetProposedNewTime ensures that no value is present for ProposedNewTime, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphAttendee) GetStatus() AnyOfmicrosoftGraphResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAttendee) GetStatusOk() (*AnyOfmicrosoftGraphResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphAttendee) SetStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphAttendee) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphAttendee) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphAttendee) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


