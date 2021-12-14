# Attendee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProposedNewTime** | Pointer to [**AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) | An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn&#39;t proposed another time, then this property is not included in a response of a GET event. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphResponseStatus**](anyOf&lt;microsoft.graph.responseStatus&gt;.md) | The attendee&#39;s response (none, accepted, declined, etc.) for the event and date-time that the response was sent. | [optional] 

## Methods

### NewAttendee

`func NewAttendee() *Attendee`

NewAttendee instantiates a new Attendee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttendeeWithDefaults

`func NewAttendeeWithDefaults() *Attendee`

NewAttendeeWithDefaults instantiates a new Attendee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposedNewTime

`func (o *Attendee) GetProposedNewTime() AnyOfmicrosoftGraphTimeSlot`

GetProposedNewTime returns the ProposedNewTime field if non-nil, zero value otherwise.

### GetProposedNewTimeOk

`func (o *Attendee) GetProposedNewTimeOk() (*AnyOfmicrosoftGraphTimeSlot, bool)`

GetProposedNewTimeOk returns a tuple with the ProposedNewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedNewTime

`func (o *Attendee) SetProposedNewTime(v AnyOfmicrosoftGraphTimeSlot)`

SetProposedNewTime sets ProposedNewTime field to given value.

### HasProposedNewTime

`func (o *Attendee) HasProposedNewTime() bool`

HasProposedNewTime returns a boolean if a field has been set.

### SetProposedNewTimeNil

`func (o *Attendee) SetProposedNewTimeNil(b bool)`

 SetProposedNewTimeNil sets the value for ProposedNewTime to be an explicit nil

### UnsetProposedNewTime
`func (o *Attendee) UnsetProposedNewTime()`

UnsetProposedNewTime ensures that no value is present for ProposedNewTime, not even an explicit nil
### GetStatus

`func (o *Attendee) GetStatus() AnyOfmicrosoftGraphResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attendee) GetStatusOk() (*AnyOfmicrosoftGraphResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attendee) SetStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attendee) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Attendee) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Attendee) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


