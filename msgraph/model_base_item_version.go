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

// BaseItemVersion struct for BaseItemVersion
type BaseItemVersion struct {
	// Identity of the user which last modified the version. Read-only.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the version was last modified. Read-only.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// Indicates the publication status of this particular version. Read-only.
	Publication AnyOfmicrosoftGraphPublicationFacet `json:"publication,omitempty"`
}

// NewBaseItemVersion instantiates a new BaseItemVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItemVersion() *BaseItemVersion {
	this := BaseItemVersion{}
	return &this
}

// NewBaseItemVersionWithDefaults instantiates a new BaseItemVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemVersionWithDefaults() *BaseItemVersion {
	this := BaseItemVersion{}
	return &this
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemVersion) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *BaseItemVersion) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *BaseItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemVersion) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemVersion) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *BaseItemVersion) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *BaseItemVersion) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *BaseItemVersion) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *BaseItemVersion) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPublicationFacet
		return ret
	}
	return o.Publication
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemVersion) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool) {
	if o == nil || o.Publication == nil {
		return nil, false
	}
	return &o.Publication, true
}

// HasPublication returns a boolean if a field has been set.
func (o *BaseItemVersion) HasPublication() bool {
	if o != nil && o.Publication != nil {
		return true
	}

	return false
}

// SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.
func (o *BaseItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet) {
	o.Publication = v
}

func (o BaseItemVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.Publication != nil {
		toSerialize["publication"] = o.Publication
	}
	return json.Marshal(toSerialize)
}

type NullableBaseItemVersion struct {
	value *BaseItemVersion
	isSet bool
}

func (v NullableBaseItemVersion) Get() *BaseItemVersion {
	return v.value
}

func (v *NullableBaseItemVersion) Set(val *BaseItemVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItemVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItemVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItemVersion(val *BaseItemVersion) *NullableBaseItemVersion {
	return &NullableBaseItemVersion{value: val, isSet: true}
}

func (v NullableBaseItemVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItemVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

