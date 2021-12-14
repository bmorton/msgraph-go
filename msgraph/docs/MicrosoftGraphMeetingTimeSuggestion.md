# MicrosoftGraphMeetingTimeSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttendeeAvailability** | Pointer to [**[]AnyOfmicrosoftGraphAttendeeAvailability**](AnyOfmicrosoftGraphAttendeeAvailability.md) | An array that shows the availability status of each attendee for this meeting suggestion. | [optional] 
**Confidence** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | A percentage that represents the likelhood of all the attendees attending. | [optional] 
**Locations** | Pointer to [**[]AnyOfmicrosoftGraphLocation**](AnyOfmicrosoftGraphLocation.md) | An array that specifies the name and geographic location of each meeting location for this meeting suggestion. | [optional] 
**MeetingTimeSlot** | Pointer to [**AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) | A time period suggested for the meeting. | [optional] 
**Order** | Pointer to **NullableInt32** | Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence. | [optional] 
**OrganizerAvailability** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) | Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown. | [optional] 
**SuggestionReason** | Pointer to **NullableString** | Reason for suggesting the meeting time. | [optional] 

## Methods

### NewMicrosoftGraphMeetingTimeSuggestion

`func NewMicrosoftGraphMeetingTimeSuggestion() *MicrosoftGraphMeetingTimeSuggestion`

NewMicrosoftGraphMeetingTimeSuggestion instantiates a new MicrosoftGraphMeetingTimeSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMeetingTimeSuggestionWithDefaults

`func NewMicrosoftGraphMeetingTimeSuggestionWithDefaults() *MicrosoftGraphMeetingTimeSuggestion`

NewMicrosoftGraphMeetingTimeSuggestionWithDefaults instantiates a new MicrosoftGraphMeetingTimeSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendeeAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetAttendeeAvailability() []*AnyOfmicrosoftGraphAttendeeAvailability`

GetAttendeeAvailability returns the AttendeeAvailability field if non-nil, zero value otherwise.

### GetAttendeeAvailabilityOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetAttendeeAvailabilityOk() (*[]*AnyOfmicrosoftGraphAttendeeAvailability, bool)`

GetAttendeeAvailabilityOk returns a tuple with the AttendeeAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeeAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetAttendeeAvailability(v []*AnyOfmicrosoftGraphAttendeeAvailability)`

SetAttendeeAvailability sets AttendeeAvailability field to given value.

### HasAttendeeAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasAttendeeAvailability() bool`

HasAttendeeAvailability returns a boolean if a field has been set.

### GetConfidence

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetConfidence() AnyOfnumberstringstring`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetConfidenceOk() (*AnyOfnumberstringstring, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetConfidence(v AnyOfnumberstringstring)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *MicrosoftGraphMeetingTimeSuggestion) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetLocations

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetLocations() []*AnyOfmicrosoftGraphLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetLocationsOk() (*[]*AnyOfmicrosoftGraphLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetLocations(v []*AnyOfmicrosoftGraphLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMeetingTimeSlot

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetMeetingTimeSlot() AnyOfmicrosoftGraphTimeSlot`

GetMeetingTimeSlot returns the MeetingTimeSlot field if non-nil, zero value otherwise.

### GetMeetingTimeSlotOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetMeetingTimeSlotOk() (*AnyOfmicrosoftGraphTimeSlot, bool)`

GetMeetingTimeSlotOk returns a tuple with the MeetingTimeSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingTimeSlot

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetMeetingTimeSlot(v AnyOfmicrosoftGraphTimeSlot)`

SetMeetingTimeSlot sets MeetingTimeSlot field to given value.

### HasMeetingTimeSlot

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasMeetingTimeSlot() bool`

HasMeetingTimeSlot returns a boolean if a field has been set.

### SetMeetingTimeSlotNil

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetMeetingTimeSlotNil(b bool)`

 SetMeetingTimeSlotNil sets the value for MeetingTimeSlot to be an explicit nil

### UnsetMeetingTimeSlot
`func (o *MicrosoftGraphMeetingTimeSuggestion) UnsetMeetingTimeSlot()`

UnsetMeetingTimeSlot ensures that no value is present for MeetingTimeSlot, not even an explicit nil
### GetOrder

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *MicrosoftGraphMeetingTimeSuggestion) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetOrganizerAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrganizerAvailability() AnyOfmicrosoftGraphFreeBusyStatus`

GetOrganizerAvailability returns the OrganizerAvailability field if non-nil, zero value otherwise.

### GetOrganizerAvailabilityOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrganizerAvailabilityOk() (*AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetOrganizerAvailabilityOk returns a tuple with the OrganizerAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrganizerAvailability(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetOrganizerAvailability sets OrganizerAvailability field to given value.

### HasOrganizerAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasOrganizerAvailability() bool`

HasOrganizerAvailability returns a boolean if a field has been set.

### SetOrganizerAvailabilityNil

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrganizerAvailabilityNil(b bool)`

 SetOrganizerAvailabilityNil sets the value for OrganizerAvailability to be an explicit nil

### UnsetOrganizerAvailability
`func (o *MicrosoftGraphMeetingTimeSuggestion) UnsetOrganizerAvailability()`

UnsetOrganizerAvailability ensures that no value is present for OrganizerAvailability, not even an explicit nil
### GetSuggestionReason

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetSuggestionReason() string`

GetSuggestionReason returns the SuggestionReason field if non-nil, zero value otherwise.

### GetSuggestionReasonOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetSuggestionReasonOk() (*string, bool)`

GetSuggestionReasonOk returns a tuple with the SuggestionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionReason

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetSuggestionReason(v string)`

SetSuggestionReason sets SuggestionReason field to given value.

### HasSuggestionReason

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasSuggestionReason() bool`

HasSuggestionReason returns a boolean if a field has been set.

### SetSuggestionReasonNil

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetSuggestionReasonNil(b bool)`

 SetSuggestionReasonNil sets the value for SuggestionReason to be an explicit nil

### UnsetSuggestionReason
`func (o *MicrosoftGraphMeetingTimeSuggestion) UnsetSuggestionReason()`

UnsetSuggestionReason ensures that no value is present for SuggestionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


