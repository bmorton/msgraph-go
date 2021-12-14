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

// DirectoryObject struct for DirectoryObject
type DirectoryObject struct {
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
}

// NewDirectoryObject instantiates a new DirectoryObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryObject() *DirectoryObject {
	this := DirectoryObject{}
	return &this
}

// NewDirectoryObjectWithDefaults instantiates a new DirectoryObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryObjectWithDefaults() *DirectoryObject {
	this := DirectoryObject{}
	return &this
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectoryObject) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectoryObject) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *DirectoryObject) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *DirectoryObject) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *DirectoryObject) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *DirectoryObject) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

func (o DirectoryObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeletedDateTime.IsSet() {
		toSerialize["deletedDateTime"] = o.DeletedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDirectoryObject struct {
	value *DirectoryObject
	isSet bool
}

func (v NullableDirectoryObject) Get() *DirectoryObject {
	return v.value
}

func (v *NullableDirectoryObject) Set(val *DirectoryObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryObject(val *DirectoryObject) *NullableDirectoryObject {
	return &NullableDirectoryObject{value: val, isSet: true}
}

func (v NullableDirectoryObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


