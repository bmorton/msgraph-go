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

// MicrosoftGraphPublicationFacet struct for MicrosoftGraphPublicationFacet
type MicrosoftGraphPublicationFacet struct {
	// The state of publication for this document. Either published or checkout. Read-only.
	Level NullableString `json:"level,omitempty"`
	// The unique identifier for the version that is visible to the current caller. Read-only.
	VersionId NullableString `json:"versionId,omitempty"`
}

// NewMicrosoftGraphPublicationFacet instantiates a new MicrosoftGraphPublicationFacet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPublicationFacet() *MicrosoftGraphPublicationFacet {
	this := MicrosoftGraphPublicationFacet{}
	return &this
}

// NewMicrosoftGraphPublicationFacetWithDefaults instantiates a new MicrosoftGraphPublicationFacet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPublicationFacetWithDefaults() *MicrosoftGraphPublicationFacet {
	this := MicrosoftGraphPublicationFacet{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicationFacet) GetLevel() string {
	if o == nil || o.Level.Get() == nil {
		var ret string
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicationFacet) GetLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicationFacet) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableString and assigns it to the Level field.
func (o *MicrosoftGraphPublicationFacet) SetLevel(v string) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *MicrosoftGraphPublicationFacet) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *MicrosoftGraphPublicationFacet) UnsetLevel() {
	o.Level.Unset()
}

// GetVersionId returns the VersionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicationFacet) GetVersionId() string {
	if o == nil || o.VersionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.VersionId.Get()
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicationFacet) GetVersionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VersionId.Get(), o.VersionId.IsSet()
}

// HasVersionId returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicationFacet) HasVersionId() bool {
	if o != nil && o.VersionId.IsSet() {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given NullableString and assigns it to the VersionId field.
func (o *MicrosoftGraphPublicationFacet) SetVersionId(v string) {
	o.VersionId.Set(&v)
}
// SetVersionIdNil sets the value for VersionId to be an explicit nil
func (o *MicrosoftGraphPublicationFacet) SetVersionIdNil() {
	o.VersionId.Set(nil)
}

// UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil
func (o *MicrosoftGraphPublicationFacet) UnsetVersionId() {
	o.VersionId.Unset()
}

func (o MicrosoftGraphPublicationFacet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.VersionId.IsSet() {
		toSerialize["versionId"] = o.VersionId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPublicationFacet struct {
	value *MicrosoftGraphPublicationFacet
	isSet bool
}

func (v NullableMicrosoftGraphPublicationFacet) Get() *MicrosoftGraphPublicationFacet {
	return v.value
}

func (v *NullableMicrosoftGraphPublicationFacet) Set(val *MicrosoftGraphPublicationFacet) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPublicationFacet) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPublicationFacet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPublicationFacet(val *MicrosoftGraphPublicationFacet) *NullableMicrosoftGraphPublicationFacet {
	return &NullableMicrosoftGraphPublicationFacet{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPublicationFacet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPublicationFacet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


