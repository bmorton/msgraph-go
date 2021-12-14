# MicrosoftGraphMeetingTimeSuggestionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmptySuggestionsReason** | Pointer to **NullableString** | A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions. | [optional] 
**MeetingTimeSuggestions** | Pointer to [**[]AnyOfmicrosoftGraphMeetingTimeSuggestion**](AnyOfmicrosoftGraphMeetingTimeSuggestion.md) | An array of meeting suggestions. | [optional] 

## Methods

### NewMicrosoftGraphMeetingTimeSuggestionsResult

`func NewMicrosoftGraphMeetingTimeSuggestionsResult() *MicrosoftGraphMeetingTimeSuggestionsResult`

NewMicrosoftGraphMeetingTimeSuggestionsResult instantiates a new MicrosoftGraphMeetingTimeSuggestionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMeetingTimeSuggestionsResultWithDefaults

`func NewMicrosoftGraphMeetingTimeSuggestionsResultWithDefaults() *MicrosoftGraphMeetingTimeSuggestionsResult`

NewMicrosoftGraphMeetingTimeSuggestionsResultWithDefaults instantiates a new MicrosoftGraphMeetingTimeSuggestionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmptySuggestionsReason

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetEmptySuggestionsReason() string`

GetEmptySuggestionsReason returns the EmptySuggestionsReason field if non-nil, zero value otherwise.

### GetEmptySuggestionsReasonOk

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetEmptySuggestionsReasonOk() (*string, bool)`

GetEmptySuggestionsReasonOk returns a tuple with the EmptySuggestionsReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySuggestionsReason

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) SetEmptySuggestionsReason(v string)`

SetEmptySuggestionsReason sets EmptySuggestionsReason field to given value.

### HasEmptySuggestionsReason

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) HasEmptySuggestionsReason() bool`

HasEmptySuggestionsReason returns a boolean if a field has been set.

### SetEmptySuggestionsReasonNil

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) SetEmptySuggestionsReasonNil(b bool)`

 SetEmptySuggestionsReasonNil sets the value for EmptySuggestionsReason to be an explicit nil

### UnsetEmptySuggestionsReason
`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) UnsetEmptySuggestionsReason()`

UnsetEmptySuggestionsReason ensures that no value is present for EmptySuggestionsReason, not even an explicit nil
### GetMeetingTimeSuggestions

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetMeetingTimeSuggestions() []*AnyOfmicrosoftGraphMeetingTimeSuggestion`

GetMeetingTimeSuggestions returns the MeetingTimeSuggestions field if non-nil, zero value otherwise.

### GetMeetingTimeSuggestionsOk

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetMeetingTimeSuggestionsOk() (*[]*AnyOfmicrosoftGraphMeetingTimeSuggestion, bool)`

GetMeetingTimeSuggestionsOk returns a tuple with the MeetingTimeSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingTimeSuggestions

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) SetMeetingTimeSuggestions(v []*AnyOfmicrosoftGraphMeetingTimeSuggestion)`

SetMeetingTimeSuggestions sets MeetingTimeSuggestions field to given value.

### HasMeetingTimeSuggestions

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) HasMeetingTimeSuggestions() bool`

HasMeetingTimeSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


