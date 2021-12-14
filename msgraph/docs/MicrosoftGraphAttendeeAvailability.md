# MicrosoftGraphAttendeeAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendee** | Pointer to [**AnyOfmicrosoftGraphAttendeeBase**](anyOf&lt;microsoft.graph.attendeeBase&gt;.md) | The email address and type of attendee - whether it&#39;s a person or a resource, and whether required or optional if it&#39;s a person. | [optional] 
**Availability** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) | The availability status of the attendee. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown. | [optional] 

## Methods

### NewMicrosoftGraphAttendeeAvailability

`func NewMicrosoftGraphAttendeeAvailability() *MicrosoftGraphAttendeeAvailability`

NewMicrosoftGraphAttendeeAvailability instantiates a new MicrosoftGraphAttendeeAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAttendeeAvailabilityWithDefaults

`func NewMicrosoftGraphAttendeeAvailabilityWithDefaults() *MicrosoftGraphAttendeeAvailability`

NewMicrosoftGraphAttendeeAvailabilityWithDefaults instantiates a new MicrosoftGraphAttendeeAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendee

`func (o *MicrosoftGraphAttendeeAvailability) GetAttendee() AnyOfmicrosoftGraphAttendeeBase`

GetAttendee returns the Attendee field if non-nil, zero value otherwise.

### GetAttendeeOk

`func (o *MicrosoftGraphAttendeeAvailability) GetAttendeeOk() (*AnyOfmicrosoftGraphAttendeeBase, bool)`

GetAttendeeOk returns a tuple with the Attendee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendee

`func (o *MicrosoftGraphAttendeeAvailability) SetAttendee(v AnyOfmicrosoftGraphAttendeeBase)`

SetAttendee sets Attendee field to given value.

### HasAttendee

`func (o *MicrosoftGraphAttendeeAvailability) HasAttendee() bool`

HasAttendee returns a boolean if a field has been set.

### SetAttendeeNil

`func (o *MicrosoftGraphAttendeeAvailability) SetAttendeeNil(b bool)`

 SetAttendeeNil sets the value for Attendee to be an explicit nil

### UnsetAttendee
`func (o *MicrosoftGraphAttendeeAvailability) UnsetAttendee()`

UnsetAttendee ensures that no value is present for Attendee, not even an explicit nil
### GetAvailability

`func (o *MicrosoftGraphAttendeeAvailability) GetAvailability() AnyOfmicrosoftGraphFreeBusyStatus`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MicrosoftGraphAttendeeAvailability) GetAvailabilityOk() (*AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *MicrosoftGraphAttendeeAvailability) SetAvailability(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *MicrosoftGraphAttendeeAvailability) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *MicrosoftGraphAttendeeAvailability) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *MicrosoftGraphAttendeeAvailability) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


