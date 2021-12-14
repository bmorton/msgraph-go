/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphListItemVersion struct for MicrosoftGraphListItemVersion
type MicrosoftGraphListItemVersion struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identity of the user which last modified the version. Read-only.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the version was last modified. Read-only.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// Indicates the publication status of this particular version. Read-only.
	Publication AnyOfmicrosoftGraphPublicationFacet `json:"publication,omitempty"`
	// A collection of the fields and values for this version of the list item.
	Fields AnyOfmicrosoftGraphFieldValueSet `json:"fields,omitempty"`
}

// NewMicrosoftGraphListItemVersion instantiates a new MicrosoftGraphListItemVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphListItemVersion() *MicrosoftGraphListItemVersion {
	this := MicrosoftGraphListItemVersion{}
	return &this
}

// NewMicrosoftGraphListItemVersionWithDefaults instantiates a new MicrosoftGraphListItemVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphListItemVersionWithDefaults() *MicrosoftGraphListItemVersion {
	this := MicrosoftGraphListItemVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphListItemVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphListItemVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphListItemVersion) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphListItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphListItemVersion) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphListItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphListItemVersion) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphListItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPublicationFacet
		return ret
	}
	return o.Publication
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphListItemVersion) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool) {
	if o == nil || o.Publication == nil {
		return nil, false
	}
	return &o.Publication, true
}

// HasPublication returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasPublication() bool {
	if o != nil && o.Publication != nil {
		return true
	}

	return false
}

// SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.
func (o *MicrosoftGraphListItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet) {
	o.Publication = v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphListItemVersion) GetFields() AnyOfmicrosoftGraphFieldValueSet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphFieldValueSet
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphListItemVersion) GetFieldsOk() (*AnyOfmicrosoftGraphFieldValueSet, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given AnyOfmicrosoftGraphFieldValueSet and assigns it to the Fields field.
func (o *MicrosoftGraphListItemVersion) SetFields(v AnyOfmicrosoftGraphFieldValueSet) {
	o.Fields = v
}

func (o MicrosoftGraphListItemVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.Publication != nil {
		toSerialize["publication"] = o.Publication
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphListItemVersion struct {
	value *MicrosoftGraphListItemVersion
	isSet bool
}

func (v NullableMicrosoftGraphListItemVersion) Get() *MicrosoftGraphListItemVersion {
	return v.value
}

func (v *NullableMicrosoftGraphListItemVersion) Set(val *MicrosoftGraphListItemVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphListItemVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphListItemVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphListItemVersion(val *MicrosoftGraphListItemVersion) *NullableMicrosoftGraphListItemVersion {
	return &NullableMicrosoftGraphListItemVersion{value: val, isSet: true}
}

func (v NullableMicrosoftGraphListItemVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphListItemVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


