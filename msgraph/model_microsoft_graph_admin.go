/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MicrosoftGraphAdmin struct for MicrosoftGraphAdmin
type MicrosoftGraphAdmin struct {
	// A container for service communications resources. Read-only.
	ServiceAnnouncement AnyOfmicrosoftGraphServiceAnnouncement `json:"serviceAnnouncement,omitempty"`
}

// NewMicrosoftGraphAdmin instantiates a new MicrosoftGraphAdmin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAdmin() *MicrosoftGraphAdmin {
	this := MicrosoftGraphAdmin{}
	return &this
}

// NewMicrosoftGraphAdminWithDefaults instantiates a new MicrosoftGraphAdmin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAdminWithDefaults() *MicrosoftGraphAdmin {
	this := MicrosoftGraphAdmin{}
	return &this
}

// GetServiceAnnouncement returns the ServiceAnnouncement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAdmin) GetServiceAnnouncement() AnyOfmicrosoftGraphServiceAnnouncement {
	if o == nil  {
		var ret AnyOfmicrosoftGraphServiceAnnouncement
		return ret
	}
	return o.ServiceAnnouncement
}

// GetServiceAnnouncementOk returns a tuple with the ServiceAnnouncement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAdmin) GetServiceAnnouncementOk() (*AnyOfmicrosoftGraphServiceAnnouncement, bool) {
	if o == nil || o.ServiceAnnouncement == nil {
		return nil, false
	}
	return &o.ServiceAnnouncement, true
}

// HasServiceAnnouncement returns a boolean if a field has been set.
func (o *MicrosoftGraphAdmin) HasServiceAnnouncement() bool {
	if o != nil && o.ServiceAnnouncement != nil {
		return true
	}

	return false
}

// SetServiceAnnouncement gets a reference to the given AnyOfmicrosoftGraphServiceAnnouncement and assigns it to the ServiceAnnouncement field.
func (o *MicrosoftGraphAdmin) SetServiceAnnouncement(v AnyOfmicrosoftGraphServiceAnnouncement) {
	o.ServiceAnnouncement = v
}

func (o MicrosoftGraphAdmin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceAnnouncement != nil {
		toSerialize["serviceAnnouncement"] = o.ServiceAnnouncement
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAdmin struct {
	value *MicrosoftGraphAdmin
	isSet bool
}

func (v NullableMicrosoftGraphAdmin) Get() *MicrosoftGraphAdmin {
	return v.value
}

func (v *NullableMicrosoftGraphAdmin) Set(val *MicrosoftGraphAdmin) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAdmin(val *MicrosoftGraphAdmin) *NullableMicrosoftGraphAdmin {
	return &NullableMicrosoftGraphAdmin{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

