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

// MicrosoftGraphFollowupFlag struct for MicrosoftGraphFollowupFlag
type MicrosoftGraphFollowupFlag struct {
	// The date and time that the follow-up was finished.
	CompletedDateTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"completedDateTime,omitempty"`
	// The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
	DueDateTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"dueDateTime,omitempty"`
	// The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
	FlagStatus AnyOfmicrosoftGraphFollowupFlagStatus `json:"flagStatus,omitempty"`
	// The date and time that the follow-up is to begin.
	StartDateTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"startDateTime,omitempty"`
}

// NewMicrosoftGraphFollowupFlag instantiates a new MicrosoftGraphFollowupFlag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphFollowupFlag() *MicrosoftGraphFollowupFlag {
	this := MicrosoftGraphFollowupFlag{}
	return &this
}

// NewMicrosoftGraphFollowupFlagWithDefaults instantiates a new MicrosoftGraphFollowupFlag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphFollowupFlagWithDefaults() *MicrosoftGraphFollowupFlag {
	this := MicrosoftGraphFollowupFlag{}
	return &this
}

// GetCompletedDateTime returns the CompletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFollowupFlag) GetCompletedDateTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.CompletedDateTime
}

// GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFollowupFlag) GetCompletedDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.CompletedDateTime == nil {
		return nil, false
	}
	return &o.CompletedDateTime, true
}

// HasCompletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphFollowupFlag) HasCompletedDateTime() bool {
	if o != nil && o.CompletedDateTime != nil {
		return true
	}

	return false
}

// SetCompletedDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the CompletedDateTime field.
func (o *MicrosoftGraphFollowupFlag) SetCompletedDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.CompletedDateTime = v
}

// GetDueDateTime returns the DueDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFollowupFlag) GetDueDateTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.DueDateTime
}

// GetDueDateTimeOk returns a tuple with the DueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFollowupFlag) GetDueDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.DueDateTime == nil {
		return nil, false
	}
	return &o.DueDateTime, true
}

// HasDueDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphFollowupFlag) HasDueDateTime() bool {
	if o != nil && o.DueDateTime != nil {
		return true
	}

	return false
}

// SetDueDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the DueDateTime field.
func (o *MicrosoftGraphFollowupFlag) SetDueDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.DueDateTime = v
}

// GetFlagStatus returns the FlagStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFollowupFlag) GetFlagStatus() AnyOfmicrosoftGraphFollowupFlagStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphFollowupFlagStatus
		return ret
	}
	return o.FlagStatus
}

// GetFlagStatusOk returns a tuple with the FlagStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFollowupFlag) GetFlagStatusOk() (*AnyOfmicrosoftGraphFollowupFlagStatus, bool) {
	if o == nil || o.FlagStatus == nil {
		return nil, false
	}
	return &o.FlagStatus, true
}

// HasFlagStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphFollowupFlag) HasFlagStatus() bool {
	if o != nil && o.FlagStatus != nil {
		return true
	}

	return false
}

// SetFlagStatus gets a reference to the given AnyOfmicrosoftGraphFollowupFlagStatus and assigns it to the FlagStatus field.
func (o *MicrosoftGraphFollowupFlag) SetFlagStatus(v AnyOfmicrosoftGraphFollowupFlagStatus) {
	o.FlagStatus = v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFollowupFlag) GetStartDateTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFollowupFlag) GetStartDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphFollowupFlag) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the StartDateTime field.
func (o *MicrosoftGraphFollowupFlag) SetStartDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.StartDateTime = v
}

func (o MicrosoftGraphFollowupFlag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompletedDateTime != nil {
		toSerialize["completedDateTime"] = o.CompletedDateTime
	}
	if o.DueDateTime != nil {
		toSerialize["dueDateTime"] = o.DueDateTime
	}
	if o.FlagStatus != nil {
		toSerialize["flagStatus"] = o.FlagStatus
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphFollowupFlag struct {
	value *MicrosoftGraphFollowupFlag
	isSet bool
}

func (v NullableMicrosoftGraphFollowupFlag) Get() *MicrosoftGraphFollowupFlag {
	return v.value
}

func (v *NullableMicrosoftGraphFollowupFlag) Set(val *MicrosoftGraphFollowupFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFollowupFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFollowupFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFollowupFlag(val *MicrosoftGraphFollowupFlag) *NullableMicrosoftGraphFollowupFlag {
	return &NullableMicrosoftGraphFollowupFlag{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFollowupFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFollowupFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


