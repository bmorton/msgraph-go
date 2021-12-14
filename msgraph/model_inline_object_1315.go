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

// InlineObject1315 struct for InlineObject1315
type InlineObject1315 struct {
	Database AnyOfobject `json:"database,omitempty"`
	Field AnyOfobject `json:"field,omitempty"`
	Criteria AnyOfobject `json:"criteria,omitempty"`
}

// NewInlineObject1315 instantiates a new InlineObject1315 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1315() *InlineObject1315 {
	this := InlineObject1315{}
	return &this
}

// NewInlineObject1315WithDefaults instantiates a new InlineObject1315 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1315WithDefaults() *InlineObject1315 {
	this := InlineObject1315{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1315) GetDatabase() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1315) GetDatabaseOk() (*AnyOfobject, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return &o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *InlineObject1315) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given AnyOfobject and assigns it to the Database field.
func (o *InlineObject1315) SetDatabase(v AnyOfobject) {
	o.Database = v
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1315) GetField() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1315) GetFieldOk() (*AnyOfobject, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return &o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *InlineObject1315) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given AnyOfobject and assigns it to the Field field.
func (o *InlineObject1315) SetField(v AnyOfobject) {
	o.Field = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1315) GetCriteria() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1315) GetCriteriaOk() (*AnyOfobject, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject1315) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfobject and assigns it to the Criteria field.
func (o *InlineObject1315) SetCriteria(v AnyOfobject) {
	o.Criteria = v
}

func (o InlineObject1315) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject1315 struct {
	value *InlineObject1315
	isSet bool
}

func (v NullableInlineObject1315) Get() *InlineObject1315 {
	return v.value
}

func (v *NullableInlineObject1315) Set(val *InlineObject1315) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1315) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1315) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1315(val *InlineObject1315) *NullableInlineObject1315 {
	return &NullableInlineObject1315{value: val, isSet: true}
}

func (v NullableInlineObject1315) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1315) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


