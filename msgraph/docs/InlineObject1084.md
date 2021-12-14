# InlineObject1084

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendees** | Pointer to [**[]AnyOfmicrosoftGraphAttendeeBase**](AnyOfmicrosoftGraphAttendeeBase.md) |  | [optional] 
**LocationConstraint** | Pointer to [**AnyOfmicrosoftGraphLocationConstraint**](anyOf&lt;microsoft.graph.locationConstraint&gt;.md) |  | [optional] 
**TimeConstraint** | Pointer to [**AnyOfmicrosoftGraphTimeConstraint**](anyOf&lt;microsoft.graph.timeConstraint&gt;.md) |  | [optional] 
**MeetingDuration** | Pointer to **NullableString** |  | [optional] 
**MaxCandidates** | Pointer to **NullableInt32** |  | [optional] 
**IsOrganizerOptional** | Pointer to **NullableBool** |  | [optional] [default to false]
**ReturnSuggestionReasons** | Pointer to **NullableBool** |  | [optional] [default to false]
**MinimumAttendeePercentage** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 

## Methods

### NewInlineObject1084

`func NewInlineObject1084() *InlineObject1084`

NewInlineObject1084 instantiates a new InlineObject1084 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1084WithDefaults

`func NewInlineObject1084WithDefaults() *InlineObject1084`

NewInlineObject1084WithDefaults instantiates a new InlineObject1084 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendees

`func (o *InlineObject1084) GetAttendees() []*AnyOfmicrosoftGraphAttendeeBase`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *InlineObject1084) GetAttendeesOk() (*[]*AnyOfmicrosoftGraphAttendeeBase, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *InlineObject1084) SetAttendees(v []*AnyOfmicrosoftGraphAttendeeBase)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *InlineObject1084) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetLocationConstraint

`func (o *InlineObject1084) GetLocationConstraint() AnyOfmicrosoftGraphLocationConstraint`

GetLocationConstraint returns the LocationConstraint field if non-nil, zero value otherwise.

### GetLocationConstraintOk

`func (o *InlineObject1084) GetLocationConstraintOk() (*AnyOfmicrosoftGraphLocationConstraint, bool)`

GetLocationConstraintOk returns a tuple with the LocationConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraint

`func (o *InlineObject1084) SetLocationConstraint(v AnyOfmicrosoftGraphLocationConstraint)`

SetLocationConstraint sets LocationConstraint field to given value.

### HasLocationConstraint

`func (o *InlineObject1084) HasLocationConstraint() bool`

HasLocationConstraint returns a boolean if a field has been set.

### SetLocationConstraintNil

`func (o *InlineObject1084) SetLocationConstraintNil(b bool)`

 SetLocationConstraintNil sets the value for LocationConstraint to be an explicit nil

### UnsetLocationConstraint
`func (o *InlineObject1084) UnsetLocationConstraint()`

UnsetLocationConstraint ensures that no value is present for LocationConstraint, not even an explicit nil
### GetTimeConstraint

`func (o *InlineObject1084) GetTimeConstraint() AnyOfmicrosoftGraphTimeConstraint`

GetTimeConstraint returns the TimeConstraint field if non-nil, zero value otherwise.

### GetTimeConstraintOk

`func (o *InlineObject1084) GetTimeConstraintOk() (*AnyOfmicrosoftGraphTimeConstraint, bool)`

GetTimeConstraintOk returns a tuple with the TimeConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeConstraint

`func (o *InlineObject1084) SetTimeConstraint(v AnyOfmicrosoftGraphTimeConstraint)`

SetTimeConstraint sets TimeConstraint field to given value.

### HasTimeConstraint

`func (o *InlineObject1084) HasTimeConstraint() bool`

HasTimeConstraint returns a boolean if a field has been set.

### SetTimeConstraintNil

`func (o *InlineObject1084) SetTimeConstraintNil(b bool)`

 SetTimeConstraintNil sets the value for TimeConstraint to be an explicit nil

### UnsetTimeConstraint
`func (o *InlineObject1084) UnsetTimeConstraint()`

UnsetTimeConstraint ensures that no value is present for TimeConstraint, not even an explicit nil
### GetMeetingDuration

`func (o *InlineObject1084) GetMeetingDuration() string`

GetMeetingDuration returns the MeetingDuration field if non-nil, zero value otherwise.

### GetMeetingDurationOk

`func (o *InlineObject1084) GetMeetingDurationOk() (*string, bool)`

GetMeetingDurationOk returns a tuple with the MeetingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingDuration

`func (o *InlineObject1084) SetMeetingDuration(v string)`

SetMeetingDuration sets MeetingDuration field to given value.

### HasMeetingDuration

`func (o *InlineObject1084) HasMeetingDuration() bool`

HasMeetingDuration returns a boolean if a field has been set.

### SetMeetingDurationNil

`func (o *InlineObject1084) SetMeetingDurationNil(b bool)`

 SetMeetingDurationNil sets the value for MeetingDuration to be an explicit nil

### UnsetMeetingDuration
`func (o *InlineObject1084) UnsetMeetingDuration()`

UnsetMeetingDuration ensures that no value is present for MeetingDuration, not even an explicit nil
### GetMaxCandidates

`func (o *InlineObject1084) GetMaxCandidates() int32`

GetMaxCandidates returns the MaxCandidates field if non-nil, zero value otherwise.

### GetMaxCandidatesOk

`func (o *InlineObject1084) GetMaxCandidatesOk() (*int32, bool)`

GetMaxCandidatesOk returns a tuple with the MaxCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCandidates

`func (o *InlineObject1084) SetMaxCandidates(v int32)`

SetMaxCandidates sets MaxCandidates field to given value.

### HasMaxCandidates

`func (o *InlineObject1084) HasMaxCandidates() bool`

HasMaxCandidates returns a boolean if a field has been set.

### SetMaxCandidatesNil

`func (o *InlineObject1084) SetMaxCandidatesNil(b bool)`

 SetMaxCandidatesNil sets the value for MaxCandidates to be an explicit nil

### UnsetMaxCandidates
`func (o *InlineObject1084) UnsetMaxCandidates()`

UnsetMaxCandidates ensures that no value is present for MaxCandidates, not even an explicit nil
### GetIsOrganizerOptional

`func (o *InlineObject1084) GetIsOrganizerOptional() bool`

GetIsOrganizerOptional returns the IsOrganizerOptional field if non-nil, zero value otherwise.

### GetIsOrganizerOptionalOk

`func (o *InlineObject1084) GetIsOrganizerOptionalOk() (*bool, bool)`

GetIsOrganizerOptionalOk returns a tuple with the IsOrganizerOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizerOptional

`func (o *InlineObject1084) SetIsOrganizerOptional(v bool)`

SetIsOrganizerOptional sets IsOrganizerOptional field to given value.

### HasIsOrganizerOptional

`func (o *InlineObject1084) HasIsOrganizerOptional() bool`

HasIsOrganizerOptional returns a boolean if a field has been set.

### SetIsOrganizerOptionalNil

`func (o *InlineObject1084) SetIsOrganizerOptionalNil(b bool)`

 SetIsOrganizerOptionalNil sets the value for IsOrganizerOptional to be an explicit nil

### UnsetIsOrganizerOptional
`func (o *InlineObject1084) UnsetIsOrganizerOptional()`

UnsetIsOrganizerOptional ensures that no value is present for IsOrganizerOptional, not even an explicit nil
### GetReturnSuggestionReasons

`func (o *InlineObject1084) GetReturnSuggestionReasons() bool`

GetReturnSuggestionReasons returns the ReturnSuggestionReasons field if non-nil, zero value otherwise.

### GetReturnSuggestionReasonsOk

`func (o *InlineObject1084) GetReturnSuggestionReasonsOk() (*bool, bool)`

GetReturnSuggestionReasonsOk returns a tuple with the ReturnSuggestionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSuggestionReasons

`func (o *InlineObject1084) SetReturnSuggestionReasons(v bool)`

SetReturnSuggestionReasons sets ReturnSuggestionReasons field to given value.

### HasReturnSuggestionReasons

`func (o *InlineObject1084) HasReturnSuggestionReasons() bool`

HasReturnSuggestionReasons returns a boolean if a field has been set.

### SetReturnSuggestionReasonsNil

`func (o *InlineObject1084) SetReturnSuggestionReasonsNil(b bool)`

 SetReturnSuggestionReasonsNil sets the value for ReturnSuggestionReasons to be an explicit nil

### UnsetReturnSuggestionReasons
`func (o *InlineObject1084) UnsetReturnSuggestionReasons()`

UnsetReturnSuggestionReasons ensures that no value is present for ReturnSuggestionReasons, not even an explicit nil
### GetMinimumAttendeePercentage

`func (o *InlineObject1084) GetMinimumAttendeePercentage() AnyOfnumberstringstring`

GetMinimumAttendeePercentage returns the MinimumAttendeePercentage field if non-nil, zero value otherwise.

### GetMinimumAttendeePercentageOk

`func (o *InlineObject1084) GetMinimumAttendeePercentageOk() (*AnyOfnumberstringstring, bool)`

GetMinimumAttendeePercentageOk returns a tuple with the MinimumAttendeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAttendeePercentage

`func (o *InlineObject1084) SetMinimumAttendeePercentage(v AnyOfnumberstringstring)`

SetMinimumAttendeePercentage sets MinimumAttendeePercentage field to given value.

### HasMinimumAttendeePercentage

`func (o *InlineObject1084) HasMinimumAttendeePercentage() bool`

HasMinimumAttendeePercentage returns a boolean if a field has been set.

### SetMinimumAttendeePercentageNil

`func (o *InlineObject1084) SetMinimumAttendeePercentageNil(b bool)`

 SetMinimumAttendeePercentageNil sets the value for MinimumAttendeePercentage to be an explicit nil

### UnsetMinimumAttendeePercentage
`func (o *InlineObject1084) UnsetMinimumAttendeePercentage()`

UnsetMinimumAttendeePercentage ensures that no value is present for MinimumAttendeePercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


