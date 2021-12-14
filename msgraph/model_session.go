/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Session struct for Session
type Session struct {
	// Endpoint that answered the session.
	Callee AnyOfmicrosoftGraphCallRecordsEndpoint `json:"callee,omitempty"`
	// Endpoint that initiated the session.
	Caller AnyOfmicrosoftGraphCallRecordsEndpoint `json:"caller,omitempty"`
	// UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Failure information associated with the session if the session failed.
	FailureInfo AnyOfmicrosoftGraphCallRecordsFailureInfo `json:"failureInfo,omitempty"`
	// List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
	Modalities *[]AnyOfmicrosoftGraphCallRecordsModality `json:"modalities,omitempty"`
	// UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The list of segments involved in the session. Read-only. Nullable.
	Segments *[]MicrosoftGraphCallRecordsSegment `json:"segments,omitempty"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession() *Session {
	this := Session{}
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetCallee returns the Callee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Session) GetCallee() AnyOfmicrosoftGraphCallRecordsEndpoint {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsEndpoint
		return ret
	}
	return o.Callee
}

// GetCalleeOk returns a tuple with the Callee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Session) GetCalleeOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool) {
	if o == nil || o.Callee == nil {
		return nil, false
	}
	return &o.Callee, true
}

// HasCallee returns a boolean if a field has been set.
func (o *Session) HasCallee() bool {
	if o != nil && o.Callee != nil {
		return true
	}

	return false
}

// SetCallee gets a reference to the given AnyOfmicrosoftGraphCallRecordsEndpoint and assigns it to the Callee field.
func (o *Session) SetCallee(v AnyOfmicrosoftGraphCallRecordsEndpoint) {
	o.Callee = v
}

// GetCaller returns the Caller field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Session) GetCaller() AnyOfmicrosoftGraphCallRecordsEndpoint {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsEndpoint
		return ret
	}
	return o.Caller
}

// GetCallerOk returns a tuple with the Caller field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Session) GetCallerOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool) {
	if o == nil || o.Caller == nil {
		return nil, false
	}
	return &o.Caller, true
}

// HasCaller returns a boolean if a field has been set.
func (o *Session) HasCaller() bool {
	if o != nil && o.Caller != nil {
		return true
	}

	return false
}

// SetCaller gets a reference to the given AnyOfmicrosoftGraphCallRecordsEndpoint and assigns it to the Caller field.
func (o *Session) SetCaller(v AnyOfmicrosoftGraphCallRecordsEndpoint) {
	o.Caller = v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *Session) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EndDateTime == nil {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *Session) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *Session) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

// GetFailureInfo returns the FailureInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Session) GetFailureInfo() AnyOfmicrosoftGraphCallRecordsFailureInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsFailureInfo
		return ret
	}
	return o.FailureInfo
}

// GetFailureInfoOk returns a tuple with the FailureInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Session) GetFailureInfoOk() (*AnyOfmicrosoftGraphCallRecordsFailureInfo, bool) {
	if o == nil || o.FailureInfo == nil {
		return nil, false
	}
	return &o.FailureInfo, true
}

// HasFailureInfo returns a boolean if a field has been set.
func (o *Session) HasFailureInfo() bool {
	if o != nil && o.FailureInfo != nil {
		return true
	}

	return false
}

// SetFailureInfo gets a reference to the given AnyOfmicrosoftGraphCallRecordsFailureInfo and assigns it to the FailureInfo field.
func (o *Session) SetFailureInfo(v AnyOfmicrosoftGraphCallRecordsFailureInfo) {
	o.FailureInfo = v
}

// GetModalities returns the Modalities field value if set, zero value otherwise.
func (o *Session) GetModalities() []AnyOfmicrosoftGraphCallRecordsModality {
	if o == nil || o.Modalities == nil {
		var ret []AnyOfmicrosoftGraphCallRecordsModality
		return ret
	}
	return *o.Modalities
}

// GetModalitiesOk returns a tuple with the Modalities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetModalitiesOk() (*[]AnyOfmicrosoftGraphCallRecordsModality, bool) {
	if o == nil || o.Modalities == nil {
		return nil, false
	}
	return o.Modalities, true
}

// HasModalities returns a boolean if a field has been set.
func (o *Session) HasModalities() bool {
	if o != nil && o.Modalities != nil {
		return true
	}

	return false
}

// SetModalities gets a reference to the given []AnyOfmicrosoftGraphCallRecordsModality and assigns it to the Modalities field.
func (o *Session) SetModalities(v []AnyOfmicrosoftGraphCallRecordsModality) {
	o.Modalities = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *Session) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *Session) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *Session) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *Session) GetSegments() []MicrosoftGraphCallRecordsSegment {
	if o == nil || o.Segments == nil {
		var ret []MicrosoftGraphCallRecordsSegment
		return ret
	}
	return *o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetSegmentsOk() (*[]MicrosoftGraphCallRecordsSegment, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *Session) HasSegments() bool {
	if o != nil && o.Segments != nil {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []MicrosoftGraphCallRecordsSegment and assigns it to the Segments field.
func (o *Session) SetSegments(v []MicrosoftGraphCallRecordsSegment) {
	o.Segments = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Callee != nil {
		toSerialize["callee"] = o.Callee
	}
	if o.Caller != nil {
		toSerialize["caller"] = o.Caller
	}
	if o.EndDateTime != nil {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if o.FailureInfo != nil {
		toSerialize["failureInfo"] = o.FailureInfo
	}
	if o.Modalities != nil {
		toSerialize["modalities"] = o.Modalities
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	return json.Marshal(toSerialize)
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

