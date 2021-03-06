/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// InlineObject1084 struct for InlineObject1084
type InlineObject1084 struct {
	Attendees *[]*AnyOfmicrosoftGraphAttendeeBase `json:"attendees,omitempty"`
	LocationConstraint AnyOfmicrosoftGraphLocationConstraint `json:"locationConstraint,omitempty"`
	TimeConstraint AnyOfmicrosoftGraphTimeConstraint `json:"timeConstraint,omitempty"`
	MeetingDuration NullableString `json:"meetingDuration,omitempty"`
	MaxCandidates NullableInt32 `json:"maxCandidates,omitempty"`
	IsOrganizerOptional NullableBool `json:"isOrganizerOptional,omitempty"`
	ReturnSuggestionReasons NullableBool `json:"returnSuggestionReasons,omitempty"`
	MinimumAttendeePercentage AnyOfnumberstringstring `json:"minimumAttendeePercentage,omitempty"`
}

// NewInlineObject1084 instantiates a new InlineObject1084 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1084() *InlineObject1084 {
	this := InlineObject1084{}
	var isOrganizerOptional bool = false
	this.IsOrganizerOptional = *NewNullableBool(&isOrganizerOptional)
	var returnSuggestionReasons bool = false
	this.ReturnSuggestionReasons = *NewNullableBool(&returnSuggestionReasons)
	return &this
}

// NewInlineObject1084WithDefaults instantiates a new InlineObject1084 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1084WithDefaults() *InlineObject1084 {
	this := InlineObject1084{}
	var isOrganizerOptional bool = false
	this.IsOrganizerOptional = *NewNullableBool(&isOrganizerOptional)
	var returnSuggestionReasons bool = false
	this.ReturnSuggestionReasons = *NewNullableBool(&returnSuggestionReasons)
	return &this
}

// GetAttendees returns the Attendees field value if set, zero value otherwise.
func (o *InlineObject1084) GetAttendees() []*AnyOfmicrosoftGraphAttendeeBase {
	if o == nil || o.Attendees == nil {
		var ret []*AnyOfmicrosoftGraphAttendeeBase
		return ret
	}
	return *o.Attendees
}

// GetAttendeesOk returns a tuple with the Attendees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1084) GetAttendeesOk() (*[]*AnyOfmicrosoftGraphAttendeeBase, bool) {
	if o == nil || o.Attendees == nil {
		return nil, false
	}
	return o.Attendees, true
}

// HasAttendees returns a boolean if a field has been set.
func (o *InlineObject1084) HasAttendees() bool {
	if o != nil && o.Attendees != nil {
		return true
	}

	return false
}

// SetAttendees gets a reference to the given []*AnyOfmicrosoftGraphAttendeeBase and assigns it to the Attendees field.
func (o *InlineObject1084) SetAttendees(v []*AnyOfmicrosoftGraphAttendeeBase) {
	o.Attendees = &v
}

// GetLocationConstraint returns the LocationConstraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetLocationConstraint() AnyOfmicrosoftGraphLocationConstraint {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLocationConstraint
		return ret
	}
	return o.LocationConstraint
}

// GetLocationConstraintOk returns a tuple with the LocationConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetLocationConstraintOk() (*AnyOfmicrosoftGraphLocationConstraint, bool) {
	if o == nil || o.LocationConstraint == nil {
		return nil, false
	}
	return &o.LocationConstraint, true
}

// HasLocationConstraint returns a boolean if a field has been set.
func (o *InlineObject1084) HasLocationConstraint() bool {
	if o != nil && o.LocationConstraint != nil {
		return true
	}

	return false
}

// SetLocationConstraint gets a reference to the given AnyOfmicrosoftGraphLocationConstraint and assigns it to the LocationConstraint field.
func (o *InlineObject1084) SetLocationConstraint(v AnyOfmicrosoftGraphLocationConstraint) {
	o.LocationConstraint = v
}

// GetTimeConstraint returns the TimeConstraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetTimeConstraint() AnyOfmicrosoftGraphTimeConstraint {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTimeConstraint
		return ret
	}
	return o.TimeConstraint
}

// GetTimeConstraintOk returns a tuple with the TimeConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetTimeConstraintOk() (*AnyOfmicrosoftGraphTimeConstraint, bool) {
	if o == nil || o.TimeConstraint == nil {
		return nil, false
	}
	return &o.TimeConstraint, true
}

// HasTimeConstraint returns a boolean if a field has been set.
func (o *InlineObject1084) HasTimeConstraint() bool {
	if o != nil && o.TimeConstraint != nil {
		return true
	}

	return false
}

// SetTimeConstraint gets a reference to the given AnyOfmicrosoftGraphTimeConstraint and assigns it to the TimeConstraint field.
func (o *InlineObject1084) SetTimeConstraint(v AnyOfmicrosoftGraphTimeConstraint) {
	o.TimeConstraint = v
}

// GetMeetingDuration returns the MeetingDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetMeetingDuration() string {
	if o == nil || o.MeetingDuration.Get() == nil {
		var ret string
		return ret
	}
	return *o.MeetingDuration.Get()
}

// GetMeetingDurationOk returns a tuple with the MeetingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetMeetingDurationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MeetingDuration.Get(), o.MeetingDuration.IsSet()
}

// HasMeetingDuration returns a boolean if a field has been set.
func (o *InlineObject1084) HasMeetingDuration() bool {
	if o != nil && o.MeetingDuration.IsSet() {
		return true
	}

	return false
}

// SetMeetingDuration gets a reference to the given NullableString and assigns it to the MeetingDuration field.
func (o *InlineObject1084) SetMeetingDuration(v string) {
	o.MeetingDuration.Set(&v)
}
// SetMeetingDurationNil sets the value for MeetingDuration to be an explicit nil
func (o *InlineObject1084) SetMeetingDurationNil() {
	o.MeetingDuration.Set(nil)
}

// UnsetMeetingDuration ensures that no value is present for MeetingDuration, not even an explicit nil
func (o *InlineObject1084) UnsetMeetingDuration() {
	o.MeetingDuration.Unset()
}

// GetMaxCandidates returns the MaxCandidates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetMaxCandidates() int32 {
	if o == nil || o.MaxCandidates.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxCandidates.Get()
}

// GetMaxCandidatesOk returns a tuple with the MaxCandidates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetMaxCandidatesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxCandidates.Get(), o.MaxCandidates.IsSet()
}

// HasMaxCandidates returns a boolean if a field has been set.
func (o *InlineObject1084) HasMaxCandidates() bool {
	if o != nil && o.MaxCandidates.IsSet() {
		return true
	}

	return false
}

// SetMaxCandidates gets a reference to the given NullableInt32 and assigns it to the MaxCandidates field.
func (o *InlineObject1084) SetMaxCandidates(v int32) {
	o.MaxCandidates.Set(&v)
}
// SetMaxCandidatesNil sets the value for MaxCandidates to be an explicit nil
func (o *InlineObject1084) SetMaxCandidatesNil() {
	o.MaxCandidates.Set(nil)
}

// UnsetMaxCandidates ensures that no value is present for MaxCandidates, not even an explicit nil
func (o *InlineObject1084) UnsetMaxCandidates() {
	o.MaxCandidates.Unset()
}

// GetIsOrganizerOptional returns the IsOrganizerOptional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetIsOrganizerOptional() bool {
	if o == nil || o.IsOrganizerOptional.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsOrganizerOptional.Get()
}

// GetIsOrganizerOptionalOk returns a tuple with the IsOrganizerOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetIsOrganizerOptionalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsOrganizerOptional.Get(), o.IsOrganizerOptional.IsSet()
}

// HasIsOrganizerOptional returns a boolean if a field has been set.
func (o *InlineObject1084) HasIsOrganizerOptional() bool {
	if o != nil && o.IsOrganizerOptional.IsSet() {
		return true
	}

	return false
}

// SetIsOrganizerOptional gets a reference to the given NullableBool and assigns it to the IsOrganizerOptional field.
func (o *InlineObject1084) SetIsOrganizerOptional(v bool) {
	o.IsOrganizerOptional.Set(&v)
}
// SetIsOrganizerOptionalNil sets the value for IsOrganizerOptional to be an explicit nil
func (o *InlineObject1084) SetIsOrganizerOptionalNil() {
	o.IsOrganizerOptional.Set(nil)
}

// UnsetIsOrganizerOptional ensures that no value is present for IsOrganizerOptional, not even an explicit nil
func (o *InlineObject1084) UnsetIsOrganizerOptional() {
	o.IsOrganizerOptional.Unset()
}

// GetReturnSuggestionReasons returns the ReturnSuggestionReasons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetReturnSuggestionReasons() bool {
	if o == nil || o.ReturnSuggestionReasons.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ReturnSuggestionReasons.Get()
}

// GetReturnSuggestionReasonsOk returns a tuple with the ReturnSuggestionReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetReturnSuggestionReasonsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReturnSuggestionReasons.Get(), o.ReturnSuggestionReasons.IsSet()
}

// HasReturnSuggestionReasons returns a boolean if a field has been set.
func (o *InlineObject1084) HasReturnSuggestionReasons() bool {
	if o != nil && o.ReturnSuggestionReasons.IsSet() {
		return true
	}

	return false
}

// SetReturnSuggestionReasons gets a reference to the given NullableBool and assigns it to the ReturnSuggestionReasons field.
func (o *InlineObject1084) SetReturnSuggestionReasons(v bool) {
	o.ReturnSuggestionReasons.Set(&v)
}
// SetReturnSuggestionReasonsNil sets the value for ReturnSuggestionReasons to be an explicit nil
func (o *InlineObject1084) SetReturnSuggestionReasonsNil() {
	o.ReturnSuggestionReasons.Set(nil)
}

// UnsetReturnSuggestionReasons ensures that no value is present for ReturnSuggestionReasons, not even an explicit nil
func (o *InlineObject1084) UnsetReturnSuggestionReasons() {
	o.ReturnSuggestionReasons.Unset()
}

// GetMinimumAttendeePercentage returns the MinimumAttendeePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1084) GetMinimumAttendeePercentage() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.MinimumAttendeePercentage
}

// GetMinimumAttendeePercentageOk returns a tuple with the MinimumAttendeePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1084) GetMinimumAttendeePercentageOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.MinimumAttendeePercentage == nil {
		return nil, false
	}
	return &o.MinimumAttendeePercentage, true
}

// HasMinimumAttendeePercentage returns a boolean if a field has been set.
func (o *InlineObject1084) HasMinimumAttendeePercentage() bool {
	if o != nil && o.MinimumAttendeePercentage != nil {
		return true
	}

	return false
}

// SetMinimumAttendeePercentage gets a reference to the given AnyOfnumberstringstring and assigns it to the MinimumAttendeePercentage field.
func (o *InlineObject1084) SetMinimumAttendeePercentage(v AnyOfnumberstringstring) {
	o.MinimumAttendeePercentage = v
}

func (o InlineObject1084) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attendees != nil {
		toSerialize["attendees"] = o.Attendees
	}
	if o.LocationConstraint != nil {
		toSerialize["locationConstraint"] = o.LocationConstraint
	}
	if o.TimeConstraint != nil {
		toSerialize["timeConstraint"] = o.TimeConstraint
	}
	if o.MeetingDuration.IsSet() {
		toSerialize["meetingDuration"] = o.MeetingDuration.Get()
	}
	if o.MaxCandidates.IsSet() {
		toSerialize["maxCandidates"] = o.MaxCandidates.Get()
	}
	if o.IsOrganizerOptional.IsSet() {
		toSerialize["isOrganizerOptional"] = o.IsOrganizerOptional.Get()
	}
	if o.ReturnSuggestionReasons.IsSet() {
		toSerialize["returnSuggestionReasons"] = o.ReturnSuggestionReasons.Get()
	}
	if o.MinimumAttendeePercentage != nil {
		toSerialize["minimumAttendeePercentage"] = o.MinimumAttendeePercentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1084 struct {
	value *InlineObject1084
	isSet bool
}

func (v NullableInlineObject1084) Get() *InlineObject1084 {
	return v.value
}

func (v *NullableInlineObject1084) Set(val *InlineObject1084) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1084) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1084) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1084(val *InlineObject1084) *NullableInlineObject1084 {
	return &NullableInlineObject1084{value: val, isSet: true}
}

func (v NullableInlineObject1084) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1084) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


