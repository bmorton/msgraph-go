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

// MicrosoftGraphPrivacy struct for MicrosoftGraphPrivacy
type MicrosoftGraphPrivacy struct {
	SubjectRightsRequests *[]MicrosoftGraphSubjectRightsRequest `json:"subjectRightsRequests,omitempty"`
}

// NewMicrosoftGraphPrivacy instantiates a new MicrosoftGraphPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrivacy() *MicrosoftGraphPrivacy {
	this := MicrosoftGraphPrivacy{}
	return &this
}

// NewMicrosoftGraphPrivacyWithDefaults instantiates a new MicrosoftGraphPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrivacyWithDefaults() *MicrosoftGraphPrivacy {
	this := MicrosoftGraphPrivacy{}
	return &this
}

// GetSubjectRightsRequests returns the SubjectRightsRequests field value if set, zero value otherwise.
func (o *MicrosoftGraphPrivacy) GetSubjectRightsRequests() []MicrosoftGraphSubjectRightsRequest {
	if o == nil || o.SubjectRightsRequests == nil {
		var ret []MicrosoftGraphSubjectRightsRequest
		return ret
	}
	return *o.SubjectRightsRequests
}

// GetSubjectRightsRequestsOk returns a tuple with the SubjectRightsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrivacy) GetSubjectRightsRequestsOk() (*[]MicrosoftGraphSubjectRightsRequest, bool) {
	if o == nil || o.SubjectRightsRequests == nil {
		return nil, false
	}
	return o.SubjectRightsRequests, true
}

// HasSubjectRightsRequests returns a boolean if a field has been set.
func (o *MicrosoftGraphPrivacy) HasSubjectRightsRequests() bool {
	if o != nil && o.SubjectRightsRequests != nil {
		return true
	}

	return false
}

// SetSubjectRightsRequests gets a reference to the given []MicrosoftGraphSubjectRightsRequest and assigns it to the SubjectRightsRequests field.
func (o *MicrosoftGraphPrivacy) SetSubjectRightsRequests(v []MicrosoftGraphSubjectRightsRequest) {
	o.SubjectRightsRequests = &v
}

func (o MicrosoftGraphPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubjectRightsRequests != nil {
		toSerialize["subjectRightsRequests"] = o.SubjectRightsRequests
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrivacy struct {
	value *MicrosoftGraphPrivacy
	isSet bool
}

func (v NullableMicrosoftGraphPrivacy) Get() *MicrosoftGraphPrivacy {
	return v.value
}

func (v *NullableMicrosoftGraphPrivacy) Set(val *MicrosoftGraphPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrivacy(val *MicrosoftGraphPrivacy) *NullableMicrosoftGraphPrivacy {
	return &NullableMicrosoftGraphPrivacy{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


