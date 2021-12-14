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

// MicrosoftGraphSharingDetail struct for MicrosoftGraphSharingDetail
type MicrosoftGraphSharingDetail struct {
	// The user who shared the document.
	SharedBy AnyOfmicrosoftGraphInsightIdentity `json:"sharedBy,omitempty"`
	// The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	SharedDateTime NullableTime `json:"sharedDateTime,omitempty"`
	SharingReference AnyOfmicrosoftGraphResourceReference `json:"sharingReference,omitempty"`
	// The subject with which the document was shared.
	SharingSubject NullableString `json:"sharingSubject,omitempty"`
	// Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
	SharingType NullableString `json:"sharingType,omitempty"`
}

// NewMicrosoftGraphSharingDetail instantiates a new MicrosoftGraphSharingDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSharingDetail() *MicrosoftGraphSharingDetail {
	this := MicrosoftGraphSharingDetail{}
	return &this
}

// NewMicrosoftGraphSharingDetailWithDefaults instantiates a new MicrosoftGraphSharingDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSharingDetailWithDefaults() *MicrosoftGraphSharingDetail {
	this := MicrosoftGraphSharingDetail{}
	return &this
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharingDetail) GetSharedBy() AnyOfmicrosoftGraphInsightIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphInsightIdentity
		return ret
	}
	return o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharingDetail) GetSharedByOk() (*AnyOfmicrosoftGraphInsightIdentity, bool) {
	if o == nil || o.SharedBy == nil {
		return nil, false
	}
	return &o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharedBy() bool {
	if o != nil && o.SharedBy != nil {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given AnyOfmicrosoftGraphInsightIdentity and assigns it to the SharedBy field.
func (o *MicrosoftGraphSharingDetail) SetSharedBy(v AnyOfmicrosoftGraphInsightIdentity) {
	o.SharedBy = v
}

// GetSharedDateTime returns the SharedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharingDetail) GetSharedDateTime() time.Time {
	if o == nil || o.SharedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SharedDateTime.Get()
}

// GetSharedDateTimeOk returns a tuple with the SharedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharingDetail) GetSharedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SharedDateTime.Get(), o.SharedDateTime.IsSet()
}

// HasSharedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharedDateTime() bool {
	if o != nil && o.SharedDateTime.IsSet() {
		return true
	}

	return false
}

// SetSharedDateTime gets a reference to the given NullableTime and assigns it to the SharedDateTime field.
func (o *MicrosoftGraphSharingDetail) SetSharedDateTime(v time.Time) {
	o.SharedDateTime.Set(&v)
}
// SetSharedDateTimeNil sets the value for SharedDateTime to be an explicit nil
func (o *MicrosoftGraphSharingDetail) SetSharedDateTimeNil() {
	o.SharedDateTime.Set(nil)
}

// UnsetSharedDateTime ensures that no value is present for SharedDateTime, not even an explicit nil
func (o *MicrosoftGraphSharingDetail) UnsetSharedDateTime() {
	o.SharedDateTime.Unset()
}

// GetSharingReference returns the SharingReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharingDetail) GetSharingReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return o.SharingReference
}

// GetSharingReferenceOk returns a tuple with the SharingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharingDetail) GetSharingReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.SharingReference == nil {
		return nil, false
	}
	return &o.SharingReference, true
}

// HasSharingReference returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharingReference() bool {
	if o != nil && o.SharingReference != nil {
		return true
	}

	return false
}

// SetSharingReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the SharingReference field.
func (o *MicrosoftGraphSharingDetail) SetSharingReference(v AnyOfmicrosoftGraphResourceReference) {
	o.SharingReference = v
}

// GetSharingSubject returns the SharingSubject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharingDetail) GetSharingSubject() string {
	if o == nil || o.SharingSubject.Get() == nil {
		var ret string
		return ret
	}
	return *o.SharingSubject.Get()
}

// GetSharingSubjectOk returns a tuple with the SharingSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharingDetail) GetSharingSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SharingSubject.Get(), o.SharingSubject.IsSet()
}

// HasSharingSubject returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharingSubject() bool {
	if o != nil && o.SharingSubject.IsSet() {
		return true
	}

	return false
}

// SetSharingSubject gets a reference to the given NullableString and assigns it to the SharingSubject field.
func (o *MicrosoftGraphSharingDetail) SetSharingSubject(v string) {
	o.SharingSubject.Set(&v)
}
// SetSharingSubjectNil sets the value for SharingSubject to be an explicit nil
func (o *MicrosoftGraphSharingDetail) SetSharingSubjectNil() {
	o.SharingSubject.Set(nil)
}

// UnsetSharingSubject ensures that no value is present for SharingSubject, not even an explicit nil
func (o *MicrosoftGraphSharingDetail) UnsetSharingSubject() {
	o.SharingSubject.Unset()
}

// GetSharingType returns the SharingType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharingDetail) GetSharingType() string {
	if o == nil || o.SharingType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SharingType.Get()
}

// GetSharingTypeOk returns a tuple with the SharingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharingDetail) GetSharingTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SharingType.Get(), o.SharingType.IsSet()
}

// HasSharingType returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharingType() bool {
	if o != nil && o.SharingType.IsSet() {
		return true
	}

	return false
}

// SetSharingType gets a reference to the given NullableString and assigns it to the SharingType field.
func (o *MicrosoftGraphSharingDetail) SetSharingType(v string) {
	o.SharingType.Set(&v)
}
// SetSharingTypeNil sets the value for SharingType to be an explicit nil
func (o *MicrosoftGraphSharingDetail) SetSharingTypeNil() {
	o.SharingType.Set(nil)
}

// UnsetSharingType ensures that no value is present for SharingType, not even an explicit nil
func (o *MicrosoftGraphSharingDetail) UnsetSharingType() {
	o.SharingType.Unset()
}

func (o MicrosoftGraphSharingDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharedBy != nil {
		toSerialize["sharedBy"] = o.SharedBy
	}
	if o.SharedDateTime.IsSet() {
		toSerialize["sharedDateTime"] = o.SharedDateTime.Get()
	}
	if o.SharingReference != nil {
		toSerialize["sharingReference"] = o.SharingReference
	}
	if o.SharingSubject.IsSet() {
		toSerialize["sharingSubject"] = o.SharingSubject.Get()
	}
	if o.SharingType.IsSet() {
		toSerialize["sharingType"] = o.SharingType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSharingDetail struct {
	value *MicrosoftGraphSharingDetail
	isSet bool
}

func (v NullableMicrosoftGraphSharingDetail) Get() *MicrosoftGraphSharingDetail {
	return v.value
}

func (v *NullableMicrosoftGraphSharingDetail) Set(val *MicrosoftGraphSharingDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSharingDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSharingDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSharingDetail(val *MicrosoftGraphSharingDetail) *NullableMicrosoftGraphSharingDetail {
	return &NullableMicrosoftGraphSharingDetail{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSharingDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSharingDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

