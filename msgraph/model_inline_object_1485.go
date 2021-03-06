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

// InlineObject1485 struct for InlineObject1485
type InlineObject1485 struct {
	OldText AnyOfobject `json:"oldText,omitempty"`
	StartNum AnyOfobject `json:"startNum,omitempty"`
	NumBytes AnyOfobject `json:"numBytes,omitempty"`
	NewText AnyOfobject `json:"newText,omitempty"`
}

// NewInlineObject1485 instantiates a new InlineObject1485 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1485() *InlineObject1485 {
	this := InlineObject1485{}
	return &this
}

// NewInlineObject1485WithDefaults instantiates a new InlineObject1485 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1485WithDefaults() *InlineObject1485 {
	this := InlineObject1485{}
	return &this
}

// GetOldText returns the OldText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1485) GetOldText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.OldText
}

// GetOldTextOk returns a tuple with the OldText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1485) GetOldTextOk() (*AnyOfobject, bool) {
	if o == nil || o.OldText == nil {
		return nil, false
	}
	return &o.OldText, true
}

// HasOldText returns a boolean if a field has been set.
func (o *InlineObject1485) HasOldText() bool {
	if o != nil && o.OldText != nil {
		return true
	}

	return false
}

// SetOldText gets a reference to the given AnyOfobject and assigns it to the OldText field.
func (o *InlineObject1485) SetOldText(v AnyOfobject) {
	o.OldText = v
}

// GetStartNum returns the StartNum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1485) GetStartNum() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StartNum
}

// GetStartNumOk returns a tuple with the StartNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1485) GetStartNumOk() (*AnyOfobject, bool) {
	if o == nil || o.StartNum == nil {
		return nil, false
	}
	return &o.StartNum, true
}

// HasStartNum returns a boolean if a field has been set.
func (o *InlineObject1485) HasStartNum() bool {
	if o != nil && o.StartNum != nil {
		return true
	}

	return false
}

// SetStartNum gets a reference to the given AnyOfobject and assigns it to the StartNum field.
func (o *InlineObject1485) SetStartNum(v AnyOfobject) {
	o.StartNum = v
}

// GetNumBytes returns the NumBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1485) GetNumBytes() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumBytes
}

// GetNumBytesOk returns a tuple with the NumBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1485) GetNumBytesOk() (*AnyOfobject, bool) {
	if o == nil || o.NumBytes == nil {
		return nil, false
	}
	return &o.NumBytes, true
}

// HasNumBytes returns a boolean if a field has been set.
func (o *InlineObject1485) HasNumBytes() bool {
	if o != nil && o.NumBytes != nil {
		return true
	}

	return false
}

// SetNumBytes gets a reference to the given AnyOfobject and assigns it to the NumBytes field.
func (o *InlineObject1485) SetNumBytes(v AnyOfobject) {
	o.NumBytes = v
}

// GetNewText returns the NewText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1485) GetNewText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NewText
}

// GetNewTextOk returns a tuple with the NewText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1485) GetNewTextOk() (*AnyOfobject, bool) {
	if o == nil || o.NewText == nil {
		return nil, false
	}
	return &o.NewText, true
}

// HasNewText returns a boolean if a field has been set.
func (o *InlineObject1485) HasNewText() bool {
	if o != nil && o.NewText != nil {
		return true
	}

	return false
}

// SetNewText gets a reference to the given AnyOfobject and assigns it to the NewText field.
func (o *InlineObject1485) SetNewText(v AnyOfobject) {
	o.NewText = v
}

func (o InlineObject1485) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OldText != nil {
		toSerialize["oldText"] = o.OldText
	}
	if o.StartNum != nil {
		toSerialize["startNum"] = o.StartNum
	}
	if o.NumBytes != nil {
		toSerialize["numBytes"] = o.NumBytes
	}
	if o.NewText != nil {
		toSerialize["newText"] = o.NewText
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1485 struct {
	value *InlineObject1485
	isSet bool
}

func (v NullableInlineObject1485) Get() *InlineObject1485 {
	return v.value
}

func (v *NullableInlineObject1485) Set(val *InlineObject1485) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1485) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1485) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1485(val *InlineObject1485) *NullableInlineObject1485 {
	return &NullableInlineObject1485{value: val, isSet: true}
}

func (v NullableInlineObject1485) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1485) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


