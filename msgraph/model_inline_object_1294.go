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

// InlineObject1294 struct for InlineObject1294
type InlineObject1294 struct {
	Database AnyOfobject `json:"database,omitempty"`
	Field AnyOfobject `json:"field,omitempty"`
	Criteria AnyOfobject `json:"criteria,omitempty"`
}

// NewInlineObject1294 instantiates a new InlineObject1294 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1294() *InlineObject1294 {
	this := InlineObject1294{}
	return &this
}

// NewInlineObject1294WithDefaults instantiates a new InlineObject1294 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1294WithDefaults() *InlineObject1294 {
	this := InlineObject1294{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1294) GetDatabase() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1294) GetDatabaseOk() (*AnyOfobject, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return &o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *InlineObject1294) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given AnyOfobject and assigns it to the Database field.
func (o *InlineObject1294) SetDatabase(v AnyOfobject) {
	o.Database = v
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1294) GetField() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1294) GetFieldOk() (*AnyOfobject, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return &o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *InlineObject1294) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given AnyOfobject and assigns it to the Field field.
func (o *InlineObject1294) SetField(v AnyOfobject) {
	o.Field = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1294) GetCriteria() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1294) GetCriteriaOk() (*AnyOfobject, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject1294) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfobject and assigns it to the Criteria field.
func (o *InlineObject1294) SetCriteria(v AnyOfobject) {
	o.Criteria = v
}

func (o InlineObject1294) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1294 struct {
	value *InlineObject1294
	isSet bool
}

func (v NullableInlineObject1294) Get() *InlineObject1294 {
	return v.value
}

func (v *NullableInlineObject1294) Set(val *InlineObject1294) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1294) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1294) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1294(val *InlineObject1294) *NullableInlineObject1294 {
	return &NullableInlineObject1294{value: val, isSet: true}
}

func (v NullableInlineObject1294) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1294) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

