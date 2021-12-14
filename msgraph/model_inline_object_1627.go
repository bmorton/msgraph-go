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

// InlineObject1627 struct for InlineObject1627
type InlineObject1627 struct {
	StartCell AnyOfobject `json:"startCell,omitempty"`
	EndCell AnyOfobject `json:"endCell,omitempty"`
}

// NewInlineObject1627 instantiates a new InlineObject1627 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1627() *InlineObject1627 {
	this := InlineObject1627{}
	return &this
}

// NewInlineObject1627WithDefaults instantiates a new InlineObject1627 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1627WithDefaults() *InlineObject1627 {
	this := InlineObject1627{}
	return &this
}

// GetStartCell returns the StartCell field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1627) GetStartCell() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StartCell
}

// GetStartCellOk returns a tuple with the StartCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1627) GetStartCellOk() (*AnyOfobject, bool) {
	if o == nil || o.StartCell == nil {
		return nil, false
	}
	return &o.StartCell, true
}

// HasStartCell returns a boolean if a field has been set.
func (o *InlineObject1627) HasStartCell() bool {
	if o != nil && o.StartCell != nil {
		return true
	}

	return false
}

// SetStartCell gets a reference to the given AnyOfobject and assigns it to the StartCell field.
func (o *InlineObject1627) SetStartCell(v AnyOfobject) {
	o.StartCell = v
}

// GetEndCell returns the EndCell field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1627) GetEndCell() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.EndCell
}

// GetEndCellOk returns a tuple with the EndCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1627) GetEndCellOk() (*AnyOfobject, bool) {
	if o == nil || o.EndCell == nil {
		return nil, false
	}
	return &o.EndCell, true
}

// HasEndCell returns a boolean if a field has been set.
func (o *InlineObject1627) HasEndCell() bool {
	if o != nil && o.EndCell != nil {
		return true
	}

	return false
}

// SetEndCell gets a reference to the given AnyOfobject and assigns it to the EndCell field.
func (o *InlineObject1627) SetEndCell(v AnyOfobject) {
	o.EndCell = v
}

func (o InlineObject1627) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartCell != nil {
		toSerialize["startCell"] = o.StartCell
	}
	if o.EndCell != nil {
		toSerialize["endCell"] = o.EndCell
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1627 struct {
	value *InlineObject1627
	isSet bool
}

func (v NullableInlineObject1627) Get() *InlineObject1627 {
	return v.value
}

func (v *NullableInlineObject1627) Set(val *InlineObject1627) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1627) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1627) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1627(val *InlineObject1627) *NullableInlineObject1627 {
	return &NullableInlineObject1627{value: val, isSet: true}
}

func (v NullableInlineObject1627) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1627) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


